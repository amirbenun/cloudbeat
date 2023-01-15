package awslib

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/elastic/elastic-agent-libs/logp"
)

type AWSClient interface {
}

type MultiRegionWrapper[T AWSClient] struct {
	clients map[string]T
}

func (w *MultiRegionWrapper[T]) Fetch(fetcher func(T) ([]AwsResource, error)) ([]AwsResource, error) {
	crossRegionResources := []AwsResource{}
	var err error
	for _, client := range w.clients {
		var results []AwsResource
		results, err = fetcher(client)
		if err != nil {
			continue
		}
		crossRegionResources = append(crossRegionResources, results...)
	}

	return crossRegionResources, err
}

func NewCrossRegionClient[T AWSClient](regions []string, cfg *aws.Config, factory func(*aws.Config) T, log *logp.Logger) (*MultiRegionWrapper[T], error) {
	var clientsMap = make(map[string]T, 0)
	for _, region := range regions {
		cfg.Region = region
		client := factory(cfg)
		clientsMap[region] = client
	}

	wrap := &MultiRegionWrapper[T]{
		clients: clientsMap,
	}

	fmt.Println(len(wrap.clients))

	return wrap, nil
}

func NewAllRegionClient[T AWSClient](cfg *aws.Config, factory func(*aws.Config) T, log *logp.Logger) (*MultiRegionWrapper[T], error) {
	c := ec2.NewFromConfig(*cfg)

	output, err := c.DescribeRegions(context.Background(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed DescribeRegions: %w", err)
	}

	regions := []string{}
	for _, region := range output.Regions {
		regions = append(regions, *region.RegionName)
	}

	return NewCrossRegionClient(regions, cfg, factory, log)
}

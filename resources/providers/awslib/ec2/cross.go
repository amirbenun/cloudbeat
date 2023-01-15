package ec2

import (
	context "context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/elastic/cloudbeat/resources/providers/awslib"
	"github.com/elastic/elastic-agent-libs/logp"
)

type crossRegionElasticCompute struct {
	*awslib.MultiRegionWrapper[ElasticCompute]
}

func (c *crossRegionElasticCompute) DescribeNetworkAcl(ctx context.Context) ([]awslib.AwsResource, error) {
	return c.Fetch(func(ec ElasticCompute) ([]awslib.AwsResource, error) {
		return ec.DescribeNetworkAcl(ctx)
	})
}

func (c *crossRegionElasticCompute) DescribeSecurityGroups(ctx context.Context) ([]awslib.AwsResource, error) {
	return c.Fetch(func(ec ElasticCompute) ([]awslib.AwsResource, error) {
		return ec.DescribeSecurityGroups(ctx)
	})
}

func NewCrossEC2Provider(log *logp.Logger, awsAccountID string, cfg *aws.Config) ElasticCompute {
	factory := func(c *aws.Config) ElasticCompute {
		return NewEC2Provider(log, awsAccountID, *cfg)
	}

	wrapper, _ := awslib.NewAllRegionClient(cfg, factory, log)
	return &crossRegionElasticCompute{
		MultiRegionWrapper: wrapper,
	}
}

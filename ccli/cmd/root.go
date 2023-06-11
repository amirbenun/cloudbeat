/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	devtools "github.com/elastic/beats/v7/dev-tools/mage"
	"github.com/elastic/cloudbeat/ccli/cmd/flag"
	"github.com/spf13/cobra"
	"golang.org/x/xerrors"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := NewApp().Execute()
	if err != nil {
		os.Exit(1)
	}
}

func NewApp() *cobra.Command {
	// rootCmd represents the base command when called without any subcommands
	var rootCmd = &cobra.Command{
		Use:   "ccli command [flags]",
		Short: "A brief description of your application",
		Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) { },
	}

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cloudbeat.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.AddCommand(
		NewBuildCommand(),
	)

	return rootCmd
}

func NewBuildCommand() *cobra.Command {
	flags := &flag.Flags{
		BuildFlagGroup: flag.NewBuildFlagGroup(),
	}
	cmd := &cobra.Command{
		Use:   "build [flags]",
		Short: "Build cloudbeat",
		Long:  `Build cloudbeat`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := flags.Bind(cmd); err != nil {
				return xerrors.Errorf("flag bind error: %w", err)
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := flags.Bind(cmd); err != nil {
				return xerrors.Errorf("flag bind error: %w", err)
			}
			options, err := flags.ToOptions(cmd.Version, args)
			if err != nil {
				return xerrors.Errorf("flag error: %w", err)
			}

			fmt.Println(options)
			fmt.Println("amiramir")
			buildArgs := devtools.BuildArgs{
				Env: map[string]string{
					"GOOS":   options.OS,
					"GOARCH": options.Architectre,
				},
				Name: "cloudbeat",
			}

			return devtools.Build(buildArgs)
		},
	}

	flags.AddFlags(cmd)
	return cmd
}

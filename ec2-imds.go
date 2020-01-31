package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/spf13/cobra"
)

var VERSION = "0.0.2"
var flagMaxRetries int

func handleResponse(response string, err error) {
	if err != nil {
		log.Fatalf("Error: ", err)
	}

	if len(strings.TrimSpace(response)) == 0 {
		log.Fatalf("Error: expected value in response, found empty string")
	}

	fmt.Print(response)
}

func target(args []string) string {
	if len(args) > 0 {
		return args[0]
	}
	return ""
}

func metadata() *ec2metadata.EC2Metadata {
	sess, err := session.NewSession()
	if err != nil {
		log.Fatalln("Failed to get session")
	}

	metadata := ec2metadata.New(sess, aws.NewConfig().WithMaxRetries(flagMaxRetries))

	if !metadata.Available() {
		log.Fatalln("Error: AWS instance metadata service not available")
	}

	return metadata
}

func main() {
	var rootCmd = &cobra.Command{
		Use:     "ec2-imds",
		Short:   "Query AWS Instance Metadata Service",
		Args:    cobra.ArbitraryArgs,
		Version: VERSION,
		Run: func(cmd *cobra.Command, args []string) {
			response, err := metadata().GetMetadata(target(args))
			handleResponse(response, err)
		},
	}

	rootCmd.PersistentFlags().IntVarP(&flagMaxRetries, "retries", "r", 3, "The number of times to retry requests to the instance metadata API")

	rootCmd.AddCommand(
		&cobra.Command{
			Use:   "user-data",
			Short: "Query the user-data of the current instance",
			Run: func(cmd *cobra.Command, args []string) {
				response, err := metadata().GetUserData()
				handleResponse(response, err)
			},
		},
		&cobra.Command{
			Use:   "region",
			Short: "Query the region of the current instance",
			Run: func(cmd *cobra.Command, args []string) {
				response, err := metadata().Region()
				handleResponse(response, err)
			},
		},
		&cobra.Command{
			Use:   "dynamic",
			Short: "Query Data from the dynamic",
			Args:  cobra.ArbitraryArgs,
			Run: func(cmd *cobra.Command, args []string) {
				response, err := metadata().GetDynamicData(target(args))
				handleResponse(response, err)
			},
		},
	)

	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}

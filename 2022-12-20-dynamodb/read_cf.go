package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudfront"
)

func readCloudfront() {
	cfClient, err := NewCloudfrontClient()

	if err != nil {
		return
	}

	resp, err := cfClient.GetDistribution(
		&cloudfront.GetDistributionInput{
			Id: aws.String(distributionID),
		},
	)

	fmt.Printf("%s", resp)
}

package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudfront"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func execute() error {
	items := read()

	cfClient, err := NewCloudfrontClient()

	if err != nil {
		return err
	}

	existingDistribution, err := cfClient.GetDistribution(&cloudfront.GetDistributionInput{
		Id: aws.String(distributionID),
	})

	input := generateUpdateDistributionInput(
		items,
		existingDistribution.Distribution.DistributionConfig,
		existingDistribution.ETag,
	)
	err = input.Validate()
	if err != nil {
		fmt.Printf("error : %s", err)
		return err
	}

	fmt.Println("running UpdateDistribution")
	resp, err := cfClient.UpdateDistribution(input)
	if err != nil {
		fmt.Printf("error : %s", err)
	}
	fmt.Printf("%s", resp)

	return nil
}

func generateUpdateDistributionInput(
	items []map[string]*dynamodb.AttributeValue,
	existingConfig *cloudfront.DistributionConfig,
	ifMatch *string,
) *cloudfront.UpdateDistributionInput {

	var cacheBehaviors []*cloudfront.CacheBehavior = make([]*cloudfront.CacheBehavior, 0)
	for _, item := range items {
		behavior := &cloudfront.CacheBehavior{
			AllowedMethods: &cloudfront.AllowedMethods{
				CachedMethods: &cloudfront.CachedMethods{
					Items:    []*string{aws.String("HEAD"), aws.String("GET")},
					Quantity: aws.Int64(2),
				},
				Items:    allMethods,
				Quantity: aws.Int64(int64(len(allMethods))),
			},
			CachePolicyId:          aws.String("4135ea2d-6df8-44a3-9df3-4b5a84be39ad"),
			Compress:               aws.Bool(true),
			FieldLevelEncryptionId: aws.String(""),
			FunctionAssociations: &cloudfront.FunctionAssociations{
				Quantity: aws.Int64(0),
			},
			LambdaFunctionAssociations: &cloudfront.LambdaFunctionAssociations{
				Quantity: aws.Int64(0),
			},
			PathPattern:          aws.String(fmt.Sprintf("/%s/*", *item["URLName"].S)),
			TargetOriginId:       aws.String(fmt.Sprintf("%s", *item["Domain"].S)),
			ViewerProtocolPolicy: aws.String("redirect-to-https"),
			SmoothStreaming:      aws.Bool(false),
		}
		cacheBehaviors = append(cacheBehaviors, behavior)
	}

	existingConfig.CacheBehaviors = &cloudfront.CacheBehaviors{
		Items:    cacheBehaviors,
		Quantity: aws.Int64(int64(len(items))),
	}

	return &cloudfront.UpdateDistributionInput{
		Id:                 aws.String(distributionID),
		DistributionConfig: existingConfig,
		IfMatch:            ifMatch,
	}
}

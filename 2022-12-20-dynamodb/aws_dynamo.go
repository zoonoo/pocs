package main

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func NewDynamoDBClient() (*dynamodb.DynamoDB, error) {
	sess, err := session.NewSessionWithOptions(session.Options{
		Profile: profile,
	})

	if err != nil {
		return nil, err
	}
	client := dynamodb.New(sess)
	return client, nil
}

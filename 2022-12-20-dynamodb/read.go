package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func read() []map[string]*dynamodb.AttributeValue {
	client, err := NewDynamoDBClient()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	resp, err := client.Scan(&dynamodb.ScanInput{
		TableName: aws.String(tableName),
	})

	fmt.Println(resp.Items)
	return resp.Items
}

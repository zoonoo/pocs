package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type CreateEvent struct {
	Name    string
	URLName string
	Domain  string
}

var createEvent CreateEvent = CreateEvent{
	URLName: "sendbird",
	Name:    "sendbrid123456",
	Domain:  "api.apne2prod1.kaldea.com",
}

func insert() {
	client, err := NewDynamoDBClient()
	// Set the key and attributes for the item to be created
	// key := map[string]*dynamodb.AttributeValue{
	// 	"Name": {
	// 		S: aws.String("sendbird123456"),
	// 	},
	// }
	attr := map[string]*dynamodb.AttributeValue{
		"Name": {
			S: aws.String(createEvent.Name),
		},
		"URLName": {
			S: aws.String(createEvent.URLName),
		},
		"Domain": {
			S: aws.String(createEvent.Domain),
		},
	}

	// Use a conditional write to ensure that the item does not already exist
	_, err = client.PutItem(&dynamodb.PutItemInput{
		TableName:           aws.String(tableName),
		Item:                attr,
		ConditionExpression: aws.String("attribute_not_exists(#name)"),
		// ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
		// 	":val": {
		// 		S: aws.String("sendbird123456"),
		// 	},
		// },
		ExpressionAttributeNames: map[string]*string{
			"#name": aws.String("Name"),
		},
		// ConditionExpression: aws.String("#name = :val"),
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Successfully added item with key 'organization-name-1'")
}

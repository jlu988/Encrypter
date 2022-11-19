package database

import (
	"Encrypter/models"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func (dh *DbHandler) AddInternalKey(keys models.Internal) error {
	_, err := dynamo.PutItem(&dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"OriginalKey": {
				S: aws.String(keys.OriginalKey),
			},
			"InternalKey": {
				S: aws.String(keys.InternalKey),
			},
		},
		TableName: aws.String(tableName),
	})

	if err != nil {
		return fmt.Errorf("failed to add internal key - %+v", err)
	}
	return nil
}

func (dh *DbHandler) GetInternalKey(key string) (*models.Internal, error) {
	result, err := dynamo.GetItem(&dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"OriginalKey": {
				S: aws.String(key),
			},
		},
		TableName: aws.String(tableName),
	})

	if err != nil {
		return nil, fmt.Errorf("404 error - %+v", err)
	}

	if result.Item == nil {
		return nil, nil
	}

	item := models.Internal{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		return nil, fmt.Errorf("unmarshal failed - %v", err)
	}
	return &item, nil
}

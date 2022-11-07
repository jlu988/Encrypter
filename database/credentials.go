package database

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Initialize a session that the SDK will use to load
// credentials from the shared credentials file ~/.aws/credentials
// and region from the shared configuration file ~/.aws/config.
type DbHandler struct{}

var (
	dynamo    *dynamodb.DynamoDB
	tableName = "Encryption"
	sess      = session.Must(session.NewSession(&aws.Config{Region: aws.String("us-east-1")}))
)

func init() {
	dynamo = connectDynamo()
}

func connectDynamo() (db *dynamodb.DynamoDB) {
	return dynamodb.New(sess)
}

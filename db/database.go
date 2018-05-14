package db

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var db *dynamodb.DynamoDB

// Initialize db connection
func InitDB(env string, awsRegion string, credentials *credentials.Credentials, awsEndpoint string) {
	awsConfig := &aws.Config{
		Region:      aws.String(awsRegion),
		Credentials: credentials,
	}
	if env == "development" {
		awsConfig.Endpoint = aws.String(awsEndpoint)
	}

	sess, err := session.NewSession(awsConfig)

	if err != nil {
		log.Fatal(err)
		return
	}
	db = dynamodb.New(sess)
}

// Get Dynamodb Instance
func GetDb() *dynamodb.DynamoDB {
	return db
}

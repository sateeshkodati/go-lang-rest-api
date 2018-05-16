package model

import (
	"log"
	"go-lang-rest-api-react-app/db"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// Product model
type Product struct {
	Name        string  `json:"name"`
	Label       string  `json:"label"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

var Tablename = "products"

func FindAll() []Product {
	db := db.GetDb()
	params := &dynamodb.ScanInput{
		TableName: aws.String(Tablename),
	}

	// Make the DynamoDB Query API call
	result, err := db.Scan(params)
	if err != nil {
		log.Println("failed to make Query API call", err)
	}

	products := []Product{}
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &products)
	if err != nil {
		log.Println(err)
	}

	return products
}

func FindBy(name string) (Product, bool) {
	db := db.GetDb()

	params := &dynamodb.GetItemInput{
		TableName: aws.String(Tablename),
		Key: map[string]*dynamodb.AttributeValue{
			"name": {
				S: aws.String(name),
			},
		},
	}

	result, err := db.GetItem(params)

	item := Product{}
	if err != nil {
		log.Println(err)
		return item, false
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, &item)

	if err != nil {
		log.Println(err)
		return item, false
	} else if item.Name == "" {
		return item, false
	}

	return item, true
}

func Create(product *Product) bool {
	db := db.GetDb()

	av, err := dynamodbattribute.MarshalMap(product)

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(Tablename),
	}

	_, err = db.PutItem(input)

	if err != nil {
		log.Println("Got error calling PutItem:")
		log.Println(err.Error())
		return false
	}
	return true
}

func Update(name string, product *Product) bool {
	db := db.GetDb()

	key := map[string]*dynamodb.AttributeValue{
		"name": {
			S: aws.String(name),
		},
	}

	expAttValues := map[string]*dynamodb.AttributeValue{
		":l": {
			S: aws.String(product.Label),
		},
		":p": {
			N: aws.String(strconv.FormatFloat(product.Price, 'f', 2, 64)),
		},
		":d": {
			S: aws.String(product.Description),
		},
	}
	ue := "set label = :l, price = :p, description = :d"

	if product.Label == "" {
		delete(expAttValues, ":l")
		ue = strings.Replace(ue, "label = :l, ", "", 1)
	}
	if product.Description == "" {
		delete(expAttValues, ":d")
		ue = strings.Replace(ue, ", description = :d", "", 1)
	}

	input := &dynamodb.UpdateItemInput{
		Key:                       key,
		TableName:                 aws.String(Tablename),
		ExpressionAttributeValues: expAttValues,
		UpdateExpression:          aws.String(ue),
		ReturnValues:              aws.String("UPDATED_NEW"),
	}

	_, err := db.UpdateItem(input)

	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func Remove(name string) bool {
	db := db.GetDb()

	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"name": {
				S: aws.String(name),
			},
		},
		TableName: aws.String(Tablename),
	}

	_, err := db.DeleteItem(input)

	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

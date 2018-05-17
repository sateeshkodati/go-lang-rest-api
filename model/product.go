package model

import (
	"go-lang-rest-api-react-app/db"
	"log"
	"strconv"
	"strings"
	"time"

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
	CreateOn    int64   `json:"createdon"`
	Year        int64   `json:"year"`
}

var Tablename = "products"

func FindAll() []Product {
	db := db.GetDb()

	currentTime := time.Now()

	currentYear := currentTime.Year()

	queryInput := &dynamodb.QueryInput{
		TableName: aws.String(Tablename),
		IndexName: aws.String("year-price-index"),
		KeyConditions: map[string]*dynamodb.Condition{
			"year": {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						N: aws.String(strconv.FormatInt(int64(currentYear), 10)),
					},
				},
			},
			"price": {
				ComparisonOperator: aws.String("GE"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						N: aws.String(strconv.FormatInt(int64(0), 10)),
					},
				},
			},
		},
		Limit:            aws.Int64(10),
		ScanIndexForward: aws.Bool(false),
	}

	result, err := db.Query(queryInput)

	// fmt.Println(result)

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
	// fmt.Println(result)
	if result.Item == nil {
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

	currentTime := time.Now()
	timestamp := currentTime.UnixNano() / int64(time.Millisecond)
	currentYear := currentTime.Year()
	product.Year = int64(currentYear)
	product.CreateOn = timestamp

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

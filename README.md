# Go lang Web application using dynamodb

## Setup steps 
- Create table ```products``` in AWS Dynamodb 

- add environment variables
```
$ export AWS_ACCESS_KEY_ID=AKIAIOSFODNN7EXAMPLE
$ export AWS_SECRET_ACCESS_KEY=wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY
$ export AWS_DEFAULT_REGION=us-west-2
$ export AWS_ENDPOINT="http://localhost:8000"  #if want to use local dynamodb
```


### Run REST api and Web App
```
export $GOPATH = path/to/path
$ cd path/to/$GOPATH
$ git clone https://github.com/sateeshkodati/go-lang-rest-api-react-app.git
$ go run server.go
```
- browse http://localhost:3000


### Data structure
```
type Product struct {
	Name        string  `json:"name"`
	Label       string  `json:"label"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}
```

### REST API Resource doc

- api base url => http://localhost:3000/api/v1

- GET - /products  => Get all products - http://localhost:3000/api/v1/products
```
Response: [
    {
        "name": "product3",
        "label": "aafdlj",
        "price": 100,
        "description": "111"
    },
    {
        "name": "product2",
        "label": "label ",
        "price": 100,
        "description": "desc"
    }
]
```
- GET - /products/{name}  => Get a product - http://localhost:3000/api/v1/products/product1
- POST - /products => Create product - http://localhost:3000/api/v1/products
```
Response: {
    "name": "product1",
    "label": "label1",
    "price": 155.45,
    "description": "product1 desc"
}
```
- PUT - /products/{name}  => Update product
```
Request: {
    "name": "product name", # primary key field
    "label": "Product Label",
    "price": 23.10, 
    "description": "product description
}
```
- DELETE - /products/{name} => Delete product





package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"shipt-coding-exercise/db"
	"shipt-coding-exercise/handler"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	// env := os.Getenv("APP_ENV")
	awsRegion := os.Getenv("AWS_DEFAULT_REGION")
	awsEndpoint := os.Getenv("AWS_ENDPOINT")

	// initialize data base
	db.InitDB(awsRegion, credentials.NewEnvCredentials(), awsEndpoint)

	router := mux.NewRouter().StrictSlash(true)
	apiRouter := router.PathPrefix("/api/v1").Subrouter()

	apiRouter.Methods("GET").Path("/products").HandlerFunc(handler.GetProducts)
	apiRouter.Methods("GET").Path("/products/{name}").HandlerFunc(handler.GetProduct)
	apiRouter.Methods("POST").Path("/products").HandlerFunc(handler.CreateProduct)

	apiRouter.Methods("PUT").Path("/products/{name}").HandlerFunc(handler.UpdateProduct)
	apiRouter.Methods("DELETE").Path("/products/{name}").HandlerFunc(handler.DeleteProduct)

	handler := cors.AllowAll().Handler(apiRouter)

	fmt.Println("Server running at http://localhost:3000")

	log.Fatal(http.ListenAndServe(":3000", handler))
}

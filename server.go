package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"go-lang-rest-api-react-app/db"
	"go-lang-rest-api-react-app/handler"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/gorilla/mux"
)

func main() {
	// env := os.Getenv("APP_ENV")
	awsRegion := os.Getenv("AWS_DEFAULT_REGION")
	awsEndpoint := os.Getenv("AWS_ENDPOINT")

	// initialize data base
	db.InitDB(awsRegion, credentials.NewEnvCredentials(), awsEndpoint)

	router := mux.NewRouter().StrictSlash(true)

	// s := http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))

	dir, _ := filepath.Abs("./web/build/")
	// router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))

	// fmt.Println(http.Dir(filepath.("~/../web/build/")))

	// fileHandler := http.StripPrefix("/static/", http.FileServer(http.Dir("/absolute/path/static")))
	// http.Handle("/static/", fileHandler)

	// router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(webAppPath))))
	// router.PathPrefix("/").Handler(http.FileServer(http.Dir(webAppPath)))

	apiRouter := router.PathPrefix("/api/v1").Subrouter()

	apiRouter.Methods("GET").Path("/products").HandlerFunc(handler.GetProducts)
	apiRouter.Methods("GET").Path("/products/{name}").HandlerFunc(handler.GetProduct)
	apiRouter.Methods("POST").Path("/products").HandlerFunc(handler.CreateProduct)

	apiRouter.Methods("PUT").Path("/products/{name}").HandlerFunc(handler.UpdateProduct)
	apiRouter.Methods("DELETE").Path("/products/{name}").HandlerFunc(handler.DeleteProduct)

	// handler := cors.AllowAll().Handler(apiRouter)

	// fmt.Println("Server running at http://localhost:3000")

	fmt.Println(dir)
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(dir))))
	log.Fatal(http.ListenAndServe(":3000", router))
	// log.Fatal(http.ListenAndServe(":3000", handler))
}

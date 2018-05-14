package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"shipt-coding-exercise/model"

	"github.com/gorilla/mux"
)

func validate(product model.Product) bool {
	return true
}

// Save product handler
func CreateProduct(res http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	// if _, ok := db.FindBy(name); ok {
	// 	http.Error(res, "Found", http.StatusConflict)
	// 	return
	// }

	product := new(model.Product)
	err = json.Unmarshal(body, product)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	if product.Name == "" {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	result := model.Create(product)
	if !result {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	res.Header().Set("Location", req.URL.Path+"/"+product.Name)
	res.WriteHeader(http.StatusCreated)
}

// Update product handler
func UpdateProduct(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	name := vars["name"]

	body, err := ioutil.ReadAll(req.Body)
	fmt.Println(body)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}

	product := new(model.Product)
	err = json.Unmarshal(body, product)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}

	result := model.Update(name, product)
	if !result {
		http.Error(res, "Bad Request", http.StatusBadRequest)
		return
	}
	// res.WriteHeader(http.StatusNoContent)
}

// Delete product handler
func DeleteProduct(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	name := vars["name"]

	result := model.Remove(name)
	if !result {
		http.Error(res, "Not Found", http.StatusNotFound)
		return
	}

	res.WriteHeader(http.StatusNoContent)
}

// Get product handler
func GetProduct(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	name := vars["name"]

	product, found := model.FindBy(name)
	if !found {
		http.NotFound(res, req)
		return
	}

	bytes, err := json.Marshal(product)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}

	writeJsonResponse(res, bytes)
}

// Get all products handler
func GetProducts(res http.ResponseWriter, req *http.Request) {
	products := model.FindAll()

	bytes, err := json.Marshal(products)

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}

	writeJsonResponse(res, bytes)
}

func writeJsonResponse(w http.ResponseWriter, bytes []byte) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(bytes)
}

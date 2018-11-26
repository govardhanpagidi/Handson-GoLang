package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Prodcut struct {
	ID          string
	Name        string
	Description string
	Version     string
}

func main() {

	router := mux.NewRouter()
	LoadProducts()
	router.HandleFunc("/products", GetProducts).Methods("GET")
	router.HandleFunc("/products/add", AddProduct).Methods("POST")
	router.HandleFunc("/products/delete", DeleteProduct).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

var products []Prodcut

func LoadProducts() {
	products = append(products, Prodcut{ID: "P101", Name: "RemoteLyncMonitor", Description: "WiFi Enabled Panel for fire and intrusion", Version: "1.2.20"})
	products = append(products, Prodcut{ID: "P102", Name: "RemoteLyncBridge", Description: "WiFi Enabled Panel for fire and intrusion", Version: "1.0.30"})
	products = append(products, Prodcut{ID: "P103", Name: "Wireless Detector", Description: "RF Wirelss Detector", Version: "2.2.30"})
	products = append(products, Prodcut{ID: "P104", Name: "RemoteLync Camera", Description: "RF Wirelss Camera", Version: "3.2.30"})
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(products)
}

func AddProduct(w http.ResponseWriter, r *http.Request) {
	var product Prodcut
	_ = json.NewDecoder(r.Body).Decode(&product)
	products = append(products, product)
	json.NewEncoder(w).Encode("Success")
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	for index, item := range products {
		if item.ID == id {
			products = append(products[:index], products[:index+1]...)
			break
		}
	}

}

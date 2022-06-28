package controllers

import "net/http"

func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get all Products"))
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get single Product"))
}
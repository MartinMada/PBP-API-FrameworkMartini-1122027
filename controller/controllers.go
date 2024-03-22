package controller

import (
	m "Martini/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-martini/martini"
	_ "github.com/mattn/go-sqlite3"
)

func GetProductByID(params martini.Params, w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	idParam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	var product m.Product
	err = db.QueryRow("SELECT ID, Name, Price FROM products WHERE ID=?", id).Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func CreateProduct(params martini.Params, w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	name := r.FormValue("name")
	priceParam := r.URL.Query().Get("price")

	price, err := strconv.Atoi(priceParam)
	if err != nil {
		http.Error(w, "Invalid price", http.StatusBadRequest)
		return
	}

	result, err := db.Exec("INSERT INTO products (name, price) VALUES (?, ?)", name, price)
	if err != nil {
		http.Error(w, "Failed to create product", http.StatusInternalServerError)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, "Failed to get product ID", http.StatusInternalServerError)
		return
	}

	var newProduct m.Product
	err = db.QueryRow("SELECT ID, Name, Price FROM products WHERE ID=?", id).Scan(&newProduct.ID, &newProduct.Name, &newProduct.Price)
	if err != nil {
		http.Error(w, "Failed to get product data", http.StatusInternalServerError)
		return
	}

	var response m.ProductResponse
	response.Status = 200
	response.Message = "Success"
	response.Data = newProduct

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func UpdateProduct(params martini.Params, w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	idParam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	priceParam := r.URL.Query().Get("price")
	price, err := strconv.Atoi(priceParam)
	if err != nil {
		http.Error(w, "Invalid price", http.StatusBadRequest)
		return
	}

	_, err = db.Exec("UPDATE products SET name=?, price=? WHERE ID=?", name, price, id)
	if err != nil {
		http.Error(w, "Failed to update product", http.StatusInternalServerError)
		return
	}

	updatedProduct := m.Product{
		ID:    id,
		Name:  name,
		Price: price,
	}

	var productResponse m.ProductResponse
	productResponse.Status = 200
	productResponse.Message = "Success"
	productResponse.Data = updatedProduct

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(productResponse)
}

func DeleteProduct(params martini.Params, w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	idParam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	var product m.Product
	err = db.QueryRow("SELECT ID, Name, Price FROM products WHERE ID=?", id).Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	_, err = db.Exec("DELETE FROM products WHERE ID=?", id)
	if err != nil {
		http.Error(w, "Failed to delete product", http.StatusInternalServerError)
		return
	}

	var productResponse m.ProductResponse
	productResponse.Status = 200
	productResponse.Message = "Success"
	productResponse.Data = product

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(productResponse)
}

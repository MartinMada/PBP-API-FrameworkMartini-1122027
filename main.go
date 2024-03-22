package main

import (
	"Martini/controller"
	"fmt"
	"log"
	"net/http"

	"github.com/go-martini/martini"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	m := martini.Classic()
	m.Get("/getproducts", controller.GetProductByID)
	m.Put("/updateproducts", controller.UpdateProduct)
	m.Post("/createproducts", controller.CreateProduct)
	m.Delete("/deleteproducts", controller.DeleteProduct)

	port := ":8080"
	fmt.Println("Connected to Port", port)
	log.Println("Connected to Port", port)
	log.Fatal(http.ListenAndServe(port, m))
}

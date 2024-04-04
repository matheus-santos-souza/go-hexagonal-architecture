package main

import (
	"database/sql"
	"fmt"

	db2 "github.com/matheus-santos-souza/go-hexagonal-architecture/adapters/db"
	"github.com/matheus-santos-souza/go-hexagonal-architecture/application"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "db.sqlite")
	productDbAdapter := db2.NewProductDb(db)
	productService := application.NewProductService(productDbAdapter)
	/* product, err := productService.Create("Product example", 30)
	if err != nil {
		log.Fatal(err.Error())
	}

	productService.Enable(product)*/

	product, _ := productService.Get("d6de3e60-4467-45d0-9140-a44238ac7e5b")
	fmt.Println(product.GetName(), product.GetStatus())
}

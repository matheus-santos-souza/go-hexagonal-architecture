package db_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/matheus-santos-souza/go-hexagonal-architecture/adapters/db"
	"github.com/matheus-santos-souza/go-hexagonal-architecture/application"
	"github.com/stretchr/testify/require"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	table := `CREATE TABLE products (
		id string,
		name string,
		price float,
		status string
	)`
	stmt, err := db.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func createProduct(db *sql.DB) {
	insert := `INSERT INTO products Values("abc","product test",0,"disabled")`
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func TestProductDb_Get(t *testing.T) {
	setUp()
	defer Db.Close()
	productDb := db.NewProductDb(Db)
	product, err := productDb.Get("abc")
	require.Nil(t, err)
	require.Equal(t, "product test", product.GetName())
	require.Equal(t, 0.0, product.GetPrice())
	require.Equal(t, "disabled", product.GetStatus())
}

func TestProductDb_Save(t *testing.T) {
	setUp()
	defer Db.Close()
	productDb := db.NewProductDb(Db)

	product := application.NewProduct()
	product.Name = "product test 2"
	product.Price = 25

	result, err := productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, "product test 2", result.GetName())
	require.Equal(t, 25.0, product.GetPrice())
	require.Equal(t, "disabled", result.GetStatus())

	product.Name = "product test 3"
	product.Price = 20
	product.Status = application.ENABLED

	result, err = productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, "product test 3", result.GetName())
	require.Equal(t, 20.0, product.GetPrice())
	require.Equal(t, "enabled", result.GetStatus())
}

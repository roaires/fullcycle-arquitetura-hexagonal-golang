package db_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/roaires/fullcycle-arquitetura-hexagonal-golang/adapters/db"
	"github.com/roaires/fullcycle-arquitetura-hexagonal-golang/application"
	"github.com/stretchr/testify/require"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	insertProduct(Db)
}

func createTable(db *sql.DB) {
	table := ` CREATE TABLE products (
			"id" string,
			"name" string,
			"price" float,
			"status" string	); `

	dml, err := db.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}
	dml.Exec()
}

func insertProduct(db *sql.DB) {
	table := ` INSERT INTO products VALUES("xpto", "Prod 1", 0, "disabled"); `

	cmd, err := db.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}
	cmd.Exec()
}

func TestProductDb_Get(t *testing.T) {
	setUp()
	defer Db.Close()

	productDb := db.NewProductDb(Db)
	product, err := productDb.Get("xpto")
	require.Nil(t, err)
	require.Equal(t, "Prod 1", product.GetName())
	require.Equal(t, 0.0, product.GetPrice())
	require.Equal(t, "disabled", product.GetStatus())
}

func TestProductDb_Save(t *testing.T) {
	setUp()
	defer Db.Close()

	productDb := db.NewProductDb(Db)

	product := application.NewProduct()
	product.Name = "Prod 2"
	product.Price = 99

	productResult, err := productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.Name, productResult.GetName())
	require.Equal(t, product.Price, productResult.GetPrice())
	require.Equal(t, product.Status, productResult.GetStatus())

	product.Name = "Prod 2 - Alterado"
	product.Price = 100
	product.Status = application.ENABLED

	productResult, err = productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.Name, productResult.GetName())
	require.Equal(t, product.Price, productResult.GetPrice())
	require.Equal(t, product.Status, productResult.GetStatus())
}

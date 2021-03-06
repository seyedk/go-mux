//main_test.go

package main_test

import (
	"log"
	"os"
	"testing"
	"github.com/seyedk/go-mux"
	// "net/http"
    //  "net/http/httptest"
    //  "strconv"
    //  "encoding/json"
    //  "bytes"
)

var a main.App

func TestMain(m *testing.M) {
	println(os.Getenv("APP_DB_USERNAME"))
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"),
	)
	ensureTableExists()
	code := m.Run()
	clearTable()
	os.Exit(code)
}

func ensureTableExists() {
	if _, err := a.DB.Exec(tableCreateQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	a.DB.Exec("DELETE FROM products")
	a.DB.Exec("ALTER SEQENCE product_id_seq RESTART WITH 1")
}

const tableCreateQuery = `CREATE TABLE IF NOT EXISTS products
(
    id SERIAL,
    name TEXT NOT NULL,
    price NUMERIC(10,2) NOT NULL DEFAULT 0.00,
    CONSTRAINT products_pkey PRIMARY KEY (id)
)`

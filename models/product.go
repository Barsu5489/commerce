package models

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float64
}

func SeedProducts() {
	db, err := sql.Open("sqlite3", "./pro.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStmt := `
	INSERT INTO products (name, description, price) VALUES (?, ?, ?);
	`
	products := []Product{
		{Name: "Ruby 1", Description: "This is product 1", Price: 10.00},
		{Name: "Diamond 2", Description: "This is product 2", Price: 20.00},
		{Name: "Product 3", Description: "This is product 3", Price: 30.00},
	}

	for _, product := range products {
		_, err = db.Exec(sqlStmt, product.Name, product.Description, product.Price)
		if err != nil {
			panic(err)
		}
	}
}

func GetProducts() []Product {
	db, err := sql.Open("sqlite3", "./pro.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, name, description, price FROM products")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		err = rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price)
		if err != nil {
			panic(err)
		}
		products = append(products, p)
	}

	return products
}

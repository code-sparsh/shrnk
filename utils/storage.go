package utils

import (
	"fmt"
	"net/url"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)


type DB struct {
	SQL *sql.DB
}


func NewDB() *DB {
	db,err := sql.Open("mysql", "sparsh:1234@tcp(localhost:3306)/shrtn")

	if err != nil {
		log.Fatal("Error opening database: ", err)
	}

	// ping 
	if err := db.Ping(); err != nil {
		log.Fatal("Database connection failed", err)
	}

	fmt.Println("Connected to the database successfully!")
	return &DB{SQL: db}
}


func (DB *DB) StoreURL(longURL string) (string, error) {

	// if the URL is invalid, return an error
	if _, err := url.ParseRequestURI(longURL); err != nil {
		return "", fmt.Errorf("invalid URL: %w", err)
	}

	shortCode := generateShortCode()

	_, err := DB.SQL.Exec("INSERT INTO url (shortCode, longURL) VALUES(?,?)", shortCode, longURL)
	
	if(err != nil) {
		log.Println("Failed to insert URL in database: ", err)
		return "", err
	}
	
	return shortCode, nil
}

func (DB *DB) RetrieveURL (shortCode string) (string, error) {

	var longURL string

	err := DB.SQL.QueryRow("SELECT longURL from url WHERE shortCode = ?", shortCode).Scan(&longURL)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("shortURL not found")
		}
		return "", fmt.Errorf("database error: %v", err)
	}

	return longURL, nil
}
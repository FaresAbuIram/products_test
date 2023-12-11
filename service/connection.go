package service

import (
	"encoding/json"
	"log"
	"os"
	"products/models"
)

var (
	Products    map[int]models.ProductData
	ErrorLogger *log.Logger
	InfoLogger  *log.Logger
	DebugLogger *log.Logger
)

func Connect() {
	ErrorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	InfoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	DebugLogger = log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)

	InfoLogger.Println("service,", "connection.go,", "Connect() Func")

	// fetch the data from the json file
	data, err := os.ReadFile("./service/products.json")
	if err != nil {
		ErrorLogger.Println(err)
		log.Fatal(err)
	}

	var products []models.Product
	Products = make(map[int]models.ProductData)
	// Unmarshal JSON data into the slice
	if err := json.Unmarshal(data, &products); err != nil {
		log.Fatal(err)
	}

	// create a map of products
	for _, item := range products {
		Products[item.ID] = models.ProductData{
			Name:     item.Name,
			Category: item.Category,
			Price:    item.Price,
			Quantity: item.Quantity,
		}
	}

	InfoLogger.Println("successfully finished fetching the data from the JSON file")

}

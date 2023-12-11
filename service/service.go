package service

import (
	"context"
	"fmt"
	"products/models"
)

type ProductServiceType interface {
	GetAllProduct(ctx context.Context) map[int]models.ProductData
	InsertNewProduct(ctx context.Context, data models.Product) error
	UpdateProduct(ctx context.Context, id int, data models.UpdateProductData) error
	DeleteProduct(ctx context.Context, id int) error
}

type ProductService struct{}

func NewProductService() *ProductService {
	return &ProductService{}
}

// GetAllProduct function to get all products
func (p *ProductService) GetAllProduct(ctx context.Context) map[int]models.ProductData {
	return Products
}

// InsertNewProduct function to insert new product to the Products map.
func (p *ProductService) InsertNewProduct(ctx context.Context, data models.Product) error {
	InfoLogger.Println("service,", "service.go,", "InsertNewProduct() Func")

	if _, ok := Products[data.ID]; ok {
		ErrorLogger.Printf("this Id: %d is already exist", data.ID)
		return fmt.Errorf("this Id: %d is already exist", data.ID)
	} else {
		Products[data.ID] = models.ProductData{
			Name:     data.Name,
			Category: data.Category,
			Price:    data.Price,
			Quantity: data.Quantity,
		}
	}

	InfoLogger.Println("successfully added the new product")
	return nil
}

// UpdateProduct function to update an existing product in Products map.
func (p *ProductService) UpdateProduct(ctx context.Context, id int, data models.UpdateProductData) error {
	InfoLogger.Println("service,", "service.go,", "UpdateProduct() Func")

	if product, ok := Products[id]; ok {
		updatedData := product
		if data.Name != nil {
			updatedData.Name = *data.Name
		}

		if data.Category != nil {
			updatedData.Category = *data.Category
		}

		if data.Price != nil {
			updatedData.Price = *data.Price
		}

		if data.Name != nil {
			updatedData.Quantity = *data.Quantity
		}

		Products[id] = updatedData
	} else {
		ErrorLogger.Printf("this Id: %d doesn't not exist", id)
		return fmt.Errorf("this Id: %d doesn't not exist", id)
	}

	InfoLogger.Println("successfully updated the product")
	return nil
}

// DeleteProduct function to delete an existing product from Products map.
func (p *ProductService) DeleteProduct(ctx context.Context, id int) error {
	InfoLogger.Println("service,", "service.go,", "DeleteProduct() Func")

	if _, ok := Products[id]; ok {
		delete(Products, id)
	} else {
		ErrorLogger.Printf("this Id: %d doesn't not exist", id)
		return fmt.Errorf("this Id: %d doesn't not exist", id)
	}

	InfoLogger.Println("successfully deleted the product")
	return nil
}

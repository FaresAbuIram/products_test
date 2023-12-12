package service

import (
	"fmt"
	"products/lib"
	"products/models"
)

type ProductServiceType interface {
	GetAllProduct() map[int]models.ProductData
	InsertNewProduct(data models.Product) error
	UpdateProduct(id int, data models.UpdateProductData) error
	DeleteProduct(id int) error
	SearchByCategoryAndPriceRange(models.SearchByCategoryAndPriceRangeModel) map[int]models.ProductData
	GetAvgPriceAndTotalQuantityByCategory(categry string) (float64, int, error)
}

type ProductService struct{}

func NewProductService() *ProductService {
	return &ProductService{}
}

// GetAllProduct function to get all products
func (p *ProductService) GetAllProduct() map[int]models.ProductData {
	return Products
}

// InsertNewProduct function to insert new product to the Products map.
func (p *ProductService) InsertNewProduct(data models.Product) error {
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

	Categories[data.Category] = append(Categories[data.Category], data.ID)

	InfoLogger.Println("successfully added the new product")
	return nil
}

// UpdateProduct function to update an existing product in Products map.
func (p *ProductService) UpdateProduct(id int, data models.UpdateProductData) error {
	InfoLogger.Println("service,", "service.go,", "UpdateProduct() Func")

	if product, ok := Products[id]; ok {
		updatedData := product
		if data.Name != nil {
			updatedData.Name = *data.Name
		}

		if data.Category != nil {
			Categories[product.Category] = lib.DeleteItemByValueFromSlice(Categories[product.Category], id)
			if len(Categories[product.Category]) == 0 {
				delete(Categories, product.Category)
			}

			Categories[*data.Category] = append(Categories[*data.Category], id)
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
func (p *ProductService) DeleteProduct(id int) error {
	InfoLogger.Println("service,", "service.go,", "DeleteProduct() Func")

	if product, ok := Products[id]; ok {
		delete(Products, id)
		Categories[product.Category] = lib.DeleteItemByValueFromSlice(Categories[product.Category], id)
		if len(Categories[product.Category]) == 0 {
			delete(Categories, product.Category)
		}
	} else {
		ErrorLogger.Printf("this Id: %d doesn't not exist", id)
		return fmt.Errorf("this Id: %d doesn't not exist", id)
	}

	InfoLogger.Println("successfully deleted the product")
	return nil
}

// SearchByCategoryAndPriceRange function to search for products by category and/or price range.
func (p *ProductService) SearchByCategoryAndPriceRange(options models.SearchByCategoryAndPriceRangeModel) map[int]models.ProductData {
	if options.Category == nil && options.MinPrice == nil && options.MaxPrice == nil {
		return Products
	}

	products := make(map[int]models.ProductData)

	if options.Category != nil {
		for _, id := range Categories[*options.Category] {
			// if the product price in the range, I added it to our result
			if (options.MinPrice == nil && options.MaxPrice == nil) ||
				(options.MinPrice == nil && options.MaxPrice != nil && Products[id].Price <= *options.MaxPrice) ||
				(options.MaxPrice == nil && options.MinPrice != nil && Products[id].Price >= *options.MinPrice) ||
				(options.MinPrice != nil && Products[id].Price >= *options.MinPrice && options.MaxPrice != nil && Products[id].Price <= *options.MaxPrice) {
				products[id] = Products[id]
			}
		}
	} else {
		for id, product := range Products {
			if (options.MinPrice == nil && options.MaxPrice == nil) ||
				(options.MinPrice == nil && options.MaxPrice != nil && product.Price <= *options.MaxPrice) ||
				(options.MaxPrice == nil && options.MinPrice != nil && product.Price >= *options.MinPrice) ||
				(options.MinPrice != nil && product.Price >= *options.MinPrice && options.MaxPrice != nil && product.Price <= *options.MaxPrice) {
				products[id] = product
			}

		}
	}

	return products
}

func (p *ProductService) GetAvgPriceAndTotalQuantityByCategory(categry string) (float64, int, error) {
	products := p.SearchByCategoryAndPriceRange(models.SearchByCategoryAndPriceRangeModel{
		Category: &categry,
		MinPrice: nil,
		MaxPrice: nil,
	})

	if len(products) == 0 {
		return 0, 0, fmt.Errorf("category doesn't exist")
	}

	totalPrice := 0.0
	totalQuantity := 0

	for _, product := range products {
		totalPrice += product.Price
		totalQuantity += product.Quantity
	}

	return totalPrice / float64(len(products)), totalQuantity, nil
}

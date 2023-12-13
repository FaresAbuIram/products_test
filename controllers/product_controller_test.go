package controllers_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"products/controllers"
	ProductServiceType "products/controllers/mocks"
	"products/models"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

var productServiceTypeMock *ProductServiceType.ProductServiceType
var productController *controllers.ProductController

func TestMain(m *testing.M) {
	// Setup data before running the tests
	productServiceTypeMock = new(ProductServiceType.ProductServiceType)
	productController = controllers.NewProductController(productServiceTypeMock)

	// Call flag.Parse() here if TestMain uses flags
	code := m.Run()

	// Exit with the status code from the tests
	os.Exit(code)
}

func TestGetAllProducts(t *testing.T) {
	// setup router
	router := gin.Default()

	// dummy result
	result := make(map[int]models.ProductData)

	result[1] = models.ProductData{
		Category: "test",
		Name:     "test",
		Price:    100,
		Quantity: 15,
	}
	productServiceTypeMock.On("GetAllProduct").Return(result)

	// prepare our request
	router.GET("/products", productController.GetAllProduct)
	req, err := http.NewRequest("GET", "/products", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)

	if response.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, response.Code)
	}

	var responseBody map[string]map[int]models.ProductData
	if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(responseBody["data"], result) {
		t.Errorf("Expected %v but got %v", result, responseBody)
	}
}

func TestAddProduct(t *testing.T) {
	// setup router
	router := gin.Default()

	// dummy data
	newProductData := models.Product{
		ID:       101,
		Name:     "test",
		Category: "test",
		Price:    100,
		Quantity: 1,
	}

	productServiceTypeMock.On("InsertNewProduct", newProductData).Return(nil)

	// prepare our request
	router.POST("/product", productController.AddProduct)

	// request data
	reqBody := `{"id": 101, "name": "test", "category": "test", "price": 100, "quantity": 1}`
	req, err := http.NewRequest("POST", "/product", strings.NewReader(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	result := httptest.NewRecorder()
	router.ServeHTTP(result, req)

	if result.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, result.Code)
	}

	expected := `{"message":"successfully stored the product"}`
	if result.Body.String() != expected {
		t.Errorf("Expected %q but got %q", expected, result.Body.String())
	}
}

func TestNegativeAddProduct(t *testing.T) {
	// setup router
	router := gin.Default()

	// dummy data
	newProductData := models.Product{
		ID:       100,
		Name:     "test",
		Category: "test",
		Price:    100,
		Quantity: 1,
	}
	// prepare our request
	router.POST("/product", productController.AddProduct)

	productServiceTypeMock.On("InsertNewProduct", newProductData).Return(fmt.Errorf("this Id: %d is already exist", newProductData.ID))

	// request data
	reqBody := `{"id": 100, "name": "test", "category": "test", "price": 100, "quantity": 1}`
	req, err := http.NewRequest("POST", "/product", strings.NewReader(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	result := httptest.NewRecorder()
	router.ServeHTTP(result, req)

	if result.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d but got %d", http.StatusInternalServerError, result.Code)
	}

	expected := `{"message":"Failed to store the product, error: this Id: 100 is already exist"}`
	if result.Body.String() != expected {
		t.Errorf("Expected %q but got %q", expected, result.Body.String())
	}
}

func TestUpdateProduct(t *testing.T) {
	// setup router
	router := gin.Default()

	// dummy data
	id := 100
	name := "test"
	category := "test"
	price := 100.0
	quantity := 1
	updatedProductData := models.UpdateProductData{
		Name:     &name,
		Category: &category,
		Price:    &price,
		Quantity: &quantity,
	}

	productServiceTypeMock.On("UpdateProduct", id, updatedProductData).Return(nil)

	// prepare our request
	router.PUT("/product/:id", productController.UpdateProduct)

	// request data
	reqBody := `{"name": "test", "category": "test", "price": 100, "quantity": 1}`
	req, err := http.NewRequest("PUT", fmt.Sprintf("/product/%d", id), strings.NewReader(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	result := httptest.NewRecorder()
	router.ServeHTTP(result, req)

	if result.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, result.Code)
	}

	expected := `{"message":"successfully updated the product"}`
	if result.Body.String() != expected {
		t.Errorf("Expected %q but got %q", expected, result.Body.String())
	}
}

func TestDeleteProduct(t *testing.T) {
	// setup router
	router := gin.Default()

	// dummy data
	id := 100

	productServiceTypeMock.On("DeleteProduct", id).Return(nil)

	// prepare our request
	router.DELETE("/product/:id", productController.DeleteProduct)

	// request data
	req, err := http.NewRequest("DELETE", fmt.Sprintf("/product/%d", id), nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	result := httptest.NewRecorder()
	router.ServeHTTP(result, req)

	if result.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, result.Code)
	}

	expected := `{"message":"successfully deleted the product"}`
	if result.Body.String() != expected {
		t.Errorf("Expected %q but got %q", expected, result.Body.String())
	}
}

func TestSearchByCategoryAndPriceRange(t *testing.T) {
	// setup router
	router := gin.Default()

	// dummy data
	category := "test"
	minPrice := 20.0
	maxPrice := 100.0

	result := make(map[int]models.ProductData)

	result[1] = models.ProductData{
		Category: "test",
		Name:     "test",
		Price:    99,
		Quantity: 35,
	}

	productServiceTypeMock.On("SearchByCategoryAndPriceRange", models.SearchByCategoryAndPriceRangeModel{
		Category: &category,
		MinPrice: &minPrice,
		MaxPrice: &maxPrice,
	}).Return(result)

	// prepare our request
	router.GET("/seach-by-categry-and-price-range", productController.SearchByCategoryAndPriceRange)

	// request data
	req, err := http.NewRequest("GET", "/seach-by-categry-and-price-range", nil)
	if err != nil {
		t.Fatal(err)
	}

	query := req.URL.Query()
	query.Add("category", category)
	query.Add("minPrice", strconv.FormatFloat(minPrice, 'f', -1, 64))
	query.Add("maxPrice", strconv.FormatFloat(maxPrice, 'f', -1, 64))
	req.URL.RawQuery = query.Encode()

	req.Header.Set("Content-Type", "application/json")

	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)

	if response.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, response.Code)
	}

	var responseBody map[string]map[int]models.ProductData
	if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(responseBody["data"], result) {
		t.Errorf("Expected %v but got %v", result, responseBody)
	}
}

func TestGetAvgPriceAndTotalQuantityByCategory(t *testing.T) {
	// setup router
	router := gin.Default()

	// dummy data
	category := "test"

	result := models.AveragePriceAndTotalQuantity{
		AveragePrice:  75,
		TotalQuantity: 20,
	}

	productServiceTypeMock.On("GetAvgPriceAndTotalQuantityByCategory", category).Return(result.AveragePrice, result.TotalQuantity, nil)

	// prepare our request
	router.GET("/get-average-price-and-total-quantity-by-category", productController.GetAvgPriceAndTotalQuantityByCategory)

	// request data
	req, err := http.NewRequest("GET", "/get-average-price-and-total-quantity-by-category", nil)
	if err != nil {
		t.Fatal(err)
	}

	query := req.URL.Query()
	query.Add("category", category)
	req.URL.RawQuery = query.Encode()

	req.Header.Set("Content-Type", "application/json")

	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)

	if response.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, response.Code)
	}

	var responseBody map[string]models.AveragePriceAndTotalQuantity
	if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
		t.Fatal(err)
	}
	
	if !reflect.DeepEqual(responseBody["data"], result) {
		t.Errorf("Expected %v but got %v", result, responseBody)
	}
}

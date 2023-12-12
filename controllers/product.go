package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"products/models"
	"products/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ErrorResponseStatus function using return response with error status
func ErrorResponseStatus(context *gin.Context, status int, message string) {
	context.JSON(status, gin.H{"message": message})
}

// SuccessResponseStatus function using return a response with Ok status
func SuccessResponseStatus(context *gin.Context, status int, data string) {
	context.JSON(status, gin.H{"message": data})
}

type LoggerLevels struct {
	ErrorLogger *log.Logger
	InfoLogger  *log.Logger
	DebugLogger *log.Logger
}

type ProductController struct {
	ProductService service.ProductServiceType
	Logger         LoggerLevels
}

func NewProductController(productServiceType service.ProductServiceType) *ProductController {
	return &ProductController{
		ProductService: productServiceType,
		Logger: LoggerLevels{
			ErrorLogger: log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
			InfoLogger:  log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime),
			DebugLogger: log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile),
		},
	}
}

// GetAllProduct Function used to get all products
// @Summary      get all products
// @Description  This route uses to get all products from the Products map
// @Produce      json
// @Success      200  {object} map[int]models.ProductData
// @Tags Products
// @Router       /products [get]
func (p *ProductController) GetAllProduct(context *gin.Context) {
	p.Logger.InfoLogger.Println("controllers,", "product.go,", "GetAllProduct() Func")

	result := p.ProductService.GetAllProduct()

	context.JSON(http.StatusOK, gin.H{"data": result})
}

// AddProduct Function used to add new product
// @Summary      add new product
// @Description  This route uses to add new product to Products map
// @Produce      json
// @Param        body  body models.Product true "new data"
// @Success      200  {object} models.MessageResponse
// @Failure      400  {object}	models.MessageResponse
// @Failure      500  {object}	models.MessageResponse
// @Tags Product Modification
// @Router       /product [post]
func (p *ProductController) AddProduct(context *gin.Context) {
	p.Logger.InfoLogger.Println("controllers,", "product.go,", "AddProduct() Func")

	var productData models.Product

	if err := context.BindJSON(&productData); err != nil {
		p.Logger.ErrorLogger.Println("missing data: ", err)
		ErrorResponseStatus(context, http.StatusBadRequest, fmt.Sprintf("missing data: %s", err.Error()))
		return
	}

	// save the resource in the Product map
	err := p.ProductService.InsertNewProduct(productData)

	if err != nil {
		p.Logger.ErrorLogger.Println("Failed to store the product: ", err)
		ErrorResponseStatus(context, http.StatusInternalServerError, fmt.Sprintf("Failed to store the product, error: %s", err.Error()))
		return
	}

	SuccessResponseStatus(context, http.StatusOK, "successfully stored the product")
}

// UpdateProduct Function used to update a product
// @Summary      update a product
// @Description  This route uses to update an existing product in Products map
// @Produce      json
// @Param        id  path int true "product id"
// @Param        body  body models.ProductData true "updated data"
// @Success      200  {object} models.MessageResponse
// @Failure      400  {object}	models.MessageResponse
// @Failure      500  {object}	models.MessageResponse
// @Tags Product Modification
// @Router       /product/{id} [put]
func (p *ProductController) UpdateProduct(context *gin.Context) {
	p.Logger.InfoLogger.Println("controllers,", "product.go,", "UpdateProduct() Func")

	// get the product id
	id := context.Param("id")

	if id == "" {
		p.Logger.ErrorLogger.Println("missing id")
		ErrorResponseStatus(context, http.StatusBadRequest, "missing id")
		return
	}

	// the new data
	var productData models.UpdateProductData
	if err := context.BindJSON(&productData); err != nil {
		p.Logger.ErrorLogger.Println("missing data: ", err)
		ErrorResponseStatus(context, http.StatusBadRequest, "missing data")
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		p.Logger.ErrorLogger.Println("id is not an interger")
		ErrorResponseStatus(context, http.StatusBadRequest, "id is not an interger")
		return
	}

	// save the resource in the Product map
	err = p.ProductService.UpdateProduct(productId, productData)

	if err != nil {
		p.Logger.ErrorLogger.Println("Failed to update the product: ", err)
		ErrorResponseStatus(context, http.StatusInternalServerError, fmt.Sprintf("Failed to update the product, error: %s", err.Error()))
		return
	}

	SuccessResponseStatus(context, http.StatusOK, "successfully updated the product")
}

// DeleteProduct Function used to delete a product
// @Summary      delete a product
// @Description  This route uses to delete an existing product from Products map
// @Produce      json
// @Param        id  path int true "product id"
// @Success      200  {object} models.MessageResponse
// @Failure      400  {object}	models.MessageResponse
// @Failure      500  {object}	models.MessageResponse
// @Tags Product Modification
// @Router       /product/{id} [delete]
func (p *ProductController) DeleteProduct(context *gin.Context) {
	p.Logger.InfoLogger.Println("controllers,", "product.go,", "DeleteProduct() Func")

	// get the product id
	id := context.Param("id")

	if id == "" {
		p.Logger.ErrorLogger.Println("missing id")
		ErrorResponseStatus(context, http.StatusBadRequest, "missing id")
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		p.Logger.ErrorLogger.Println("id is not an interger")
		ErrorResponseStatus(context, http.StatusBadRequest, "id is not an interger")
		return
	}

	// save the resource in the Product map
	err = p.ProductService.DeleteProduct(productId)

	if err != nil {
		p.Logger.ErrorLogger.Println("Failed to delete the product: ", err)
		ErrorResponseStatus(context, http.StatusInternalServerError, fmt.Sprintf("Failed to delete the product, error: %s", err.Error()))
		return
	}

	SuccessResponseStatus(context, http.StatusOK, "successfully deleted the product")
}

// SearchByCategoryAndPriceRange Function used to to search for products by category and/or price range.
// @Summary      search for products by category and/or price range
// @Description  This route uses to search for products by category and/or price range from Products map
// @Produce      json
// @Param        query  query models.SearchByCategoryAndPriceRangeModel true "search query"
// @Success      200  {object}  map[int]models.ProductData
// @Failure      400  {object}	models.MessageResponse
// @Failure      500  {object}	models.MessageResponse
// @Tags Product Functionality
// @Router       /seach-by-categry-and-price-range [get]
func (p *ProductController) SearchByCategoryAndPriceRange(context *gin.Context) {
	p.Logger.InfoLogger.Println("controllers,", "product.go,", "SearchByCategoryAndPriceRange() Func")

	var query models.SearchByCategoryAndPriceRangeModel
	category := context.Request.URL.Query().Get("category")
	minPrice := context.Request.URL.Query().Get("minPrice")
	maxPrice := context.Request.URL.Query().Get("maxPrice")

	if category != "" {
		query.Category = &category
	}

	if minPrice != "" {
		minPriceFloat, err := strconv.ParseFloat(minPrice, 64)
		if err != nil {
			p.Logger.ErrorLogger.Println("min price is not a number")
			ErrorResponseStatus(context, http.StatusBadRequest, "min price is not a numer")
			return
		}
		query.MinPrice = &minPriceFloat
	}

	if maxPrice != "" {
		maxPriceFloat, err := strconv.ParseFloat(maxPrice, 64)
		if err != nil {
			p.Logger.ErrorLogger.Println("max price is not a number")
			ErrorResponseStatus(context, http.StatusBadRequest, "max price is not a number")
			return
		}
		query.MaxPrice = &maxPriceFloat
	}

	result := p.ProductService.SearchByCategoryAndPriceRange(query)

	context.JSON(http.StatusOK, gin.H{"data": result})
}

// GetAvgPriceAndTotalQuantityByCategory Function used to get average price and total quantity of all products in a category
// @Summary      get average price and total quantity of all products in a category
// @Description  This route uses to get average price and total quantity of all products in a category
// @Produce      json
// @Param        category  query string true "category name"
// @Success      200  {object} models.AveragePriceAndTotalQuantity
// @Failure      400  {object}	models.MessageResponse
// @Failure      500  {object}	models.MessageResponse
// @Tags Data Aggregation
// @Router       /get-average-price-and-total-quantity-by-category [get]
func (p *ProductController) GetAvgPriceAndTotalQuantityByCategory(context *gin.Context) {
	p.Logger.InfoLogger.Println("controllers,", "product.go,", "GetAvgPriceAndTotalQuantityByCategory() Func")

	category := context.Request.URL.Query().Get("category")

	if category == "" {
		p.Logger.ErrorLogger.Println("category is a required field")
		ErrorResponseStatus(context, http.StatusBadRequest, "category is a required field")
		return
	}

	avgPrice, totalQuantity, err := p.ProductService.GetAvgPriceAndTotalQuantityByCategory(category)

	if err != nil {
		p.Logger.ErrorLogger.Println(err)
		ErrorResponseStatus(context, http.StatusBadRequest, err.Error())
		return
	}

	result := models.AveragePriceAndTotalQuantity{
		AveragePrice:  avgPrice,
		TotalQuantity: totalQuantity,
	}

	context.JSON(http.StatusOK, gin.H{"data": result})
}

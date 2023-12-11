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

	result := p.ProductService.GetAllProduct(context)

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
// @Tags Products
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
	err := p.ProductService.InsertNewProduct(context, productData)

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
// @Tags Products
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
	err = p.ProductService.UpdateProduct(context, productId, productData)

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
// @Tags Products
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
	err = p.ProductService.DeleteProduct(context, productId)

	if err != nil {
		p.Logger.ErrorLogger.Println("Failed to delete the product: ", err)
		ErrorResponseStatus(context, http.StatusInternalServerError, fmt.Sprintf("Failed to delete the product, error: %s", err.Error()))
		return
	}

	SuccessResponseStatus(context, http.StatusOK, "successfully deleted the product")
}

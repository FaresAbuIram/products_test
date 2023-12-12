package routes

import (
	"products/controllers"
	"products/docs"
	"products/service"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Setup(router *gin.Engine) {
	docs.SwaggerInfo.Title = "Swagger Product API"
	docs.SwaggerInfo.Description = "This is a Product server."
	docs.SwaggerInfo.Version = "2.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Schemes = []string{"http"}

	productService := service.NewProductService()
	productController := controllers.NewProductController(productService)

	router.Use(static.Serve("/", static.LocalFile("./website/dist", true)))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Product Modification
	router.GET("/products", productController.GetAllProduct)
	router.POST("/product", productController.AddProduct)
	router.PUT("/product/:id", productController.UpdateProduct)
	router.DELETE("/product/:id", productController.DeleteProduct)

	// Product Functionality
	router.GET("/seach-by-categry-and-price-range", productController.SearchByCategoryAndPriceRange)

	// Data Aggregation
	router.GET("/get-average-price-and-total-quantity-by-category", productController.GetAvgPriceAndTotalQuantityByCategory)
}

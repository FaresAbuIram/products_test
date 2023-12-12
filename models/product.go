package models

type Product struct {
	ID       int     `bson:"id" binding:"required"`
	Name     string  `bson:"name" binding:"required"`
	Category string  `bson:"category" binding:"required"`
	Price    float64 `bson:"price" binding:"required"`
	Quantity int     `bson:"quantity" binding:"required"`
}

type ProductData struct {
	Name     string  `bson:"name"`
	Category string  `bson:"category"`
	Price    float64 `bson:"price"`
	Quantity int     `bson:"quantity"`
}

type UpdateProductData struct {
	Name     *string
	Category *string
	Price    *float64
	Quantity *int
}

type MessageResponse struct {
	Message string `bson:"message"`
}

type SearchByCategoryAndPriceRangeModel struct {
	Category *string  `bson:"category"`
	MinPrice *float64 `bson:"minPrice"`
	MaxPrice *float64 `bson:"maxPrice"`
}

type AveragePriceAndTotalQuantity struct {
	AveragePrice  float64
	TotalQuantity int
}

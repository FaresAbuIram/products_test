// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/get-average-price-and-total-quantity-by-category": {
            "get": {
                "description": "This route uses to get average price and total quantity of all products in a category",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Data Aggregation"
                ],
                "summary": "get average price and total quantity of all products in a category",
                "parameters": [
                    {
                        "type": "string",
                        "description": "category name",
                        "name": "category",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.AveragePriceAndTotalQuantity"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.MessageResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.MessageResponse"
                        }
                    }
                }
            }
        },
        "/product": {
            "post": {
                "description": "This route uses to add new product to Products map",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product Modification"
                ],
                "summary": "add new product",
                "parameters": [
                    {
                        "description": "new data",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Product"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.MessageResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.MessageResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.MessageResponse"
                        }
                    }
                }
            }
        },
        "/product/{id}": {
            "put": {
                "description": "This route uses to update an existing product in Products map",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product Modification"
                ],
                "summary": "update a product",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "product id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "updated data",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ProductData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.MessageResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.MessageResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.MessageResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "This route uses to delete an existing product from Products map",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product Modification"
                ],
                "summary": "delete a product",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "product id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.MessageResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.MessageResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.MessageResponse"
                        }
                    }
                }
            }
        },
        "/products": {
            "get": {
                "description": "This route uses to get all products from the Products map",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "get all products",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "$ref": "#/definitions/models.ProductData"
                            }
                        }
                    }
                }
            }
        },
        "/seach-by-categry-and-price-range": {
            "get": {
                "description": "This route uses to search for products by category and/or price range from Products map",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product Functionality"
                ],
                "summary": "search for products by category and/or price range",
                "parameters": [
                    {
                        "type": "string",
                        "name": "category",
                        "in": "query"
                    },
                    {
                        "type": "number",
                        "name": "maxPrice",
                        "in": "query"
                    },
                    {
                        "type": "number",
                        "name": "minPrice",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "$ref": "#/definitions/models.ProductData"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.MessageResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.MessageResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.AveragePriceAndTotalQuantity": {
            "type": "object",
            "properties": {
                "averagePrice": {
                    "type": "number"
                },
                "totalQuantity": {
                    "type": "integer"
                }
            }
        },
        "models.MessageResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "models.Product": {
            "type": "object",
            "required": [
                "category",
                "id",
                "name",
                "price",
                "quantity"
            ],
            "properties": {
                "category": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                },
                "quantity": {
                    "type": "integer"
                }
            }
        },
        "models.ProductData": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                },
                "quantity": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

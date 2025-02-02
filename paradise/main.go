package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var products = []map[string]interface{}{
	{"id": 1, "name": "Coconut", "price": 20},
	{"id": 2, "name": "Pineapple Juice", "price": 7},
}

func main() {
	r := gin.Default()

	r.GET("/message", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to Paradise 🌴"})
	})

	r.GET("/products", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"products": products})
	})

	r.POST("/products", func(c *gin.Context) {
		var newProduct map[string]interface{}
		if err := c.BindJSON(&newProduct); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}
		newProduct["id"] = len(products) + 1 // Auto-increment ID
		products = append(products, newProduct)
		c.JSON(http.StatusCreated, gin.H{"message": "Product added", "product": newProduct})
	})

	r.PUT("/products/:id", func(c *gin.Context) {
		var updatedProduct map[string]interface{}
		id := c.Param("id") // Get the product ID from the URL parameter
		found := false

		println("Updating Product ID:", id)

		if err := c.BindJSON(&updatedProduct); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}
		for i, product := range products {
			if product["id"] == id {
				products[i] = updatedProduct
				updatedProduct["id"] = id // Make sure the ID remains the same
				found = true
				break
			}
		}

		if !found {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Product updated", "product": updatedProduct})
	})

	r.DELETE("/products/:id", func(c *gin.Context) {
		id := c.Param("id")
		found := false

		// Debug: Check the ID being passed
		println("Deleting Product ID:", id)

		for i, product := range products {
			if product["id"] == id {
				products = append(products[:i], products[i+1:]...) // Remove product from slice
				found = true
				break
			}
		}

		if !found {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
	})

	r.Run(":8000")
}

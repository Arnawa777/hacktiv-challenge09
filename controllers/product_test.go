package controllers

import (
	"challenge-08/database"
	"challenge-08/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Test case for when products are found
func TestGetAllProducts(t *testing.T) {

	// Initialize Gin context and database connection
	gin.SetMode(gin.TestMode)
	gin.Default()
	database.GetDB()

	// Create some test products
	testProducts := []models.Product{
		{UserID: 2, Title: "Edit sama user", Description: "Biar 2313 ssama user"},
		{UserID: 1, Title: "Test Product", Description: "Test Description"},
	}

	// Set up mock request
	req, _ := http.NewRequest("GET", "/products", nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// Call controller function
	GetAllProducts(c)

	// Verify response
	assert.Equal(t, http.StatusOK, w.Code)

	expected := testProducts
	var actual []models.Product
	err := c.ShouldBindJSON(&actual)
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)

}

// Test case for when no products are found
func TestGetAllProductsNotFound(t *testing.T) {

	// Create a new HTTP request
	req, err := http.NewRequest("GET", "/products", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new response recorder to record the HTTP response
	res := httptest.NewRecorder()

	// Set up the Gin context
	ctx, _ := gin.CreateTestContext(res)
	ctx.Request = req

	// Call the GetAllProducts function
	GetAllProducts(ctx)

	// Assert that the response status code is 404 Not Found
	assert.Equal(t, http.StatusNotFound, res.Code)
}

func TestGetProductFoundById(t *testing.T) {
	database.MockDB()
	// Create a new Gin router instance
	r := gin.Default()

	// Register the handler function
	r.GET("/products/:productID", GetProduct)

	// Create a new HTTP request to get the product by ID
	req, _ := http.NewRequest("GET", "/products/8", nil)

	// Create a new HTTP response recorder
	rec := httptest.NewRecorder()

	// Call the Gin router to handle the request
	r.ServeHTTP(rec, req)

	// Assert that the response status code is OK (200)
	assert.Equal(t, http.StatusOK, rec.Code)

	// Assert that the response body is the JSON representation of the product
	expected := `{"id": 8, "created_at": "2023-04-10T04:20:49.6960942+07:00","updated_at": "2023-04-10T04:20:49.6960942+07:00", "user_id": 1, "title": "Test Product", "description": "Test Description"}`
	assert.JSONEq(t, expected, rec.Body.String())
}

func TestGetProductNotFoundById(t *testing.T) {
	// Create a new Gin router instance
	r := gin.Default()

	// Register the handler function
	r.GET("/products/:productID", GetProduct)

	// Create a new HTTP request to get a non-existing product by ID
	req, _ := http.NewRequest("GET", "/products/999", nil)

	// Create a new HTTP response recorder
	rec := httptest.NewRecorder()

	// Call the Gin router to handle the request
	r.ServeHTTP(rec, req)

	// Assert that the response status code is Not Found (404)
	assert.Equal(t, http.StatusNotFound, rec.Code)

	// Assert that the response body is the error message "Product not found"
	assert.JSONEq(t, `{"message":"Product not found"}`, rec.Body.String())
}

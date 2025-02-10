package server

import (
	"ecomm-backend/internal/database"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (s *Server) allProducts (c *gin.Context) {
    var products []database.Product

    result := db.Select("name", "description", "summary", "cover", "id", "created_at").Find(&products)
    if result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
    }

    c.JSON(http.StatusFound, products)
}

func (s *Server) newProduct (c *gin.Context) {
    var product database.Product

    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    result := db.Create(&product)
    if result.Error != nil {
        c.JSON(http.StatusNotModified, gin.H{"error": result.Error.Error()})
    }
    c.JSON(http.StatusCreated, gin.H{"success": "product added"})
}

func (s *Server) getProductById (c *gin.Context) {
    productId := c.Param("id")
    var product database.Product
    
    result := db.Where("id = ?", productId).Take(&product)
    if errors.Is(result.Error, gorm.ErrRecordNotFound) {
        c.JSON(http.StatusNotFound, gin.H{"error": gorm.ErrRecordNotFound.Error()})
    }
    
    c.JSON(http.StatusFound, product)
}

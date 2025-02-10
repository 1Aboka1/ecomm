package server

import (
	"ecomm-backend/internal/database"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (s *Server) allCategories (c *gin.Context) {
    var categories []database.Category

    result := db.Select("name", "id", "created_at").Find(&categories)
    if result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
    }

    c.JSON(http.StatusOK, categories)
}

func (s *Server) newCategory (c *gin.Context) {
    var category database.Category

    if err := c.ShouldBindJSON(&category); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    result := db.Create(&category)
    if result.Error != nil {
        c.JSON(http.StatusNotModified, gin.H{"error": result.Error.Error()})
    }
    c.JSON(http.StatusCreated, gin.H{"success": "category added"})
}

func (s *Server) getCategoryById (c *gin.Context) {
    categoryId, ok := c.GetQuery("id")
    if !ok {
        c.JSON(http.StatusBadRequest, gin.H{"error": "input id of category"})
        return
    }
    var category database.Category
    
    result := db.Where("id = ?", categoryId).Take(&category)
    if errors.Is(result.Error, gorm.ErrRecordNotFound) {
        c.JSON(http.StatusNotFound, gin.H{"error": gorm.ErrRecordNotFound.Error()})
    }
    
    c.JSON(http.StatusOK, category)
}

func (s *Server) allSubCategoriesByCategoryId(c *gin.Context) {
    categoryId := c.Param("id")
    var categories []database.SubCategory

    result := db.Select("name", "id", "created_at").Where("category_id = ?", categoryId).Find(&categories)
    if result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
    }

    c.JSON(http.StatusOK, categories)
}

func (s *Server) newSubCategory (c *gin.Context) {
    var category database.SubCategory

    if err := c.ShouldBindJSON(&category); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    result := db.Create(&category)
    if result.Error != nil {
        c.JSON(http.StatusNotModified, gin.H{"error": result.Error.Error()})
    }
    c.JSON(http.StatusCreated, gin.H{"success": "category added"})
}

func (s *Server) getSubCategoryById (c *gin.Context) {
    categoryId, ok := c.GetQuery("id")
    if !ok {
        c.JSON(http.StatusBadRequest, gin.H{"error": "input id of subcategory"})
        return
    }
    var category database.SubCategory
    
    result := db.Where("id = ?", categoryId).Take(&category)
    if errors.Is(result.Error, gorm.ErrRecordNotFound) {
        c.JSON(http.StatusNotFound, gin.H{"error": gorm.ErrRecordNotFound.Error()})
    }
    
    c.JSON(http.StatusOK, category)
}

package server

import (
	"ecomm-backend/internal/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) allProducts (c *gin.Context) {
    var products []database.Product

    result := db.Select("name", "description", "summary", "cover", "id", "created_at").Find(&products)
    if result.Error != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
    }

    c.JSON(http.StatusOK, products)
}

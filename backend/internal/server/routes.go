package server

import (
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

var (
    port = os.Getenv("FRONTEND_PORT")
    app_env = os.Getenv("APP_ENV")
)

func (s *Server) RegisterRoutes() http.Handler {
    r := gin.Default()

    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:" + port}, // Add your frontend URL
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
        AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
        AllowCredentials: true, // Enable cookies/auth
    }))

    v1 := r.Group("/v1")
    {
        v1.GET("/auth", s.SignUp)
    }


    return r
}

package server

import (
	"ecomm-backend/internal/auth"
	"ecomm-backend/internal/database"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)



var (
    appTimeout = time.Second * 10
    frontendPort = os.Getenv("FRONTEND_PORT")
    redisPort = os.Getenv("REDIS_PORT")
    appEnv = os.Getenv("APP_ENV")

    db *gorm.DB

    allowOnlyAdmin = auth.Authorizer([]uint{database.Admin})
    allowOnlyCustomerAndGreater = auth.Authorizer([]uint{database.Customer, database.Admin})
    allowOnlyGuest = auth.Authorizer([]uint{database.Guest})
    allowEveryone = auth.Authorizer([]uint{database.Guest, database.Admin, database.Customer})
)

func (s *Server) RegisterRoutes() http.Handler {
    r := gin.Default()

    db = database.OrmDb

    store, _ := redis.NewStore(10, "tcp", "localhost:" + redisPort, "", []byte("secret"))
    r.Use(sessions.Sessions("session0", store))


    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:" + frontendPort}, // Add your frontend URL
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
        AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
        AllowCredentials: true, // Enable cookies/auth
    }))

    v1 := r.Group("/v1")

    {
        authRoute := v1.Group("/auth")
        authRoute.POST("/sign_out", allowOnlyCustomerAndGreater, s.signOut)
        authRoute.Use(allowOnlyGuest)
        authRoute.POST("/send_otp", s.sendSMS)
        authRoute.POST("/verify_otp", s.verifySMS)
        authRoute.POST("/sign_in", s.signIn)
        authRoute.POST("/sign_up", s.signUp)
    }

    {
        productRoute := v1.Group("/product")
        productRoute.GET("/all", allowOnlyCustomerAndGreater, s.allProducts)
        productRoute.POST("/new", allowOnlyAdmin, s.newProduct)
        productRoute.GET("/by_id", allowEveryone, s.getProductById)
    }

    {
        categoryRoute := v1.Group("/category")
        categoryRoute.GET("/all", allowOnlyCustomerAndGreater, s.allCategories)
        categoryRoute.POST("/new", allowOnlyAdmin, s.newCategory)
        categoryRoute.GET("/by_id", allowEveryone, s.getCategoryById)
    }

    {
        subCategoryRoute := v1.Group("/sub_category")
        subCategoryRoute.GET("/all", allowOnlyCustomerAndGreater, s.allSubCategoriesByCategoryId)
        subCategoryRoute.POST("/new", allowOnlyAdmin, s.newSubCategory)
        subCategoryRoute.GET("/by_id", allowEveryone, s.getSubCategoryById)
    }

    return r
}

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
)

func (s *Server) RegisterRoutes() http.Handler {
  r := gin.Default()

  // get DB
  db = database.OrmDb

  // init redis
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
    authRoute.POST("/sign_out", auth.AllowOnlyCustomerAndGreater, s.signOut)
    authRoute.Use(auth.AllowOnlyGuest)
    authRoute.POST("/send_otp", s.sendSMS)
    authRoute.POST("/verify_otp", s.verifySMS)
        authRoute.POST("/sign_in", s.signIn)
        authRoute.POST("/sign_up", s.signUp)
    }

    {
      // TODO: need to find way to attach auth middleware to specific routes like cart etc.
        graphRoute := v1.Group("/graph")
        graphRoute.POST("/query_product", productGraphHandler())
        graphRoute.POST("/query_cart", cartGraphHandler())
        graphRoute.POST("/query_category", categoryGraphHandler())
        graphRoute.POST("/query_sku_attribute", skuAttributeGraphHandler())
        graphRoute.GET("/", playgroundHandler())
    }

    return r
}

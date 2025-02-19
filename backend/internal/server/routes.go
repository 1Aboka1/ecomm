package server

import (
	"context"
	"ecomm-backend/graph"
	"ecomm-backend/internal/auth"
	"ecomm-backend/internal/database"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/vektah/gqlparser/v2/ast"
	"gorm.io/gorm"
)



var (
  appTimeout = time.Second * 10
  frontendPort = os.Getenv("FRONTEND_PORT")
  redisPort = os.Getenv("REDIS_PORT")
  appEnv = os.Getenv("APP_ENV")

  db *gorm.DB
)

func graphqlHandler() gin.HandlerFunc {
  h := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
  h.AddTransport(transport.Options{})
  h.AddTransport(transport.GET{})
  h.AddTransport(transport.POST{})

  h.SetQueryCache(lru.New[*ast.QueryDocument](1000))

  h.Use(extension.Introspection{})
  h.Use(extension.AutomaticPersistedQuery{
    Cache: lru.New[string](100),
  })
  return func(c *gin.Context) {
    ctx := context.WithValue(c.Request.Context(), "ginContextKey", c)
    c.Request = c.Request.WithContext(ctx)

    h.ServeHTTP(c.Writer, c.Request)
  }
}
// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
  h := playground.Handler("GraphQL", "/v1/graph/query")
  return func(c *gin.Context) {
    h.ServeHTTP(c.Writer, c.Request)
  }
}

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
        graphRoute.POST("/query", graphqlHandler())
        graphRoute.GET("/", playgroundHandler())
    }

    return r
}

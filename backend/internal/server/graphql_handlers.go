package server

import (
	"context"
	product_graph "ecomm-backend/graphqls/product/graph"
	cart_graph "ecomm-backend/graphqls/cart/graph"
	category_graph "ecomm-backend/graphqls/category/graph"
	sku_attribute_graph "ecomm-backend/graphqls/sku_attribute/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/vektah/gqlparser/v2/ast"
)

func productGraphHandler() gin.HandlerFunc {
  h := handler.New(product_graph.NewExecutableSchema(product_graph.Config{Resolvers: &product_graph.Resolver{}}))
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

func cartGraphHandler() gin.HandlerFunc {
  h := handler.New(cart_graph.NewExecutableSchema(cart_graph.Config{Resolvers: &cart_graph.Resolver{}}))
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

func categoryGraphHandler() gin.HandlerFunc {
  h := handler.New(category_graph.NewExecutableSchema(category_graph.Config{Resolvers: &category_graph.Resolver{}}))
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

func skuAttributeGraphHandler() gin.HandlerFunc {
  h := handler.New(sku_attribute_graph.NewExecutableSchema(sku_attribute_graph.Config{Resolvers: &sku_attribute_graph.Resolver{}}))
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

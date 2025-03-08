package utils

import (
	"context"
	"errors"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func ctxToGinCtx(ctx context.Context) (*gin.Context, error) {
  ginContext := ctx.Value("ginContextKey")
	if ginContext == nil {
		err := errors.New("could not retrieve gin.Context")
		return nil, err
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		err := errors.New("gin.Context has wrong type")
		return nil, err
	}
  return gc, nil
}

func GetSession(ctx context.Context) (sessions.Session, error) {
  c, err := ctxToGinCtx(ctx);
  if err != nil {
    return nil, err
  }

  session := sessions.Default(c)
  return session, nil
}

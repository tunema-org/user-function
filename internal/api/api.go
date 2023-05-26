package api

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tunema-org/user-function/internal/backend"
)

type handler struct {
	backend *backend.Backend
}

func NewHandler(ctx context.Context, backend *backend.Backend) *gin.Engine {
	h := &handler{
		backend: backend,
	}

	r := gin.Default()

	r.POST("/users/register", h.Register)
	r.POST("/users/login", h.Login)
	r.GET("/users/me", h.Me)
	r.GET("/users/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "healthy",
		})
	})

	return r

	// return func(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	// 	switch req.Resource {
	// 	case "/users/register":
	// 		if req.HTTPMethod != http.MethodPost {
	// 			return JSONMethodNotAllowed(http.MethodPost)
	// 		}

	// 		return h.Register(ctx, req)
	// 	case "/users/login":
	// 		if req.HTTPMethod != http.MethodPost {
	// 			return JSONMethodNotAllowed(http.MethodPost)
	// 		}

	// 		return h.Login(ctx, req)
	// 	case "/users/me":
	// 		if req.HTTPMethod != http.MethodGet {
	// 			return JSONMethodNotAllowed(http.MethodGet)
	// 		}

	// 		return h.Me(ctx, req)
	// 	default:
	// 		return JSONNotFound()
	// 	}
	// }
}

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

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	})

	r.POST("/users/register", h.Register)
	r.POST("/users/login", h.Login)
	r.GET("/users/me", h.Me)
	r.PATCH("/users", h.UpdateProfile)
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

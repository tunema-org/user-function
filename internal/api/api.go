package api

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
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

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))

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

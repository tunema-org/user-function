package api

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/tunema-org/user-function/internal/backend"
)

type handler struct {
	backend *backend.Backend
}

func NewHandler(ctx context.Context, backend *backend.Backend) func(events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	h := &handler{
		backend: backend,
	}

	return func(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
		switch req.Resource {
		case "/users/register":
			if req.HTTPMethod != http.MethodPost {
				return JSONMethodNotAllowed(http.MethodPost)
			}

			return h.Register(ctx, req)
		case "/users/login":
			if req.HTTPMethod != http.MethodPost {
				return JSONMethodNotAllowed(http.MethodPost)
			}

			return h.Login(ctx, req)
		case "/users/me":
			if req.HTTPMethod != http.MethodGet {
				return JSONMethodNotAllowed(http.MethodGet)
			}

			return h.Me(ctx, req)
		default:
			return JSONNotFound()
		}
	}
}

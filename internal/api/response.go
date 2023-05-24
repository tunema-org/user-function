package api

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type M map[string]string

func JSON(status int, body any) (*events.APIGatewayProxyResponse, error) {
	resp := events.APIGatewayProxyResponse{
		Headers:    map[string]string{"Content-Type": "application/json"},
		StatusCode: status,
	}

	jsonBody, _ := json.Marshal(body)
	resp.Body = string(jsonBody)
	return &resp, nil
}

func JSONMethodNotAllowed(allowed string) (*events.APIGatewayProxyResponse, error) {
	return &events.APIGatewayProxyResponse{
		Headers:    map[string]string{"Content-Type": "application/json"},
		StatusCode: http.StatusMethodNotAllowed,
		Body:       `{"message": "Method Not Allowed", "allowed_method": ` + allowed + `}`,
	}, nil
}

func JSONNotFound() (*events.APIGatewayProxyResponse, error) {
	return &events.APIGatewayProxyResponse{
		Headers:    map[string]string{"Content-Type": "application/json"},
		StatusCode: http.StatusNotFound,
		Body:       `{"message": "Not Found"}`,
	}, nil
}

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Response struct {
	Message string `json:"message"`
}

type User struct {
	Name  *string `json:"name"`
	Email *string `json:"email"`
}

// Use REST API not work => HTTP working
func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var user User
	log.Println("##request:", request)

	if err := json.Unmarshal([]byte(request.Body), &user); err != nil {
		log.Println("Error parsing request body:", err)
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Invalid request body",
		}, err
	}

	response := Response{
		Message: fmt.Sprintf("Name, %s, Email, %s", *user.Name, *user.Email),
	}

	body, err := json.Marshal(response)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Internal Server Error",
		}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(body),
	}, nil
}

func main() {
	lambda.Start(handleRequest)
}

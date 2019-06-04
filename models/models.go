package models

import "github.com/aws/aws-lambda-go/events"

// Response type for API Gateway
type Response events.APIGatewayProxyResponse

// Request type for API Gateway
type Request events.APIGatewayProxyRequest

// Quote type
type Quote struct {
	ID   string    `json:"id"`
	Text QuoteText `json:"text"`
}

// QuoteText string
type QuoteText struct {
	string `json:"text"`
}

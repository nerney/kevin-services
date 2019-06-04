package main

import (
	"context"
	"fmt"
	"kevin_services/models"
	"kevin_services/util"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(func(ctx context.Context, req models.Request) (models.Response, error) {
		method := util.NormalizeString(req.HTTPMethod)
		if notPermitted := util.MethodNotPermitted(method); notPermitted != nil {
			return *notPermitted, nil
		}
		return models.Response{
			StatusCode: http.StatusOK,
			Body:       fmt.Sprintf("Congrats! You made a %v request", method),
		}, nil
	})
}

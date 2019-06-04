package main

import (
	"context"
	"errors"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(func(ctx context.Context) error {
		return errors.New(("someday i'll do something"))
	})
}

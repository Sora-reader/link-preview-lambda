package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("Processing request data for request %s.\n", request.RequestContext.RequestID)
	body := EncodeImage(ConvertImage(request.QueryStringParameters["image"]))
	headers := map[string]string{"Content-Type": "image/png"}

	return events.APIGatewayProxyResponse{
		Body:            body,
		StatusCode:      200,
		Headers:         headers,
		IsBase64Encoded: true,
	}, nil
}

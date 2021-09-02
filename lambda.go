package main

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/disintegration/imaging"
	"image"
	"image/color"
	"net/http"
	"time"
)

const width = 1200
const height = 630
const minHeight = 250

func convertImage(url string) image.Image {
	fmt.Printf("====> Should convert url: %s.\n", url)
	asset, err := Asset("logo.png")
	if err != nil {
		return imaging.New(1, 1, color.Black)
	}
	background, _ := imaging.Decode(bytes.NewReader(asset))

	fmt.Printf("Started request.\n")
	tr := &http.Transport{}
	tr.TLSClientConfig = &tls.Config{
		NextProtos: []string{"h1"},
	}

	client := &http.Client{
		Timeout:   5 * time.Second,
		Transport: tr,
	}
	imageResp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Println(imageResp)
	//imageResp, _ := http.Get(url)
	//defer imageResp.Body.Close()

	fmt.Printf("Finished request.\n")
	image, _ := imaging.Decode(imageResp.Body)
	image_h := image.Bounds().Max.Y

	if image_h < minHeight {
		image = imaging.Resize(image, 0, minHeight, imaging.Lanczos)
		image_h = image.Bounds().Max.Y
	}

	output := imaging.Resize(background, 0, image_h, imaging.Lanczos)
	output = imaging.PasteCenter(output, image)
	fmt.Printf("====> Returning request.\n")
	return image
}

func encodeImage(image image.Image) string {
	var buf bytes.Buffer
	fmt.Printf("====> Start encoding.\n")
	b64encoder := base64.NewEncoder(base64.StdEncoding, &buf)
	imaging.Encode(b64encoder, image, imaging.JPEG)
	b64encoder.Close()
	encoded := buf.String()
	fmt.Printf("====> Ended encoding.\n")
	return encoded
}

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("Processing request data for request %s.\n", request.RequestContext.RequestID)
	body := encodeImage(convertImage(request.QueryStringParameters["image"]))
	headers := map[string]string{"content-type": "image/png"}

	return events.APIGatewayProxyResponse{Body: body, StatusCode: 200, Headers: headers}, nil
}

func main() {
	lambda.Start(HandleRequest)
}

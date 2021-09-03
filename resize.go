package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	"net/http"

	"github.com/disintegration/imaging"
)

const width = 1200
const height = 630
const minHeight = 250

func ConvertImage(url string) image.Image {
	fmt.Printf("====> Should convert url: %s.\n", url)
	asset, err := Asset("assets/logo.png")
	if err != nil {
		fmt.Println("Can't load asset, falling back to black square")
		return imaging.New(width, height, color.Black)
	}
	background, _ := imaging.Decode(bytes.NewReader(asset))

	fmt.Printf("Started request.\n")
	imageResp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer imageResp.Body.Close()

	fmt.Printf("Finished request.\n")
	image, _ := imaging.Decode(imageResp.Body)
	image_h := image.Bounds().Max.Y

	if image_h < minHeight {
		image = imaging.Resize(image, 0, minHeight, imaging.Lanczos)
		image_h = image.Bounds().Max.Y
	}

	output := imaging.Resize(background, 0, image_h, imaging.Lanczos)
	output = imaging.PasteCenter(output, image)
	fmt.Printf("====> Returning image %s.\n", output.Bounds().Max)
	return output
}

func EncodeImage(image image.Image) string {
	var buf bytes.Buffer
	fmt.Printf("====> Start encoding.\n")
	b64encoder := base64.NewEncoder(base64.StdEncoding, &buf)
	imaging.Encode(b64encoder, image, imaging.JPEG)
	b64encoder.Close()
	encoded := buf.String()
	fmt.Printf("====> Ended encoding.\n")
	return encoded
}

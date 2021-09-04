package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"net/http"

	"github.com/disintegration/imaging"
)

const width = 1200
const height = 630
const minHeight = 250

func ConvertImage(url string) image.Image {
	fmt.Printf("====> Should convert url: %s.\n", url)

	fmt.Printf("Started request.\n")
	imageResp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer imageResp.Body.Close()

	fmt.Printf("Finished request.\n")
	image, _ := imaging.Decode(imageResp.Body)
	imageH, imageW := image.Bounds().Max.Y, image.Bounds().Max.Y

	if imageH < minHeight {
		image = imaging.Resize(image, 0, minHeight, imaging.Lanczos)
		imageH = image.Bounds().Max.Y
	} else if imageH > height || imageW > width {
		image = imaging.Fit(image, width, height, imaging.Lanczos)
		imageH = image.Bounds().Max.Y
	}

	bgWidth := int(float64(width) / (float64(height) / float64(imageH)))
	background := imaging.Resize(image, bgWidth, imageH, imaging.Lanczos)
	background = imaging.Blur(background, 20)

	output := imaging.Resize(background, 0, imageH, imaging.Lanczos)
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

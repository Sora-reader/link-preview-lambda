package main

import (
  "net/http"
  "github.com/disintegration/imaging"
  "bytes"
)

const width = 1200
const height = 630
const minHeight = 250

const mangaUrl = "https://static.readmanga.live/uploads/pics/01/41/131_p.jpg"

func main() {
  asset, _ := Asset("logo.png")
  background, _ := imaging.Decode(bytes.NewReader(asset))

  imageResp, _ := http.Get(mangaUrl)
  defer imageResp.Body.Close()

  image, _ := imaging.Decode(imageResp.Body)
  image_h := image.Bounds().Max.Y

  if image_h < minHeight {
    image = imaging.Resize(image, 0, minHeight, imaging.Lanczos)
    image_h = image.Bounds().Max.Y
  }

  output := imaging.Resize(background, 0, image_h, imaging.Lanczos)
  output = imaging.PasteCenter(output, image)
  imaging.Save(output, "test.png")
}

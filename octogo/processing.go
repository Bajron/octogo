package octogo

import (
	"image"
	//_ "golang.org/x/image/bmp"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	"log"
	"os"
)

func Copy(src image.Image) image.Image {
	bounds := src.Bounds()
	ret := image.NewRGBA(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			ret.Set(x, y, src.At(x, y))
		}
	}
	return ret
}

type ProcessFunction func(image.Image) image.Image

func Process(src, dst string, f ProcessFunction) {
	reader, err := os.Open(src)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	img, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	processed := f(img)
	if processed == nil {
		return
	}
	writer, err := os.Create(dst)
	if err != nil {
		log.Fatal(err)
	}
	defer writer.Close()

	err = png.Encode(writer, processed)
	if err != nil {
		log.Fatal(err)
	}
	return
}

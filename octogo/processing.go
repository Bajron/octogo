package octogo

import (
	"image"
	//_ "golang.org/x/image/bmp"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type ProcessFunction func(image.Image) image.Image
type Encoder interface {
	Encode(writer io.Writer, img image.Image) error
}

var encoders map[string]Encoder

func init() {
	encoders = make(map[string]Encoder)
	encoders[".png"] = PngEncoder{}
}

type PngEncoder struct{}

func (PngEncoder) Encode(writer io.Writer, img image.Image) error {
	return png.Encode(writer, img)
}

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
		log.Println("Processing failed :C")
		return
	}
	writer, err := os.Create(dst)
	if err != nil {
		log.Fatal(err)
	}
	defer writer.Close()

	ext := strings.ToLower(filepath.Ext(dst))

	log.Printf("Encoding [%s] to %s", ext, dst)
	err = encoders[ext].Encode(writer, processed)
	if err != nil {
		log.Fatal(err)
	}
	return
}

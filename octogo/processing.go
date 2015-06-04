package octogo

import (
	"image"
	//_ "golang.org/x/image/bmp"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type ProcessingFunction func(image.Image) image.Image
type Encoder interface {
	Encode(writer io.Writer, img image.Image) error
}

var encoders map[string]Encoder
var processors map[string]ProcessingFunction

func init() {
	encoders = make(map[string]Encoder)
	encoders[".png"] = PngEncoder{}

	processors = make(map[string]ProcessingFunction)
	processors["copy"] = Copy
	processors["gray"] = Grayscale
	processors["mean_horizontal"] = MeanHorizontal
}

type PngEncoder struct{}

func (PngEncoder) Encode(writer io.Writer, img image.Image) error {
	return png.Encode(writer, img)
}

func GetProcessingFunction(name string) ProcessingFunction {
	f, ok := processors[name]
	if !ok {
		f = Copy
	}
	return f
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

func Grayscale(src image.Image) image.Image {
	bounds := src.Bounds()
	ret := image.NewGray16(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := src.At(x, y).RGBA()
			mean := (r + g + b) / 3
			ret.Set(x, y, color.Gray16{uint16(mean)})
		}
	}
	return ret
}

func Process(src, dst string, f ProcessingFunction) {
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
	enc, ok := encoders[ext]
	if !ok {
		log.Printf("Encoder for %s not found", dst)
		return
	}
	err = enc.Encode(writer, processed)
	if err != nil {
		log.Fatal(err)
	}
	return
}

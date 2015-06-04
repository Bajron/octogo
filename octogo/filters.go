package octogo

import (
	"image"
	"image/color"
)

func meanByComponent(colors ...color.Color) color.Color {
	l := uint32(len(colors))
	if l == 0 {
		return image.Black
	}
	r, g, b, a := colors[0].RGBA()
	for i := uint32(1); i < l; i++ {
		rr, gg, bb, aa := colors[i].RGBA()
		r, g, b, a = r+rr, g+gg, b+bb, a+aa
	}
	r, g, b, a = r/l, g/l, b/l, a/l
	return color.RGBA64{
		uint16(r), uint16(g), uint16(b), uint16(a)}
}

func MeanHorizontal(src image.Image) image.Image {
	bounds := src.Bounds()
	minX := bounds.Min.X
	maxX := bounds.Max.X - 1

	if minX == maxX {
		return Copy(src)
	}

	ret := image.NewRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		ret.Set(minX, y, meanByComponent(
			src.At(minX, y), src.At(minX+1, y)))
		for x := minX + 1; x <= maxX-1; x++ {
			ret.Set(x, y, meanByComponent(
				src.At(x-1, y), src.At(x, y), src.At(x+1, y)))
		}
		ret.Set(maxX, y, meanByComponent(
			src.At(maxX, y), src.At(maxX-1, y)))
	}
	return ret
}

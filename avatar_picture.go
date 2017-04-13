package avatar

import (
	"bytes"
	"image"
	"image/color"
)

type PictureOptions struct {
	BgColor color.Color
	Size    int
}

type Picture struct {
	source        []byte
	options       *PictureOptions
	originalImage image.Image
}

func (p Picture) originalImg() image.Image {
	return p.originalImage
}

func (p Picture) Source() []byte {
	return p.source
}

// Generates the square avatar
// It returns the avatar image in []byte or an error something went wrong
func (p Picture) Square() ([]byte, error) {
	return square(p, p.options)
}

func (p Picture) Circle() ([]byte, error) {
	return circle(p, p.options)
}

func (p Picture) loadOriginalImage() (image.Image, error) {
	srcReader := bytes.NewReader(p.source)

	// Convert source file to image.Image
	originalImage, _, err := image.Decode(srcReader)
	if err != nil {
		return nil, err
	}

	return originalImage, nil
}

func (o PictureOptions) bgColor() color.Color {
	return bgColor(o.BgColor)
}

func (o PictureOptions) size() int {
	return size(o.Size)
}

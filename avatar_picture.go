package avatar

import (
	"bytes"
	"image"
)

type Picture struct {
	Type
}

func (p *Picture) loadOriginalImage() error {
	srcReader := bytes.NewReader(p.source)

	// Convert source file to image.Image
	originalImage, _, err := image.Decode(srcReader)
	if err != nil {
		return err
	}

	p.originalImg = originalImage

	return nil
}

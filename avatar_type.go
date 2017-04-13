package avatar

import (
	"bytes"
	"errors"
	"image"
	"image/color"
	"image/draw"
	"image/png"

	"fmt"
	"github.com/nfnt/resize"
)

type Type struct {
	Avatar
	source      []byte
	size        int
	originalImg image.Image
	squareImg   image.Image
	circleImg   image.Image
}

func (t Type) Source() []byte {
	return t.source
}

func (t *Type) Square() ([]byte, error) {
	var imgOutput []byte
	buf := bytes.NewBuffer(imgOutput)

	m := resize.Resize(uint(t.size), uint(t.size), t.squareImg, resize.Lanczos2)
	err := png.Encode(buf, m)
	if err != nil {
		return imgOutput, err
	}

	return buf.Bytes(), nil
}

func (t *Type) Circle() ([]byte, error) {
	var imgOutput []byte
	buf := bytes.NewBuffer(imgOutput)

	m := resize.Resize(uint(t.size), uint(t.size), t.circleImg, resize.Lanczos2)
	err := png.Encode(buf, m)
	if err != nil {
		return imgOutput, err
	}

	return buf.Bytes(), nil
}

func (t *Type) loadSquareImage() error {
	if t.originalImg == nil {
		return errors.New("Cannot create square image without original image")
	}
	oImage := t.originalImg

	minSize := getMinSide(oImage.Bounds())

	dstRect := image.Rect(0, 0, minSize, minSize)
	dstImg := image.NewRGBA(dstRect)
	draw.Draw(dstImg, dstImg.Bounds(), &image.Uniform{color.Transparent}, image.ZP, draw.Src)

	squareRect := getMaskCenterBounds(oImage.Bounds())
	draw.DrawMask(dstImg, dstImg.Bounds(), oImage, image.Point{X: squareRect.Min.X, Y: squareRect.Min.Y}, dstRect, image.ZP, draw.Over)

	t.squareImg = dstImg

	return nil
}

func (p *Type) loadCircleImage() error {
	if p.squareImg == nil {
		return p.loadSquareImage()
	}
	sImage := p.squareImg

	size := getMinSide(sImage.Bounds())

	dstRect := image.Rect(0, 0, size, size)
	dstImg := image.NewRGBA(dstRect)
	draw.Draw(dstImg, dstRect, &image.Uniform{color.Transparent}, image.ZP, draw.Src)

	circleMask := &circle{
		p: image.Pt(int(float64(size)/2), int(float64(size)/2)),
		r: size / 2,
	}

	b := circleMask.Bounds()
	fmt.Println(b)

	at := circleMask.At(b.Max.X/2, b.Max.Y)
	fmt.Println(at)

	draw.DrawMask(dstImg, sImage.Bounds(), sImage, image.ZP, circleMask, image.ZP, draw.Over)

	p.circleImg = dstImg

	return nil
}

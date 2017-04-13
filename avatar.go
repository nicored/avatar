package avatar

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	_ "image/jpeg"
	"image/png"
	"math"

	"github.com/nfnt/resize"
)

type Avatar interface {
	Source() []byte          // get source content for avatar generation
	Square() ([]byte, error) // generates the square avatar
	Circle() ([]byte, error) // generates the round avatar
	loadOriginalImage() (image.Image, error)
	originalImg() image.Image
}

type AvatarOptions interface {
	bgColor() color.Color
	size() int
}

type dimensions struct {
	w int
	h int
}

var (
	defaultSize      = 300
	defaultBgColor   = color.Transparent
	defaultTxtColor  = color.White
	defaultNInitials = 2
)

func NewAvatarFromInitials(text []byte, options *InitialsOptions) (*Initials, error) {
	if options == nil {
		options = &InitialsOptions{}
	}

	newAvatar := Initials{
		source:  text,
		options: options,
	}

	original, err := newAvatar.loadOriginalImage()
	if err != nil {
		return nil, err
	}

	newAvatar.originalImage = original
	return &newAvatar, nil
}

func NewAvatarFromPic(pic []byte, options *PictureOptions) (*Picture, error) {
	if options == nil {
		options = &PictureOptions{}
	}

	newAvatar := Picture{
		source:  pic,
		options: options,
	}

	original, err := newAvatar.loadOriginalImage()
	if err != nil {
		return nil, err
	}

	newAvatar.originalImage = original
	return &newAvatar, nil
}

func bgColor(bgColor color.Color) color.Color {
	if bgColor == nil {
		return defaultBgColor
	}

	return bgColor
}

func size(size int) int {
	if size > 0 {
		return size
	}

	return defaultSize
}

func square(a Avatar, options AvatarOptions) ([]byte, error) {
	squareImg, err := generateSquareImage(a)
	if err != nil {
		return []byte{}, err
	}

	size := defaultSize
	if options != nil {
		size = options.size()
	}

	return encode(squareImg, size)
}

func circle(a Avatar, options AvatarOptions) ([]byte, error) {
	circleImg, err := generateCircleImage(a, options.bgColor())
	if err != nil {
		return []byte{}, err
	}

	size := defaultSize
	if options != nil {
		size = options.size()
	}

	return encode(circleImg, size)
}

func encode(img image.Image, size int) ([]byte, error) {
	var imgOutput []byte
	buf := bytes.NewBuffer(imgOutput)

	resized := resize.Resize(uint(size), uint(size), img, resize.Lanczos2)

	err := png.Encode(buf, resized)
	if err != nil {
		return imgOutput, err
	}

	return buf.Bytes(), nil
}

// Computes the original image to generate the square image
// from the center of the original image
func generateSquareImage(a Avatar) (image.Image, error) {
	if a.originalImg() == nil {
		return nil, errors.New("Cannot create square image without original image")
	}
	oImage := a.originalImg()

	minSize := getMinSide(oImage.Bounds())

	dstRect := image.Rect(0, 0, minSize, minSize)
	dstImg := image.NewRGBA(dstRect)
	draw.Draw(dstImg, dstImg.Bounds(), &image.Uniform{color.Transparent}, image.ZP, draw.Src)

	squareRect := getMaskCenterBounds(oImage.Bounds())
	draw.DrawMask(dstImg, dstImg.Bounds(), oImage, image.Point{X: squareRect.Min.X, Y: squareRect.Min.Y}, dstRect, image.ZP, draw.Over)

	return dstImg, nil
}

// Computes the square image to generate the round image
// from the center of the original image
func generateCircleImage(a Avatar, bgColor color.Color) (image.Image, error) {
	sqImg, err := generateSquareImage(a)
	if err != nil {
		return nil, err
	}

	size := getMinSide(sqImg.Bounds())

	dstRect := image.Rect(0, 0, size, size)
	dstImg := image.NewRGBA(dstRect)
	draw.Draw(dstImg, dstRect, &image.Uniform{color.Transparent}, image.ZP, draw.Src)

	circleMask := &Circle{
		p: image.Pt(int(float64(size)/2), int(float64(size)/2)),
		r: size / 2,
	}

	b := circleMask.Bounds()
	fmt.Println(b)

	at := circleMask.At(b.Max.X/2, b.Max.Y)
	fmt.Println(at)

	draw.DrawMask(dstImg, sqImg.Bounds(), sqImg, image.ZP, circleMask, image.ZP, draw.Over)

	return dstImg, nil
}

func getMaskCenterBounds(rect image.Rectangle) image.Rectangle {
	srcBounds := rect.Bounds()

	center := getCenter(srcBounds)
	minSize := getMinSide(srcBounds)

	minPt := image.Point{}
	minPt.X = center.X - (int(minSize) / 2)
	minPt.Y = center.Y - (int(minSize) / 2)

	maxPt := image.Point{}
	maxPt.X = center.X + (int(minSize) / 2)
	maxPt.Y = center.Y + (int(minSize) / 2)

	return image.Rectangle{
		Min: minPt,
		Max: maxPt,
	}
}

func getCenter(rect image.Rectangle) image.Point {
	srcBounds := rect.Bounds()

	return image.Point{
		X: int(float32(srcBounds.Max.X) / float32(2)),
		Y: int(float32(srcBounds.Max.Y) / float32(2)),
	}
}

func getMinSide(rect image.Rectangle) int {
	srcBounds := rect.Bounds()

	height := srcBounds.Max.Y - srcBounds.Min.Y
	width := srcBounds.Max.X - srcBounds.Min.X

	return int(math.Min(float64(width), float64(height)))
}

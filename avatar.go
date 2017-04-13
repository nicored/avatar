package avatar

import (
	"image"
	"image/color"
	_ "image/jpeg"
	"math"
)

type Avatar interface {
	Source() []byte
	Square() []byte
	Circle() []byte
}

type dimensions struct {
	w int
	h int
}

func NewAvatarFromInitials(initials []byte, size int, fontPathTTF string, bgColor color.RGBA, textColor color.RGBA) (*Initials, error) {
	newAvatar := Initials{
		fontPath: fontPathTTF,
		bgRGBA:   bgColor,
		txtRGBA:  textColor,
	}
	newAvatar.source = initials
	newAvatar.size = size

	if err := newAvatar.loadOriginalImage(); err != nil {
		return nil, err
	}

	if err := newAvatar.loadCircleImage(); err != nil {
		return nil, err
	}

	return &newAvatar, nil
}

func NewAvatarFromPic(pic []byte, size int) (*Picture, error) {
	newAvatar := Picture{}
	newAvatar.size = size
	newAvatar.source = pic

	if err := newAvatar.loadOriginalImage(); err != nil {
		return nil, err
	}

	if err := newAvatar.loadSquareImage(); err != nil {
		return nil, err
	}

	if err := newAvatar.loadCircleImage(); err != nil {
		return nil, err
	}

	return &newAvatar, nil
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

//
//func DrawAvatar(srcFilePath string, outputPath string) error {
//	// We read the source image
//	srcFile, err := os.Open(srcFilePath)
//	if err != nil {
//		return err
//	}
//	defer srcFile.Close()
//
//	// Convert source file to image.Image
//	src, _, err := image.Decode(srcFile)
//	if err != nil {
//		return err
//	}
//
//	// Create a new empty file for output image
//	// at destination path
//	out, err := os.Create(outputPath)
//	if err != nil {
//		return err
//	}
//	defer out.Close()
//
//	// Get the center of the source image
//	srcBounds := src.Bounds()
//	center := image.Point{
//		X: int(float32(srcBounds.Max.X) / float32(2)),
//		Y: int(float32(srcBounds.Max.Y) / float32(2)),
//	}
//
//	// Create the mask to be applied to the source image
//	circleMask := &circle{
//		p: center,
//		r: int(math.Min(float64(center.X), float64(center.Y))),
//	}
//
//	// Create the new image, completely transparent
//	imgRect := image.Rect(0, 0, circleMask.r*2, circleMask.r*2)
//	dst := image.NewRGBA(imgRect)
//	draw.Draw(dst, dst.Bounds(), &image.Uniform{color.Transparent}, image.ZP, draw.Src)
//
//	// Apply the mask to the source image, and draw it in the new transparent image
//	draw.DrawMask(dst, dst.Bounds(), src, image.ZP, circleMask, image.Point{X: center.X - circleMask.r, Y: center.Y - circleMask.r}, draw.Over)
//
//	// We resize it to be smaller
//	m := resize.Resize(200, 200, dst, resize.Lanczos2)
//
//	// We create the new PNG, which is stored in file system
//	err = png.Encode(out, m)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}

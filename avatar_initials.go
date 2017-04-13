package avatar

import (
	"image"
	"image/color"
	"image/draw"
	"io/ioutil"
	"regexp"
	"unicode"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

type Initials struct {
	Type
	fontPath string
	bgRGBA   color.RGBA
	txtRGBA  color.RGBA
}

func (i *Initials) loadOriginalImage() error {
	size := i.size * 3

	// Draw background img
	imgRect := image.Rect(0, 0, size, size)
	dst := image.NewRGBA(imgRect)
	draw.Draw(
		dst,
		dst.Bounds(),
		image.NewUniform(i.bgRGBA),
		image.ZP,
		draw.Src)

	// Prepare Font
	fontBytes, err := ioutil.ReadFile(i.fontPath)
	if err != nil {
		return err
	}

	ftFont, err := freetype.ParseFont(fontBytes)
	if err != nil {
		return err
	}

	fontFace := truetype.NewFace(ftFont, &truetype.Options{
		Size: getFontSizeThatFits(i.source, float64(size), ftFont),
	})

	fd := font.Drawer{
		Dst:  dst,
		Src:  image.NewUniform(i.txtRGBA),
		Face: fontFace,
	}

	// Figure out baseline and adv for string in img
	txtWidth := fd.MeasureBytes(i.source)
	txtWidthInt := int(txtWidth >> 6)

	bounds, _ := fd.BoundBytes(i.source)
	txtHeight := bounds.Max.Y - bounds.Min.Y
	txtHeightInt := int(txtHeight >> 6)

	advLine := (size / 2) - (txtWidthInt / 2)
	baseline := (size / 2) + (txtHeightInt / 2)

	fd.Dot = fixed.Point26_6{X: fixed.Int26_6(advLine << 6), Y: fixed.Int26_6(baseline << 6)}

	fd.DrawBytes(i.source)

	i.squareImg = dst

	return nil
}

func getFontSizeThatFits(text []byte, imgWidth float64, ftFont *truetype.Font) float64 {
	fontSize := float64(100)

	drawer := font.Drawer{
		Face: truetype.NewFace(ftFont, &truetype.Options{
			Size: fontSize,
		}),
	}

	tw := float64(drawer.MeasureBytes(text) >> 6)

	ratio := fontSize / tw

	return ratio * (imgWidth - (40./100)*imgWidth)
}

func getInitials(text []byte, nChars int) []byte {
	if len(text) == 0 {
		return []byte("")
	}

	var initials = []byte{}
	var previous = byte(' ')

	regEmail := regexp.MustCompile("^(((([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+(\\.([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+)*)|((\\x22)((((\\x20|\\x09)*(\\x0d\\x0a))?(\\x20|\\x09)+)?(([\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x7f]|\\x21|[\\x23-\\x5b]|[\\x5d-\\x7e]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(\\([\\x01-\\x09\\x0b\\x0c\\x0d-\\x7f]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}]))))*(((\\x20|\\x09)*(\\x0d\\x0a))?(\\x20|\\x09)+)?(\\x22)))@((([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])([a-zA-Z]|\\d|-|\\.|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.)+(([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])([a-zA-Z]|\\d|-|\\.|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.?$")
	skipFromAt := regEmail.Match(text)

	for _, ch := range text[:] {
		if skipFromAt == true && rune(ch) == '@' {
			break
		}

		if isSymbol(rune(ch)) {
			previous = ch
			continue
		}

		if ((unicode.IsUpper(rune(ch)) && unicode.IsLower(rune(previous))) || (unicode.IsLower(rune(ch)) && len(initials) == 0)) && !isSymbol(rune(ch)) {
			initials = append(initials, ch)
			previous = ch
		}
	}

	for i := len(initials); i < nChars && len(text) > i; i++ {
		if isSymbol(rune(text[i])) {
			continue
		}

		initials = append(initials, text[i])
	}

	return initials
}

func isSymbol(ch rune) bool {
	return unicode.IsSymbol(ch) || unicode.IsSpace(ch) || unicode.IsPunct(ch)
}

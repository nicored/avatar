package avatar

import (
	"image/color"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitials_SquareCircle(t *testing.T) {
	size := 100

	newAvatar, err := NewAvatarFromInitials([]byte("John Smith"), &InitialsOptions{
		Size:      size,
		NInitials: 2,
		FontPath:  getTestResource("test_data", "Arial.ttf"),
		TextColor: color.White,
		BgColor:   color.RGBA{0, 0, 255, 255},
	})
	assert.NoError(t, err)

	round, err := newAvatar.Circle()
	assert.NoError(t, err)

	roundOutputPath := getTestResource("output", "round_john_smith_initials.png")
	roundFile, err := os.Create(roundOutputPath)
	assert.NoError(t, err)
	roundFile.Write(round)

	square, err := newAvatar.Square()
	assert.NoError(t, err)

	squareOutputPath := getTestResource("output", "square_john_smith_initials.png")
	squareFile, err := os.Create(squareOutputPath)
	assert.NoError(t, err)
	squareFile.Write(square)
}

package avatar

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPicture_Circle(t *testing.T) {
	size := 200

	mascotPath := getTestResource("test_data", "super_mascot.jpg")
	fileBytes, err := ioutil.ReadFile(mascotPath)
	assert.NoError(t, err)

	newAvatar, err := NewAvatarFromPic(fileBytes, &PictureOptions{
		Size: size,
	})
	assert.NoError(t, err)

	round, err := newAvatar.Circle()

	roundMascotPath := getTestResource("output", "round_super_mascot.png")
	roundFile, err := os.Create(roundMascotPath)
	assert.NoError(t, err)
	roundFile.Write(round)
}

func TestPicture_Square(t *testing.T) {
	size := 200

	mascotPath := getTestResource("test_data", "super_mascot.jpg")
	fileBytes, err := ioutil.ReadFile(mascotPath)
	assert.NoError(t, err)

	newAvatar, _ := NewAvatarFromPic(fileBytes, &PictureOptions{
		Size: size,
	})

	square, err := newAvatar.Square()
	squareMascotPath := getTestResource("output", "square_super_mascot.png")
	squareFile, err := os.Create(squareMascotPath)
	assert.NoError(t, err)
	squareFile.Write(square)
}

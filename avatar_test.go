package avatar

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"path/filepath"
	"runtime"
)

func TestNewAvatarFromPic(t *testing.T) {
	size := 300

	mascotPath := getTestResource("test_data", "super_mascot.jpg")
	fileBytes, err := ioutil.ReadFile(mascotPath)
	assert.NoError(t, err)

	newAvatar, err := NewAvatarFromPic(fileBytes, &PictureOptions{
		Size: size,
	})
	assert.NoError(t, err)

	assert.Equal(t, fileBytes, newAvatar.Source())
}

func getTestResource(dir, resource string) string {
	_, currTestPath, _, _ := runtime.Caller(0)
	testDataPath := filepath.Join(filepath.Dir(currTestPath), dir, resource)

	return testDataPath
}

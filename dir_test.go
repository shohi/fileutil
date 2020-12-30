package fileutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllSubfolders(t *testing.T) {
	assert := assert.New(t)

	folders, err := AllSubfolders(".")
	assert.Nil(err)
	assert.Equal(1, len(folders))
	assert.Equal("testdata", folders[0])
}

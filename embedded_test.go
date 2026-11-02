package tabler_test

import (
	"io/fs"
	"testing"

	"github.com/go-app-blazar/tabler-font"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEmbedded(t *testing.T) {
	files, err := fs.Glob(tabler.FS, "*.*")
	require.NoError(t, err)
	assert.NotEmpty(t, files)
	assert.Contains(t, files, "ta-filled.css")
	assert.Contains(t, files, "ta-filled.woff2")
	assert.Contains(t, files, "ta-outline.css")
	assert.Contains(t, files, "ta-outline.woff2")
}

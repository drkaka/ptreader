package ptreader

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadMeta(t *testing.T) {
	file := "db.PT"
	r, cleanup, err := newReader(file)
	if !assert.NoError(t, err, "should have no error.") {
		t.Fatal(err)
	}
	defer cleanup()

	version, err := r.readVersion()
	assert.NoError(t, err, "read version wrong")
	assert.Equal(t, 10, version, "version wrong")
	t.Log("version:", version)

	magic, err := r.readMagic()
	assert.NoError(t, err, "read magic wrong")
	assert.Equal(t, 1347700294, magic, "magic wrong")
	t.Log("magic:", magic)

	tags, colors, err := r.readTags()
	assert.NoError(t, err, "read tags wrong")
	assert.Equal(t, len(tags), len(colors), "tags or colors length wrong")
	for i := 0; i < len(tags); i++ {
		t.Logf("tag [%d]: %s", i, tags[i])
	}
	t.Log("colors:", colors)

	minfilter, err := r.readMinfilter()
	assert.NoError(t, err, "read minfilter wrong")
	t.Log("minfilter:", minfilter)

	foldLevel, err := r.readFoldLevel()
	assert.NoError(t, err, "read foldLevel wrong")
	t.Log("foldLevel:", foldLevel)

	prefs, err := r.readPrefs()
	assert.NoError(t, err, "read prefs wrong")
	t.Log("prefs:", prefs)
}

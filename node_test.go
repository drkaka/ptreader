package ptreader

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadNode(t *testing.T) {
	file := "db.PT"
	r, cleanup, err := newReader(file)
	if !assert.NoError(t, err, "should have no error.") {
		t.Fatal(err)
	}
	defer cleanup()

	assert.NoError(t, r.readMeta(), "error to read meta.")

	n, err := r.readNode()
	assert.NoError(t, err, "error to read node.")
	t.Log(n)
}

package ptreader

import (
	"compress/gzip"
	"os"
)

// Reader to read bytes
type Reader struct {
	r *gzip.Reader
}

// Read a db file
func Read(filePath string) (*Node, error) {
	r, cleanup, err := newReader(filePath)
	if err != nil {
		return nil, err
	}
	defer cleanup()

	// read meta data.
	if err := r.readMeta(); err != nil {
		return nil, err
	}

	n, err := r.readNode()
	if err != nil {
		return nil, err
	}

	return &n, nil
}

// get a reader and its cleanup func
func newReader(filePath string) (*Reader, func(), error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, nil, err
	}

	fz, err := gzip.NewReader(f)
	if err != nil {
		f.Close()
		return nil, nil, err
	}

	cleanup := func() {
		f.Close()
		fz.Close()
	}
	return &Reader{r: fz}, cleanup, nil
}

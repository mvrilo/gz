package gz

import (
	"bytes"
	"compress/gzip"
	"os"
)

// Copied from compress/gzip package for easy access
const (
	NoCompression      = gzip.NoCompression
	BestSpeed          = gzip.BestSpeed
	BestCompression    = gzip.BestCompression
	DefaultCompression = gzip.DefaultCompression
)

// Read opens the file and return the uncompressed data
func Read(path string) ([]byte, int64, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, 0, err
	}

	g, err := gzip.NewReader(f)
	if err != nil {
		return nil, 0, err
	}
	defer g.Close()

	b := new(bytes.Buffer)
	i, err := b.ReadFrom(g)
	if err != nil {
		return nil, 0, err
	}
	return b.Bytes(), i, nil
}

// Write writes compressed data to a file with given compression ratio. Default compression is 0 (no compression).
func Write(path string, data []byte, compression int) (int, error) {
	f, err := os.Create(path)
	if err != nil {
		return 0, err
	}

	w, err := gzip.NewWriterLevel(f, compression)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	i, err := w.Write(data)
	if err != nil {
		return 0, err
	}
	defer w.Close()

	return i, nil
}

// WriteBest writes compressed data to a file using BestCompression ratio.
func WriteBest(path string, data []byte) (int, error) {
	return Write(path, data, BestCompression)
}

// WriteFast writes compressed data to a file using BestSpeed ratio.
func WriteFast(path string, data []byte) (int, error) {
	return Write(path, data, BestSpeed)
}

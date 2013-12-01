package gz

import (
	"compress/gzip"
	"io"
	"os"
)

func reader(path string) (*os.File, io.ReadCloser, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}

	g, err := gzip.NewReader(f)
	return f, g, err
}

func Read(path string) ([]byte, int, error) {
	f, g, err := reader(path)
	if err != nil {
		return nil, 0, err
	}
	defer g.Close()
	defer f.Close()

	s, err := f.Stat()
	if err != nil {
		return nil, 0, err
	}

	body := make([]byte, s.Size())
	i, err := g.Read(body)
	if err != nil {
		return nil, 0, err
	}
	return body, i, nil
}

func writer(path string, data []byte) (*os.File, io.WriteCloser, error) {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return nil, nil, err
	}

	w, err := gzip.NewWriterLevel(f, gzip.BestCompression)
	return f, w, err
}

func Write(path string, data []byte) (int, error) {
	f, w, err := writer(path, data)
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

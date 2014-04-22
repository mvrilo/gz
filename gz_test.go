package gz

import (
	"bytes"
	"os"
	"testing"
)

var (
	file = "test.txt.gz"
	data = []byte("this is going into the file")
)

func TestWrite(t *testing.T) {
	i, err := Write(file, data, BestSpeed)
	if err != nil {
		t.Error("Write returned an error: ", err.Error())
	}
	if i == 0 {
		t.Error("Bytes returned should be more than zero")
	}

	i, err = Write(".", data, BestSpeed)
	if err == nil {
		t.Error("Writing to a directory should return an error")
	}
	if i > 0 {
		t.Error("Writing to a directory should return zero bytes")
	}
}

func TestWriteBest(t *testing.T) {
	_, err := WriteBest(file, data)
	if err != nil {
		t.Error("WriteBest returned an error: ", err.Error())
	}
}

func TestWriteFast(t *testing.T) {
	_, err := WriteFast(file, data)
	if err != nil {
		t.Error("WriteFast returned an error: ", err.Error())
	}
}

func TestRead(t *testing.T) {
	b, i, err := Read(file)
	if err != nil {
		t.Error("Read return an error: ", err.Error())
	}
	if i == 0 {
		t.Error("Bytes returned should be more than zero")
	}
	defer os.Remove(file)

	if !bytes.Equal(b, data) {
		t.Error("Data should be the same")
	}

	b, i, err = Read("missing.txt.gz")
	if err == nil {
		t.Error("Should return an error")
	}
	if b != nil {
		t.Error("Data should be nil")
	}
	if i > 0 {
		t.Error("Bytes length should be zero")
	}
}

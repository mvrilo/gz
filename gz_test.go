package gz

import (
	"bytes"
	"io"
	"os"
	"testing"
)

const path = "test.txt.gz"

var data = []byte("this is going into the file")

func checkErr(err error, t *testing.T) {
	if err != nil {
		t.Fatal(err)
	}
}

func TestWrite(t *testing.T) {
	i, err := Write(path, []byte(data))
	checkErr(err, t)
	if i < 1 {
		t.Fatal("bytes length returned is less than one")
	}
	_, err = os.Open(path)
	checkErr(err, t)

	// if everything is okay run Write again to
	// make sure it's not appending
	Write(path, []byte(data))
}

func TestRead(t *testing.T) {
	b, i, err := Read(path)
	checkErr(err, t)
	if i < 1 {
		t.Fatal("bytes length returned is less than one")
	}
	if len(b) < 1 {
		t.Fatal("len(data) returned less than one")
	}

	buf := new(bytes.Buffer)
	f, r, err := reader(path)
	checkErr(err, t)
	defer func() {
		f.Close()
		os.Remove(path)
	}()

	io.Copy(buf, r)
	if len(buf.Bytes()) != len(data) {
		t.Fatal("data length comparison is wrong")
	}

	if buf.String() != string(data) {
		t.Fatal("data string comparison is wrong")
	}
}

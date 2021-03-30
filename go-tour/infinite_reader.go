package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (MyReader) Read(b []byte) (int, error) {
	return copy(b, "A"), nil
}

func main() {
	reader.Validate(MyReader{})
}

package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

// 65-90
// 97-122
func (rot13 rot13Reader) Read(b []byte) (int, error) {
	count, err := rot13.r.Read(b)
	if err != nil {
		return count, err
	}
	middle_of_big := byte(65 + 13)
	middle_of_small := byte(97 + 13)

	for i := 0; i < len(b); i++ {
		if b[i] >= 65 && b[i] <= 90 {
			if b[i] >= middle_of_big {
				b[i] -= 13
			} else {
				b[i] += 13
			}
		}

		if b[i] >= 97 && b[i] <= 122 {
			if b[i] >= middle_of_small {
				b[i] -= 13
			} else {
				b[i] += 13
			}
		}
	}

	return count, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(b []byte) (int, error) {
	pre_read:= make([]byte, len(b))
	num, err := rot.r.Read(pre_read)
	for indx := range b {
		b[indx] = perform_cypher_lowercase(pre_read[indx])
	}
	return num, err
} 

func perform_cypher_lowercase (b byte) byte{
	current_ASCII := int(b)
	if ((current_ASCII < 97 || current_ASCII > 122) && (current_ASCII < 65 || current_ASCII > 90)) {
		return b
	}
	
	if (current_ASCII >= 97 || current_ASCII <= 122) {
		new_ASCII := ((current_ASCII - 97 + 13) % 26) + 97
		return byte(new_ASCII)
	}
	
	//Case where we have uppercase letters	
	new_ASCII := ((current_ASCII - 65 + 13) % 26) + 65
	return byte(new_ASCII)
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

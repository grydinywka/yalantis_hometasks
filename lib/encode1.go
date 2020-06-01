package encode1

import (
	"bytes"
	"fmt"
)

var CODER = map[byte]int{'a': 1, 'e': 2, 'i': 3, 'o': 4, 'u': 5}
var DECODER = map[byte]string{'1': "a", '2': "e", '3': "i", '4': "o", '5': "u"}

func Encode(s *string) *string {
	fmt.Println("The encode func")
	var result string
	var buf bytes.Buffer

	buf.Write([]byte(""))

	for _, char := range *s {
		if val, ok := CODER[byte(char)]; ok {
			fmt.Fprintf(&buf, "%d", int(val))
		} else {
			fmt.Fprintf(&buf, "%s", string(char))
		}
	}

	result = buf.String()
	return &result
}

func Decode(s *string) *string {
	var result string
	var buf bytes.Buffer

	buf.Write([]byte(""))
	for _, char := range *s {
		if val, ok := DECODER[byte(char)]; ok {
			fmt.Fprintf(&buf, "%s", val)
		} else {
			fmt.Fprintf(&buf, "%s", string(char))
		}
	}
	result = buf.String()

	return &result
}

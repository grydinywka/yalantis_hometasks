package main

import (
	"fmt"

	L "./lib"
)

func main() {
	var phrase string = "eat0 ! ((())) is e, G`oo0d dog!) `ert `rtjo 55 44 33 22 11"

	pLatin := L.ToPigLatin(&phrase)
	fmt.Println("pLatin = " + *pLatin)

	encoded := L.Encode(&phrase)
	fmt.Printf("Encode = %s\n", *encoded)

	decoded := L.Decode(&phrase)
	fmt.Printf("Encode = %s\n", *decoded)
}

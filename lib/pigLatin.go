package symbols

import (
	"bytes"
	"fmt"
	"strings"
)

func isVowelChar(char byte) bool {
	for _, vowel := range [6]byte{'a', 'e', 'i', 'o', 'u', '`'} {
		if vowel == char {
			return true
		}
	}
	return false
}

func isLetter(char byte) bool {
	if ('A' <= char && char <= 'Z') || ('`' <= char && char <= 'z') {
		return true
	}
	return false
}

func ToPigLatin(phrase *string) *string {
	var result string
	var resBytes bytes.Buffer

	resBytes.Write([]byte(""))

	if len(*phrase) == 0 {
		fmt.Println("THe empty string,,,")
		return phrase
	}
	fmt.Println("phrase in English is : " + *phrase)

	for _, word := range strings.SplitN(strings.ToLower(*phrase), " ", -1) {
		length := len(word)
		j := 0                 // line indicates start vowel
		if isLetter(word[j]) { // check if the word starts with letter
			for ; j < length && !isVowelChar(word[j]); j++ {
			}
			k := length - 1 // line that indicate not letter
			// notLetterFlag := false
			for ; k > j && !isLetter(word[k]); k-- {
			}
			k += 1 // make shift because now k is before non letter symbol

			if j == 0 { // it means that the word starts with vowel
				fmt.Fprintf(&resBytes, "%syay", word[:k])
			} else {
				fmt.Fprintf(&resBytes, "%s%say", word[j:k], word[:j])
			}
			if k < length {

				fmt.Fprintf(&resBytes, "%s", word[k:])
			}
		} else {
			fmt.Fprintf(&resBytes, "%s", word)
		}

		fmt.Fprintf(&resBytes, " ")
	}

	result = strings.TrimSpace(resBytes.String())

	return &result
}

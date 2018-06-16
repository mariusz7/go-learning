package main

import (
	"fmt"
)

func main() {

	var permutatedString []rune = []rune("Kęd")

	permutation(permutatedString)
}

func permutation(str []rune) {

	permutationImpl(str, nil)
}

func permutationImpl(str []rune, prefix []rune) {

	if len(str) == 0 {
		if nil != prefix {
			fmt.Println(string(prefix))
		}
	} else {
		for i := 0; i < len(str); i++ {

			var str2 []rune = make([]rune, len(str) - 1)
			copy(str2, str[0:i])
			copy(str2[i:], str[i + 1:len(str)])

			var prefix2 []rune = make([]rune, len(prefix))
			copy(prefix2, prefix)
			prefix2 = append(prefix2, str[i])

			permutationImpl(str2, prefix2)
		}
	}
}

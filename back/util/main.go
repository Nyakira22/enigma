package util

var ALPHABET string = GenerateAlphabet()

func GenerateAlphabet() string {
	alphabet := make([]byte, 0, 26)
	var ch byte
	for ch = 'A'; ch <= 'Z'; ch++ {
		alphabet = append(alphabet, ch)
	}
	return string(alphabet)
}

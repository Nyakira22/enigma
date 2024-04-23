package main

import (
	"fmt"
	"strings"
)

var ALPHABET string = generateAlphabet()

func main() {
	mapAlphabet := "BADC"
	if len(mapAlphabet) > len(ALPHABET) {
		mapAlphabet = mapAlphabet[:len(ALPHABET)]
	}
	plugboard := NewAlphabetMapper(mapAlphabet)

	encrypted_index := plugboard.forward(strings.Index(ALPHABET, "C"))
	fmt.Println(string(ALPHABET[encrypted_index]))
	decrypted_index := plugboard.backward(encrypted_index)
	fmt.Println(string(ALPHABET[decrypted_index]))
}

type PlugBoard struct {
	alphabet    string
	forwardMap  map[string]string
	backwardMap map[string]string
}

func NewAlphabetMapper(mapAlphabet string) *PlugBoard {
	p := &PlugBoard{
		alphabet:    ALPHABET[:len(mapAlphabet)],
		forwardMap:  make(map[string]string),
		backwardMap: make(map[string]string),
	}
	p.mapping(mapAlphabet)
	return p
}

func (p *PlugBoard) mapping(mapAlphabet string) {
	for i, char := range p.alphabet {
		p.forwardMap[string(char)] = string(mapAlphabet[i])
		p.backwardMap[string(mapAlphabet[i])] = string(char)
	}
}

func (p *PlugBoard) forward(index_num int) int {
	//n[0]はbyteを返すから、一旦runeスライスを作成する
	char := string(getRuneAt(p.alphabet, index_num))
	char = p.forwardMap[char]
	return strings.Index(p.alphabet, char)
}

func (p *PlugBoard) backward(index_num int) int {
	//n[0]はbyteを返すから、一旦runeスライスを作成する
	char := string(getRuneAt(p.alphabet, index_num))
	char = p.backwardMap[char]
	return strings.Index(p.alphabet, char)
}

func generateAlphabet() string {
	alphabet := make([]byte, 0, 26)
	var ch byte
	for ch = 'A'; ch <= 'Z'; ch++ {
		alphabet = append(alphabet, ch)
	}
	return string(alphabet)
}

func getRuneAt(s string, i int) rune {
	rs := []rune(s)
	return rs[i]
}

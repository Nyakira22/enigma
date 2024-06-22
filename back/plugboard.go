package main

import (
	"strings"
)

// PlugBoard
type PlugBoard struct {
	alphabet    string
	forwardMap  map[string]string
	backwardMap map[string]string
}

func NewPlugBoard(mapAlphabet string) *PlugBoard {
	p := &PlugBoard{
		alphabet:    ALPHABET,
		forwardMap:  make(map[string]string),
		backwardMap: make(map[string]string),
	}
	p.mapping(mapAlphabet)
	return p
}

func (p *PlugBoard) mapping(mapAlphabet string) {
	//渡された文字列の長さ分だけのALPHABETを取得しループ
	for i, char := range p.alphabet[:len(mapAlphabet)] {
		p.forwardMap[string(char)] = string(mapAlphabet[i])
		p.backwardMap[string(mapAlphabet[i])] = string(char)
	}
}

func (p *PlugBoard) forward(index_num int) int {
	//n[i]はbyteを返すので、一旦runeスライスを作成する
	char := string(getRuneAt(p.alphabet, index_num))
	char = p.forwardMap[char]
	return strings.Index(p.alphabet, char)
}

func (p *PlugBoard) backward(index_num int) int {
	//n[0]はbyteを返すので、一旦runeスライスを作成する
	char := string(getRuneAt(p.alphabet, index_num))
	char = p.backwardMap[char]
	return strings.Index(p.alphabet, char)
}

func getRuneAt(s string, i int) rune {
	rs := []rune(s)
	return rs[i]
}

func generateAlphabet() string {
	alphabet := make([]byte, 0, 26)
	var ch byte
	for ch = 'A'; ch <= 'Z'; ch++ {
		alphabet = append(alphabet, ch)
	}
	return string(alphabet)
}

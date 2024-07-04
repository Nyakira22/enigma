package plugboard

import (
	"strings"

	"github.com/akiradomi/workspace/go-enigma/enigma/back/util"
)

// PlugBoard
type PlugBoard struct {
	Alphabet    string
	forwardMap  map[string]string
	backwardMap map[string]string
}

func NewPlugBoard(mapAlphabet string) *PlugBoard {
	p := &PlugBoard{
		Alphabet:    util.ALPHABET,
		forwardMap:  make(map[string]string),
		backwardMap: make(map[string]string),
	}
	p.Mapping(mapAlphabet)
	return p
}

func (p *PlugBoard) Mapping(mapAlphabet string) {
	//渡された文字列の長さ分だけのALPHABETを取得しループ
	for i, char := range p.Alphabet[:len(mapAlphabet)] {
		p.forwardMap[string(char)] = string(mapAlphabet[i])
		p.backwardMap[string(mapAlphabet[i])] = string(char)
	}
}

func (p *PlugBoard) Forward(index_num int) int {
	//n[i]はbyteを返すので、一旦runeスライスを作成する
	char := string(GetRuneAt(p.Alphabet, index_num))
	char = p.forwardMap[char]
	return strings.Index(p.Alphabet, char)
}

func (p *PlugBoard) Backward(index_num int) int {
	//n[0]はbyteを返すので、一旦runeスライスを作成する
	char := string(GetRuneAt(p.Alphabet, index_num))
	char = p.backwardMap[char]
	return strings.Index(p.Alphabet, char)
}

func GetRuneAt(s string, i int) rune {
	rs := []rune(s)
	return rs[i]
}

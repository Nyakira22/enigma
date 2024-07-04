package roter

import (
	"github.com/akiradomi/workspace/go-enigma/enigma/back/cmd/plugboard"
	"github.com/akiradomi/workspace/go-enigma/enigma/back/util"
)

// Roter
type Roter struct {
	plugboard.PlugBoard
	Offset    int
	Rotations int
}

func NewRoter(mapAlphabet string, offset int) *Roter {
	parent := plugboard.NewPlugBoard(mapAlphabet)
	r := &Roter{
		PlugBoard: *parent,
		Offset:    offset,
		Rotations: 0,
	}
	r.Mapping(mapAlphabet)
	return r
}

// アルファベットの順をoffsetの数値だけローテーションする
func (r *Roter) Rotate(offset int) int {
	r.Alphabet = r.Alphabet[offset:] + r.Alphabet[:offset]
	r.Rotations += offset
	return r.Rotations
}

// ローターの設定を初期値に戻す
func (r *Roter) Reset() {
	r.Alphabet = util.ALPHABET
	r.Rotations = 0
}

package main

// Roter
type Roter struct {
	PlugBoard
	offset    int
	rotations int
}

func NewRoter(mapAlphabet string, offset int) *Roter {
	parent := NewPlugBoard(mapAlphabet)
	r := &Roter{
		PlugBoard: *parent,
		offset:    offset,
		rotations: 0,
	}
	r.mapping(mapAlphabet)
	return r
}

// アルファベットの順をoffsetの数値だけローテーションする
func (r *Roter) rotate(offset int) int {
	r.alphabet = r.alphabet[offset:] + r.alphabet[:offset]
	r.rotations += offset
	return r.rotations
}

// ローターの設定を初期値に戻す
func (r *Roter) reset() {
	r.alphabet = ALPHABET
	r.rotations = 0
}

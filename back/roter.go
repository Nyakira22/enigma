package main

// Roter
type Roter struct {
	PlugBoard
	offset    int
	rotations int
}

// PlugBoardを埋め込んで(embedded)初期化
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

func (r *Roter) rotate(offset int) int {
	//offsetの値を元にALPHABETの先頭からn文字を後ろに移動させる
	r.alphabet = r.alphabet[offset:] + r.alphabet[:offset]
	r.rotations += offset
	return r.rotations
}

func (r *Roter) reset() {
	r.alphabet = ALPHABET
	r.rotations = 0
}

package main

import (
	"fmt"
	"os"
	"strings"
)

var ALPHABET string = generateAlphabet()

func main() {
	mapAlphabet := "BADC"
	if len(mapAlphabet) > len(ALPHABET) {
		mapAlphabet = mapAlphabet[:len(ALPHABET)]
	}
	// plugboard := NewAlphabetMapper(mapAlphabet)
	// encrypted_index := plugboard.forward(strings.Index(ALPHABET, "A"))
	// fmt.Println(string(ALPHABET[encrypted_index]))
	// decrypted_index := plugboard.backward(encrypted_index)
	// fmt.Println(string(ALPHABET[decrypted_index]))

	// roter := NewRoter(mapAlphabet, 1)
	// encrypted_index := roter.forward(strings.Index(ALPHABET, "A"))
	// fmt.Println(string(ALPHABET[encrypted_index]))
	// decrypted_index := roter.backward(encrypted_index)
	// fmt.Println(string(ALPHABET[decrypted_index]))

	// roter.rotate(1)

	// encrypted_index_r := roter.forward(strings.Index(ALPHABET, "A"))
	// fmt.Println(string(ALPHABET[encrypted_index_r]))
	// decrypted_index_r := roter.backward(encrypted_index_r)
	// fmt.Println(string(ALPHABET[decrypted_index_r]))

	refrector := NewReflector(mapAlphabet)
	i := refrector.reflect(2)
	fmt.Println(string(ALPHABET[i]))
}

// PlugBoard
type PlugBoard struct {
	alphabet    string
	forwardMap  map[string]string
	backwardMap map[string]string
}

// コンストラクタ
func NewAlphabetMapper(mapAlphabet string) *PlugBoard {
	//mapの初期化
	p := &PlugBoard{
		alphabet:    ALPHABET,
		forwardMap:  make(map[string]string),
		backwardMap: make(map[string]string),
	}
	p.mapping(mapAlphabet)
	return p
}

func (p *PlugBoard) mapping(mapAlphabet string) {
	//渡された文字列の長さ分だけのALPHABETを取得しループ、ALPHABETの分だけループしようとするとエラー
	for i, char := range p.alphabet[:len(mapAlphabet)] {
		p.forwardMap[string(char)] = string(mapAlphabet[i])
		p.backwardMap[string(mapAlphabet[i])] = string(char)
	}
}

func (p *PlugBoard) forward(index_num int) int {
	//n[0]はbyteを返すので、一旦runeスライスを作成する
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

// Roter
type Roter struct {
	PlugBoard
	offset    int
	rotations int
}

// PlugBoardを埋め込んで(embedded)初期化
func NewRoter(mapAlphabet string, offset int) *Roter {
	parent := NewAlphabetMapper(mapAlphabet)
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

// Reflector
type Reflector struct {
	reflectorMap map[string]string
}

func NewReflector(mapAlphabet string) *Reflector {
	ref := &Reflector{
		reflectorMap: make(map[string]string),
	}

	//渡されたtextとALPHABETの文字列からmapを生成
	for i, char := range ALPHABET[:len(mapAlphabet)] {
		ref.reflectorMap[string(char)] = string(mapAlphabet[i])
	}

	//生成したmapのkey:valueが対の関係になっているかチェック
	for x, y := range ref.reflectorMap {
		if x != ref.reflectorMap[y] {
			os.Exit(1)
		}
	}
	return ref
}

func (ref *Reflector) reflect(index_num int) int {
	//渡されたindex番号の文字列をキーとする文字列をmapから取得
	reflected_char := ref.reflectorMap[string(ALPHABET[index_num])]
	for i, v := range ALPHABET {
		//ALPHABETの文字列から指定の文字列のindex番号を取得しreturn
		if string(v) == reflected_char {
			return i
		}
	}
	panic("Error")
}

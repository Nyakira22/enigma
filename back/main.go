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
	//インスタンスを作成
	plugboard := NewPlugBoard(mapAlphabet)
	roter1 := NewRoter(mapAlphabet, 3)
	roter2 := NewRoter(mapAlphabet, 2)
	roter3 := NewRoter(mapAlphabet, 1)
	refrector := NewReflector(mapAlphabet)
	//roterのスライスを作成
	roters := []*Roter{
		roter1,
		roter2,
		roter3,
	}
	EnigmaMachine := NewEnigmaMachine(*plugboard, *refrector, roters)
	encryptedText := EnigmaMachine.encript("T")
	fmt.Println(encryptedText)

}

// PlugBoard
type PlugBoard struct {
	alphabet    string
	forwardMap  map[string]string
	backwardMap map[string]string
}

// コンストラクタ
func NewPlugBoard(mapAlphabet string) *PlugBoard {
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

// enigmamachine roterは配列で複数定義できるようにする
type EnigmaMachine struct {
	PlugBoard
	Reflector
	Roters []*Roter
}

func NewEnigmaMachine(PlugBoard PlugBoard, Reflector Reflector, Roters []*Roter) *EnigmaMachine {
	e := &EnigmaMachine{
		PlugBoard: PlugBoard,
		Reflector: Reflector,
		Roters:    Roters,
	}
	return e
}

func (e *EnigmaMachine) encript(text string) string {
	s := make([]string, 0)
	for _, char := range text {
		//一連の変換処理を実行する
		fmt.Println(string(char))
		encryptedChar := e.goThrough(string(char))
		s = append(s, encryptedChar)
	}
	//変換処理で取得した文字列のスライスをjoinで結合してreturn
	return strings.Join(s, "")
}

func (e *EnigmaMachine) decript(text string) string {
	s := make([]string, 0)
	//ローテーションしたroterを初期位置に戻す
	for _, roter := range e.Roters {
		roter.reset()
	}
	for _, char := range text {
		//一連の変換処理を実行する
		fmt.Println(string(char))
		encryptedChar := e.goThrough(string(char))
		s = append(s, encryptedChar)
	}
	return strings.Join(s, "")
}

func (e *EnigmaMachine) goThrough(char string) string {
	char = strings.ToUpper(char)
	//文字列がアルファベットにないものだったらそのまま返す
	if !strings.Contains(ALPHABET, char) {
		return char
	}
	indexNum := strings.Index(ALPHABET, char)
	// fmt.Println(indexNum, "first")

	indexNum = e.PlugBoard.forward(indexNum)
	// fmt.Println(indexNum, "plugboard_forward")

	for _, roter := range e.Roters {
		indexNum = roter.forward(indexNum)
	}

	indexNum = e.Reflector.reflect(indexNum)
	// fmt.Println(indexNum, "reflector")

	//roterを逆順で回してbackwardする
	for i := 0; i < len(e.Roters)/2; i++ {
		e.Roters[i], e.Roters[len(e.Roters)-i-1] = e.Roters[len(e.Roters)-i-1], e.Roters[i]
	}
	//逆順になった各ローターでbackward
	for _, roter := range e.Roters {
		indexNum = roter.backward(indexNum)
	}
	// fmt.Println(indexNum, "roter_backward")

	indexNum = e.PlugBoard.backward(indexNum)
	// fmt.Println(indexNum, "plugboard_backward")

	//逆順になったローターをローテーションする
	for _, roter := range e.Roters {
		if roter.rotate(0)%len(ALPHABET) != 0 {
			break
		}
	}

	//逆順になったroterを元に戻す
	for i := 0; i < len(e.Roters)/2; i++ {
		e.Roters[i], e.Roters[len(e.Roters)-i-1] = e.Roters[len(e.Roters)-i-1], e.Roters[i]
	}

	char = string(ALPHABET[indexNum])
	// fmt.Println(string(ALPHABET[indexNum]))
	return char

}

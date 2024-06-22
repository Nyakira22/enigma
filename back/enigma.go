package main

import (
	"strings"
)

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
		encryptedChar := e.goThrough(string(char))
		s = append(s, encryptedChar)
	}
	return strings.Join(s, "")
}

func (e *EnigmaMachine) goThrough(char string) string {
	char = strings.ToUpper(char)
	if !strings.Contains(ALPHABET, char) {
		return char
	}
	indexNum := strings.Index(ALPHABET, char)
	indexNum = e.PlugBoard.forward(indexNum)
	for _, roter := range e.Roters {
		indexNum = roter.forward(indexNum)
	}
	indexNum = e.Reflector.reflect(indexNum)
	//roterを逆順にする
	for i := 0; i < len(e.Roters)/2; i++ {
		e.Roters[i], e.Roters[len(e.Roters)-i-1] = e.Roters[len(e.Roters)-i-1], e.Roters[i]
	}
	//逆順になったローターでbackwardを実行
	for _, roter := range e.Roters {
		indexNum = roter.backward(indexNum)
	}
	indexNum = e.PlugBoard.backward(indexNum)
	//ローターをローテーションする
	for _, roter := range e.Roters {
		if roter.rotate(roter.offset)%len(ALPHABET) != 0 {
			break
		}
	}
	//逆順になったroterを元に戻す
	for i := 0; i < len(e.Roters)/2; i++ {
		e.Roters[i], e.Roters[len(e.Roters)-i-1] = e.Roters[len(e.Roters)-i-1], e.Roters[i]
	}
	char = string(ALPHABET[indexNum])
	return char
}

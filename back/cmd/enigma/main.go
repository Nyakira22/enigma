package enigma

import (
	"fmt"
	"strings"

	"github.com/akiradomi/workspace/go-enigma/enigma/back/cmd/plugboard"
	"github.com/akiradomi/workspace/go-enigma/enigma/back/cmd/reflector"
	"github.com/akiradomi/workspace/go-enigma/enigma/back/cmd/roter"
	"github.com/akiradomi/workspace/go-enigma/enigma/back/util"
)

// enigmamachine roterは配列で複数定義できるようにする
type EnigmaMachine struct {
	plugboard.PlugBoard
	reflector.Reflector
	Roters []*roter.Roter
}

func NewEnigmaMachine(PlugBoard plugboard.PlugBoard, Reflector reflector.Reflector, Roters []*roter.Roter) *EnigmaMachine {
	e := &EnigmaMachine{
		PlugBoard: PlugBoard,
		Reflector: Reflector,
		Roters:    Roters,
	}
	return e
}

func (e *EnigmaMachine) Encript(text string) string {
	fmt.Println(text)
	s := make([]string, 0)
	for _, char := range text {
		//一連の変換処理を実行する
		encryptedChar := e.GoThrough(string(char))
		s = append(s, encryptedChar)
	}
	//変換処理で取得した文字列のスライスをjoinで結合してreturn
	return strings.Join(s, "")
}

func (e *EnigmaMachine) Decript(text string) string {
	s := make([]string, 0)
	//ローテーションしたroterを初期位置に戻す
	for _, roter := range e.Roters {
		roter.Reset()
	}
	for _, char := range text {
		//一連の変換処理を実行する
		encryptedChar := e.GoThrough(string(char))
		s = append(s, encryptedChar)
	}
	return strings.Join(s, "")
}

func (e *EnigmaMachine) GoThrough(char string) string {
	char = strings.ToUpper(char)
	if !strings.Contains(util.ALPHABET, char) {
		return char
	}
	indexNum := strings.Index(util.ALPHABET, char)
	indexNum = e.PlugBoard.Forward(indexNum)
	for _, roter := range e.Roters {
		indexNum = roter.Forward(indexNum)
	}
	indexNum = e.Reflector.Reflect(indexNum)
	//roterを逆順にする
	for i := 0; i < len(e.Roters)/2; i++ {
		e.Roters[i], e.Roters[len(e.Roters)-i-1] = e.Roters[len(e.Roters)-i-1], e.Roters[i]
	}
	//逆順になったローターでbackwardを実行
	for _, roter := range e.Roters {
		indexNum = roter.Backward(indexNum)
	}
	indexNum = e.PlugBoard.Backward(indexNum)
	//ローターをローテーションする
	for _, roter := range e.Roters {
		if roter.Rotate(roter.Offset)%len(util.ALPHABET) != 0 {
			break
		}
	}
	//逆順になったroterを元に戻す
	for i := 0; i < len(e.Roters)/2; i++ {
		e.Roters[i], e.Roters[len(e.Roters)-i-1] = e.Roters[len(e.Roters)-i-1], e.Roters[i]
	}
	char = string(util.ALPHABET[indexNum])
	return char
}

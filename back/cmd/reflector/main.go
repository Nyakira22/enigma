package reflector

import (
	"fmt"
	"os"

	"github.com/akiradomi/workspace/go-enigma/enigma/back/util"
)

// Reflector
type Reflector struct {
	reflectorMap map[string]string
}

func NewReflector(mapAlphabet string) *Reflector {
	ref := &Reflector{
		reflectorMap: make(map[string]string),
	}

	//渡されたtextとALPHABETの文字列からmapを生成
	for i, char := range util.ALPHABET[:len(mapAlphabet)] {
		ref.reflectorMap[string(char)] = string(mapAlphabet[i])
	}

	//生成したmapのkey:valueが対の関係になっているかチェック
	for x, y := range ref.reflectorMap {
		if x != ref.reflectorMap[y] {
			fmt.Println("ValueError", x, y)
			os.Exit(1)
		}
	}
	return ref
}

func (ref *Reflector) Reflect(index_num int) int {
	//渡されたindex番号の文字列をキーとする文字列をmapから取得
	reflected_char := ref.reflectorMap[string(util.ALPHABET[index_num])]
	for i, v := range util.ALPHABET {
		//ALPHABETの文字列から指定の文字列のindex番号を取得しreturn
		if string(v) == reflected_char {
			return i
		}
	}
	panic("Error")
}

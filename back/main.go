package main

import (
	"fmt"
	"math/rand"

	"github.com/akiradomi/workspace/go-enigma/enigma/back/cmd/enigma"
	"github.com/akiradomi/workspace/go-enigma/enigma/back/cmd/plugboard"
	"github.com/akiradomi/workspace/go-enigma/enigma/back/cmd/reflector"
	"github.com/akiradomi/workspace/go-enigma/enigma/back/cmd/roter"
	"github.com/akiradomi/workspace/go-enigma/enigma/back/util"
)

// var util.ALPHABET string = generateAlphabet()

func main() {
	pulugboard := plugboard.NewPlugBoard(getRandomAlphabet())
	roter1 := roter.NewRoter(getRandomAlphabet(), 3)
	roter2 := roter.NewRoter(getRandomAlphabet(), 2)
	roter3 := roter.NewRoter(getRandomAlphabet(), 1)
	roters := []*roter.Roter{
		roter1,
		roter2,
		roter3,
	}

	//アルファベットのスライスを生成
	alphabetList := make([]string, 0, len(util.ALPHABET))
	for _, char := range util.ALPHABET {
		alphabetList = append(alphabetList, string(char))
	}

	//アルファベットリストのインデックスを取得
	indexes := make([]int, len(alphabetList))
	for i := range indexes {
		indexes[i] = i
	}

	// ランダムなアルファベットののペアを生成
	r := []rune(util.ALPHABET)
	for i := 0; i < int(len(r)/2); i++ {
		x := rand.Intn(len(indexes))
		index_x := indexes[x]
		indexes = append(indexes[:x], indexes[x+1:]...)
		y := rand.Intn(len(indexes))
		index_y := indexes[y]
		indexes = append(indexes[:y], indexes[y+1:]...)
		r[index_x], r[index_y] = r[index_y], r[index_x]
	}

	reflector := reflector.NewReflector(string(r))
	enigma := enigma.NewEnigmaMachine(*pulugboard, *reflector, roters)
	s := "enigma with golang test"
	e := enigma.Encript(s)
	fmt.Println(e)
	d := enigma.Decript(e)
	fmt.Println(d)

}

// ランダムなアルファベット文字列を生成
func getRandomAlphabet() string {
	randIndices := rand.Perm(len(util.ALPHABET))
	randomString := make([]byte, len(util.ALPHABET))
	for i, idx := range randIndices {
		randomString[i] = util.ALPHABET[idx]
	}
	return string(randomString)
}

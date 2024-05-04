package main

import (
	"fmt"
	"math/rand"
)

var ALPHABET string = generateAlphabet()

func main() {
	//インスタンスを作成
	pulugboard := NewPlugBoard(getRandomAlphabet())
	roter1 := NewRoter(getRandomAlphabet(), 3)
	roter2 := NewRoter(getRandomAlphabet(), 2)
	roter3 := NewRoter(getRandomAlphabet(), 1)
	roters := []*Roter{
		roter1,
		roter2,
		roter3,
	}

	//アルファベットのスライスを生成
	alphabetList := make([]string, 0, len(ALPHABET))
	for _, char := range ALPHABET {
		alphabetList = append(alphabetList, string(char))
	}

	//アルファベットリストのインデックスを取得
	indexes := make([]int, len(alphabetList))
	for i := range indexes {
		indexes[i] = i
	}

	r := []rune(ALPHABET)
	// ランダムな要素のペアを選んで交換
	for i := 0; i < int(len(r)/2); i++ {
		x := rand.Intn(len(indexes))
		index_x := indexes[x]
		indexes = append(indexes[:x], indexes[x+1:]...)
		y := rand.Intn(len(indexes))
		index_y := indexes[y]
		indexes = append(indexes[:y], indexes[y+1:]...)
		r[index_x], r[index_y] = r[index_y], r[index_x]
	}

	reflector := NewReflector(string(r))
	enigma := NewEnigmaMachine(*pulugboard, *reflector, roters)
	s := "GOLANG PRACTICE"
	e := enigma.encript(s)
	fmt.Println(e)
	d := enigma.decript(e)
	fmt.Println(d)

}

func getRandomAlphabet() string {
	randIndices := rand.Perm(len(ALPHABET))
	randomString := make([]byte, len(ALPHABET))
	for i, idx := range randIndices {
		randomString[i] = ALPHABET[idx]
	}
	return string(randomString)
}

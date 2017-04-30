package main

import(
	"fmt"
	"os"
) 

func main() {
	origin := os.Args[1]
	fmt.Println(anagramList(origin))
}

func anagramList(word string) []string {
	if len(word) == 1 {
		return []string{word}
	}

	ret := []string{}
	for i := 0; i < len(word); i++ {
	  words := anagramList(word[0:i] + word[i+1:len(word)])
		for _, w := range(words) {
			ret = append(ret, word[i:i+1] + w)
		}
	}
	return ret
}

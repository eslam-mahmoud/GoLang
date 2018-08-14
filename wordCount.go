//https://tour.golang.org/moretypes/23

package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	words := make(map[string]int)
	stringFields := strings.Fields(s);
	for i := 0; i < len(stringFields) ; i++ {
		if words[stringFields[i]] != 0 {
			words[stringFields[i]] = words[stringFields[i]]+1
		} else {
			words[stringFields[i]] = 1
		}
	}
	return words
}

func main() {
	wc.Test(WordCount)
}

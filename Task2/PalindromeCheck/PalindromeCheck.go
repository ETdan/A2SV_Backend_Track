package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your Word:")
	word, _ := reader.ReadString('\n')
	word = strings.TrimSpace(word)

	i := 0
	j := len(word) - 1
	notPalindrom := false
	if j > 1 {

		for i < j {
			if word[i] != word[j] {
				notPalindrom = true
				break
			}
			i++
			j--
		}
	}
	if notPalindrom == true {
		fmt.Println(word, " is Not Palindrom")
	} else {
		fmt.Println(word, " is Palindrom")
	}
}

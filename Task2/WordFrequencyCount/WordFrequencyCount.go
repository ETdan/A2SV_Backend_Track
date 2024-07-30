package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your Word: ")
	word, _ := reader.ReadString('\n')
	// fmt.Println(word)
	// fmt.Printf("%T %v", word, word)
	// arr := []string{}
	arr := strings.Split(word, " ")
	counter := make(map[string]int)
	for _, val := range arr {
		counter[val] += 1
	}
	fmt.Println("")
	for key, val := range counter {

		fmt.Printf("%-35v %v\n", key, val)
	}
}

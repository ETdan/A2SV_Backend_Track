package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func stripRegex(in string) string {
	reg, _ := regexp.Compile("[^a-zA-Z0-9 ]+")
	return reg.ReplaceAllString(in, "")
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your Word: ")
	word, _ := reader.ReadString('\n')
	arr := strings.Split(strings.ToLower(word), " ")
	for index, val := range arr {

		arr[index] = stripRegex(val)

	}

	counter := make(map[string]int)
	for _, val := range arr {
		counter[val] += 1
	}
	fmt.Println("")
	for key, val := range counter {
		fmt.Printf("%-35v %v\n", key, val)
	}
}

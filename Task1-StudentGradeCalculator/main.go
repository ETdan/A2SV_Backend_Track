package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter name: ")
	name, _ := reader.ReadString('\n')

	s := newStudent(name)
	s.accept_subjects()
	s.calculate_average()

	fmt.Println(s.get_result())

}

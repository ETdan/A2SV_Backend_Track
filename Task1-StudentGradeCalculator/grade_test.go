package main

import (
	"fmt"
	"testing"
)

func TestNewStudent(t *testing.T) {
	s := newStudent("Daniel")

	if s.name != "Daniel" {
		t.Error("object not created with expected name")
		return
	}

	s.accept_subjects()
	fmt.Println("2")
	fmt.Println("a")
	fmt.Println("5")
	fmt.Println("b")
	fmt.Println("5")
}

package main

import "fmt"

type student struct {
	name     string
	subjects map[string]float64
	average  float64
}

func newStudent(name string) student {
	s := student{
		name:     name,
		subjects: map[string]float64{},
		average:  0.0,
	}
	return s
}

func (s student) get_name() string {
	return s.name
}

func (s student) get_subjects() map[string]float64 {
	return s.subjects
}

func (s *student) accept_subjects() {
	var n int

	fmt.Print("how many subjects:")
	fmt.Scan(&n)

	for ; n > 0; n-- {
		var subject string
		var grade float64

		fmt.Print("Enter name of the subjects: ")
		fmt.Scan(&subject)

		fmt.Print("Enter grade(0-100): ")
		fmt.Scan(&grade)

		for grade < 0 || grade > 100 {
			fmt.Print("Enter grade(0-100): ")
			fmt.Scan(&grade)
		}

		s.subjects[subject] = grade
	}
}

func (s *student) calculate_average() {
	average := 0.0
	for _, val := range s.subjects {
		average += val
	}
	s.average = (float64(average) / float64(len(s.subjects)))
}
func (s student) get_average() float64 {
	return s.average
}

func (s student) get_result() string {
	f := "\n"
	f += "*******************************\n"
	f += fmt.Sprintf("%-8v %v\n", "", s.name)
	f += "*******************************\n"
	f += fmt.Sprintf("%-25v %v\n", "Subjects", "Grade")
	f += "\n"
	for key, val := range s.subjects {
		f += fmt.Sprintf("%-25v %v\n", key, val)
	}
	f += fmt.Sprintf("%-25v %v\n", "Average:", s.average)
	f += "\n*******************************\n"
	return f
}

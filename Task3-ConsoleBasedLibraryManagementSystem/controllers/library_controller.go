package controllers

import (
	"BookKeeper/models"
	"BookKeeper/service"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var l = service.Library{
	Curr_book_ID:   0,
	Curr_member_ID: 0,
	Books:          map[int]models.Book{},
	Meambers:       map[int]models.Member{},
}

func Choose() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%-10v %v \n", "", "*********************************************")
	fmt.Printf("%-10v %v \n", "", "Welcome to BookKeper your Library Manager")
	fmt.Printf("%-10v %v \n", "", "*********************************************")
	fmt.Printf("%-20v %v \n", "", "Press 1 to Add Book")
	fmt.Printf("%-20v %v \n", "", "Press 2 to Add Member")
	fmt.Printf("%-20v %v \n", "", "Press 3 to Remove Book")
	fmt.Printf("%-20v %v \n", "", "Press 4 to Return Book")
	fmt.Printf("%-20v %v \n", "", "Press 5 to Borrow Book")
	fmt.Printf("%-20v %v \n", "", "Press 6 to List Available Book")
	fmt.Printf("%-20v %v \n", "", "Press 7 to List Borrowed Book")
	fmt.Printf("%-20v %v \n", "", "Press 8 to List All Book")
	fmt.Printf("%-20v %v \n", "", "Press 9 to List All Memberes")
	fmt.Printf("%-20v %v \n", "", "Press 10 to List Stat")
	fmt.Printf("%-20v %v \n", "", "Press 0 to Exit")
	fmt.Printf("%-10v %v \n", "", "*********************************************")
	fmt.Print("Enter your choice: ")
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)
	breaked := false
	switch choice {
	case "1":
		AddBookInput(&l, reader)
	case "2":
		AddMemberInput(&l, reader)
	case "3":
		RemoveBookInput(&l, reader)
	case "4":
		ReturnBookInput(&l, reader)
	case "5":
		BorrowBookInput(&l, reader)
	case "6":
		ListAvailableBooksInput(&l, reader)
	case "7":
		ListBorrowedBooksInput(&l, reader)
	case "8":
		ListBooks(l)
	case "9":
		ListMembers(l)
	case "10":
		ListStat(l)
	case "0":
		breaked = true
		break
	default:
		Choose()
	}
	if !breaked {
		Choose()
	} else {
		fmt.Println("Thnak! come again")
	}
}

func ListStat(l service.Library) {
	fmt.Println("you choosed to See Stat")
	fmt.Printf("No of Book:%v No of Members%v\n", l.Curr_book_ID, l.Curr_member_ID)
}
func ListMembers(l service.Library) {
	fmt.Println("you choosed List members ")
	fmt.Println(l.ListMembers())
}
func ListBooks(l service.Library) {
	fmt.Println("you choosed to list books ")
	fmt.Println(l.ListBooks())
}
func AddBookInput(l *service.Library, reader *bufio.Reader) {
	fmt.Println("you choosed to add A book pls fill the following")
	fmt.Print("Title: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Print("Author: ")
	author, _ := reader.ReadString('\n')
	author = strings.TrimSpace(author)
	b := models.Book{
		ID:     l.Curr_book_ID,
		Title:  title,
		Author: author,
		Status: "Available",
	}
	l.Curr_book_ID++
	fmt.Printf("ID: %v Title: %v Author: %v Status:%v\n", b.ID, b.Title, b.Author, b.Status)
	l.AddBook(b)
}
func AddMemberInput(l *service.Library, reader *bufio.Reader) {
	fmt.Println("you choosed to add a Member pls fill the following")
	fmt.Print("Name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	m := models.Member{
		ID:            l.Curr_member_ID,
		Name:          name,
		BorrowedBooks: []models.Book{},
	}
	l.Curr_member_ID++
	l.AddMeamber(m)
}

func RemoveBookInput(l *service.Library, reader *bufio.Reader) {
	fmt.Println("you choosed to Remove a Book pls fill the following")
	fmt.Print("Book ID: ")
	BookId, _ := reader.ReadString('\n')
	BookId = strings.TrimSpace(BookId)
	bookId, _ := strconv.Atoi(BookId)
	l.RemoveBook(bookId)
	fmt.Println("you successfuly Removed a Book")
}
func ReturnBookInput(l *service.Library, reader *bufio.Reader) {
	fmt.Println("you choosed to Return a Book pls fill the following")
	fmt.Print("Member ID: ")
	MemberId, _ := reader.ReadString('\n')
	MemberId = strings.TrimSpace(MemberId)
	memberID, _ := strconv.Atoi(MemberId)
	fmt.Print("Book ID: ")
	BookId, _ := reader.ReadString('\n')
	BookId = strings.TrimSpace(BookId)
	bookId, _ := strconv.Atoi(BookId)
	if l.ReturnBook(bookId, memberID) == nil {

		fmt.Println("you successfuly Returned a Book")
	} else {
		fmt.Println("No Data found with the given input")
	}
}
func BorrowBookInput(l *service.Library, reader *bufio.Reader) {
	fmt.Println("you choosed to Borrow a Book pls fill the following")
	fmt.Print("Member ID: ")
	MemberId, _ := reader.ReadString('\n')
	MemberId = strings.TrimSpace(MemberId)
	memberID, _ := strconv.Atoi(MemberId)
	fmt.Print("Book ID: ")
	BookId, _ := reader.ReadString('\n')
	BookId = strings.TrimSpace(BookId)
	bookId, _ := strconv.Atoi(BookId)
	l.BorrowBook(bookId, memberID)

	if l.BorrowBook(bookId, memberID) == nil {
		fmt.Println("you successfuly Borrowed a Book")
	} else {
		fmt.Println("No Data found with the given input")
	}
}
func ListAvailableBooksInput(l *service.Library, reader *bufio.Reader) {
	fmt.Println("you choosed to see list of Available pls fill the following")
	books := l.ListAvailableBooks()
	fmt.Println(books)
}
func ListBorrowedBooksInput(l *service.Library, reader *bufio.Reader) {
	fmt.Println("you choosed to see LIst of Borrowed Books pls fill the following")
	fmt.Print("Member ID: ")

	MemberId, _ := reader.ReadString('\n')
	MemberId = strings.TrimSpace(MemberId)
	memberID, _ := strconv.Atoi(MemberId)

	books := l.ListBorrowedBooks(memberID)
	fmt.Println(books)
}

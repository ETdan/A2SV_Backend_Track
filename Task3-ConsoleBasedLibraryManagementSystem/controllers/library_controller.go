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
	Curr_book_ID:   4,
	Curr_member_ID: 4,
	Books: map[int]models.Book{
		1: {ID: 1, Title: "To Kill a Mockingbird", Author: "Harper Lee", Status: "Borrowed"},
		2: {ID: 2, Title: "1984", Author: "George Orwell", Status: "Borrowed"},
		3: {ID: 3, Title: "Moby Dick", Author: "Herman Melville", Status: "Available"},
	},
	Meambers: map[int]models.Member{
		1: {
			ID:   1,
			Name: "Alice Smith",
			BorrowedBooks: []models.Book{
				{ID: 2, Title: "1984", Author: "George Orwell", Status: "Borrowed"},
			},
		},
		2: {
			ID:   2,
			Name: "Bob Johnson",
			BorrowedBooks: []models.Book{
				{ID: 1, Title: "To Kill a Mockingbird", Author: "Harper Lee", Status: "Borrowed"},
			},
		}},
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

// Done
func ListStat(l service.Library) {
	fmt.Println("you choosed to See Stat")
	fmt.Printf("\tNo of Book:%v \n\tNo of Members%v\n", l.Curr_book_ID, l.Curr_member_ID)
}

// Done
func ListMembers(l service.Library) {
	fmt.Println("you choosed List members ")

	fmt.Printf("%v %v \n", "", "*********************************************")
	fmt.Printf("%v %v \n", "", "Members")
	fmt.Printf("%v %v \n", "", "*********************************************")
	fmt.Printf("\t%v \t%-20v \t%v\n", "ID", "Name", "Borrowed Books")
	fmt.Println("-----------------------------------------------------------------------------------")
	for _, member := range l.ListMembers() {
		fmt.Printf("\t%v \t%-20v \t%v\n", member.ID, member.Name, formatUserBorrowedBooks(member.BorrowedBooks))
		fmt.Println("-----------------------------------------------------------------------------------")

	}
}

// Done
func ListBooks(l service.Library) {
	fmt.Printf("%v %v \n", "", "*********************************************")
	fmt.Printf("%v %v \n", "", "Books")
	fmt.Printf("%v %v \n", "", "*********************************************")
	fmt.Printf("\t%v \t%-20v \t%v \t%v\n", "ID", "Title", "Author", "Status")
	fmt.Println("-----------------------------------------------------------------------------------")
	for _, book := range l.ListBooks() {
		fmt.Printf("\t%v \t%-25v \t%-20v \t%v\n", book.ID, book.Title, book.Author, book.Status)
		fmt.Println("-----------------------------------------------------------------------------------")

	}
}

// Done
func AddBookInput(l *service.Library, reader *bufio.Reader) {
	var author string
	var title string

	fmt.Println("you choosed to add A book pls fill the following")

	fmt.Print("Title: ")

	for {
		title, _ = reader.ReadString('\n')
		title = strings.TrimSpace(title)
		if title == "" {
			fmt.Println("Title must not be empty ")
			fmt.Print("Title: ")
		} else {
			break
		}
	}
	fmt.Print("Author: ")
	for {
		author, _ = reader.ReadString('\n')
		author = strings.TrimSpace(author)
		if author == "" {
			fmt.Println("Author must not be empty ")
			fmt.Print("Author: ")
		} else {
			break
		}
	}

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

// Done
func RemoveBookInput(l *service.Library, reader *bufio.Reader) {
	fmt.Println("you choosed to Remove a Book pls fill the following")
	fmt.Print("Book ID: ")
	BookId, _ := reader.ReadString('\n')
	BookId = strings.TrimSpace(BookId)
	bookId, _ := strconv.Atoi(BookId)
	l.RemoveBook(bookId)

}

// Done
func ReturnBookInput(l *service.Library, reader *bufio.Reader) {
	fmt.Println("you choosed to Return a Book pls fill the following")

	fmt.Print("Member ID: ")
	MemberId := ""
	for {
		MemberId, _ = reader.ReadString('\n')
		if isInteger(MemberId) {
			break
		}
		fmt.Println("Member ID Must be a Number")
		fmt.Print("Member ID: ")

	}
	MemberId = strings.TrimSpace(MemberId)
	memberID, _ := strconv.Atoi(MemberId)

	fmt.Print("Book ID: ")
	BookId := ""
	for {
		BookId, _ = reader.ReadString('\n')
		if isInteger(BookId) {
			break
		}
		fmt.Println("Book ID Must be a Number")
		fmt.Print("Member ID: ")

	}

	BookId = strings.TrimSpace(BookId)
	bookId, _ := strconv.Atoi(BookId)
	s := l.ReturnBook(bookId, memberID)
	if s == nil {
		fmt.Println("you successfuly Returned a Book")
	} else {
		fmt.Println(s)
	}
}

// Done
func BorrowBookInput(l *service.Library, reader *bufio.Reader) {
	fmt.Println("you choosed to Borrow a Book pls fill the following")

	fmt.Print("Member ID: ")
	MemberId := ""
	for {
		MemberId, _ = reader.ReadString('\n')
		if isInteger(MemberId) {
			break
		}
		fmt.Println("Member ID Must be a Number")
		fmt.Print("Member ID: ")

	}
	MemberId = strings.TrimSpace(MemberId)
	memberID, _ := strconv.Atoi(MemberId)

	fmt.Print("Book ID: ")
	BookId := ""
	for {
		BookId, _ = reader.ReadString('\n')
		if isInteger(BookId) {
			break
		}
		fmt.Println("Book ID Must be a Number")
		fmt.Print("Member ID: ")

	}

	BookId = strings.TrimSpace(BookId)
	bookId, _ := strconv.Atoi(BookId)

	s := l.BorrowBook(bookId, memberID)
	// fmt.Printf("%T %v %T %v", bookId, bookId, memberID, memberID)
	if s == nil {
		fmt.Println("you successfuly Borrowed a Book")
	} else {
		fmt.Println(s)
	}
}

// Done
func ListAvailableBooksInput(l *service.Library, reader *bufio.Reader) {
	fmt.Println("you choosed to see list of Available pls fill the following")
	books := l.ListAvailableBooks()
	fmt.Printf("%v %v \n", "", "*********************************************")
	fmt.Printf("%v %v %v\n", "", "Available Books")
	fmt.Printf("%v %v \n", "", "*********************************************")
	fmt.Printf("\t%v \t%-20v \t%-20v \t%v\n", "ID", "Title", "Author", "Status")
	fmt.Println("-----------------------------------------------------------------------------------")
	for _, book := range books {
		fmt.Printf("\t%v \t%-20v \t%-20v \t%v\n", book.ID, book.Title, book.Author, book.Status)
		fmt.Println("-----------------------------------------------------------------------------------")
	}
}

// Done
func ListBorrowedBooksInput(l *service.Library, reader *bufio.Reader) {
	fmt.Println("you choosed to see LIst of Borrowed Books pls fill the following")
	fmt.Print("Member ID: ")
	MemberId := ""
	for {
		MemberId, _ = reader.ReadString('\n')
		if isInteger(MemberId) {
			break
		}
		fmt.Println("Member ID Must be a Number")
		fmt.Print("Member ID: ")

	}
	MemberId = strings.TrimSpace(MemberId)
	memberID, _ := strconv.Atoi(MemberId)

	books := l.ListBorrowedBooks(memberID)

	fmt.Printf("%v %v \n", "", "*********************************************")
	fmt.Printf("%v %v %v\n", "", "Books Borrowed By ", l.Meambers[memberID].Name)
	fmt.Printf("%v %v \n", "", "*********************************************")
	fmt.Printf("\t%v \t%-20v \t%-20v \t%v\n", "ID", "Title", "Author", "Status")
	fmt.Println("-----------------------------------------------------------------------------------")
	for _, book := range books {
		fmt.Printf("\t%v \t%-20v \t%-20v \t%v\n", book.ID, book.Title, book.Author, book.Status)
		fmt.Println("-----------------------------------------------------------------------------------")

	}
}

func formatUserBorrowedBooks(books []models.Book) string {
	f := ""
	for _, val := range books {
		f += val.Title + " by " + val.Author + " | "
	}
	return f
}

func isInteger(s string) bool {
	s = strings.TrimSpace(s)
	_, err := strconv.Atoi(s)
	return err == nil
}

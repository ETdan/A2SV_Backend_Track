package service

import (
	"BookKeeper/models"
	"errors"
	"fmt"
	"slices"
)

type Library struct {
	Curr_book_ID   int
	Curr_member_ID int
	Books          map[int]models.Book
	Meambers       map[int]models.Member
}

func (l *Library) AddBook(book models.Book) {
	l.Books[book.ID] = book
}
func (l *Library) AddMeamber(meamber models.Member) {
	l.Meambers[meamber.ID] = meamber
}
func (l *Library) RemoveBook(bookID int) {
	_, ok := l.Books[bookID]
	if ok {
		delete(l.Books, bookID)
		fmt.Println("you successfuly Removed a Book")
	} else {
		fmt.Println("No Book with this ID")
	}
}

func (l *Library) BorrowBook(bookID int, memberID int) error {
	_, okb := l.Books[bookID]
	_, okm := l.Meambers[memberID]
	if !okb || !okm || l.Books[bookID].Status == "Borrowed" {
		if !okb {
			return errors.New("Book ID Does Not Exist")
		} else if !okm {
			return errors.New("Member ID Does Not Exist")
		} else {
			return errors.New("Book is Borrowed")
		}
	} else {
		b := l.Books[bookID]
		// fmt.Println("found book", b)
		b.Status = "Borrowed"
		l.Books[bookID] = b
		// fmt.Println("after book edit", b)

		m := l.Meambers[memberID]
		m.BorrowedBooks = append(m.BorrowedBooks, b)
		l.Meambers[memberID] = m

		// fmt.Println("Borrowed Books successfully")
		return nil
	}
}
func (l *Library) ReturnBook(bookID int, memberID int) error {
	_, okb := l.Books[bookID]
	_, okm := l.Meambers[memberID]
	if !okb || !okm || l.Books[bookID].Status == "Available" {
		if !okb {
			return errors.New("Book ID Does Not Exist")
		} else if !okm {
			return errors.New("Member ID Does Not Exist")
		} else {
			return errors.New("Book is is NOT Borrowed")
		}
	}

	if slices.Contains(l.Meambers[memberID].BorrowedBooks, l.Books[bookID]) {
		m := l.Meambers[memberID]
		b := l.Books[bookID]
		m.BorrowedBooks = RemoveBookFromMemberBorrowedList(m.BorrowedBooks, b)

		b.Status = "Available"
		l.Meambers[memberID] = m
		l.Books[bookID] = b
		return nil
	}
	return errors.New("Book NOt Borrowed by  Member")
}
func (l *Library) ListAvailableBooks() []models.Book {
	bookList := []models.Book{}
	for _, val := range l.Books {
		if val.Status == "Available" {
			bookList = append(bookList, val)
		}
	}
	return bookList
}
func (l *Library) ListBorrowedBooks(memberID int) []models.Book {
	bookList := []models.Book{}
	m := l.Meambers[memberID]
	for _, val := range m.BorrowedBooks {
		if val.Status == "Borrowed" {
			bookList = append(bookList, val)
		}
	}
	return bookList
}
func (l *Library) ListMembers() map[int]models.Member {
	return l.Meambers
}
func (l *Library) ListBooks() map[int]models.Book {
	return l.Books
}

func RemoveBookFromMemberBorrowedList(bookList []models.Book, book models.Book) []models.Book {
	newBook := []models.Book{}
	for _, val := range bookList {
		if val == book {
			continue
		}
		newBook = append(newBook, val)
	}
	return newBook
}

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
	delete(l.Books, bookID)
}

func (l *Library) BorrowBook(bookID int, memberID int) error {
	_, okb := l.Books[bookID]
	_, okm := l.Meambers[memberID]
	if !okb || !okm || l.Books[bookID].Status == "Borrowed" {
		return errors.New("models.Book is Borrowed")
	} else {
		b := l.Books[bookID]
		fmt.Println("found book", b)
		b.Status = "Borrowed"
		l.Books[bookID] = b
		fmt.Println("after book edit", b)

		m := l.Meambers[memberID]
		m.BorrowedBooks = append(m.BorrowedBooks, b)
		l.Meambers[memberID] = m
		fmt.Println("Borrowed Books", m)

		return nil
	}
}
func (l *Library) ReturnBook(bookID int, memberID int) error {
	_, okb := l.Books[bookID]
	_, okm := l.Meambers[memberID]

	if okb || okm || l.Books[bookID].Status == "Borrowed" {
		if slices.Contains(l.Meambers[memberID].BorrowedBooks, l.Books[bookID]) {
			m := l.Meambers[memberID]
			b := l.Books[bookID]
			m.BorrowedBooks = RemoveBookFromMemberBorrowedList(m.BorrowedBooks, b)
			b.Status = "Available"
			l.Meambers[memberID] = m
			l.Books[bookID] = b
		}
		return nil
	}
	return errors.New("Book or Member not Found")
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

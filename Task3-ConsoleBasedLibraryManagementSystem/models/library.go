package models

import (
	"errors"
	"slices"
)

type Library struct {
	books    map[int]Book
	meambers map[int]Member
}

func (l *Library) AddBook(book Book) {
	l.books[book.ID] = book
}
func (l *Library) RemoveBook(bookID int) {
	delete(l.books, bookID)
}

func (l *Library) BorrowBook(bookID int, memberID int) error {
	if l.books[bookID].Status == "Borrowed" {
		return errors.New("Book is Borrowed")
	} else {
		b := l.books[bookID]
		b.Status = "Borrowed"
		m := l.meambers[memberID]
		m.BorrowedBooks = append(m.BorrowedBooks, b)
		return nil
	}
}
func (l *Library) ReturnBook(bookID int, memberID int) error {
	if l.books[bookID].Status == "Borrowed" {
		if slices.Contains(l.meambers[memberID].BorrowedBooks, l.books[bookID]) {
			m := l.meambers[memberID]
			b := l.books[bookID]
			m.BorrowedBooks = RemoveBookFromMemberBorrowedList(m.BorrowedBooks, b)
			b.Status = ""
		}
	}
	return nil
}
func (l *Library) ListAvailableBooks() []Book {
	bookList := []Book{}
	for _, val := range l.books {
		if val.Status == "Available" {
			bookList = append(bookList, val)
		}
	}
	return bookList
}
func (l *Library) ListBorrowedBooks(memberID int) []Book {
	bookList := []Book{}
	for _, val := range l.books {
		if val.Status == "Borrowed" {
			bookList = append(bookList, val)
		}
	}
	return bookList
}

func RemoveBookFromMemberBorrowedList(bookList []Book, book Book) []Book {
	newBook := []Book{}
	for _, val := range bookList {
		if val == book {
			continue
		}
		newBook = append(newBook, val)
	}
	return newBook
}

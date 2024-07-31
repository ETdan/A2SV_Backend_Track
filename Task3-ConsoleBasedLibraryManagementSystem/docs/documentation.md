# 📚 BookKeeper Library Management System Documentation 📚

Welcome to the **BookKeeper** Library Management System! This console-based app, written in Go, is your go-to solution for managing a library's books and members. 📖👩‍💻

## 🏛️ Overview

**BookKeeper** allows you to:

- Add and remove books 📚
- Add and manage members 🧑‍🤝‍🧑
- Borrow and return books 🔄
- List available and borrowed books 📋
- View library statistics 📊

## 📂 Structure

### 📜 Controllers

The `controllers` package manages user interactions and handles the main menu.

#### Functions

- **`Choose()`**: Displays the main menu and handles user choices. 📝
- **`AddBookInput(l *service.Library, reader *bufio.Reader)`**: Adds a new book to the library. 📘
- **`AddMemberInput(l *service.Library, reader *bufio.Reader)`**: Adds a new member. 🧑‍🤝‍🧑
- **`RemoveBookInput(l *service.Library, reader *bufio.Reader)`**: Removes a book from the library. ❌📚
- **`ReturnBookInput(l *service.Library, reader *bufio.Reader)`**: Returns a borrowed book. 🔄📖
- **`BorrowBookInput(l *service.Library, reader *bufio.Reader)`**: Borrows a book for a member. 📚➡️🧑‍🤝‍🧑
- **`ListAvailableBooksInput(l *service.Library, reader *bufio.Reader)`**: Lists all available books. 📜📘
- **`ListBorrowedBooksInput(l *service.Library, reader *bufio.Reader)`**: Lists books borrowed by a specific member. 📋🧑‍🤝‍🧑
- **`ListBooks(l service.Library)`**: Lists all books in the library. 📚
- **`ListMembers(l service.Library)`**: Lists all members of the library. 🧑‍🤝‍🧑
- **`ListStat(l service.Library)`**: Displays the number of books and members. 📊

#### Utility Functions

- **`formatUserBorrowedBooks(books []models.Book) string`**: Formats borrowed books for display. 📚✨
- **`isInteger(s string) bool`**: Checks if a string is an integer. 🔢

### 🛠️ Service

The `service` package is the backbone of the library, handling core functionalities and data management.

#### Types

- **`Library`**: Represents the library with books and members. 🏛️

#### Methods

- **`AddBook(book models.Book)`**: Adds a book to the library. 📚
- **`AddMeamber(meamber models.Member)`**: Adds a member to the library. 🧑‍🤝‍🧑
- **`RemoveBook(bookID int)`**: Removes a book from the library. ❌📚
- **`BorrowBook(bookID int, memberID int) error`**: Allows a member to borrow a book. 📚➡️🧑‍🤝‍🧑
- **`ReturnBook(bookID int, memberID int) error`**: Allows a member to return a book. 🔄📖
- **`ListAvailableBooks() []models.Book`**: Lists all available books. 📜📘
- **`ListBorrowedBooks(memberID int) []models.Book`**: Lists all books borrowed by a member. 📋🧑‍🤝‍🧑
- **`ListMembers() map[int]models.Member`**: Lists all members. 🧑‍🤝‍🧑
- **`ListBooks() map[int]models.Book`**: Lists all books. 📚

#### Utility Functions

- **`RemoveBookFromMemberBorrowedList(bookList []models.Book, book models.Book) []models.Book`**: Removes a book from a member's borrowed list. ❌📚

## 🚀 Usage

1. **Run the Application**: Start the main function to launch the app. 🚀
2. **Select an Option**: Choose from the menu to perform different actions. 📝
3. **Follow Prompts**: Provide the necessary input as requested by the app. ✍️

## 📚 Example

1. Start the application. 🏁
2. Choose `1` to add a book. 📚
3. Enter the book title and author. ✍️
4. The book is added to the library. 🎉

## ⚠️ Notes

- Ensure the `models` package includes the `Book` and `Member` data structures. 📦
- The `service` package manages the state and operations of the library. 🔧

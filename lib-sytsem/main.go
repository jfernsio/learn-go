package main

import (
	"fmt"
)

type Book struct {
	Id        int
	Title     string
	Available bool
}

func listBooks(books []Book) {
	for _, b := range books {
		fmt.Printf("ID: %d, Title: %s, Available: %t\n", b.Id, b.Title, b.Available)
	}
}

func updateAvailability(book *Book) error {
	if book.Available {
		book.Available = false
		return nil
	}
	return fmt.Errorf("book is not available")
}

func findBook(id int, books []Book) (*Book, error) {
	for i, b := range books {
		if b.Id == id {
			fmt.Println(&books[i])
			return &books[i], nil
		}
	}
	return nil, fmt.Errorf("book not found")
}

func brrowBook (user string,id int, books []Book, userBooks map[string][]int) error{
	b,err := findBook(id,books) 
	if err != nil {
		return err
	}else {
		bID := b.Id
		if bID != id {
			return fmt.Errorf("book ID mismatch")
		}
		
		err := updateAvailability(b)
		if err != nil {
			return err
		}
		userBooks[user] = append(userBooks[user], id)
		for user,booksId := range userBooks {
			fmt.Printf("User: %s, Borrowed Book IDs: %v\n", user, booksId)
		}
	}
		
	
	return nil
}

func returnBook(user string, id int, books []Book, userBooks map[string][]int) error {
	b, err := findBook(id, books)
	if err != nil {
		return err
	}
	if !b.Available {
		b.Available = true
	} else {
		return fmt.Errorf("book is already available")
	}
	//check if book id is in userBooks
	bookIds, exists := userBooks[user]
	if !exists {
		return fmt.Errorf("user has not borrowed any books")
	}
	//remove book id from userBooks
	for i, bookId := range bookIds {
		if bookId == id {
			userBooks[user] = append(bookIds[:i], bookIds[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("user has not borrowed this book")
}
func main() {
	//create a slice of book struct
	var books []Book

	//add books to the slice
	books = append(books, Book{Id: 1, Title: "The Great Gatsby", Available: true})
	books = append(books, Book{Id: 2, Title: "1984", Available: false})
	books = append(books, Book{Id: 3, Title: "To Kill a Mockingbird", Available: true})
	books = append(books, Book{Id: 4, Title: "Pride and Prejudice", Available: true})

	fmt.Println("Library Mangaement System")
	fmt.Println("Library Books List !")
	fmt.Println("-----------------------")
	listBooks(books)

	// a map w key = username, value = slice of borrowed book IDs.
	userBooks := make(map[string][]int)

	// userBooks["alice"] = []int{1, 2}
	// userBooks["bob"] = []int{3}

	err := brrowBook("alice", 1, books, userBooks)
	if err != nil {
		fmt.Println("Error borrowing book:", err)
	} else {
		fmt.Println("\nAfter Borrowing Book:")
		listBooks(books)
	}
	err = returnBook("k", 2, books, userBooks)
	if err != nil {
		fmt.Println("Error returning book:", err)
	} else {
		fmt.Println("\nAfter Returning Book:")
		listBooks(books)
	}
	// err := updateAvailability(&books[2])
	// if err != nil {
	// 	fmt.Println("Error updating book availability:", err)
	// } else {
	// 	fmt.Println("\nUpdated Book Availability:")
	// 	listBooks(books)
	// }
	// b, er := findBook(2, books) //-> must throw err
	// if er != nil {
	// 	fmt.Println("Error finding book:", er)
	// } else {
	// 	fmt.Println("Found Book:")
	// 	fmt.Printf("ID: %d, Title: %s, Available: %t\n", b.Id, b.Title, b.Available)
	// }

	// fmt.Println("\nUser Borrowed Books:")
	// for user, bookIDs := range userBooks {
	// 	fmt.Printf("User: %s, Borrowed Book IDs: %v\n", user, bookIDs)
	// }
}

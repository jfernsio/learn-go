package main

import ( "fmt")

type Book struct {
	Id int
	Title string
	Available bool
}

func listBooks(books []Book)  {
	for _, b := range books {
		fmt.Printf("ID: %d, Title: %s, Available: %t\n", b.Id, b.Title, b.Available)
	}
}

func updateAvailability (book *Book) error {
	if book.Available {
		book.Available = false
		return nil
	}
	return fmt.Errorf("book is not available")
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
	

	err := updateAvailability(&books[2])
	if err != nil {
		fmt.Println("Error updating book availability:", err)
	}else  {
		fmt.Println("\nUpdated Book Availability:")
		listBooks(books)
	}
	
}	
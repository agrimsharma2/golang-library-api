package models

import (
	"sync"
)

type book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

var (
	books   []book
	booksMu sync.Mutex
)

func Init() {
	// Initialize your books or database connections here
	books = []book{
		{ID: "1", Title: "Atomic Habits", Author: "James Clear", Price: 15.72},
		{ID: "2", Title: "The Tattooist of Auschwitz", Author: "Heather Morris", Price: 19.25},
		{ID: "3", Title: "Harry Potter and the Chamber of Secrets", Author: "JK Rowling", Price: 10.00},
	}
}

// Implement functions like GetBooks, GetBookByID, CreateBook, UpdateBook, and DeleteBook here

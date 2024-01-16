package main

import (
	"time"
)

// Book is Interface for book details.
type Book struct {
	Book_id        string `json:"book_id"`
	Title          string `json:"title"`
	Author         string `json:"author"`
	Publisher      string `json:"publisher"`
	Year_published int    `json:"year_published"`
	Genre          string `json:"genre"`
	Is_loaned      bool   `json:"is_loaned"`
}

// Customers is Interface for customer details.
type Customer struct {
	Customer_id string `json:"customer_id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
}

// Loans is Interface for load details.
type Loan struct {
	Loan_id     string    `json:"loan_id"`
	Customer_id string    `json:"customer_id"`
	Book_id     string    `json:"book_id"`
	Issued_date time.Time `json:"issued_date"`
	Due_date    time.Time `json:"due_date"`
}

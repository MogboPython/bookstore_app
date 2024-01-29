package main

// Book is Interface for book details.
type Book struct {
	Book_id        int    `json:"book_id"`
	Title          string `json:"title"`
	Author         string `json:"author"`
	Publisher      string `json:"publisher"`
	Year_published int    `json:"year_published"`
	Genre          string `json:"genre"`
}

// Customers is Interface for customer details.
type Customer struct {
	Customer_id int    `json:"customer_id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
}

// Loans is Interface for load details.
type Loan struct {
	Loan_id     int    `json:"loan_id"`
	Customer_id int    `json:"customer_id"`
	Book_id     int    `json:"book_id"`
	Issued_date string `json:"issued_date"`
	Time_frame  int    `json:"time_frame"`
	Due_date    string `json:"due_date"`
}

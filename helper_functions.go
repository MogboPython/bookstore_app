package main

import (
	"database/sql"
	"log"
)

func CreateBook(db *sql.DB, title, author, publisher, genre string, year int) error {
	query := "INSERT INTO books (title, author, publisher, genre, year_published) VALUES (?, ?, ?, ?, ?)"
	_, err := db.Exec(query, title, author, publisher, genre, year)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func GetBooks(db *sql.DB) ([]Book, error) {
	var (
		books []Book
		book  Book
	)

	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(&book.Book_id, &book.Title, &book.Author, &book.Publisher, &book.Genre, &book.Year_published)
		if err != nil {
			log.Println(err)
			continue
		}
		books = append(books, book)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	defer rows.Close()
	return books, nil

}

func GetBook(db *sql.DB, id int) (*Book, error) {
	query := "SELECT * FROM books WHERE book_id = ?"
	row := db.QueryRow(query, id)

	book := &Book{}
	err := row.Scan(&book.Book_id, &book.Title, &book.Author, &book.Publisher, &book.Genre, &book.Year_published)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return book, nil
}

func DeleteBook(db *sql.DB, id int) error {
	query := "DELETE FROM books WHERE book_id = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

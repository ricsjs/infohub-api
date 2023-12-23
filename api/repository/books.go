package repository

import (
	"log"

	"github.com/ricsjs/infohub-api/api/db"
	"github.com/ricsjs/infohub-api/api/models"
)

func GetAllBooks() ([]models.Book, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	rows, err := conn.Query("SELECT * FROM BOOK")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Gender, &book.Publication, &book.Resume); err != nil {
			log.Fatal(err)
		}
		books = append(books, book)
	}

	return books, nil
}

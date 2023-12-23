package repository

import (
	"log"

	"github.com/ricsjs/infohub-api/api/db"
	"github.com/ricsjs/infohub-api/api/models"
)

func GetAllMovies() ([]models.Movie, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	rows, err := conn.Query("SELECT * FROM movie")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var movies []models.Movie
	for rows.Next() {
		var movie models.Movie
		if err := rows.Scan(&movie.ID, &movie.Title, &movie.Director, &movie.Gender, &movie.Launch, &movie.Synopsis); err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}

	return movies, nil
}

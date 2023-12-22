package repository

import (
	"log"

	"github.com/ricsjs/infohub-api/api/db"
	"github.com/ricsjs/infohub-api/api/models"
)

func GetAllBeers() ([]models.Beer, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	rows, err := conn.Query("SELECT * FROM BEER")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var beers []models.Beer
	for rows.Next() {
		var beer models.Beer
		if err := rows.Scan(&beer.ID, &beer.Name, &beer.Style, &beer.Alcohol, &beer.Descrition, &beer.Origin); err != nil {
			log.Fatal(err)
		}
		beers = append(beers, beer)
	}
	return beers, nil
}

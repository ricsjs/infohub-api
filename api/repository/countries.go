package repository

import (
	"log"

	"github.com/ricsjs/infohub-api/api/db"
	"github.com/ricsjs/infohub-api/api/models"
)

func GetAllCountries() ([]models.Country, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	rows, err := conn.Query("SELECT * FROM country")
	if err != nil {
		log.Fatal(err)
	}

	var countries []models.Country
	for rows.Next() {
		var country models.Country
		if err := rows.Scan(&country.ID, &country.Name, &country.Capital, &country.Population, &country.Area, &country.Language); err != nil {
			log.Fatal(err)
		}
		countries = append(countries, country)
	}

	return countries, nil
}

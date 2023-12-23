package repository

import (
	"log"

	"github.com/ricsjs/infohub-api/api/db"
	"github.com/ricsjs/infohub-api/api/models"
)

func GetAllCoffee() ([]models.Coffee, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	rows, err := conn.Query("SELECT * FROM COFFEE")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var coffees []models.Coffee
	for rows.Next() {
		var coffee models.Coffee
		if err := rows.Scan(&coffee.ID, &coffee.Name, &coffee.Origin, &coffee.Roasting, &coffee.Descrition, &coffee.Price); err != nil {
			log.Fatal(err)
		}
		coffees = append(coffees, coffee)
	}

	return coffees, nil
}

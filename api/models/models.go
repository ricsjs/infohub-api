package models

type Beer struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Style      string `json:"style"`
	Alcohol    string `json:"alcohol"`
	Descrition string `json:"description"`
	Origin     string `json:"origin"`
}

type Book struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Gender      string `json:"gender"`
	Publication int    `json:"publication"`
	Resume      string `json:"resume"`
}

type Coffee struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Origin     string  `json:"origin"`
	Roasting   string  `json:"roasting"`
	Descrition string  `json:"description"`
	Price      float32 `json:"price"`
}

type Country struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Capital    string `json:"capital"`
	Population string `json:"population"`
	Area       string `json:"area"`
	Language   string `json:"language"`
}

type Movie struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Director string `json:"director"`
	Gender   string `json:"gender"`
	Launch   int    `json:"launch"`
	Synopsis string `json:"synopsis"`
}

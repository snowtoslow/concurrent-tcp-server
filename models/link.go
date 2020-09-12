package models

type Link struct {
	Route1 string `json:"route1"`
	Route2 string `json:"route2"`
	Route3 string `json:"route3,omitempty"`
	Route4 string `json:"route4,omitempty"`
}

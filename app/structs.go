package app

import (
	"time"
)

type Place struct {
	Id       string
	Country  string
	City     string
	Street   string
	Building int
}

type Person struct {
	Id        string
	FirstName string
	LastName  string
	Position  string
}

type Event struct {
	Id        string
	Name      string
	DateTime  time.Time
	Place     *Place
	Members   []*Person
	CreatedAt time.Time
}

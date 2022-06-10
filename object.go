package sccrawler

type Reservation struct{
	Id		string
	Name	string
	Tel		string
	Place	string
	Date	string
}

func NewObject() Reservation{
	return Reservation{}
}
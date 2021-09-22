package models

type Error struct {
	Status  int
	Message string
}

type Success struct {
	Token  string
	Status int
	Url    string
}

package models

type Item struct {
	ID    int
	Name  string
	Count int
}

type Soda struct {
	Item
}

type Coffe struct {
	Item
}

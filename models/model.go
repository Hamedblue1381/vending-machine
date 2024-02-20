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

type VendingMachineRequest struct {
	Count int `json:"count"`
}
type VendingMachineResponse struct {
	Id    int    `json:"id"`
	Stock int    `json:"stock"`
	Coins int    `json:"coins"`
	Items []Item `json:"items"`
}

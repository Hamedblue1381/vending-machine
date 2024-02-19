package state

import (
	"fmt"
)

type NoStockState struct {
	vendingMachine *VendingMachine
}

func (n *NoStockState) InsertCoin() error {
	return fmt.Errorf("items out of stock")
}

func (n *NoStockState) ChooseItem(int) error {
	return fmt.Errorf("items out of stock")
}

func (n *NoStockState) Dispense(id int) error {
	return fmt.Errorf("items out of stock")

}

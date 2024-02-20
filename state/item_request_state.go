package state

import (
	"fmt"
)

type ItemRequestState struct {
	vendingMachine *VendingMachine
}

func (i *ItemRequestState) InsertCoin() error {
	return fmt.Errorf("can't insert coin now")
}

func (i *ItemRequestState) ChooseItem(id int) error {
	itemExists := false
	for _, item := range i.vendingMachine.Items {
		if item.ID == id {
			itemExists = true
			break
		}
	}

	if !itemExists {
		return fmt.Errorf("item with ID %d does not exist", id)
	}

	if i.vendingMachine.Stock > 0 {
		i.vendingMachine.SetState(i.vendingMachine.Dispensing)
		return nil
	} else {
		return fmt.Errorf("out of stock")
	}

}

func (i *ItemRequestState) Dispense(id int) error {
	return fmt.Errorf("requesting")

}

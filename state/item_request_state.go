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

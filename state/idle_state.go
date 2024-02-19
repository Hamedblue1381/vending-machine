package state

import (
	"fmt"
)

type IdleState struct {
	vendingMachine *VendingMachine
}

func (i *IdleState) InsertCoin() error {
	i.vendingMachine.SetState(i.vendingMachine.ItemRequesting)
	i.vendingMachine.Coins++
	return nil
}

func (i *IdleState) ChooseItem(int) error {
	return fmt.Errorf("idle")
}

func (i *IdleState) Dispense(id int) error {
	return fmt.Errorf("idle")

}

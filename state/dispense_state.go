package state

import (
	"fmt"
)

type DispensingState struct {
	vendingMachine *VendingMachine
}

func (d *DispensingState) InsertCoin() error {
	return fmt.Errorf("dispensing")
}

func (d *DispensingState) ChooseItem(int) error {
	return fmt.Errorf("dispensing")
}

func (d *DispensingState) Dispense(id int) error {
	err := d.vendingMachine.DecreaseCountById(id)
	if err != nil {
		d.vendingMachine.SetState(d.vendingMachine.ItemRequesting)
		return err
	}

	d.vendingMachine.Stock = d.vendingMachine.Stock - 1
	if d.vendingMachine.Stock == 0 {
		d.vendingMachine.SetState(d.vendingMachine.NoStock)
	} else {
		d.vendingMachine.SetState(d.vendingMachine.Idle)
	}

	return nil
}

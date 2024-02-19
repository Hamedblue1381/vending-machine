package state

import (
	"fmt"

	"github.com/hamedblue1381/vending-machine/models"
)

type VendingMachine struct {
	Idle           State
	ItemRequesting State
	NoStock        State
	Dispensing     State

	CurrentState State
	Stock        int
	Coins        int
	Items        []models.Item
}

// TODO: optimize
func (v *VendingMachine) DecreaseCountById(id int) error {
	for i, item := range v.Items {
		if item.ID == id {
			if v.Items[i].Count > 0 {
				v.Items[i].Count--
				return nil
			} else {
				return fmt.Errorf("item is out of stock")
			}
		}
	}
	return fmt.Errorf("item with ID %d not found", id)
}

func NewVendingMachine(itemCount int) *VendingMachine {
	cola := models.Soda{
		Item: models.Item{
			ID: 1, Name: "Soda", Count: itemCount,
		}}

	coffee := models.Coffe{
		Item: models.Item{
			ID: 2, Name: "Coffee", Count: itemCount,
		}}

	v := &VendingMachine{
		Coins: 0,
		Items: []models.Item{cola.Item, coffee.Item},
	}
	noStockState := &NoStockState{
		vendingMachine: v,
	}
	itemRequestState := &ItemRequestState{
		vendingMachine: v,
	}
	IdleState := &IdleState{
		vendingMachine: v,
	}
	dispensingState := &DispensingState{
		vendingMachine: v,
	}
	v.SetState(IdleState)
	v.Idle = IdleState
	v.NoStock = noStockState
	v.ItemRequesting = itemRequestState
	v.Dispensing = dispensingState
	for _, item := range v.Items {
		v.Stock += item.Count
	}

	return v
}

func (v *VendingMachine) InsertCoin() error {
	return v.CurrentState.InsertCoin()
}
func (v *VendingMachine) ChooseItem(id int) error {
	return v.CurrentState.ChooseItem(id)
}
func (v *VendingMachine) Dispense(id int) error {
	return v.CurrentState.Dispense(id)
}

func (v *VendingMachine) SetState(s State) {
	v.CurrentState = s
}

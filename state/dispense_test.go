package state

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDispenseItemIdle(t *testing.T) {
	v, _ := MockData()
	err := v.Dispense(v.Items[0].ID)
	if err == nil {
		t.Error("expected error, but got nil")
	}
	assert.EqualError(t, err, "idle")
}
func TestDispenseItemItemRequesting(t *testing.T) {
	v, _ := MockData()
	v.SetState(v.ItemRequesting)
	err := v.Dispense(v.Items[0].ID)
	if err == nil {
		t.Error("expected error, but got nil")
	}
	assert.EqualError(t, err, "requesting")
}
func TestDispenseNoStock(t *testing.T) {
	v, _ := MockData()
	v.SetState(v.NoStock)
	err := v.Dispense(v.Items[0].ID)
	if err == nil {
		t.Error("expected error, but got nil")
	}
	assert.EqualError(t, err, "items out of stock")
}
func TestDispenseItemDispensing(t *testing.T) {
	v, _ := MockData()
	if err := v.InsertCoin(); err != nil {
		t.Error(err)
	}
	if err := v.ChooseItem(v.Items[0].ID); err != nil {
		t.Error(err)
	}
	err := v.Dispense(v.Items[0].ID)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, v.CurrentState, v.Idle)
}
func TestDispenseItemDispensingInvalidId(t *testing.T) {
	v, _ := MockData()
	startingStock := v.Stock
	if err := v.InsertCoin(); err != nil {
		t.Error(err)
	}
	if err := v.ChooseItem(v.Items[0].ID); err != nil {
		t.Error(err)
	}
	err := v.Dispense(999)
	if err == nil {
		t.Error("expected error, but got nil")
	}
	assert.Equal(t, v.CurrentState, v.ItemRequesting)
	assert.Equal(t, v.Stock, startingStock)
}

func TestDispenseItemDispensingStockEmpty(t *testing.T) {
	v, _ := MockData()
	v.Stock = 1
	if err := v.InsertCoin(); err != nil {
		t.Error(err)
	}
	if err := v.ChooseItem(v.Items[0].ID); err != nil {
		t.Error(err)
	}
	err := v.Dispense(v.Items[0].ID)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, v.CurrentState, v.NoStock)

}

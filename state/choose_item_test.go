package state

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChooseItemIdle(t *testing.T) {
	v, _ := MockData()
	err := v.ChooseItem(v.Items[0].ID)
	if err == nil {
		t.Error("expected error, but got nil")
	}
	assert.EqualError(t, err, "idle")
}
func TestChooseItemItemRequesting(t *testing.T) {
	v, _ := MockData()
	if err := v.InsertCoin(); err != nil {
		t.Error(err)
	}
	err := v.ChooseItem(v.Items[0].ID)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, v.CurrentState, v.Dispensing)
}
func TestChooseNoStock(t *testing.T) {
	v, _ := MockData()
	v.SetState(v.NoStock)
	err := v.ChooseItem(v.Items[0].ID)
	if err == nil {
		t.Error("expected error, but got nil")
	}
	assert.EqualError(t, err, "items out of stock")
}
func TestChooseItemDispensing(t *testing.T) {
	v, _ := MockData()
	v.SetState(v.Dispensing)
	err := v.ChooseItem(v.Items[0].ID)
	if err == nil {
		t.Error("expected error, but got nil")
	}
	assert.EqualError(t, err, "dispensing")
}

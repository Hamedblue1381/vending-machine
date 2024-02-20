package state

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertCoinIdle(t *testing.T) {
	v, _ := MockData()
	startingCoin := v.Coins
	if err := v.InsertCoin(); err != nil {
		t.Error(err)
	}
	assert.NotEqual(t, v.Coins, startingCoin)
	assert.Equal(t, v.Coins, startingCoin+1)
	assert.Equal(t, v.CurrentState, v.ItemRequesting)

}
func TestInsertCoinItemRequesting(t *testing.T) {
	v, _ := MockData()
	v.SetState(v.ItemRequesting)
	err := v.InsertCoin()
	if err == nil {
		t.Error("expected error, but got nil")
	}
	assert.EqualError(t, err, "can't insert coin now")
}
func TestInsertCoinNoStock(t *testing.T) {
	v, _ := MockData()
	v.SetState(v.NoStock)
	err := v.InsertCoin()
	if err == nil {
		t.Error("expected error, but got nil")
	}
	assert.EqualError(t, err, "items out of stock")
}
func TestInsertCoinDispensing(t *testing.T) {
	v, _ := MockData()
	v.SetState(v.Dispensing)
	err := v.InsertCoin()
	if err == nil {
		t.Error("expected error, but got nil")
	}
	assert.EqualError(t, err, "dispensing")
}

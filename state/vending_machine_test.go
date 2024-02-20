package state

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeNewVendingMachine(t *testing.T) {
	v, testItemCount := MockData()

	assert.Equal(t, v.CurrentState, v.Idle)
	assert.Equal(t, v.Stock, testItemCount*2)
	assert.Equal(t, v.Coins, 0)
}
func TestSetState(t *testing.T) {
	v, _ := MockData()

	v.SetState(v.ItemRequesting)

	assert.Equal(t, v.CurrentState, v.ItemRequesting)

}
func TestDecreaseCount(t *testing.T) {
	v, _ := MockData()
	testItem := v.Items[0]
	err := v.DecreaseCountById(v.Items[0].ID)
	if err != nil {
		t.Error(err)
	}
	assert.NotEqual(t, testItem.Count, v.Items[0].Count)
}
func TestDecreaseCountOfZeroStock(t *testing.T) {
	v, _ := MockData()
	v.Items[0].Count = 0
	err := v.DecreaseCountById(v.Items[0].ID)
	if err == nil {
		t.Error("expected error, but got nil")
	}

	assert.EqualError(t, err, "item is out of stock")
}

func MockData() (*VendingMachine, int) {
	testItemCount := 20
	v := NewVendingMachine(testItemCount)
	return v, testItemCount
}

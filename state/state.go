package state

type State interface {
	InsertCoin() error
	ChooseItem(int) error
	Dispense(int) error
}

package cli

import "fmt"

func (c *CliService) InsertCoin() error {
	err := c.vendingMachine.InsertCoin()
	if err != nil {
		return err
	}
	c.ShowItems()
	return nil
}

func (c *CliService) ShowItems() error {
	fmt.Printf("Choose from items: \n%+v \n", c.vendingMachine.Items)
	c.ChooseItem()
	return nil
}
func (c *CliService) ChooseItem() error {
	var id int
	fmt.Scanln(&id)
	if err := c.vendingMachine.ChooseItem(id); err != nil {
		return err
	}
	c.DispenseItem(id)
	return nil
}
func (c *CliService) DispenseItem(id int) error {
	if err := c.vendingMachine.Dispense(id); err != nil {
		return err
	}
	fmt.Println(c.vendingMachine.Items)
	fmt.Println("Item dispensed successfuly")
	return nil
}

package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/hamedblue1381/vending-machine/state"
)

type CliServicer interface {
	InsertCoin() error
	ShowItems() error
	ChooseItem() error
	DispenseItem(id int) error
}
type CliService struct {
	CliServicer
	vendingMachine  *state.VendingMachine
	vendingMachines []state.VendingMachine
	updateChan      chan []state.VendingMachine
	mutex           *sync.Mutex
}

func NewCliService(v *state.VendingMachine, vendingMachines []state.VendingMachine, mutex *sync.Mutex, updateChan chan []state.VendingMachine) *CliService {
	return &CliService{
		vendingMachine:  v,
		vendingMachines: vendingMachines,
		mutex:           mutex,
		updateChan:      updateChan,
	}
}

func (c *CliService) Start() {
	fmt.Println("Vending machine started. Type 'coin' to insert a coin.")

	scanner := bufio.NewScanner(os.Stdin)
	go func() {
		for {
			if !scanner.Scan() {
				fmt.Println("Error reading input:", scanner.Err())
				break
			}
			input := strings.ToLower(scanner.Text())
			switch input {
			case "run":
				c.Run()
			case "coin":
				err := c.InsertCoin()
				if err != nil {
					fmt.Println(err)
				}
			default:
				fmt.Println("Unknown command")
			}
		}
	}()
	<-make(chan struct{})

}

func (s *CliService) Run() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	fmt.Println("Set the stock number : ")
	var itemCount int
	fmt.Scanln(&itemCount)

	vmachine := state.NewVendingMachine(itemCount)

	s.vendingMachines = append(s.vendingMachines, *vmachine)

	s.updateChan <- s.vendingMachines

}

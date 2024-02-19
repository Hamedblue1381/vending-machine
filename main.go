package main

import (
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/hamedblue1381/vending-machine/api"
	"github.com/hamedblue1381/vending-machine/cli"
	"github.com/hamedblue1381/vending-machine/state"
)

var (
	vendingMachines []state.VendingMachine
	mutex           sync.Mutex
)

func main() {
	updateChan := make(chan []state.VendingMachine)

	v := state.NewVendingMachine(20)
	h := api.NewVendingMachineHandler(v, vendingMachines, updateChan, &mutex)
	apiService := api.NewApiService()

	cliService := cli.NewCliService(v, vendingMachines, &mutex, updateChan)

	go apiService.Start(h)
	go cliService.Start()
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	<-interrupt

}

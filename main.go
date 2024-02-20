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
	cliiEnabled     string = os.Getenv("CLI_ENABLED")
)

func main() {
	updateChan := make(chan []state.VendingMachine)

	v := state.NewVendingMachine(20)
	h := api.NewVendingMachineHandler(v, vendingMachines, updateChan, &mutex)
	apiService := api.NewApiService()

	if cliiEnabled != "" && cliiEnabled == "true" {
		cliService := cli.NewCliService(v, vendingMachines, &mutex, updateChan)
		go cliService.Start()
	}

	go apiService.Start(h)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	<-interrupt

}

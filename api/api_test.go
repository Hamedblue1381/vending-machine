package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/hamedblue1381/vending-machine/state"
	"github.com/stretchr/testify/assert"
)

func TestInsertHandler(t *testing.T) {
	handler := start()
	req, err := http.NewRequest("GET", "/insert", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler.InsertCoin(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code, "handler returned wrong status code")

}
func TestChooseItemHandler(t *testing.T) {
	handler := start()
	a := NewApiService()
	a.router.HandleFunc("/item/{id}", handler.ChooseItem).Methods("GET")

	handler.vendingMachine.SetState(handler.vendingMachine.ItemRequesting)

	req, err := http.NewRequest("GET", "/item/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	a.router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "handler returned wrong status code")

}
func TestCreateVendingMachine(t *testing.T) {
	handler := start()

	jsonBody, _ := json.Marshal(&struct{ Count int }{Count: 10})
	req, err := http.NewRequest("POST", "/add-machine", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler.CreateVendingMachine(rr, req)
	assert.Equal(t, http.StatusCreated, rr.Code, "handler returned wrong status code")

}
func TestGetMachines(t *testing.T) {
	handler := start()

	req, err := http.NewRequest("GET", "/machines", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler.GetVendingMachines(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code, "handler returned wrong status code")
}

var (
	vendingMachines []state.VendingMachine
	mutex           sync.Mutex
)

func start() *VendingMachineHandle {
	updateChan := make(chan []state.VendingMachine)

	v := state.NewVendingMachine(20)
	h := NewVendingMachineHandler(v, vendingMachines, updateChan, &mutex)
	return h

}

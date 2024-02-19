package api

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
	"github.com/hamedblue1381/vending-machine/models"
	"github.com/hamedblue1381/vending-machine/state"
)

type VendingMachineHandler interface {
	InsertCoin(w http.ResponseWriter, r *http.Request)
	GetVendingMachines(w http.ResponseWriter, r *http.Request)
	ChooseItem(w http.ResponseWriter, r *http.Request)

	getItems(w http.ResponseWriter, r *http.Request)
	CreateVendingMachine(w http.ResponseWriter, r *http.Request)
}

type VendingMachineHandle struct {
	VendingMachineHandler
	vendingMachine  *state.VendingMachine
	vendingMachines []state.VendingMachine
	updateChan      chan []state.VendingMachine
	mutex           *sync.Mutex
}
type VendingMachineRequest struct {
	Count int `json:"count"`
}
type VendingMachineResponse struct {
	Id    int
	Stock int
	Coins int
	Items []models.Item
}

func NewVendingMachineHandler(v *state.VendingMachine, vendingMachines []state.VendingMachine, updateChan chan []state.VendingMachine, mutex *sync.Mutex) *VendingMachineHandle {
	handler := &VendingMachineHandle{
		vendingMachine:  v,
		vendingMachines: vendingMachines,
		updateChan:      updateChan,
		mutex:           mutex,
	}

	go handler.startListeningUpdates()

	return handler
}

func (v *VendingMachineHandle) InsertCoin(w http.ResponseWriter, r *http.Request) {
	err := v.vendingMachine.InsertCoin()
	if err != nil {
		return
	}
	v.getItems(w, r)
}

func (v *VendingMachineHandle) GetVendingMachines(w http.ResponseWriter, r *http.Request) {
	vendingMachines := make([]VendingMachineResponse, 0)

	for id, vmachine := range v.vendingMachines {
		vendingMachines = append(vendingMachines, VendingMachineResponse{
			Id:    id,
			Stock: vmachine.Stock,
			Coins: vmachine.Coins,
			Items: vmachine.Items,
		})
	}

	jsonData, err := json.Marshal(vendingMachines)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func (v *VendingMachineHandle) ChooseItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemIDStr := vars["id"]
	itemID, err := strconv.Atoi(itemIDStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = v.vendingMachine.ChooseItem(itemID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = v.vendingMachine.Dispense(itemID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Item dispensed successfully!"})
}

func (v *VendingMachineHandle) CreateVendingMachine(w http.ResponseWriter, r *http.Request) {
	var req VendingMachineRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if v.vendingMachines == nil {
		v.vendingMachines = []state.VendingMachine{}
	}

	vmachine := state.NewVendingMachine(req.Count)
	v.vendingMachines = append(v.vendingMachines, *vmachine)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	response := VendingMachineResponse{
		Stock: vmachine.Stock,
		Coins: vmachine.Coins,
		Items: vmachine.Items,
	}
	json.NewEncoder(w).Encode(response)
}

func (v *VendingMachineHandle) getItems(w http.ResponseWriter, r *http.Request) {
	items := v.vendingMachine.Items
	jsonData, err := json.Marshal(items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func (h *VendingMachineHandle) startListeningUpdates() {
	for {
		select {
		case updatedVendingMachines := <-h.updateChan:
			h.mutex.Lock()
			h.vendingMachines = updatedVendingMachines
			h.mutex.Unlock()
		case <-context.Background().Done():
			return
		}
	}
}

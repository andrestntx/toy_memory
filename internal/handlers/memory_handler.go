package handlers

import (
	"encoding/json"
	"flink/pgk/memory"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	oId, _ := vars["order_id"]

	max, _ := strconv.Atoi(r.FormValue("max"))
	locations := memory.ToyMemory.Get(oId, max)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(History{
		OrderId:   oId,
		Locations: locations,
	})
}

func Put(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	oId, _ := vars["order_id"]

	var l memory.Location
	_ = json.NewDecoder(r.Body).Decode(&l)

	memory.ToyMemory.Put(oId, l)

	w.WriteHeader(http.StatusOK)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	oId, _ := vars["order_id"]
	memory.ToyMemory.Delete(oId)

	w.WriteHeader(http.StatusOK)
}

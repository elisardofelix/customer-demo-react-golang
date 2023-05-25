package handlers

import (
	"customer-demo/pkg/providers"
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

type CustomerHandler interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	GetById(w http.ResponseWriter, r *http.Request)
}

type customerHandler struct {
	customerProvider providers.CustomerProvider
}

type ErrResponse struct {
	Error string `json:"error,omitempty"`
}

type SuccessResponse struct {
	Data interface{} `json:"data,omitempty"`
}

func (c *customerHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	customers, err := c.customerProvider.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := ErrResponse{
			Error: err.Error(),
		}

		json.NewEncoder(w).Encode(response)
		return
	}

	response := SuccessResponse{
		Data: customers,
	}

	json.NewEncoder(w).Encode(response)
}

func (c *customerHandler) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//Get the id from the url
	id := chi.URLParam(r, "id")

	//Try to convert to int the ID
	TempId, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		response := ErrResponse{
			Error: "Invalid ID",
		}

		json.NewEncoder(w).Encode(response)
		return
	}

	customer, err := c.customerProvider.GetById(TempId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := ErrResponse{
			Error: err.Error(),
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	response := SuccessResponse{
		Data: customer,
	}

	json.NewEncoder(w).Encode(response)
}

func NewCustomerHandler(customerProvider providers.CustomerProvider) CustomerHandler {
	return &customerHandler{
		customerProvider: customerProvider,
	}
}

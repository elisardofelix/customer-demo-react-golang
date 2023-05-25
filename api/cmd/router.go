package main

import (
	"customer-demo/pkg/handlers"
	"github.com/go-chi/chi"
)

func SetCustomerRouter(r *chi.Mux, customerHandler handlers.CustomerHandler) {
	r.Get("/api/customers", customerHandler.GetAll)
	r.Get("/api/customers/{id}", customerHandler.GetById)
}

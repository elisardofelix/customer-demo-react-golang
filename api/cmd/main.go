package main

import (
	"customer-demo/pkg/handlers"
	"customer-demo/pkg/providers"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"net/http"
)

func main() {

	db := MySQLConnect()

	customerProvider := providers.NewCustomerProvider(db)

	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	}))

	customerHandler := handlers.NewCustomerHandler(customerProvider)

	SetCustomerRouter(r, customerHandler)

	http.ListenAndServe(":8080", r)

}

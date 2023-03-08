package main

import (
	"context"
	"net/http"

	"github.com/MaiconGiehl/TechTurAPI/internal/infra/webserver/handlers"
	"github.com/go-chi/chi/v5"
)

var ctx = context.Background()

func main() {
	http.ListenAndServe(":8000", router())
}

func router() http.Handler{
	router := chi.NewRouter()
	
	router.Route("/customer", func(r chi.Router) {
		customerHandler := handlers.NewCustomerHandler(ctx)
		r.Get("/{id}", customerHandler.CustomerHandler)
	})

	router.Route("/bus", func(r chi.Router) {
		busHandler := handlers.NewBusHandler(ctx)
		r.Get("/{id}", busHandler.BusHandler)
	})

	router.Route("/ticket", func(r chi.Router) {
		ticketHandler := handlers.NewTicketHandler(ctx)
		r.Get("/{id}", ticketHandler.TicketHandler)
	})

	router.Route("/route", func(r chi.Router) {
		routeHandler := handlers.NewRouteHandler(ctx)
		r.Get("/{id}", routeHandler.RouteHandler)
	})

	return router
}
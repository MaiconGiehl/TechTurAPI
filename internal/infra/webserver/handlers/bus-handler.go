package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type BusHandlerType struct {
	ctx context.Context
}

func NewBusHandler(ctx context.Context) *BusHandlerType {
	return &BusHandlerType{
		ctx: ctx,
	}
}

func (handler *BusHandlerType) BusHandler(w http.ResponseWriter,  r *http.Request) {
	fmt.Print(chi.URLParam(r, "id"))
}
package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type CustomerHandlerType struct {
	Ctx context.Context
}

func NewCustomerHandler(ctx context.Context) *CustomerHandlerType {
	return &CustomerHandlerType{
		Ctx: ctx,
	}
}

func (handler *CustomerHandlerType) CustomerHandler(w http.ResponseWriter,  r *http.Request) {
	fmt.Print(chi.URLParam(r, "id"))
}
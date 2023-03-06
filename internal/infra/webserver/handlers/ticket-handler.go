package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type TicketHandlerType struct {
	Ctx context.Context
}

func NewTicketHandler(ctx context.Context) *TicketHandlerType {
	return &TicketHandlerType{
		Ctx: ctx,
	}
}

func (handler *TicketHandlerType) TicketHandler(w http.ResponseWriter,  r *http.Request) {
	fmt.Print(chi.URLParam(r, "id"))
}
package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type RouteHandlerType struct {
	Ctx context.Context
}

func NewRouteHandler(ctx context.Context) *RouteHandlerType {
	return &RouteHandlerType{
		Ctx: ctx,
	}
}

func (handler *RouteHandlerType) RouteHandler(w http.ResponseWriter,  r *http.Request) {
	fmt.Print(chi.URLParam(r, "id"))
}
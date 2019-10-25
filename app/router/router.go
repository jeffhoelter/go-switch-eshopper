package router

import (
	"github.com/jeffhoelter/go-switch-eshopper/app/app"
	"github.com/jeffhoelter/go-switch-eshopper/app/requestlog"

	"github.com/go-chi/chi"
)

// New creates a new Chi logger
func New(a *app.App) *chi.Mux {
	l := a.Logger()
	r := chi.NewRouter()
	r.Method("GET", "/", requestlog.NewHandler(a.HandleIndex, l))
	return r
}

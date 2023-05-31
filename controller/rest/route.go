package rest

import (
	"github.com/go-chi/chi/v5"
	"github.com/lapwingcloud/lapwingwire/controller/ent"
)

type Handler interface {
	AgentHandler
	RegisterRoutes(*chi.Mux)
}

type handler struct {
	db *ent.Client
}

func NewHandler(db *ent.Client) Handler {
	return &handler{
		db: db,
	}
}

func (t *handler) RegisterRoutes(r *chi.Mux) {
	r.Put("/v1/agent", t.PutAgent)
}

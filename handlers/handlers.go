package handlers

import (
	"github.com/gastrader/gohouse/config"
	"github.com/gastrader/gohouse/ent"
)

type Handler struct {
	Client *ent.Client
	Config *config.Config
}

func NewHandlers(client *ent.Client, config *config.Config) *Handler{
	return &Handler{
		Client: client,
		Config: config,
	}
}
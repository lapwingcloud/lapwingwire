package rest

import (
	"encoding/json"
	"net/http"

	"github.com/lapwingcloud/lapwingwire/controller/ent"
)

type Agent struct {
	ID       string   `json:"id"`
	Hostname string   `json:"hostname"`
	Tags     []string `json:"tags"`
}

type AgentHandler interface {
	PutAgent(http.ResponseWriter, *http.Request)
}

type PutAgentRequest struct {
	Agent
}

func (t *PutAgentRequest) Bind(r *http.Request) error {
	return nil
}

func (t *handler) PutAgent(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	data := &PutAgentRequest{}
	if err := json.NewDecoder(r.Body).Decode(data); err != nil {
		NewResponse(w, r).BadRequest().Err(err).Msg("failed to decode request body to json")
		return
	}
	if data.ID == "" {
		tagCreates := make([]*ent.TagCreate, len(data.Tags))
		for i, tagName := range data.Tags {
			tagCreates[i] = t.db.Tag.Create().SetName(tagName)
		}
		tags, err := t.db.Tag.CreateBulk(tagCreates...).Save(ctx)
		if err != nil {
			NewResponse(w, r).BadRequest().Err(err).Msg("failed to create tags")
			return
		}
		agent, err := t.db.Agent.
			Create().
			SetHostname(data.Hostname).
			AddTags(tags...).
			Save(ctx)
		if err != nil {
			NewResponse(w, r).BadRequest().Err(err).Msg("failed to create the agent")
			return
		}
		NewResponse(w, r).OK().Data(agent)
		return
	} else {
		NewResponse(w, r).NotImplemented().Msg("not implemented")
		return
	}
}

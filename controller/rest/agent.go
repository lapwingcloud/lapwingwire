package rest

import (
	"net/http"

	"github.com/rs/zerolog/hlog"
)

type Agent struct {
	ID       string `json:"id"`
	Hostname string `json:"hostname"`
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
	// ctx := r.Context()
	// data := &PutAgentRequest{}
	// if err := render.Bind(r, data); err != nil {
	// 	if err := render.Render(w, r, ErrBadRequest(err)); err != nil {
	// 		_ = render.Render(w, r, ErrUnprocessableContent(err))
	// 	}
	// 	return
	// }
	// if data.ID == "" {
	// 	t.db.Agent.Create().SetHostname(data.Hostname).Save(ctx)
	// } else {
	// 	agentID, err := strconv.Atoi(data.ID)
	// 	// if err != nil {
	// 	// 	return ErrBadRequest(err)
	// 	// }
	// 	_, err = t.db.Agent.Query().Where(agent.ID(agentID)).First(ctx)
	// 	if err != nil {

	// 	}
	// }
	hlog.FromRequest(r).Warn().Msg("asdasdqlkwejqlwkej")
	NewResponse(w, r).BadRequest().Msg("hello")
}

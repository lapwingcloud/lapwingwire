package rest

import (
	"encoding/json"
	"net/http"

	"github.com/lapwingcloud/lapwingwire/controller/ent"
	"github.com/lapwingcloud/lapwingwire/controller/ent/agent"
	"github.com/lapwingcloud/lapwingwire/controller/ent/predicate"
	"github.com/lapwingcloud/lapwingwire/controller/ent/tag"
)

type Agent struct {
	ID       int    `json:"id"`
	Hostname string `json:"hostname"`
	Tags     []Tag  `json:"tags"`
}

func (t *handler) PutAgent(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	rw := NewResponder(w, r)

	var req struct {
		Agent
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		rw.BadRequest().Err(err).Msg("failed to decode request body to json")
		return
	}
	tagAnds := make([]predicate.Tag, len(req.Tags))
	for i, tagData := range req.Tags {
		tagAnds[i] = tag.And(tag.Name(tagData.Name), tag.Value(tagData.Value))
	}
	existingTagEnts, err := t.db.Tag.Query().Where(tag.Or(tagAnds...)).All(ctx)
	if err != nil {
		rw.InternalServerError().Err(err).Msg("failed to query existing tags")
		return
	}
	existingTagMap := make(map[string]string, len(existingTagEnts))
	for _, tg := range existingTagEnts {
		existingTagMap[tg.Name] = tg.Value
	}
	tagCreates := []*ent.TagCreate{}
	for _, tagData := range req.Tags {
		if existingValue, ok := existingTagMap[tagData.Name]; ok && tagData.Value == existingValue {
			continue
		}
		tagCreates = append(tagCreates, t.db.Tag.Create().SetName(tagData.Name).SetValue(tagData.Value))
	}
	newTagEnts, err := t.db.Tag.CreateBulk(tagCreates...).Save(ctx)
	if err != nil {
		rw.InternalServerError().Err(err).Msg("failed to create tags")
		return
	}
	var agentEnt *ent.Agent
	if req.ID == 0 {
		agentEnt, err = t.db.Agent.
			Create().
			SetHostname(req.Hostname).
			AddTags(existingTagEnts...).
			AddTags(newTagEnts...).
			Save(ctx)
		if err != nil {
			rw.BadRequest().Err(err).Msg("failed to create the agent")
			return
		}
	} else {
		agentEnt, err = t.db.Agent.Query().Where(agent.ID(req.ID)).First(ctx)
		if err != nil {
			rw.InternalServerError().Err(err).Msg("failed to query agent")
			return
		}
		if agentEnt == nil {
			rw.NotFound().Msg("agent not found")
			return
		}
		agentEnt, err = t.db.Agent.
			UpdateOne(agentEnt).
			SetHostname(req.Hostname).
			ClearTags().
			AddTags(existingTagEnts...).
			AddTags(newTagEnts...).
			Save(ctx)
		if err != nil {
			rw.BadRequest().Err(err).Msg("failed to update the agent")
			return
		}
	}
	// rw.OK().Data(PutAgentResponse{
	// 	Agent: Agent{
	// 		ID:       agentEnt.ID,
	// 		Hostname: agentEnt.Hostname,
	// 	},
	// })
	rw.OK().Data(agentEnt)
}

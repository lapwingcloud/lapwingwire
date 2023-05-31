package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Agent struct {
	ent.Schema
}

func (Agent) Fields() []ent.Field {
	return []ent.Field{
		field.String("hostname").NotEmpty(),
		field.Time("created_at").Default(time.Now),
	}
}

func (Agent) Edges() []ent.Edge {
	return []ent.Edge{}
}

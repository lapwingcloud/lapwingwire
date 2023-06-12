package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Tag holds the schema definition for the Tag entity.
type Tag struct {
	ent.Schema
}

// Fields of the Tag.
func (Tag) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("value").NotEmpty(),
	}
}

// Edges of the Tag.
func (Tag) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("agents", Agent.Type).Ref("tags"),
	}
}

func (Tag) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name", "value").Unique(),
	}
}

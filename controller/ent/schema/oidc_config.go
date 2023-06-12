package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// OIDCConfig holds the schema definition for the OIDCConfig entity.
type OIDCConfig struct {
	ent.Schema
}

// Fields of the OIDCConfig.
func (OIDCConfig) Fields() []ent.Field {
	return []ent.Field{
		field.String("provider_key").NotEmpty().Unique(),
		field.String("provider_name").NotEmpty(),
		field.String("discovery_uri").NotEmpty(),
		field.String("client_id").NotEmpty(),
		field.String("client_secret").NotEmpty(),
		field.String("redirect_uri").NotEmpty(),
	}
}

// Edges of the OIDCConfig.
func (OIDCConfig) Edges() []ent.Edge {
	return nil
}

// Code generated by ent, DO NOT EDIT.

package tag

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the tag type in the database.
	Label = "tag"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeAgents holds the string denoting the agents edge name in mutations.
	EdgeAgents = "agents"
	// Table holds the table name of the tag in the database.
	Table = "tags"
	// AgentsTable is the table that holds the agents relation/edge. The primary key declared below.
	AgentsTable = "agent_tags"
	// AgentsInverseTable is the table name for the Agent entity.
	// It exists in this package in order to avoid circular dependency with the "agent" package.
	AgentsInverseTable = "agents"
)

// Columns holds all SQL columns for tag fields.
var Columns = []string{
	FieldID,
	FieldName,
}

var (
	// AgentsPrimaryKey and AgentsColumn2 are the table columns denoting the
	// primary key for the agents relation (M2M).
	AgentsPrimaryKey = []string{"agent_id", "tag_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
)

// OrderOption defines the ordering options for the Tag queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByAgentsCount orders the results by agents count.
func ByAgentsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAgentsStep(), opts...)
	}
}

// ByAgents orders the results by agents terms.
func ByAgents(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAgentsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newAgentsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AgentsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, AgentsTable, AgentsPrimaryKey...),
	)
}

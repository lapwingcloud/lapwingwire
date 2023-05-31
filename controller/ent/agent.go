// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/lapwingcloud/lapwingwire/controller/ent/agent"
)

// Agent is the model entity for the Agent schema.
type Agent struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Hostname holds the value of the "hostname" field.
	Hostname string `json:"hostname,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt    time.Time `json:"created_at,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Agent) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case agent.FieldID:
			values[i] = new(sql.NullInt64)
		case agent.FieldHostname:
			values[i] = new(sql.NullString)
		case agent.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Agent fields.
func (a *Agent) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case agent.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = int(value.Int64)
		case agent.FieldHostname:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field hostname", values[i])
			} else if value.Valid {
				a.Hostname = value.String
			}
		case agent.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				a.CreatedAt = value.Time
			}
		default:
			a.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Agent.
// This includes values selected through modifiers, order, etc.
func (a *Agent) Value(name string) (ent.Value, error) {
	return a.selectValues.Get(name)
}

// Update returns a builder for updating this Agent.
// Note that you need to call Agent.Unwrap() before calling this method if this Agent
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Agent) Update() *AgentUpdateOne {
	return NewAgentClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the Agent entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Agent) Unwrap() *Agent {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Agent is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Agent) String() string {
	var builder strings.Builder
	builder.WriteString("Agent(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("hostname=")
	builder.WriteString(a.Hostname)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(a.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Agents is a parsable slice of Agent.
type Agents []*Agent
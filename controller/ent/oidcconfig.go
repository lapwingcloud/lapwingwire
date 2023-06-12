// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/lapwingcloud/lapwingwire/controller/ent/oidcconfig"
)

// OIDCConfig is the model entity for the OIDCConfig schema.
type OIDCConfig struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// ProviderKey holds the value of the "provider_key" field.
	ProviderKey string `json:"provider_key,omitempty"`
	// ProviderName holds the value of the "provider_name" field.
	ProviderName string `json:"provider_name,omitempty"`
	// DiscoveryURI holds the value of the "discovery_uri" field.
	DiscoveryURI string `json:"discovery_uri,omitempty"`
	// ClientID holds the value of the "client_id" field.
	ClientID string `json:"client_id,omitempty"`
	// ClientSecret holds the value of the "client_secret" field.
	ClientSecret string `json:"client_secret,omitempty"`
	// RedirectURI holds the value of the "redirect_uri" field.
	RedirectURI  string `json:"redirect_uri,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OIDCConfig) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case oidcconfig.FieldID:
			values[i] = new(sql.NullInt64)
		case oidcconfig.FieldProviderKey, oidcconfig.FieldProviderName, oidcconfig.FieldDiscoveryURI, oidcconfig.FieldClientID, oidcconfig.FieldClientSecret, oidcconfig.FieldRedirectURI:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OIDCConfig fields.
func (oc *OIDCConfig) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case oidcconfig.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			oc.ID = int(value.Int64)
		case oidcconfig.FieldProviderKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field provider_key", values[i])
			} else if value.Valid {
				oc.ProviderKey = value.String
			}
		case oidcconfig.FieldProviderName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field provider_name", values[i])
			} else if value.Valid {
				oc.ProviderName = value.String
			}
		case oidcconfig.FieldDiscoveryURI:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field discovery_uri", values[i])
			} else if value.Valid {
				oc.DiscoveryURI = value.String
			}
		case oidcconfig.FieldClientID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field client_id", values[i])
			} else if value.Valid {
				oc.ClientID = value.String
			}
		case oidcconfig.FieldClientSecret:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field client_secret", values[i])
			} else if value.Valid {
				oc.ClientSecret = value.String
			}
		case oidcconfig.FieldRedirectURI:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field redirect_uri", values[i])
			} else if value.Valid {
				oc.RedirectURI = value.String
			}
		default:
			oc.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the OIDCConfig.
// This includes values selected through modifiers, order, etc.
func (oc *OIDCConfig) Value(name string) (ent.Value, error) {
	return oc.selectValues.Get(name)
}

// Update returns a builder for updating this OIDCConfig.
// Note that you need to call OIDCConfig.Unwrap() before calling this method if this OIDCConfig
// was returned from a transaction, and the transaction was committed or rolled back.
func (oc *OIDCConfig) Update() *OIDCConfigUpdateOne {
	return NewOIDCConfigClient(oc.config).UpdateOne(oc)
}

// Unwrap unwraps the OIDCConfig entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (oc *OIDCConfig) Unwrap() *OIDCConfig {
	_tx, ok := oc.config.driver.(*txDriver)
	if !ok {
		panic("ent: OIDCConfig is not a transactional entity")
	}
	oc.config.driver = _tx.drv
	return oc
}

// String implements the fmt.Stringer.
func (oc *OIDCConfig) String() string {
	var builder strings.Builder
	builder.WriteString("OIDCConfig(")
	builder.WriteString(fmt.Sprintf("id=%v, ", oc.ID))
	builder.WriteString("provider_key=")
	builder.WriteString(oc.ProviderKey)
	builder.WriteString(", ")
	builder.WriteString("provider_name=")
	builder.WriteString(oc.ProviderName)
	builder.WriteString(", ")
	builder.WriteString("discovery_uri=")
	builder.WriteString(oc.DiscoveryURI)
	builder.WriteString(", ")
	builder.WriteString("client_id=")
	builder.WriteString(oc.ClientID)
	builder.WriteString(", ")
	builder.WriteString("client_secret=")
	builder.WriteString(oc.ClientSecret)
	builder.WriteString(", ")
	builder.WriteString("redirect_uri=")
	builder.WriteString(oc.RedirectURI)
	builder.WriteByte(')')
	return builder.String()
}

// OIDCConfigs is a parsable slice of OIDCConfig.
type OIDCConfigs []*OIDCConfig

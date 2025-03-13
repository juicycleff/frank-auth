package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/juicycleff/frank/pkg/entity"
)

// WebhookEvent holds the schema definition for the WebhookEvent entity.
type WebhookEvent struct {
	ent.Schema
}

// Fields of the WebhookEvent.
func (WebhookEvent) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			Unique(),
		field.String("webhook_id").
			NotEmpty(),
		field.String("event_type").
			NotEmpty(),
		entity.JSONMapStringField("headers", true),
		entity.JSONMapField("payload", true),
		field.Bool("delivered").
			Default(false),
		field.Time("delivered_at").
			Optional().
			Nillable(),
		field.Int("attempts").
			Default(0),
		field.Time("next_retry").
			Optional().
			Nillable(),
		field.Int("status_code").
			Optional().
			Nillable(),
		field.String("response_body").
			Optional(),
		field.String("error").
			Optional(),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the WebhookEvent.
func (WebhookEvent) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("webhook", Webhook.Type).
			Ref("events").
			Unique().
			Field("webhook_id").
			Required(),
	}
}

// Indexes of the WebhookEvent.
func (WebhookEvent) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("webhook_id"),
		index.Fields("event_type"),
		index.Fields("delivered"),
		index.Fields("next_retry"),
	}
}

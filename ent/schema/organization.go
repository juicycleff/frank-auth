package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/juicycleff/frank/pkg/entity"
)

// Organization holds the schema definition for the Organization entity.
type Organization struct {
	ent.Schema
}

// Fields of the Organization.
func (Organization) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty(),
		field.String("slug").
			Unique().
			NotEmpty(),
		field.String("domain").
			Optional(),
		field.String("logo_url").
			Optional(),
		field.String("plan").
			Default("free"),
		field.Bool("active").
			Default(true),
		entity.JSONMapField("metadata", true),
		field.Time("trial_ends_at").
			Optional().
			Nillable(),
		field.Bool("trial_used").
			Default(false),
	}
}

// Edges of the Organization.
func (Organization) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", User.Type),
		edge.To("api_keys", ApiKey.Type),
		edge.To("webhooks", Webhook.Type),
		edge.To("feature_flags", OrganizationFeature.Type),
		edge.To("identity_providers", IdentityProvider.Type),
		edge.To("oauth_clients", OAuthClient.Type),
	}
}

// Indexes of the Organization.
func (Organization) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("domain"),
		index.Fields("slug"),
	}
}

// Mixin of the Organization.
func (Organization) Mixin() []ent.Mixin {
	return []ent.Mixin{
		ModelBaseMixin{},
	}
}

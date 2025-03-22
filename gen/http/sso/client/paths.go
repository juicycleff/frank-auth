// Code generated by goa v3.20.0, DO NOT EDIT.
//
// HTTP request path constructors for the sso service.
//
// Command:
// $ goa gen github.com/juicycleff/frank/design -o .

package client

import (
	"fmt"
)

// ListProvidersSsoPath returns the URL path to the sso service list_providers HTTP endpoint.
func ListProvidersSsoPath() string {
	return "/v1/auth/sso/providers"
}

// ProviderAuthSsoPath returns the URL path to the sso service provider_auth HTTP endpoint.
func ProviderAuthSsoPath(provider string) string {
	return fmt.Sprintf("/v1/auth/sso/providers/%v", provider)
}

// ProviderCallbackSsoPath returns the URL path to the sso service provider_callback HTTP endpoint.
func ProviderCallbackSsoPath(provider string) string {
	return fmt.Sprintf("/v1/auth/sso/callback/%v", provider)
}

// ListIdentityProvidersSsoPath returns the URL path to the sso service list_identity_providers HTTP endpoint.
func ListIdentityProvidersSsoPath() string {
	return "/v1/auth/sso/identity-providers"
}

// CreateIdentityProviderSsoPath returns the URL path to the sso service create_identity_provider HTTP endpoint.
func CreateIdentityProviderSsoPath() string {
	return "/v1/auth/sso/identity-providers"
}

// GetIdentityProviderSsoPath returns the URL path to the sso service get_identity_provider HTTP endpoint.
func GetIdentityProviderSsoPath(id string) string {
	return fmt.Sprintf("/v1/auth/sso/identity-providers/%v", id)
}

// UpdateIdentityProviderSsoPath returns the URL path to the sso service update_identity_provider HTTP endpoint.
func UpdateIdentityProviderSsoPath(id string) string {
	return fmt.Sprintf("/v1/auth/sso/identity-providers/%v", id)
}

// DeleteIdentityProviderSsoPath returns the URL path to the sso service delete_identity_provider HTTP endpoint.
func DeleteIdentityProviderSsoPath(id string) string {
	return fmt.Sprintf("/v1/auth/sso/identity-providers/%v", id)
}

// SamlMetadataSsoPath returns the URL path to the sso service saml_metadata HTTP endpoint.
func SamlMetadataSsoPath(id string) string {
	return fmt.Sprintf("/v1/auth/sso/saml/%v/metadata", id)
}

// SamlAcsSsoPath returns the URL path to the sso service saml_acs HTTP endpoint.
func SamlAcsSsoPath(id string) string {
	return fmt.Sprintf("/v1/auth/sso/saml/%v/acs", id)
}

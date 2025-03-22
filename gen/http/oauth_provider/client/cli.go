// Code generated by goa v3.20.0, DO NOT EDIT.
//
// oauth_provider HTTP client CLI support package
//
// Command:
// $ goa gen github.com/juicycleff/frank/design -o .

package client

import (
	"encoding/json"
	"fmt"
	"strconv"

	oauthprovider "github.com/juicycleff/frank/gen/oauth_provider"
	goa "goa.design/goa/v3/pkg"
)

// BuildAuthorizePayload builds the payload for the oauth_provider authorize
// endpoint from CLI flags.
func BuildAuthorizePayload(oauthProviderAuthorizeClientID string, oauthProviderAuthorizeResponseType string, oauthProviderAuthorizeRedirectURI string, oauthProviderAuthorizeScope string, oauthProviderAuthorizeState string, oauthProviderAuthorizeCodeChallenge string, oauthProviderAuthorizeCodeChallengeMethod string, oauthProviderAuthorizeOauth2 string, oauthProviderAuthorizeXAPIKey string, oauthProviderAuthorizeJWT string) (*oauthprovider.AuthorizePayload, error) {
	var clientID string
	{
		clientID = oauthProviderAuthorizeClientID
	}
	var responseType string
	{
		responseType = oauthProviderAuthorizeResponseType
	}
	var redirectURI string
	{
		redirectURI = oauthProviderAuthorizeRedirectURI
	}
	var scope *string
	{
		if oauthProviderAuthorizeScope != "" {
			scope = &oauthProviderAuthorizeScope
		}
	}
	var state *string
	{
		if oauthProviderAuthorizeState != "" {
			state = &oauthProviderAuthorizeState
		}
	}
	var codeChallenge *string
	{
		if oauthProviderAuthorizeCodeChallenge != "" {
			codeChallenge = &oauthProviderAuthorizeCodeChallenge
		}
	}
	var codeChallengeMethod *string
	{
		if oauthProviderAuthorizeCodeChallengeMethod != "" {
			codeChallengeMethod = &oauthProviderAuthorizeCodeChallengeMethod
		}
	}
	var oauth2 *string
	{
		if oauthProviderAuthorizeOauth2 != "" {
			oauth2 = &oauthProviderAuthorizeOauth2
		}
	}
	var xAPIKey *string
	{
		if oauthProviderAuthorizeXAPIKey != "" {
			xAPIKey = &oauthProviderAuthorizeXAPIKey
		}
	}
	var jwt *string
	{
		if oauthProviderAuthorizeJWT != "" {
			jwt = &oauthProviderAuthorizeJWT
		}
	}
	v := &oauthprovider.AuthorizePayload{}
	v.ClientID = clientID
	v.ResponseType = responseType
	v.RedirectURI = redirectURI
	v.Scope = scope
	v.State = state
	v.CodeChallenge = codeChallenge
	v.CodeChallengeMethod = codeChallengeMethod
	v.Oauth2 = oauth2
	v.XAPIKey = xAPIKey
	v.JWT = jwt

	return v, nil
}

// BuildTokenPayload builds the payload for the oauth_provider token endpoint
// from CLI flags.
func BuildTokenPayload(oauthProviderTokenGrantType string, oauthProviderTokenCode string, oauthProviderTokenRedirectURI string, oauthProviderTokenClientID string, oauthProviderTokenClientSecret string, oauthProviderTokenRefreshToken string, oauthProviderTokenCodeVerifier string, oauthProviderTokenUsername string, oauthProviderTokenPassword string, oauthProviderTokenScope string, oauthProviderTokenOauth2 string, oauthProviderTokenXAPIKey string, oauthProviderTokenJWT string) (*oauthprovider.TokenPayload, error) {
	var grantType string
	{
		grantType = oauthProviderTokenGrantType
	}
	var code *string
	{
		if oauthProviderTokenCode != "" {
			code = &oauthProviderTokenCode
		}
	}
	var redirectURI *string
	{
		if oauthProviderTokenRedirectURI != "" {
			redirectURI = &oauthProviderTokenRedirectURI
		}
	}
	var clientID *string
	{
		if oauthProviderTokenClientID != "" {
			clientID = &oauthProviderTokenClientID
		}
	}
	var clientSecret *string
	{
		if oauthProviderTokenClientSecret != "" {
			clientSecret = &oauthProviderTokenClientSecret
		}
	}
	var refreshToken *string
	{
		if oauthProviderTokenRefreshToken != "" {
			refreshToken = &oauthProviderTokenRefreshToken
		}
	}
	var codeVerifier *string
	{
		if oauthProviderTokenCodeVerifier != "" {
			codeVerifier = &oauthProviderTokenCodeVerifier
		}
	}
	var username *string
	{
		if oauthProviderTokenUsername != "" {
			username = &oauthProviderTokenUsername
		}
	}
	var password *string
	{
		if oauthProviderTokenPassword != "" {
			password = &oauthProviderTokenPassword
		}
	}
	var scope *string
	{
		if oauthProviderTokenScope != "" {
			scope = &oauthProviderTokenScope
		}
	}
	var oauth2 *string
	{
		if oauthProviderTokenOauth2 != "" {
			oauth2 = &oauthProviderTokenOauth2
		}
	}
	var xAPIKey *string
	{
		if oauthProviderTokenXAPIKey != "" {
			xAPIKey = &oauthProviderTokenXAPIKey
		}
	}
	var jwt *string
	{
		if oauthProviderTokenJWT != "" {
			jwt = &oauthProviderTokenJWT
		}
	}
	v := &oauthprovider.TokenPayload{}
	v.GrantType = grantType
	v.Code = code
	v.RedirectURI = redirectURI
	v.ClientID = clientID
	v.ClientSecret = clientSecret
	v.RefreshToken = refreshToken
	v.CodeVerifier = codeVerifier
	v.Username = username
	v.Password = password
	v.Scope = scope
	v.Oauth2 = oauth2
	v.XAPIKey = xAPIKey
	v.JWT = jwt

	return v, nil
}

// BuildIntrospectPayload builds the payload for the oauth_provider introspect
// endpoint from CLI flags.
func BuildIntrospectPayload(oauthProviderIntrospectToken string, oauthProviderIntrospectTokenTypeHint string, oauthProviderIntrospectOauth2 string, oauthProviderIntrospectXAPIKey string, oauthProviderIntrospectJWT string) (*oauthprovider.IntrospectPayload, error) {
	var token string
	{
		token = oauthProviderIntrospectToken
	}
	var tokenTypeHint *string
	{
		if oauthProviderIntrospectTokenTypeHint != "" {
			tokenTypeHint = &oauthProviderIntrospectTokenTypeHint
		}
	}
	var oauth2 *string
	{
		if oauthProviderIntrospectOauth2 != "" {
			oauth2 = &oauthProviderIntrospectOauth2
		}
	}
	var xAPIKey *string
	{
		if oauthProviderIntrospectXAPIKey != "" {
			xAPIKey = &oauthProviderIntrospectXAPIKey
		}
	}
	var jwt *string
	{
		if oauthProviderIntrospectJWT != "" {
			jwt = &oauthProviderIntrospectJWT
		}
	}
	v := &oauthprovider.IntrospectPayload{}
	v.Token = token
	v.TokenTypeHint = tokenTypeHint
	v.Oauth2 = oauth2
	v.XAPIKey = xAPIKey
	v.JWT = jwt

	return v, nil
}

// BuildRevokePayload builds the payload for the oauth_provider revoke endpoint
// from CLI flags.
func BuildRevokePayload(oauthProviderRevokeToken string, oauthProviderRevokeTokenTypeHint string, oauthProviderRevokeClientID string, oauthProviderRevokeClientSecret string, oauthProviderRevokeOauth2 string, oauthProviderRevokeXAPIKey string, oauthProviderRevokeJWT string) (*oauthprovider.RevokePayload, error) {
	var token string
	{
		token = oauthProviderRevokeToken
	}
	var tokenTypeHint *string
	{
		if oauthProviderRevokeTokenTypeHint != "" {
			tokenTypeHint = &oauthProviderRevokeTokenTypeHint
		}
	}
	var clientID *string
	{
		if oauthProviderRevokeClientID != "" {
			clientID = &oauthProviderRevokeClientID
		}
	}
	var clientSecret *string
	{
		if oauthProviderRevokeClientSecret != "" {
			clientSecret = &oauthProviderRevokeClientSecret
		}
	}
	var oauth2 *string
	{
		if oauthProviderRevokeOauth2 != "" {
			oauth2 = &oauthProviderRevokeOauth2
		}
	}
	var xAPIKey *string
	{
		if oauthProviderRevokeXAPIKey != "" {
			xAPIKey = &oauthProviderRevokeXAPIKey
		}
	}
	var jwt *string
	{
		if oauthProviderRevokeJWT != "" {
			jwt = &oauthProviderRevokeJWT
		}
	}
	v := &oauthprovider.RevokePayload{}
	v.Token = token
	v.TokenTypeHint = tokenTypeHint
	v.ClientID = clientID
	v.ClientSecret = clientSecret
	v.Oauth2 = oauth2
	v.XAPIKey = xAPIKey
	v.JWT = jwt

	return v, nil
}

// BuildConsentPayload builds the payload for the oauth_provider consent
// endpoint from CLI flags.
func BuildConsentPayload(oauthProviderConsentBody string, oauthProviderConsentJWT string) (*oauthprovider.ConsentPayload, error) {
	var err error
	var body ConsentRequestBody
	{
		err = json.Unmarshal([]byte(oauthProviderConsentBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"approved\": false,\n      \"client_id\": \"Aut iure.\",\n      \"redirect_uri\": \"Autem temporibus doloremque quo.\",\n      \"scope\": \"Provident quis voluptatem facilis.\",\n      \"state\": \"Quibusdam necessitatibus earum rerum veniam.\"\n   }'")
		}
	}
	var jwt *string
	{
		if oauthProviderConsentJWT != "" {
			jwt = &oauthProviderConsentJWT
		}
	}
	v := &oauthprovider.ConsentPayload{
		ClientID:    body.ClientID,
		Scope:       body.Scope,
		RedirectURI: body.RedirectURI,
		State:       body.State,
		Approved:    body.Approved,
	}
	{
		var zero bool
		if v.Approved == zero {
			v.Approved = false
		}
	}
	v.JWT = jwt

	return v, nil
}

// BuildUserinfoPayload builds the payload for the oauth_provider userinfo
// endpoint from CLI flags.
func BuildUserinfoPayload(oauthProviderUserinfoJWT string) (*oauthprovider.UserinfoPayload, error) {
	var jwt *string
	{
		if oauthProviderUserinfoJWT != "" {
			jwt = &oauthProviderUserinfoJWT
		}
	}
	v := &oauthprovider.UserinfoPayload{}
	v.JWT = jwt

	return v, nil
}

// BuildListClientsPayload builds the payload for the oauth_provider
// list_clients endpoint from CLI flags.
func BuildListClientsPayload(oauthProviderListClientsOffset string, oauthProviderListClientsLimit string, oauthProviderListClientsOrganizationID string, oauthProviderListClientsJWT string) (*oauthprovider.ListClientsPayload, error) {
	var err error
	var offset int
	{
		if oauthProviderListClientsOffset != "" {
			var v int64
			v, err = strconv.ParseInt(oauthProviderListClientsOffset, 10, strconv.IntSize)
			offset = int(v)
			if err != nil {
				return nil, fmt.Errorf("invalid value for offset, must be INT")
			}
			if offset < 0 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("offset", offset, 0, true))
			}
			if err != nil {
				return nil, err
			}
		}
	}
	var limit int
	{
		if oauthProviderListClientsLimit != "" {
			var v int64
			v, err = strconv.ParseInt(oauthProviderListClientsLimit, 10, strconv.IntSize)
			limit = int(v)
			if err != nil {
				return nil, fmt.Errorf("invalid value for limit, must be INT")
			}
			if limit < 1 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("limit", limit, 1, true))
			}
			if limit > 100 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("limit", limit, 100, false))
			}
			if err != nil {
				return nil, err
			}
		}
	}
	var organizationID *string
	{
		if oauthProviderListClientsOrganizationID != "" {
			organizationID = &oauthProviderListClientsOrganizationID
		}
	}
	var jwt *string
	{
		if oauthProviderListClientsJWT != "" {
			jwt = &oauthProviderListClientsJWT
		}
	}
	v := &oauthprovider.ListClientsPayload{}
	v.Offset = offset
	v.Limit = limit
	v.OrganizationID = organizationID
	v.JWT = jwt

	return v, nil
}

// BuildCreateClientPayload builds the payload for the oauth_provider
// create_client endpoint from CLI flags.
func BuildCreateClientPayload(oauthProviderCreateClientBody string, oauthProviderCreateClientJWT string) (*oauthprovider.CreateClientPayload, error) {
	var err error
	var body CreateClientRequestBody
	{
		err = json.Unmarshal([]byte(oauthProviderCreateClientBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"allowed_cors_origins\": [\n         \"Asperiores ad.\",\n         \"Consectetur temporibus deserunt ut.\",\n         \"Iusto beatae et odit.\",\n         \"Debitis repudiandae aut ullam praesentium numquam alias.\"\n      ],\n      \"allowed_grant_types\": [\n         \"authorization_code\",\n         \"refresh_token\"\n      ],\n      \"client_description\": \"Qui placeat quasi nisi esse rerum.\",\n      \"client_name\": \"My App\",\n      \"client_uri\": \"Nemo nihil esse laudantium illum.\",\n      \"logo_uri\": \"Ut sapiente dignissimos est incidunt in architecto.\",\n      \"post_logout_redirect_uris\": [\n         \"Consectetur culpa eveniet exercitationem eligendi autem aut.\",\n         \"Blanditiis eos.\",\n         \"Ut magni rem perspiciatis dolores.\",\n         \"Dolore voluptate repellendus sit.\"\n      ],\n      \"public\": false,\n      \"redirect_uris\": [\n         \"https://example.com/callback\"\n      ],\n      \"refresh_token_expiry_seconds\": 8282348841375723111,\n      \"requires_consent\": false,\n      \"requires_pkce\": false,\n      \"token_expiry_seconds\": 4969521611687589696\n   }'")
		}
		if body.RedirectUris == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("redirect_uris", "body"))
		}
		if err != nil {
			return nil, err
		}
	}
	var jwt *string
	{
		if oauthProviderCreateClientJWT != "" {
			jwt = &oauthProviderCreateClientJWT
		}
	}
	v := &oauthprovider.CreateClientPayload{
		ClientName:                body.ClientName,
		ClientDescription:         body.ClientDescription,
		ClientURI:                 body.ClientURI,
		LogoURI:                   body.LogoURI,
		Public:                    body.Public,
		TokenExpirySeconds:        body.TokenExpirySeconds,
		RefreshTokenExpirySeconds: body.RefreshTokenExpirySeconds,
		RequiresPkce:              body.RequiresPkce,
		RequiresConsent:           body.RequiresConsent,
	}
	if body.RedirectUris != nil {
		v.RedirectUris = make([]string, len(body.RedirectUris))
		for i, val := range body.RedirectUris {
			v.RedirectUris[i] = val
		}
	} else {
		v.RedirectUris = []string{}
	}
	if body.PostLogoutRedirectUris != nil {
		v.PostLogoutRedirectUris = make([]string, len(body.PostLogoutRedirectUris))
		for i, val := range body.PostLogoutRedirectUris {
			v.PostLogoutRedirectUris[i] = val
		}
	}
	if body.AllowedCorsOrigins != nil {
		v.AllowedCorsOrigins = make([]string, len(body.AllowedCorsOrigins))
		for i, val := range body.AllowedCorsOrigins {
			v.AllowedCorsOrigins[i] = val
		}
	}
	if body.AllowedGrantTypes != nil {
		v.AllowedGrantTypes = make([]string, len(body.AllowedGrantTypes))
		for i, val := range body.AllowedGrantTypes {
			v.AllowedGrantTypes[i] = val
		}
	}
	{
		var zero bool
		if v.Public == zero {
			v.Public = false
		}
	}
	{
		var zero int
		if v.TokenExpirySeconds == zero {
			v.TokenExpirySeconds = 3600
		}
	}
	{
		var zero int
		if v.RefreshTokenExpirySeconds == zero {
			v.RefreshTokenExpirySeconds = 2592000
		}
	}
	{
		var zero bool
		if v.RequiresPkce == zero {
			v.RequiresPkce = true
		}
	}
	{
		var zero bool
		if v.RequiresConsent == zero {
			v.RequiresConsent = true
		}
	}
	v.JWT = jwt

	return v, nil
}

// BuildGetClientPayload builds the payload for the oauth_provider get_client
// endpoint from CLI flags.
func BuildGetClientPayload(oauthProviderGetClientID string, oauthProviderGetClientJWT string) (*oauthprovider.GetClientPayload, error) {
	var id string
	{
		id = oauthProviderGetClientID
	}
	var jwt *string
	{
		if oauthProviderGetClientJWT != "" {
			jwt = &oauthProviderGetClientJWT
		}
	}
	v := &oauthprovider.GetClientPayload{}
	v.ID = id
	v.JWT = jwt

	return v, nil
}

// BuildUpdateClientPayload builds the payload for the oauth_provider
// update_client endpoint from CLI flags.
func BuildUpdateClientPayload(oauthProviderUpdateClientBody string, oauthProviderUpdateClientID string, oauthProviderUpdateClientJWT string) (*oauthprovider.UpdateClientPayload, error) {
	var err error
	var body UpdateClientRequestBody
	{
		err = json.Unmarshal([]byte(oauthProviderUpdateClientBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"client\": {\n         \"active\": true,\n         \"allowed_cors_origins\": [\n            \"Non et illum.\",\n            \"Porro harum veniam velit.\",\n            \"Accusantium sint.\"\n         ],\n         \"allowed_grant_types\": [\n            \"Molestiae placeat ratione quas sit.\",\n            \"Et sequi libero ut unde dolor ad.\",\n            \"Sit necessitatibus.\"\n         ],\n         \"client_description\": \"Est qui pariatur dolor.\",\n         \"client_name\": \"Quia voluptas doloremque optio voluptas est.\",\n         \"client_uri\": \"Vel sunt amet.\",\n         \"logo_uri\": \"Ipsam iusto.\",\n         \"post_logout_redirect_uris\": [\n            \"Totam doloremque enim ad minima.\",\n            \"Quia ut.\"\n         ],\n         \"public\": true,\n         \"redirect_uris\": [\n            \"Delectus alias unde voluptas reprehenderit.\",\n            \"Id reiciendis nesciunt ut soluta asperiores.\"\n         ],\n         \"refresh_token_expiry_seconds\": 2471342395392452438,\n         \"requires_consent\": true,\n         \"requires_pkce\": true,\n         \"token_expiry_seconds\": 3283658336607788687\n      }\n   }'")
		}
		if body.Client == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("client", "body"))
		}
		if err != nil {
			return nil, err
		}
	}
	var id string
	{
		id = oauthProviderUpdateClientID
	}
	var jwt *string
	{
		if oauthProviderUpdateClientJWT != "" {
			jwt = &oauthProviderUpdateClientJWT
		}
	}
	v := &oauthprovider.UpdateClientPayload{}
	if body.Client != nil {
		v.Client = marshalUpdateOAuthClientRequestRequestBodyToOauthproviderUpdateOAuthClientRequest(body.Client)
	}
	v.ID = id
	v.JWT = jwt

	return v, nil
}

// BuildDeleteClientPayload builds the payload for the oauth_provider
// delete_client endpoint from CLI flags.
func BuildDeleteClientPayload(oauthProviderDeleteClientID string, oauthProviderDeleteClientJWT string) (*oauthprovider.DeleteClientPayload, error) {
	var id string
	{
		id = oauthProviderDeleteClientID
	}
	var jwt *string
	{
		if oauthProviderDeleteClientJWT != "" {
			jwt = &oauthProviderDeleteClientJWT
		}
	}
	v := &oauthprovider.DeleteClientPayload{}
	v.ID = id
	v.JWT = jwt

	return v, nil
}

// BuildRotateClientSecretPayload builds the payload for the oauth_provider
// rotate_client_secret endpoint from CLI flags.
func BuildRotateClientSecretPayload(oauthProviderRotateClientSecretID string, oauthProviderRotateClientSecretJWT string) (*oauthprovider.RotateClientSecretPayload, error) {
	var id string
	{
		id = oauthProviderRotateClientSecretID
	}
	var jwt *string
	{
		if oauthProviderRotateClientSecretJWT != "" {
			jwt = &oauthProviderRotateClientSecretJWT
		}
	}
	v := &oauthprovider.RotateClientSecretPayload{}
	v.ID = id
	v.JWT = jwt

	return v, nil
}

// BuildListScopesPayload builds the payload for the oauth_provider list_scopes
// endpoint from CLI flags.
func BuildListScopesPayload(oauthProviderListScopesOffset string, oauthProviderListScopesLimit string, oauthProviderListScopesJWT string) (*oauthprovider.ListScopesPayload, error) {
	var err error
	var offset int
	{
		if oauthProviderListScopesOffset != "" {
			var v int64
			v, err = strconv.ParseInt(oauthProviderListScopesOffset, 10, strconv.IntSize)
			offset = int(v)
			if err != nil {
				return nil, fmt.Errorf("invalid value for offset, must be INT")
			}
			if offset < 0 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("offset", offset, 0, true))
			}
			if err != nil {
				return nil, err
			}
		}
	}
	var limit int
	{
		if oauthProviderListScopesLimit != "" {
			var v int64
			v, err = strconv.ParseInt(oauthProviderListScopesLimit, 10, strconv.IntSize)
			limit = int(v)
			if err != nil {
				return nil, fmt.Errorf("invalid value for limit, must be INT")
			}
			if limit < 1 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("limit", limit, 1, true))
			}
			if limit > 100 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("limit", limit, 100, false))
			}
			if err != nil {
				return nil, err
			}
		}
	}
	var jwt *string
	{
		if oauthProviderListScopesJWT != "" {
			jwt = &oauthProviderListScopesJWT
		}
	}
	v := &oauthprovider.ListScopesPayload{}
	v.Offset = offset
	v.Limit = limit
	v.JWT = jwt

	return v, nil
}

// BuildCreateScopePayload builds the payload for the oauth_provider
// create_scope endpoint from CLI flags.
func BuildCreateScopePayload(oauthProviderCreateScopeBody string, oauthProviderCreateScopeJWT string) (*oauthprovider.CreateScopePayload, error) {
	var err error
	var body CreateScopeRequestBody
	{
		err = json.Unmarshal([]byte(oauthProviderCreateScopeBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"default_scope\": true,\n      \"description\": \"Read user information\",\n      \"name\": \"read:users\",\n      \"public\": true\n   }'")
		}
	}
	var jwt *string
	{
		if oauthProviderCreateScopeJWT != "" {
			jwt = &oauthProviderCreateScopeJWT
		}
	}
	v := &oauthprovider.CreateScopePayload{
		Name:         body.Name,
		Description:  body.Description,
		DefaultScope: body.DefaultScope,
		Public:       body.Public,
	}
	{
		var zero bool
		if v.DefaultScope == zero {
			v.DefaultScope = false
		}
	}
	{
		var zero bool
		if v.Public == zero {
			v.Public = true
		}
	}
	v.JWT = jwt

	return v, nil
}

// BuildGetScopePayload builds the payload for the oauth_provider get_scope
// endpoint from CLI flags.
func BuildGetScopePayload(oauthProviderGetScopeID string, oauthProviderGetScopeJWT string) (*oauthprovider.GetScopePayload, error) {
	var id string
	{
		id = oauthProviderGetScopeID
	}
	var jwt *string
	{
		if oauthProviderGetScopeJWT != "" {
			jwt = &oauthProviderGetScopeJWT
		}
	}
	v := &oauthprovider.GetScopePayload{}
	v.ID = id
	v.JWT = jwt

	return v, nil
}

// BuildUpdateScopePayload builds the payload for the oauth_provider
// update_scope endpoint from CLI flags.
func BuildUpdateScopePayload(oauthProviderUpdateScopeBody string, oauthProviderUpdateScopeID string, oauthProviderUpdateScopeJWT string) (*oauthprovider.UpdateScopePayload, error) {
	var err error
	var body UpdateScopeRequestBody
	{
		err = json.Unmarshal([]byte(oauthProviderUpdateScopeBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"default_scope\": true,\n      \"description\": \"Ut provident est vero.\",\n      \"public\": false\n   }'")
		}
	}
	var id string
	{
		id = oauthProviderUpdateScopeID
	}
	var jwt *string
	{
		if oauthProviderUpdateScopeJWT != "" {
			jwt = &oauthProviderUpdateScopeJWT
		}
	}
	v := &oauthprovider.UpdateScopePayload{
		Description:  body.Description,
		DefaultScope: body.DefaultScope,
		Public:       body.Public,
	}
	v.ID = id
	v.JWT = jwt

	return v, nil
}

// BuildDeleteScopePayload builds the payload for the oauth_provider
// delete_scope endpoint from CLI flags.
func BuildDeleteScopePayload(oauthProviderDeleteScopeID string, oauthProviderDeleteScopeJWT string) (*oauthprovider.DeleteScopePayload, error) {
	var id string
	{
		id = oauthProviderDeleteScopeID
	}
	var jwt *string
	{
		if oauthProviderDeleteScopeJWT != "" {
			jwt = &oauthProviderDeleteScopeJWT
		}
	}
	v := &oauthprovider.DeleteScopePayload{}
	v.ID = id
	v.JWT = jwt

	return v, nil
}

// Code generated by goa v3.20.0, DO NOT EDIT.
//
// passwordless HTTP client CLI support package
//
// Command:
// $ goa gen github.com/juicycleff/frank/design -o .

package client

import (
	"encoding/json"
	"fmt"

	passwordless "github.com/juicycleff/frank/gen/passwordless"
	goa "goa.design/goa/v3/pkg"
)

// BuildEmailPayload builds the payload for the passwordless email endpoint
// from CLI flags.
func BuildEmailPayload(passwordlessEmailBody string, passwordlessEmailOauth2 string, passwordlessEmailXAPIKey string, passwordlessEmailJWT string) (*passwordless.EmailPayload, error) {
	var err error
	var body EmailRequestBody
	{
		err = json.Unmarshal([]byte(passwordlessEmailBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"email\": \"user@example.com\",\n      \"redirect_url\": \"Rerum nesciunt qui.\"\n   }'")
		}
		err = goa.MergeErrors(err, goa.ValidateFormat("body.email", body.Email, goa.FormatEmail))
		if err != nil {
			return nil, err
		}
	}
	var oauth2 *string
	{
		if passwordlessEmailOauth2 != "" {
			oauth2 = &passwordlessEmailOauth2
		}
	}
	var xAPIKey *string
	{
		if passwordlessEmailXAPIKey != "" {
			xAPIKey = &passwordlessEmailXAPIKey
		}
	}
	var jwt *string
	{
		if passwordlessEmailJWT != "" {
			jwt = &passwordlessEmailJWT
		}
	}
	v := &passwordless.EmailPayload{
		Email:       body.Email,
		RedirectURL: body.RedirectURL,
	}
	v.Oauth2 = oauth2
	v.XAPIKey = xAPIKey
	v.JWT = jwt

	return v, nil
}

// BuildSmsPayload builds the payload for the passwordless sms endpoint from
// CLI flags.
func BuildSmsPayload(passwordlessSmsBody string, passwordlessSmsOauth2 string, passwordlessSmsXAPIKey string, passwordlessSmsJWT string) (*passwordless.SmsPayload, error) {
	var err error
	var body SmsRequestBody
	{
		err = json.Unmarshal([]byte(passwordlessSmsBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"phone_number\": \"+12345678901\",\n      \"redirect_url\": \"Veniam suscipit est esse ut omnis est.\"\n   }'")
		}
	}
	var oauth2 *string
	{
		if passwordlessSmsOauth2 != "" {
			oauth2 = &passwordlessSmsOauth2
		}
	}
	var xAPIKey *string
	{
		if passwordlessSmsXAPIKey != "" {
			xAPIKey = &passwordlessSmsXAPIKey
		}
	}
	var jwt *string
	{
		if passwordlessSmsJWT != "" {
			jwt = &passwordlessSmsJWT
		}
	}
	v := &passwordless.SmsPayload{
		PhoneNumber: body.PhoneNumber,
		RedirectURL: body.RedirectURL,
	}
	v.Oauth2 = oauth2
	v.XAPIKey = xAPIKey
	v.JWT = jwt

	return v, nil
}

// BuildVerifyPayload builds the payload for the passwordless verify endpoint
// from CLI flags.
func BuildVerifyPayload(passwordlessVerifyBody string, passwordlessVerifyOauth2 string, passwordlessVerifyXAPIKey string, passwordlessVerifyJWT string) (*passwordless.VerifyPayload, error) {
	var err error
	var body VerifyRequestBody
	{
		err = json.Unmarshal([]byte(passwordlessVerifyBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"auth_type\": \"email\",\n      \"code\": \"Dolorum sequi consequuntur rerum corrupti aliquid.\",\n      \"phone_number\": \"Voluptate et qui itaque enim quo aut.\",\n      \"token\": \"Voluptatibus esse.\"\n   }'")
		}
		if !(body.AuthType == "email" || body.AuthType == "sms") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.auth_type", body.AuthType, []any{"email", "sms"}))
		}
		if err != nil {
			return nil, err
		}
	}
	var oauth2 *string
	{
		if passwordlessVerifyOauth2 != "" {
			oauth2 = &passwordlessVerifyOauth2
		}
	}
	var xAPIKey *string
	{
		if passwordlessVerifyXAPIKey != "" {
			xAPIKey = &passwordlessVerifyXAPIKey
		}
	}
	var jwt *string
	{
		if passwordlessVerifyJWT != "" {
			jwt = &passwordlessVerifyJWT
		}
	}
	v := &passwordless.VerifyPayload{
		Token:       body.Token,
		PhoneNumber: body.PhoneNumber,
		Code:        body.Code,
		AuthType:    body.AuthType,
	}
	v.Oauth2 = oauth2
	v.XAPIKey = xAPIKey
	v.JWT = jwt

	return v, nil
}

// BuildMethodsPayload builds the payload for the passwordless methods endpoint
// from CLI flags.
func BuildMethodsPayload(passwordlessMethodsOauth2 string, passwordlessMethodsXAPIKey string, passwordlessMethodsJWT string) (*passwordless.MethodsPayload, error) {
	var oauth2 *string
	{
		if passwordlessMethodsOauth2 != "" {
			oauth2 = &passwordlessMethodsOauth2
		}
	}
	var xAPIKey *string
	{
		if passwordlessMethodsXAPIKey != "" {
			xAPIKey = &passwordlessMethodsXAPIKey
		}
	}
	var jwt *string
	{
		if passwordlessMethodsJWT != "" {
			jwt = &passwordlessMethodsJWT
		}
	}
	v := &passwordless.MethodsPayload{}
	v.Oauth2 = oauth2
	v.XAPIKey = xAPIKey
	v.JWT = jwt

	return v, nil
}

// BuildMagicLinkPayload builds the payload for the passwordless magic_link
// endpoint from CLI flags.
func BuildMagicLinkPayload(passwordlessMagicLinkBody string, passwordlessMagicLinkJWT string) (*passwordless.MagicLinkPayload, error) {
	var err error
	var body MagicLinkRequestBody
	{
		err = json.Unmarshal([]byte(passwordlessMagicLinkBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"email\": \"user@example.com\",\n      \"expires_in\": 316998,\n      \"redirect_url\": \"https://example.com/dashboard\",\n      \"user_id\": \"usr_123456789\"\n   }'")
		}
		err = goa.MergeErrors(err, goa.ValidateFormat("body.email", body.Email, goa.FormatEmail))
		if body.ExpiresIn < 60 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.expires_in", body.ExpiresIn, 60, true))
		}
		if body.ExpiresIn > 604800 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.expires_in", body.ExpiresIn, 604800, false))
		}
		if err != nil {
			return nil, err
		}
	}
	var jwt *string
	{
		if passwordlessMagicLinkJWT != "" {
			jwt = &passwordlessMagicLinkJWT
		}
	}
	v := &passwordless.MagicLinkPayload{
		Email:       body.Email,
		UserID:      body.UserID,
		RedirectURL: body.RedirectURL,
		ExpiresIn:   body.ExpiresIn,
	}
	{
		var zero int
		if v.ExpiresIn == zero {
			v.ExpiresIn = 86400
		}
	}
	v.JWT = jwt

	return v, nil
}

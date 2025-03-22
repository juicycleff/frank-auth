// Code generated by goa v3.20.0, DO NOT EDIT.
//
// HTTP request path constructors for the passwordless service.
//
// Command:
// $ goa gen github.com/juicycleff/frank/design -o .

package client

// EmailPasswordlessPath returns the URL path to the passwordless service email HTTP endpoint.
func EmailPasswordlessPath() string {
	return "/v1/auth/passwordless/email"
}

// SmsPasswordlessPath returns the URL path to the passwordless service sms HTTP endpoint.
func SmsPasswordlessPath() string {
	return "/v1/auth/passwordless/sms"
}

// VerifyPasswordlessPath returns the URL path to the passwordless service verify HTTP endpoint.
func VerifyPasswordlessPath() string {
	return "/v1/auth/passwordless/verify"
}

// MethodsPasswordlessPath returns the URL path to the passwordless service methods HTTP endpoint.
func MethodsPasswordlessPath() string {
	return "/v1/auth/passwordless/methods"
}

// MagicLinkPasswordlessPath returns the URL path to the passwordless service magic_link HTTP endpoint.
func MagicLinkPasswordlessPath() string {
	return "/v1/auth/passwordless/magic-link"
}

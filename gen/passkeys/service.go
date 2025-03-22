// Code generated by goa v3.20.0, DO NOT EDIT.
//
// passkeys service
//
// Command:
// $ goa gen github.com/juicycleff/frank/design -o .

package passkeys

import (
	"context"

	"goa.design/goa/v3/security"
)

// Passkey (WebAuthn) authentication service
type Service interface {
	// Begin passkey registration
	RegisterBegin(context.Context, *RegisterBeginPayload) (res *RegisterBeginResult, err error)
	// Complete passkey registration
	RegisterComplete(context.Context, *RegisterCompletePayload) (res *RegisteredPasskey, err error)
	// Begin passkey authentication
	LoginBegin(context.Context, *LoginBeginPayload) (res *LoginBeginResult, err error)
	// Complete passkey authentication
	LoginComplete(context.Context, *LoginCompletePayload) (res *LoginCompleteResult, err error)
	// List registered passkeys
	List(context.Context, *ListPayload) (res *ListResult, err error)
	// Update passkey
	Update(context.Context, *UpdatePayload) (res *UpdateResult, err error)
	// Delete passkey
	Delete(context.Context, *DeletePayload) (err error)
}

// Auther defines the authorization functions to be implemented by the service.
type Auther interface {
	// JWTAuth implements the authorization logic for the JWT security scheme.
	JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error)
	// OAuth2Auth implements the authorization logic for the OAuth2 security scheme.
	OAuth2Auth(ctx context.Context, token string, schema *security.OAuth2Scheme) (context.Context, error)
	// APIKeyAuth implements the authorization logic for the APIKey security scheme.
	APIKeyAuth(ctx context.Context, key string, schema *security.APIKeyScheme) (context.Context, error)
}

// APIName is the name of the API as defined in the design.
const APIName = "frank"

// APIVersion is the version of the API as defined in the design.
const APIVersion = "1.0.0"

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "passkeys"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [7]string{"register_begin", "register_complete", "login_begin", "login_complete", "list", "update", "delete"}

// Bad request response
type BadRequestError struct {
	// Error code
	Code string
	// Error message
	Message string
	// Additional error details
	Details any
	// Unique error ID
	ID *string
}

// Conflict response
type ConflictError struct {
	// Error code
	Code string
	// Error message
	Message string
	// Additional error details
	Details any
	// Unique error ID
	ID *string
}

// DeletePayload is the payload type of the passkeys service delete method.
type DeletePayload struct {
	JWT *string
	// Passkey ID
	ID string
}

// Forbidden response
type ForbiddenError struct {
	// Error code
	Code string
	// Error message
	Message string
	// Additional error details
	Details any
	// Unique error ID
	ID *string
}

// Internal server error response
type InternalServerError struct {
	// Error code
	Code string
	// Error message
	Message string
	// Additional error details
	Details any
	// Unique error ID
	ID *string
}

// ListPayload is the payload type of the passkeys service list method.
type ListPayload struct {
	JWT *string
}

// ListResult is the result type of the passkeys service list method.
type ListResult struct {
	Passkeys []*RegisteredPasskey
}

// LoginBeginPayload is the payload type of the passkeys service login_begin
// method.
type LoginBeginPayload struct {
	// OAuth2 access token
	Oauth2 *string
	// API key
	XAPIKey *string
	// JWT token
	JWT *string
}

// LoginBeginResult is the result type of the passkeys service login_begin
// method.
type LoginBeginResult struct {
	// WebAuthn credential request options
	Options any
	// Authentication session ID
	SessionID string
}

// LoginCompletePayload is the payload type of the passkeys service
// login_complete method.
type LoginCompletePayload struct {
	// OAuth2 access token
	Oauth2 *string
	// API key
	XAPIKey *string
	// JWT token
	JWT *string
	// Authentication session ID
	SessionID string
	// WebAuthn assertion response
	Response any
}

// LoginCompleteResult is the result type of the passkeys service
// login_complete method.
type LoginCompleteResult struct {
	// Whether authentication was successful
	Authenticated bool
	// User ID
	UserID string
}

// Not found response
type NotFoundError struct {
	// Error code
	Code string
	// Error message
	Message string
	// Additional error details
	Details any
	// Unique error ID
	ID *string
}

// RegisterBeginPayload is the payload type of the passkeys service
// register_begin method.
type RegisterBeginPayload struct {
	JWT *string
	// Name of the device
	DeviceName *string
	// Type of the device
	DeviceType *string
}

// RegisterBeginResult is the result type of the passkeys service
// register_begin method.
type RegisterBeginResult struct {
	// WebAuthn credential creation options
	Options any
	// Registration session ID
	SessionID string
}

// RegisterCompletePayload is the payload type of the passkeys service
// register_complete method.
type RegisterCompletePayload struct {
	JWT *string
	// Registration session ID
	SessionID string
	// WebAuthn credential creation response
	Response any
	// Name of the device
	DeviceName *string
	// Type of the device
	DeviceType *string
}

// RegisteredPasskey is the result type of the passkeys service
// register_complete method.
type RegisteredPasskey struct {
	// Passkey ID
	ID string
	// Passkey name
	Name string
	// Device type
	DeviceType string
	// Registration timestamp
	RegisteredAt string
	// Last usage timestamp
	LastUsed *string
}

// Unauthorized response
type UnauthorizedError struct {
	// Error code
	Code string
	// Error message
	Message string
	// Additional error details
	Details any
	// Unique error ID
	ID *string
}

// Update passkey request
type UpdatePasskeyRequest struct {
	// New passkey name
	Name string
}

// UpdatePayload is the payload type of the passkeys service update method.
type UpdatePayload struct {
	JWT *string
	// Passkey ID
	ID      string
	Request *UpdatePasskeyRequest
}

// UpdateResult is the result type of the passkeys service update method.
type UpdateResult struct {
	// Success message
	Message string
}

// Error returns an error description.
func (e *BadRequestError) Error() string {
	return "Bad request response"
}

// ErrorName returns "BadRequestError".
//
// Deprecated: Use GoaErrorName - https://github.com/goadesign/goa/issues/3105
func (e *BadRequestError) ErrorName() string {
	return e.GoaErrorName()
}

// GoaErrorName returns "BadRequestError".
func (e *BadRequestError) GoaErrorName() string {
	return "bad_request"
}

// Error returns an error description.
func (e *ConflictError) Error() string {
	return "Conflict response"
}

// ErrorName returns "ConflictError".
//
// Deprecated: Use GoaErrorName - https://github.com/goadesign/goa/issues/3105
func (e *ConflictError) ErrorName() string {
	return e.GoaErrorName()
}

// GoaErrorName returns "ConflictError".
func (e *ConflictError) GoaErrorName() string {
	return "conflict"
}

// Error returns an error description.
func (e *ForbiddenError) Error() string {
	return "Forbidden response"
}

// ErrorName returns "ForbiddenError".
//
// Deprecated: Use GoaErrorName - https://github.com/goadesign/goa/issues/3105
func (e *ForbiddenError) ErrorName() string {
	return e.GoaErrorName()
}

// GoaErrorName returns "ForbiddenError".
func (e *ForbiddenError) GoaErrorName() string {
	return "forbidden"
}

// Error returns an error description.
func (e *InternalServerError) Error() string {
	return "Internal server error response"
}

// ErrorName returns "InternalServerError".
//
// Deprecated: Use GoaErrorName - https://github.com/goadesign/goa/issues/3105
func (e *InternalServerError) ErrorName() string {
	return e.GoaErrorName()
}

// GoaErrorName returns "InternalServerError".
func (e *InternalServerError) GoaErrorName() string {
	return "internal_error"
}

// Error returns an error description.
func (e *NotFoundError) Error() string {
	return "Not found response"
}

// ErrorName returns "NotFoundError".
//
// Deprecated: Use GoaErrorName - https://github.com/goadesign/goa/issues/3105
func (e *NotFoundError) ErrorName() string {
	return e.GoaErrorName()
}

// GoaErrorName returns "NotFoundError".
func (e *NotFoundError) GoaErrorName() string {
	return "not_found"
}

// Error returns an error description.
func (e *UnauthorizedError) Error() string {
	return "Unauthorized response"
}

// ErrorName returns "UnauthorizedError".
//
// Deprecated: Use GoaErrorName - https://github.com/goadesign/goa/issues/3105
func (e *UnauthorizedError) ErrorName() string {
	return e.GoaErrorName()
}

// GoaErrorName returns "UnauthorizedError".
func (e *UnauthorizedError) GoaErrorName() string {
	return "unauthorized"
}

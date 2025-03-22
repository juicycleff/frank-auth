// Code generated by goa v3.20.0, DO NOT EDIT.
//
// users service
//
// Command:
// $ goa gen github.com/juicycleff/frank/design -o .

package users

import (
	"context"

	designtypes "github.com/juicycleff/frank/gen/designtypes"
	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// User management service
type Service interface {
	// List users
	List(context.Context, *ListPayload) (res *ListResult, err error)
	// Create a new user
	Create(context.Context, *CreatePayload) (res *User, err error)
	// Get user by ID
	Get(context.Context, *GetPayload) (res *User, err error)
	// Update user
	Update(context.Context, *UpdatePayload) (res *User, err error)
	// Delete user
	Delete(context.Context, *DeletePayload) (err error)
	// Update current user
	UpdateMe(context.Context, *UpdateMePayload) (res *User, err error)
	// Update current user password
	UpdatePassword(context.Context, *UpdatePasswordPayload) (res *UpdatePasswordResult, err error)
	// Get current user sessions
	GetSessions(context.Context, *GetSessionsPayload) (res *GetSessionsResult, err error)
	// Delete user session
	DeleteSession(context.Context, *DeleteSessionPayload) (err error)
	// Get user organizations
	GetOrganizations(context.Context, *GetOrganizationsPayload) (res *GetOrganizationsResult, err error)
}

// Auther defines the authorization functions to be implemented by the service.
type Auther interface {
	// JWTAuth implements the authorization logic for the JWT security scheme.
	JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error)
}

// APIName is the name of the API as defined in the design.
const APIName = "frank"

// APIVersion is the version of the API as defined in the design.
const APIVersion = "1.0.0"

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "users"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [10]string{"list", "create", "get", "update", "delete", "update_me", "update_password", "get_sessions", "delete_session", "get_organizations"}

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

// CreatePayload is the payload type of the users service create method.
type CreatePayload struct {
	JWT *string
	// User email
	Email string
	// User password
	Password *string
	// User phone number
	PhoneNumber *string
	// User first name
	FirstName *string
	// User last name
	LastName *string
	// User metadata
	Metadata map[string]any
	// Profile image URL
	ProfileImageURL *string
	// User locale
	Locale string
	// Organization ID to add user to
	OrganizationID *string
}

// DeletePayload is the payload type of the users service delete method.
type DeletePayload struct {
	JWT *string
	// User ID
	ID string
}

// DeleteSessionPayload is the payload type of the users service delete_session
// method.
type DeleteSessionPayload struct {
	JWT *string
	// Session ID
	SessionID string
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

// GetOrganizationsPayload is the payload type of the users service
// get_organizations method.
type GetOrganizationsPayload struct {
	JWT *string
	// User ID
	ID string
}

// GetOrganizationsResult is the result type of the users service
// get_organizations method.
type GetOrganizationsResult struct {
	Organizations []*OrganizationResponse
}

// GetPayload is the payload type of the users service get method.
type GetPayload struct {
	JWT *string
	// User ID
	ID string
}

// GetSessionsPayload is the payload type of the users service get_sessions
// method.
type GetSessionsPayload struct {
	JWT *string
}

// GetSessionsResult is the result type of the users service get_sessions
// method.
type GetSessionsResult struct {
	Sessions []*designtypes.UserSessionResponse
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

// ListPayload is the payload type of the users service list method.
type ListPayload struct {
	JWT *string
	// Pagination offset
	Offset int
	// Number of items to return
	Limit int
	// Search term
	Search *string
	// Filter by organization ID
	OrganizationID *string
}

// ListResult is the result type of the users service list method.
type ListResult struct {
	Data       []*User
	Pagination *designtypes.PaginationResponse
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

// Organization information
type OrganizationResponse struct {
	// Organization ID
	ID string
	// Organization name
	Name string
	// Organization slug
	Slug string
	// Organization domain
	Domain *string
	// Organization logo URL
	LogoURL *string
	// Organization plan
	Plan *string
	// Whether organization is active
	Active bool
	// Organization metadata
	Metadata map[string]any
	// Trial end date
	TrialEndsAt *string
	// Whether trial has been used
	TrialUsed *bool
	// Creation timestamp
	CreatedAt string
	// Last update timestamp
	UpdatedAt string
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

// UpdateMePayload is the payload type of the users service update_me method.
type UpdateMePayload struct {
	JWT *string
	// User phone number
	PhoneNumber *string
	// User first name
	FirstName *string
	// User last name
	LastName *string
	// User metadata
	Metadata map[string]any
	// Profile image URL
	ProfileImageURL *string
	// User locale
	Locale *string
	// Whether user is active
	Active *bool
	// Primary organization ID
	PrimaryOrganizationID *string
}

// UpdatePasswordPayload is the payload type of the users service
// update_password method.
type UpdatePasswordPayload struct {
	JWT *string
	// Current password
	CurrentPassword string
	// New password
	NewPassword string
}

// UpdatePasswordResult is the result type of the users service update_password
// method.
type UpdatePasswordResult struct {
	// Success message
	Message string
}

// UpdatePayload is the payload type of the users service update method.
type UpdatePayload struct {
	JWT *string
	// User ID
	ID   string
	User *designtypes.UpdateUserRequest
}

// User is the result type of the users service create method.
type User struct {
	// User ID
	ID string
	// User email
	Email string
	// User first name
	FirstName *string
	// User last name
	LastName *string
	// Whether email is verified
	EmailVerified bool
	// User phone number
	PhoneNumber *string
	// Whether phone is verified
	PhoneVerified *bool
	// URL to user's profile image
	ProfileImageURL *string
	// User's locale preference
	Locale *string
	// User metadata
	Metadata map[string]any
	// Whether account is active
	Active bool
	// Account creation timestamp
	CreatedAt string
	// Account last update timestamp
	UpdatedAt string
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

// MakeNotFound builds a goa.ServiceError from an error.
func MakeNotFound(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "not_found", false, false, false)
}

// MakeUnauthorized builds a goa.ServiceError from an error.
func MakeUnauthorized(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "unauthorized", false, false, false)
}

// MakeForbidden builds a goa.ServiceError from an error.
func MakeForbidden(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "forbidden", false, false, false)
}

// MakeBadRequest builds a goa.ServiceError from an error.
func MakeBadRequest(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "bad_request", false, false, false)
}

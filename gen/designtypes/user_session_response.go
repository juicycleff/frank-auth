// Code generated by goa v3.20.0, DO NOT EDIT.
//
// User types
//
// Command:
// $ goa gen github.com/juicycleff/frank/design -o .

package designtypes

// User session information
type UserSessionResponse struct {
	// Session ID
	ID string
	// Device ID
	DeviceID *string
	// IP address
	IPAddress *string
	// User agent string
	UserAgent *string
	// Location
	Location *string
	// Last activity timestamp
	LastActiveAt *string
	// Creation timestamp
	CreatedAt string
}

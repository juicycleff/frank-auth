// Code generated by goa v3.20.0, DO NOT EDIT.
//
// User types
//
// Command:
// $ goa gen github.com/juicycleff/frank/design -o .

package designtypes

// HealthResponse is the result type of the health service check method.
type HealthResponse struct {
	// Overall health status
	Status string
	// Timestamp of health check
	Timestamp string
	// Status of individual services
	Services []*HealthStatus
}

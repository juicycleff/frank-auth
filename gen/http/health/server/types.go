// Code generated by goa v3.20.0, DO NOT EDIT.
//
// health HTTP server types
//
// Command:
// $ goa gen github.com/juicycleff/frank/design -o .

package server

import (
	designtypes "github.com/juicycleff/frank/gen/designtypes"
	health "github.com/juicycleff/frank/gen/health"
)

// CheckOKResponseBody is the type of the "health" service "check" endpoint
// HTTP response body.
type CheckOKResponseBody struct {
	// Overall health status
	Status string `form:"status" json:"status" xml:"status"`
	// Timestamp of health check
	Timestamp string `form:"timestamp" json:"timestamp" xml:"timestamp"`
	// Status of individual services
	Services []*HealthStatusResponseBody `form:"services,omitempty" json:"services,omitempty" xml:"services,omitempty"`
}

// ReadyOKResponseBody is the type of the "health" service "ready" endpoint
// HTTP response body.
type ReadyOKResponseBody struct {
	// Readiness status
	Status string `form:"status" json:"status" xml:"status"`
	// Timestamp of health check
	Timestamp string `form:"timestamp" json:"timestamp" xml:"timestamp"`
}

// VersionResponseBody is the type of the "health" service "version" endpoint
// HTTP response body.
type VersionResponseBody struct {
	// System version
	Version string `form:"version" json:"version" xml:"version"`
	// Build date
	BuildDate string `form:"build_date" json:"build_date" xml:"build_date"`
	// Git commit hash
	GitCommit *string `form:"git_commit,omitempty" json:"git_commit,omitempty" xml:"git_commit,omitempty"`
	// Go version
	GoVersion *string `form:"go_version,omitempty" json:"go_version,omitempty" xml:"go_version,omitempty"`
}

// MetricsResponseBody is the type of the "health" service "metrics" endpoint
// HTTP response body.
type MetricsResponseBody struct {
	// System uptime in seconds
	Uptime int64 `form:"uptime" json:"uptime" xml:"uptime"`
	// Memory usage in bytes
	MemoryUsage int64 `form:"memory_usage" json:"memory_usage" xml:"memory_usage"`
	// Number of goroutines
	Goroutines int `form:"goroutines" json:"goroutines" xml:"goroutines"`
	// Total request count
	Requests *int64 `form:"requests,omitempty" json:"requests,omitempty" xml:"requests,omitempty"`
	// Total error count
	Errors *int64 `form:"errors,omitempty" json:"errors,omitempty" xml:"errors,omitempty"`
	// Requests per second
	RequestRate *float32 `form:"request_rate,omitempty" json:"request_rate,omitempty" xml:"request_rate,omitempty"`
}

// HealthStatusResponseBody is used to define fields on response body types.
type HealthStatusResponseBody struct {
	// Service name
	Service string `form:"service" json:"service" xml:"service"`
	// Service status
	Status string `form:"status" json:"status" xml:"status"`
	// Additional message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// NewCheckOKResponseBody builds the HTTP response body from the result of the
// "check" endpoint of the "health" service.
func NewCheckOKResponseBody(res *designtypes.HealthResponse) *CheckOKResponseBody {
	body := &CheckOKResponseBody{
		Status:    res.Status,
		Timestamp: res.Timestamp,
	}
	if res.Services != nil {
		body.Services = make([]*HealthStatusResponseBody, len(res.Services))
		for i, val := range res.Services {
			body.Services[i] = marshalDesigntypesHealthStatusToHealthStatusResponseBody(val)
		}
	}
	return body
}

// NewReadyOKResponseBody builds the HTTP response body from the result of the
// "ready" endpoint of the "health" service.
func NewReadyOKResponseBody(res *designtypes.ReadyResponse) *ReadyOKResponseBody {
	body := &ReadyOKResponseBody{
		Status:    res.Status,
		Timestamp: res.Timestamp,
	}
	return body
}

// NewVersionResponseBody builds the HTTP response body from the result of the
// "version" endpoint of the "health" service.
func NewVersionResponseBody(res *health.VersionResult) *VersionResponseBody {
	body := &VersionResponseBody{
		Version:   res.Version,
		BuildDate: res.BuildDate,
		GitCommit: res.GitCommit,
		GoVersion: res.GoVersion,
	}
	return body
}

// NewMetricsResponseBody builds the HTTP response body from the result of the
// "metrics" endpoint of the "health" service.
func NewMetricsResponseBody(res *health.MetricsResult) *MetricsResponseBody {
	body := &MetricsResponseBody{
		Uptime:      res.Uptime,
		MemoryUsage: res.MemoryUsage,
		Goroutines:  res.Goroutines,
		Requests:    res.Requests,
		Errors:      res.Errors,
		RequestRate: res.RequestRate,
	}
	return body
}

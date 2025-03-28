// Code generated by goa v3.20.0, DO NOT EDIT.
//
// health HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/juicycleff/frank/design -o .

package server

import (
	"context"
	"net/http"

	designtypes "github.com/juicycleff/frank/gen/designtypes"
	health "github.com/juicycleff/frank/gen/health"
	goahttp "goa.design/goa/v3/http"
)

// EncodeCheckResponse returns an encoder for responses returned by the health
// check endpoint.
func EncodeCheckResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.(*designtypes.HealthResponse)
		ctx = context.WithValue(ctx, goahttp.ContentTypeKey, "application/json")
		enc := encoder(ctx, w)
		body := NewCheckOKResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// EncodeReadyResponse returns an encoder for responses returned by the health
// ready endpoint.
func EncodeReadyResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.(*designtypes.ReadyResponse)
		ctx = context.WithValue(ctx, goahttp.ContentTypeKey, "application/json")
		enc := encoder(ctx, w)
		body := NewReadyOKResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// EncodeVersionResponse returns an encoder for responses returned by the
// health version endpoint.
func EncodeVersionResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.(*health.VersionResult)
		enc := encoder(ctx, w)
		body := NewVersionResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// EncodeMetricsResponse returns an encoder for responses returned by the
// health metrics endpoint.
func EncodeMetricsResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.(*health.MetricsResult)
		enc := encoder(ctx, w)
		body := NewMetricsResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// EncodeDebugResponse returns an encoder for responses returned by the health
// debug endpoint.
func EncodeDebugResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.(any)
		enc := encoder(ctx, w)
		body := res
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// marshalDesigntypesHealthStatusToHealthStatusResponseBody builds a value of
// type *HealthStatusResponseBody from a value of type
// *designtypes.HealthStatus.
func marshalDesigntypesHealthStatusToHealthStatusResponseBody(v *designtypes.HealthStatus) *HealthStatusResponseBody {
	if v == nil {
		return nil
	}
	res := &HealthStatusResponseBody{
		Service: v.Service,
		Status:  v.Status,
		Message: v.Message,
	}

	return res
}

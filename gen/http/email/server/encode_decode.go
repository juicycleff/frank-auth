// Code generated by goa v3.20.0, DO NOT EDIT.
//
// email HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/juicycleff/frank/design -o .

package server

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"

	designtypes "github.com/juicycleff/frank/gen/designtypes"
	email "github.com/juicycleff/frank/gen/email"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeListTemplatesResponse returns an encoder for responses returned by the
// email list_templates endpoint.
func EncodeListTemplatesResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.(*email.ListTemplatesResult)
		enc := encoder(ctx, w)
		body := NewListTemplatesResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeListTemplatesRequest returns a decoder for requests sent to the email
// list_templates endpoint.
func DecodeListTemplatesRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			offset         int
			limit          int
			type_          *string
			organizationID *string
			locale         *string
			jwt            *string
			err            error
		)
		qp := r.URL.Query()
		{
			offsetRaw := qp.Get("offset")
			if offsetRaw != "" {
				v, err2 := strconv.ParseInt(offsetRaw, 10, strconv.IntSize)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("offset", offsetRaw, "integer"))
				}
				offset = int(v)
			}
		}
		if offset < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("offset", offset, 0, true))
		}
		{
			limitRaw := qp.Get("limit")
			if limitRaw == "" {
				limit = 20
			} else {
				v, err2 := strconv.ParseInt(limitRaw, 10, strconv.IntSize)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("limit", limitRaw, "integer"))
				}
				limit = int(v)
			}
		}
		if limit < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("limit", limit, 1, true))
		}
		if limit > 100 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("limit", limit, 100, false))
		}
		type_Raw := qp.Get("type")
		if type_Raw != "" {
			type_ = &type_Raw
		}
		organizationIDRaw := qp.Get("organization_id")
		if organizationIDRaw != "" {
			organizationID = &organizationIDRaw
		}
		localeRaw := qp.Get("locale")
		if localeRaw != "" {
			locale = &localeRaw
		}
		jwtRaw := r.Header.Get("Authorization")
		if jwtRaw != "" {
			jwt = &jwtRaw
		}
		if err != nil {
			return nil, err
		}
		payload := NewListTemplatesPayload(offset, limit, type_, organizationID, locale, jwt)
		if payload.JWT != nil {
			if strings.Contains(*payload.JWT, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.JWT, " ", 2)[1]
				payload.JWT = &cred
			}
		}

		return payload, nil
	}
}

// EncodeListTemplatesError returns an encoder for errors returned by the
// list_templates email endpoint.
func EncodeListTemplatesError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "bad_request":
			var res *email.BadRequestError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewListTemplatesBadRequestResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "forbidden":
			var res *email.ForbiddenError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewListTemplatesForbiddenResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusForbidden)
			return enc.Encode(body)
		case "internal_error":
			var res *email.InternalServerError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewListTemplatesInternalErrorResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		case "not_found":
			var res *email.NotFoundError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewListTemplatesNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "unauthorized":
			var res *email.UnauthorizedError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewListTemplatesUnauthorizedResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeCreateTemplateResponse returns an encoder for responses returned by
// the email create_template endpoint.
func EncodeCreateTemplateResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.(*email.EmailTemplateResponse)
		enc := encoder(ctx, w)
		body := NewCreateTemplateResponseBody(res)
		w.WriteHeader(http.StatusCreated)
		return enc.Encode(body)
	}
}

// DecodeCreateTemplateRequest returns a decoder for requests sent to the email
// create_template endpoint.
func DecodeCreateTemplateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			body CreateTemplateRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			var gerr *goa.ServiceError
			if errors.As(err, &gerr) {
				return nil, gerr
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateCreateTemplateRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			jwt *string
		)
		jwtRaw := r.Header.Get("Authorization")
		if jwtRaw != "" {
			jwt = &jwtRaw
		}
		payload := NewCreateTemplatePayload(&body, jwt)
		if payload.JWT != nil {
			if strings.Contains(*payload.JWT, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.JWT, " ", 2)[1]
				payload.JWT = &cred
			}
		}

		return payload, nil
	}
}

// EncodeCreateTemplateError returns an encoder for errors returned by the
// create_template email endpoint.
func EncodeCreateTemplateError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "bad_request":
			var res *email.BadRequestError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewCreateTemplateBadRequestResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "forbidden":
			var res *email.ForbiddenError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewCreateTemplateForbiddenResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusForbidden)
			return enc.Encode(body)
		case "internal_error":
			var res *email.InternalServerError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewCreateTemplateInternalErrorResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		case "not_found":
			var res *email.NotFoundError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewCreateTemplateNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "unauthorized":
			var res *email.UnauthorizedError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewCreateTemplateUnauthorizedResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeGetTemplateResponse returns an encoder for responses returned by the
// email get_template endpoint.
func EncodeGetTemplateResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.(*email.EmailTemplateResponse)
		enc := encoder(ctx, w)
		body := NewGetTemplateResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeGetTemplateRequest returns a decoder for requests sent to the email
// get_template endpoint.
func DecodeGetTemplateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			id  string
			jwt *string

			params = mux.Vars(r)
		)
		id = params["id"]
		jwtRaw := r.Header.Get("Authorization")
		if jwtRaw != "" {
			jwt = &jwtRaw
		}
		payload := NewGetTemplatePayload(id, jwt)
		if payload.JWT != nil {
			if strings.Contains(*payload.JWT, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.JWT, " ", 2)[1]
				payload.JWT = &cred
			}
		}

		return payload, nil
	}
}

// EncodeGetTemplateError returns an encoder for errors returned by the
// get_template email endpoint.
func EncodeGetTemplateError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "bad_request":
			var res *email.BadRequestError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewGetTemplateBadRequestResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "forbidden":
			var res *email.ForbiddenError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewGetTemplateForbiddenResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusForbidden)
			return enc.Encode(body)
		case "internal_error":
			var res *email.InternalServerError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewGetTemplateInternalErrorResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		case "not_found":
			var res *email.NotFoundError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewGetTemplateNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "unauthorized":
			var res *email.UnauthorizedError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewGetTemplateUnauthorizedResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeGetTemplateByTypeResponse returns an encoder for responses returned by
// the email get_template_by_type endpoint.
func EncodeGetTemplateByTypeResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.(*email.EmailTemplateResponse)
		enc := encoder(ctx, w)
		body := NewGetTemplateByTypeResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeGetTemplateByTypeRequest returns a decoder for requests sent to the
// email get_template_by_type endpoint.
func DecodeGetTemplateByTypeRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			type_          string
			organizationID *string
			locale         string
			jwt            *string

			params = mux.Vars(r)
		)
		type_ = params["type"]
		qp := r.URL.Query()
		organizationIDRaw := qp.Get("organization_id")
		if organizationIDRaw != "" {
			organizationID = &organizationIDRaw
		}
		localeRaw := qp.Get("locale")
		if localeRaw != "" {
			locale = localeRaw
		} else {
			locale = "en"
		}
		jwtRaw := r.Header.Get("Authorization")
		if jwtRaw != "" {
			jwt = &jwtRaw
		}
		payload := NewGetTemplateByTypePayload(type_, organizationID, locale, jwt)
		if payload.JWT != nil {
			if strings.Contains(*payload.JWT, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.JWT, " ", 2)[1]
				payload.JWT = &cred
			}
		}

		return payload, nil
	}
}

// EncodeGetTemplateByTypeError returns an encoder for errors returned by the
// get_template_by_type email endpoint.
func EncodeGetTemplateByTypeError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "bad_request":
			var res *email.BadRequestError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewGetTemplateByTypeBadRequestResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "forbidden":
			var res *email.ForbiddenError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewGetTemplateByTypeForbiddenResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusForbidden)
			return enc.Encode(body)
		case "internal_error":
			var res *email.InternalServerError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewGetTemplateByTypeInternalErrorResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		case "not_found":
			var res *email.NotFoundError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewGetTemplateByTypeNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "unauthorized":
			var res *email.UnauthorizedError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewGetTemplateByTypeUnauthorizedResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeUpdateTemplateResponse returns an encoder for responses returned by
// the email update_template endpoint.
func EncodeUpdateTemplateResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.(*email.EmailTemplateResponse)
		enc := encoder(ctx, w)
		body := NewUpdateTemplateResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeUpdateTemplateRequest returns a decoder for requests sent to the email
// update_template endpoint.
func DecodeUpdateTemplateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			body UpdateTemplateRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			var gerr *goa.ServiceError
			if errors.As(err, &gerr) {
				return nil, gerr
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateUpdateTemplateRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			id  string
			jwt *string

			params = mux.Vars(r)
		)
		id = params["id"]
		jwtRaw := r.Header.Get("Authorization")
		if jwtRaw != "" {
			jwt = &jwtRaw
		}
		payload := NewUpdateTemplatePayload(&body, id, jwt)
		if payload.JWT != nil {
			if strings.Contains(*payload.JWT, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.JWT, " ", 2)[1]
				payload.JWT = &cred
			}
		}

		return payload, nil
	}
}

// EncodeUpdateTemplateError returns an encoder for errors returned by the
// update_template email endpoint.
func EncodeUpdateTemplateError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "bad_request":
			var res *email.BadRequestError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewUpdateTemplateBadRequestResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "forbidden":
			var res *email.ForbiddenError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewUpdateTemplateForbiddenResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusForbidden)
			return enc.Encode(body)
		case "internal_error":
			var res *email.InternalServerError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewUpdateTemplateInternalErrorResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		case "not_found":
			var res *email.NotFoundError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewUpdateTemplateNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "unauthorized":
			var res *email.UnauthorizedError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewUpdateTemplateUnauthorizedResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeDeleteTemplateResponse returns an encoder for responses returned by
// the email delete_template endpoint.
func EncodeDeleteTemplateResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		w.WriteHeader(http.StatusNoContent)
		return nil
	}
}

// DecodeDeleteTemplateRequest returns a decoder for requests sent to the email
// delete_template endpoint.
func DecodeDeleteTemplateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			id  string
			jwt *string

			params = mux.Vars(r)
		)
		id = params["id"]
		jwtRaw := r.Header.Get("Authorization")
		if jwtRaw != "" {
			jwt = &jwtRaw
		}
		payload := NewDeleteTemplatePayload(id, jwt)
		if payload.JWT != nil {
			if strings.Contains(*payload.JWT, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.JWT, " ", 2)[1]
				payload.JWT = &cred
			}
		}

		return payload, nil
	}
}

// EncodeDeleteTemplateError returns an encoder for errors returned by the
// delete_template email endpoint.
func EncodeDeleteTemplateError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "bad_request":
			var res *email.BadRequestError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewDeleteTemplateBadRequestResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "forbidden":
			var res *email.ForbiddenError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewDeleteTemplateForbiddenResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusForbidden)
			return enc.Encode(body)
		case "internal_error":
			var res *email.InternalServerError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewDeleteTemplateInternalErrorResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		case "not_found":
			var res *email.NotFoundError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewDeleteTemplateNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "unauthorized":
			var res *email.UnauthorizedError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewDeleteTemplateUnauthorizedResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeSendResponse returns an encoder for responses returned by the email
// send endpoint.
func EncodeSendResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.(*email.SendResult)
		enc := encoder(ctx, w)
		body := NewSendResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeSendRequest returns a decoder for requests sent to the email send
// endpoint.
func DecodeSendRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			body SendRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			var gerr *goa.ServiceError
			if errors.As(err, &gerr) {
				return nil, gerr
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateSendRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			jwt *string
		)
		jwtRaw := r.Header.Get("Authorization")
		if jwtRaw != "" {
			jwt = &jwtRaw
		}
		payload := NewSendPayload(&body, jwt)
		if payload.JWT != nil {
			if strings.Contains(*payload.JWT, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.JWT, " ", 2)[1]
				payload.JWT = &cred
			}
		}

		return payload, nil
	}
}

// EncodeSendError returns an encoder for errors returned by the send email
// endpoint.
func EncodeSendError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "bad_request":
			var res *email.BadRequestError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewSendBadRequestResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "forbidden":
			var res *email.ForbiddenError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewSendForbiddenResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusForbidden)
			return enc.Encode(body)
		case "internal_error":
			var res *email.InternalServerError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewSendInternalErrorResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		case "not_found":
			var res *email.NotFoundError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewSendNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "unauthorized":
			var res *email.UnauthorizedError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewSendUnauthorizedResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeSendTemplateResponse returns an encoder for responses returned by the
// email send_template endpoint.
func EncodeSendTemplateResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.(*email.SendTemplateResult)
		enc := encoder(ctx, w)
		body := NewSendTemplateResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeSendTemplateRequest returns a decoder for requests sent to the email
// send_template endpoint.
func DecodeSendTemplateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			body SendTemplateRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			var gerr *goa.ServiceError
			if errors.As(err, &gerr) {
				return nil, gerr
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateSendTemplateRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			jwt *string
		)
		jwtRaw := r.Header.Get("Authorization")
		if jwtRaw != "" {
			jwt = &jwtRaw
		}
		payload := NewSendTemplatePayload(&body, jwt)
		if payload.JWT != nil {
			if strings.Contains(*payload.JWT, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.JWT, " ", 2)[1]
				payload.JWT = &cred
			}
		}

		return payload, nil
	}
}

// EncodeSendTemplateError returns an encoder for errors returned by the
// send_template email endpoint.
func EncodeSendTemplateError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "bad_request":
			var res *email.BadRequestError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewSendTemplateBadRequestResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "forbidden":
			var res *email.ForbiddenError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewSendTemplateForbiddenResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusForbidden)
			return enc.Encode(body)
		case "internal_error":
			var res *email.InternalServerError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewSendTemplateInternalErrorResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		case "not_found":
			var res *email.NotFoundError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewSendTemplateNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "unauthorized":
			var res *email.UnauthorizedError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewSendTemplateUnauthorizedResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// marshalEmailEmailTemplateResponseToEmailTemplateResponseResponseBody builds
// a value of type *EmailTemplateResponseResponseBody from a value of type
// *email.EmailTemplateResponse.
func marshalEmailEmailTemplateResponseToEmailTemplateResponseResponseBody(v *email.EmailTemplateResponse) *EmailTemplateResponseResponseBody {
	res := &EmailTemplateResponseResponseBody{
		ID:             v.ID,
		Name:           v.Name,
		Subject:        v.Subject,
		Type:           v.Type,
		HTMLContent:    v.HTMLContent,
		TextContent:    v.TextContent,
		OrganizationID: v.OrganizationID,
		Active:         v.Active,
		System:         v.System,
		Locale:         v.Locale,
		CreatedAt:      v.CreatedAt,
		UpdatedAt:      v.UpdatedAt,
	}
	if v.Metadata != nil {
		res.Metadata = make(map[string]any, len(v.Metadata))
		for key, val := range v.Metadata {
			tk := key
			tv := val
			res.Metadata[tk] = tv
		}
	}

	return res
}

// marshalDesigntypesPaginationResponseToPaginationResponseResponseBody builds
// a value of type *PaginationResponseResponseBody from a value of type
// *designtypes.PaginationResponse.
func marshalDesigntypesPaginationResponseToPaginationResponseResponseBody(v *designtypes.PaginationResponse) *PaginationResponseResponseBody {
	res := &PaginationResponseResponseBody{
		Total:  v.Total,
		Offset: v.Offset,
		Limit:  v.Limit,
	}

	return res
}

// unmarshalUpdateEmailTemplateRequestRequestBodyToEmailUpdateEmailTemplateRequest
// builds a value of type *email.UpdateEmailTemplateRequest from a value of
// type *UpdateEmailTemplateRequestRequestBody.
func unmarshalUpdateEmailTemplateRequestRequestBodyToEmailUpdateEmailTemplateRequest(v *UpdateEmailTemplateRequestRequestBody) *email.UpdateEmailTemplateRequest {
	res := &email.UpdateEmailTemplateRequest{
		Name:        v.Name,
		Subject:     v.Subject,
		HTMLContent: v.HTMLContent,
		TextContent: v.TextContent,
		Active:      v.Active,
		Locale:      v.Locale,
	}
	if v.Metadata != nil {
		res.Metadata = make(map[string]any, len(v.Metadata))
		for key, val := range v.Metadata {
			tk := key
			tv := val
			res.Metadata[tk] = tv
		}
	}

	return res
}

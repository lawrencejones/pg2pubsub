// Code generated by goa v3.2.6, DO NOT EDIT.
//
// Imports HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/lawrencejones/pgsink/api/design -o api

package server

import (
	"context"
	"net/http"

	imports "github.com/lawrencejones/pgsink/api/gen/imports"
	goahttp "goa.design/goa/v3/http"
)

// EncodeListResponse returns an encoder for responses returned by the Imports
// List endpoint.
func EncodeListResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.([]*imports.Import)
		enc := encoder(ctx, w)
		body := NewListResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// marshalImportsImportToImportResponse builds a value of type *ImportResponse
// from a value of type *imports.Import.
func marshalImportsImportToImportResponse(v *imports.Import) *ImportResponse {
	res := &ImportResponse{
		ID:             v.ID,
		SubscriptionID: v.SubscriptionID,
		Schema:         v.Schema,
		TableName:      v.TableName,
		CompletedAt:    v.CompletedAt,
		CreatedAt:      v.CreatedAt,
		UpdatedAt:      v.UpdatedAt,
		ExpiredAt:      v.ExpiredAt,
		Error:          v.Error,
		ErrorCount:     v.ErrorCount,
		LastErrorAt:    v.LastErrorAt,
	}

	return res
}
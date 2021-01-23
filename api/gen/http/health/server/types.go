// Code generated by goa v3.2.6, DO NOT EDIT.
//
// Health HTTP server types
//
// Command:
// $ goa gen github.com/lawrencejones/pgsink/api/design -o api

package server

import (
	health "github.com/lawrencejones/pgsink/api/gen/health"
)

// CheckResponseBody is the type of the "Health" service "Check" endpoint HTTP
// response body.
type CheckResponseBody struct {
	// Status of the API
	Status string `form:"status" json:"status" xml:"status"`
}

// NewCheckResponseBody builds the HTTP response body from the result of the
// "Check" endpoint of the "Health" service.
func NewCheckResponseBody(res *health.CheckResult) *CheckResponseBody {
	body := &CheckResponseBody{
		Status: res.Status,
	}
	return body
}

// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
)

// NewPostEchoParams creates a new PostEchoParams object
//
// There are no default values defined in the spec.
func NewPostEchoParams() PostEchoParams {

	return PostEchoParams{}
}

// PostEchoParams contains all the bound params for the post echo operation
// typically these are obtained from a http.Request
//
// swagger:parameters PostEcho
type PostEchoParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: body
	*/
	Message string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPostEchoParams() beforehand.
func (o *PostEchoParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body string
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("message", "body", ""))
			} else {
				res = append(res, errors.NewParseError("message", "body", "", err))
			}
		} else {
			// no validation required on inline body
			o.Message = body
		}
	} else {
		res = append(res, errors.Required("message", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeletePolicyHandlerFunc turns a function with the right signature into a delete policy handler
type DeletePolicyHandlerFunc func(DeletePolicyParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeletePolicyHandlerFunc) Handle(params DeletePolicyParams) middleware.Responder {
	return fn(params)
}

// DeletePolicyHandler interface for that can handle valid delete policy params
type DeletePolicyHandler interface {
	Handle(DeletePolicyParams) middleware.Responder
}

// NewDeletePolicy creates a new http.Handler for the delete policy operation
func NewDeletePolicy(ctx *middleware.Context, handler DeletePolicyHandler) *DeletePolicy {
	return &DeletePolicy{Context: ctx, Handler: handler}
}

/*DeletePolicy swagger:route DELETE /policy policy deletePolicy

Delete a policy (sub)tree

*/
type DeletePolicy struct {
	Context *middleware.Context
	Handler DeletePolicyHandler
}

func (o *DeletePolicy) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeletePolicyParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

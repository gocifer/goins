// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PostIpamHandlerFunc turns a function with the right signature into a post ipam handler
type PostIpamHandlerFunc func(PostIpamParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostIpamHandlerFunc) Handle(params PostIpamParams) middleware.Responder {
	return fn(params)
}

// PostIpamHandler interface for that can handle valid post ipam params
type PostIpamHandler interface {
	Handle(PostIpamParams) middleware.Responder
}

// NewPostIpam creates a new http.Handler for the post ipam operation
func NewPostIpam(ctx *middleware.Context, handler PostIpamHandler) *PostIpam {
	return &PostIpam{Context: ctx, Handler: handler}
}

/*PostIpam swagger:route POST /ipam ipam postIpam

Allocate an IP address

*/
type PostIpam struct {
	Context *middleware.Context
	Handler PostIpamHandler
}

func (o *PostIpam) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostIpamParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

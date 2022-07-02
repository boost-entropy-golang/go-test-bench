// Code generated by go-swagger; DO NOT EDIT.

package unvalidated_redirect

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UnvalidatedRedirectGetQueryRedirectHandlerFunc turns a function with the right signature into a unvalidated redirect get query redirect handler
type UnvalidatedRedirectGetQueryRedirectHandlerFunc func(UnvalidatedRedirectGetQueryRedirectParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UnvalidatedRedirectGetQueryRedirectHandlerFunc) Handle(params UnvalidatedRedirectGetQueryRedirectParams) middleware.Responder {
	return fn(params)
}

// UnvalidatedRedirectGetQueryRedirectHandler interface for that can handle valid unvalidated redirect get query redirect params
type UnvalidatedRedirectGetQueryRedirectHandler interface {
	Handle(UnvalidatedRedirectGetQueryRedirectParams) middleware.Responder
}

// NewUnvalidatedRedirectGetQueryRedirect creates a new http.Handler for the unvalidated redirect get query redirect operation
func NewUnvalidatedRedirectGetQueryRedirect(ctx *middleware.Context, handler UnvalidatedRedirectGetQueryRedirectHandler) *UnvalidatedRedirectGetQueryRedirect {
	return &UnvalidatedRedirectGetQueryRedirect{Context: ctx, Handler: handler}
}

/* UnvalidatedRedirectGetQueryRedirect swagger:route GET /unvalidatedRedirect/http.Redirect/query/{safety} unvalidated_redirect unvalidatedRedirectGetQueryRedirect

demonstrates Unvalidated Redirect via query, with vulnerable function http.Redirect

*/
type UnvalidatedRedirectGetQueryRedirect struct {
	Context *middleware.Context
	Handler UnvalidatedRedirectGetQueryRedirectHandler
}

func (o *UnvalidatedRedirectGetQueryRedirect) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUnvalidatedRedirectGetQueryRedirectParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

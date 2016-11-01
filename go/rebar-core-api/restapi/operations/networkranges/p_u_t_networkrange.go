package networkranges

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PUTNetworkrangeHandlerFunc turns a function with the right signature into a p u t networkrange handler
type PUTNetworkrangeHandlerFunc func(PUTNetworkrangeParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PUTNetworkrangeHandlerFunc) Handle(params PUTNetworkrangeParams) middleware.Responder {
	return fn(params)
}

// PUTNetworkrangeHandler interface for that can handle valid p u t networkrange params
type PUTNetworkrangeHandler interface {
	Handle(PUTNetworkrangeParams) middleware.Responder
}

// NewPUTNetworkrange creates a new http.Handler for the p u t networkrange operation
func NewPUTNetworkrange(ctx *middleware.Context, handler PUTNetworkrangeHandler) *PUTNetworkrange {
	return &PUTNetworkrange{Context: ctx, Handler: handler}
}

/*PUTNetworkrange swagger:route PUT /networkranges/{id} Networkranges pUTNetworkrange

Update NetworkRange

*/
type PUTNetworkrange struct {
	Context *middleware.Context
	Handler PUTNetworkrangeHandler
}

func (o *PUTNetworkrange) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewPUTNetworkrangeParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
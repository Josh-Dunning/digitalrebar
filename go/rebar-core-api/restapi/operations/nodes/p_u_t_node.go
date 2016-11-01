package nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PUTNodeHandlerFunc turns a function with the right signature into a p u t node handler
type PUTNodeHandlerFunc func(PUTNodeParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PUTNodeHandlerFunc) Handle(params PUTNodeParams) middleware.Responder {
	return fn(params)
}

// PUTNodeHandler interface for that can handle valid p u t node params
type PUTNodeHandler interface {
	Handle(PUTNodeParams) middleware.Responder
}

// NewPUTNode creates a new http.Handler for the p u t node operation
func NewPUTNode(ctx *middleware.Context, handler PUTNodeHandler) *PUTNode {
	return &PUTNode{Context: ctx, Handler: handler}
}

/*PUTNode swagger:route PUT /nodes/{id} Nodes pUTNode

Update Node

*/
type PUTNode struct {
	Context *middleware.Context
	Handler PUTNodeHandler
}

func (o *PUTNode) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewPUTNodeParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
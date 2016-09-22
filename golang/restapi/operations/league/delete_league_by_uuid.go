package league

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteLeagueByUUIDHandlerFunc turns a function with the right signature into a delete league by UUID handler
type DeleteLeagueByUUIDHandlerFunc func(DeleteLeagueByUUIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteLeagueByUUIDHandlerFunc) Handle(params DeleteLeagueByUUIDParams) middleware.Responder {
	return fn(params)
}

// DeleteLeagueByUUIDHandler interface for that can handle valid delete league by UUID params
type DeleteLeagueByUUIDHandler interface {
	Handle(DeleteLeagueByUUIDParams) middleware.Responder
}

// NewDeleteLeagueByUUID creates a new http.Handler for the delete league by UUID operation
func NewDeleteLeagueByUUID(ctx *middleware.Context, handler DeleteLeagueByUUIDHandler) *DeleteLeagueByUUID {
	return &DeleteLeagueByUUID{Context: ctx, Handler: handler}
}

/*DeleteLeagueByUUID swagger:route DELETE /leagues/{UUID} League deleteLeagueByUuid

删除特定联赛信息

删除特定联赛信息

*/
type DeleteLeagueByUUID struct {
	Context *middleware.Context
	Handler DeleteLeagueByUUIDHandler
}

func (o *DeleteLeagueByUUID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewDeleteLeagueByUUIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

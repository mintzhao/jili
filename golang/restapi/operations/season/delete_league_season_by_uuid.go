package season

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteLeagueSeasonByUUIDHandlerFunc turns a function with the right signature into a delete league season by UUID handler
type DeleteLeagueSeasonByUUIDHandlerFunc func(DeleteLeagueSeasonByUUIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteLeagueSeasonByUUIDHandlerFunc) Handle(params DeleteLeagueSeasonByUUIDParams) middleware.Responder {
	return fn(params)
}

// DeleteLeagueSeasonByUUIDHandler interface for that can handle valid delete league season by UUID params
type DeleteLeagueSeasonByUUIDHandler interface {
	Handle(DeleteLeagueSeasonByUUIDParams) middleware.Responder
}

// NewDeleteLeagueSeasonByUUID creates a new http.Handler for the delete league season by UUID operation
func NewDeleteLeagueSeasonByUUID(ctx *middleware.Context, handler DeleteLeagueSeasonByUUIDHandler) *DeleteLeagueSeasonByUUID {
	return &DeleteLeagueSeasonByUUID{Context: ctx, Handler: handler}
}

/*DeleteLeagueSeasonByUUID swagger:route DELETE /seasons/{SUUID} Season deleteLeagueSeasonByUuid

删除联赛赛季

删除联赛赛季

*/
type DeleteLeagueSeasonByUUID struct {
	Context *middleware.Context
	Handler DeleteLeagueSeasonByUUIDHandler
}

func (o *DeleteLeagueSeasonByUUID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewDeleteLeagueSeasonByUUIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

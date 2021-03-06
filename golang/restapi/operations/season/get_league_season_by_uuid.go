package season

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetLeagueSeasonByUUIDHandlerFunc turns a function with the right signature into a get league season by UUID handler
type GetLeagueSeasonByUUIDHandlerFunc func(GetLeagueSeasonByUUIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetLeagueSeasonByUUIDHandlerFunc) Handle(params GetLeagueSeasonByUUIDParams) middleware.Responder {
	return fn(params)
}

// GetLeagueSeasonByUUIDHandler interface for that can handle valid get league season by UUID params
type GetLeagueSeasonByUUIDHandler interface {
	Handle(GetLeagueSeasonByUUIDParams) middleware.Responder
}

// NewGetLeagueSeasonByUUID creates a new http.Handler for the get league season by UUID operation
func NewGetLeagueSeasonByUUID(ctx *middleware.Context, handler GetLeagueSeasonByUUIDHandler) *GetLeagueSeasonByUUID {
	return &GetLeagueSeasonByUUID{Context: ctx, Handler: handler}
}

/*GetLeagueSeasonByUUID swagger:route GET /seasons/{SUUID} Season getLeagueSeasonByUuid

联赛赛季信息

返回联赛赛季信息, 仅是支持这样的操作,但是不必要, app不会用到

*/
type GetLeagueSeasonByUUID struct {
	Context *middleware.Context
	Handler GetLeagueSeasonByUUIDHandler
}

func (o *GetLeagueSeasonByUUID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetLeagueSeasonByUUIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

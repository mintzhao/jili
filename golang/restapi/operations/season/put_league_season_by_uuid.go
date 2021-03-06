package season

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PutLeagueSeasonByUUIDHandlerFunc turns a function with the right signature into a put league season by UUID handler
type PutLeagueSeasonByUUIDHandlerFunc func(PutLeagueSeasonByUUIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PutLeagueSeasonByUUIDHandlerFunc) Handle(params PutLeagueSeasonByUUIDParams) middleware.Responder {
	return fn(params)
}

// PutLeagueSeasonByUUIDHandler interface for that can handle valid put league season by UUID params
type PutLeagueSeasonByUUIDHandler interface {
	Handle(PutLeagueSeasonByUUIDParams) middleware.Responder
}

// NewPutLeagueSeasonByUUID creates a new http.Handler for the put league season by UUID operation
func NewPutLeagueSeasonByUUID(ctx *middleware.Context, handler PutLeagueSeasonByUUIDHandler) *PutLeagueSeasonByUUID {
	return &PutLeagueSeasonByUUID{Context: ctx, Handler: handler}
}

/*PutLeagueSeasonByUUID swagger:route PUT /seasons/{SUUID} Season putLeagueSeasonByUuid

更新联赛赛季

根据参数,更新联赛赛季

*/
type PutLeagueSeasonByUUID struct {
	Context *middleware.Context
	Handler PutLeagueSeasonByUUIDHandler
}

func (o *PutLeagueSeasonByUUID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewPutLeagueSeasonByUUIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

package season

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetLeagueSeasonsHandlerFunc turns a function with the right signature into a get league seasons handler
type GetLeagueSeasonsHandlerFunc func(GetLeagueSeasonsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetLeagueSeasonsHandlerFunc) Handle(params GetLeagueSeasonsParams) middleware.Responder {
	return fn(params)
}

// GetLeagueSeasonsHandler interface for that can handle valid get league seasons params
type GetLeagueSeasonsHandler interface {
	Handle(GetLeagueSeasonsParams) middleware.Responder
}

// NewGetLeagueSeasons creates a new http.Handler for the get league seasons operation
func NewGetLeagueSeasons(ctx *middleware.Context, handler GetLeagueSeasonsHandler) *GetLeagueSeasons {
	return &GetLeagueSeasons{Context: ctx, Handler: handler}
}

/*GetLeagueSeasons swagger:route GET /leagues/{{UUID}}/seasons Season getLeagueSeasons

联赛赛季列表

返回系统支持的一系列联赛赛季信息, 仅是支持这样的操作,但是不必要, app不会用到

*/
type GetLeagueSeasons struct {
	Context *middleware.Context
	Handler GetLeagueSeasonsHandler
}

func (o *GetLeagueSeasons) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetLeagueSeasonsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

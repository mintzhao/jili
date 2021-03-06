package season

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PostLeagueSeasonHandlerFunc turns a function with the right signature into a post league season handler
type PostLeagueSeasonHandlerFunc func(PostLeagueSeasonParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostLeagueSeasonHandlerFunc) Handle(params PostLeagueSeasonParams) middleware.Responder {
	return fn(params)
}

// PostLeagueSeasonHandler interface for that can handle valid post league season params
type PostLeagueSeasonHandler interface {
	Handle(PostLeagueSeasonParams) middleware.Responder
}

// NewPostLeagueSeason creates a new http.Handler for the post league season operation
func NewPostLeagueSeason(ctx *middleware.Context, handler PostLeagueSeasonHandler) *PostLeagueSeason {
	return &PostLeagueSeason{Context: ctx, Handler: handler}
}

/*PostLeagueSeason swagger:route POST /leagues/{{UUID}}/seasons Season postLeagueSeason

新建联赛赛季

根据参数,新建联赛赛季

*/
type PostLeagueSeason struct {
	Context *middleware.Context
	Handler PostLeagueSeasonHandler
}

func (o *PostLeagueSeason) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewPostLeagueSeasonParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

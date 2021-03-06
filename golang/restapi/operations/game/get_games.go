package game

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetGamesHandlerFunc turns a function with the right signature into a get games handler
type GetGamesHandlerFunc func(GetGamesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetGamesHandlerFunc) Handle(params GetGamesParams) middleware.Responder {
	return fn(params)
}

// GetGamesHandler interface for that can handle valid get games params
type GetGamesHandler interface {
	Handle(GetGamesParams) middleware.Responder
}

// NewGetGames creates a new http.Handler for the get games operation
func NewGetGames(ctx *middleware.Context, handler GetGamesHandler) *GetGames {
	return &GetGames{Context: ctx, Handler: handler}
}

/*GetGames swagger:route GET /games Game getGames

比赛列表

比赛列表

*/
type GetGames struct {
	Context *middleware.Context
	Handler GetGamesHandler
}

func (o *GetGames) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetGamesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

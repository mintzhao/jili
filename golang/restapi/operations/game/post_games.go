package game

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PostGamesHandlerFunc turns a function with the right signature into a post games handler
type PostGamesHandlerFunc func(PostGamesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostGamesHandlerFunc) Handle(params PostGamesParams) middleware.Responder {
	return fn(params)
}

// PostGamesHandler interface for that can handle valid post games params
type PostGamesHandler interface {
	Handle(PostGamesParams) middleware.Responder
}

// NewPostGames creates a new http.Handler for the post games operation
func NewPostGames(ctx *middleware.Context, handler PostGamesHandler) *PostGames {
	return &PostGames{Context: ctx, Handler: handler}
}

/*PostGames swagger:route POST /games Game postGames

新增比赛

新增一场比赛

*/
type PostGames struct {
	Context *middleware.Context
	Handler PostGamesHandler
}

func (o *PostGames) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewPostGamesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

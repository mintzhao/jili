package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/mintzhao/jili/golang/restapi/operations/game"
	"github.com/mintzhao/jili/golang/restapi/operations/gamer"
	"github.com/mintzhao/jili/golang/restapi/operations/league"
	"github.com/mintzhao/jili/golang/restapi/operations/season"
)

// NewJiliAPI creates a new Jili instance
func NewJiliAPI(spec *loads.Document) *JiliAPI {
	return &JiliAPI{
		handlers:        make(map[string]map[string]http.Handler),
		formats:         strfmt.Default,
		defaultConsumes: "application/json",
		defaultProduces: "application/json",
		ServerShutdown:  func() {},
		spec:            spec,
	}
}

/*JiliAPI 集历是一款赛事赛程日历APP,提供最新最精准的赛事直播日历,支持一键导入系统日历,不再错过任何一场喜欢的比赛,更不用忍受APP广告带来的痛苦 */
type JiliAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	defaultConsumes string
	defaultProduces string
	// UrlformConsumer registers a consumer for a "application/x-www-form-urlencoded" mime type
	UrlformConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer

	// GameDeleteGameByUUIDHandler sets the operation handler for the delete game by UUID operation
	GameDeleteGameByUUIDHandler game.DeleteGameByUUIDHandler
	// LeagueDeleteLeagueByUUIDHandler sets the operation handler for the delete league by UUID operation
	LeagueDeleteLeagueByUUIDHandler league.DeleteLeagueByUUIDHandler
	// SeasonDeleteLeagueSeasonByUUIDHandler sets the operation handler for the delete league season by UUID operation
	SeasonDeleteLeagueSeasonByUUIDHandler season.DeleteLeagueSeasonByUUIDHandler
	// GameGetGameByUUIDHandler sets the operation handler for the get game by UUID operation
	GameGetGameByUUIDHandler game.GetGameByUUIDHandler
	// GameGetGamesHandler sets the operation handler for the get games operation
	GameGetGamesHandler game.GetGamesHandler
	// LeagueGetLeagueByUUIDHandler sets the operation handler for the get league by UUID operation
	LeagueGetLeagueByUUIDHandler league.GetLeagueByUUIDHandler
	// SeasonGetLeagueSeasonByUUIDHandler sets the operation handler for the get league season by UUID operation
	SeasonGetLeagueSeasonByUUIDHandler season.GetLeagueSeasonByUUIDHandler
	// SeasonGetLeagueSeasonsHandler sets the operation handler for the get league seasons operation
	SeasonGetLeagueSeasonsHandler season.GetLeagueSeasonsHandler
	// LeagueGetLeaguesHandler sets the operation handler for the get leagues operation
	LeagueGetLeaguesHandler league.GetLeaguesHandler
	// GamerPostGamersHandler sets the operation handler for the post gamers operation
	GamerPostGamersHandler gamer.PostGamersHandler
	// GamePostGamesHandler sets the operation handler for the post games operation
	GamePostGamesHandler game.PostGamesHandler
	// SeasonPostLeagueSeasonHandler sets the operation handler for the post league season operation
	SeasonPostLeagueSeasonHandler season.PostLeagueSeasonHandler
	// LeaguePostLeaguesHandler sets the operation handler for the post leagues operation
	LeaguePostLeaguesHandler league.PostLeaguesHandler
	// GamePutGameByUUIDHandler sets the operation handler for the put game by UUID operation
	GamePutGameByUUIDHandler game.PutGameByUUIDHandler
	// LeaguePutLeagueByUUIDHandler sets the operation handler for the put league by UUID operation
	LeaguePutLeagueByUUIDHandler league.PutLeagueByUUIDHandler
	// SeasonPutLeagueSeasonByUUIDHandler sets the operation handler for the put league season by UUID operation
	SeasonPutLeagueSeasonByUUIDHandler season.PutLeagueSeasonByUUIDHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *JiliAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *JiliAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *JiliAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *JiliAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *JiliAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *JiliAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *JiliAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the JiliAPI
func (o *JiliAPI) Validate() error {
	var unregistered []string

	if o.UrlformConsumer == nil {
		unregistered = append(unregistered, "UrlformConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.GameDeleteGameByUUIDHandler == nil {
		unregistered = append(unregistered, "game.DeleteGameByUUIDHandler")
	}

	if o.LeagueDeleteLeagueByUUIDHandler == nil {
		unregistered = append(unregistered, "league.DeleteLeagueByUUIDHandler")
	}

	if o.SeasonDeleteLeagueSeasonByUUIDHandler == nil {
		unregistered = append(unregistered, "season.DeleteLeagueSeasonByUUIDHandler")
	}

	if o.GameGetGameByUUIDHandler == nil {
		unregistered = append(unregistered, "game.GetGameByUUIDHandler")
	}

	if o.GameGetGamesHandler == nil {
		unregistered = append(unregistered, "game.GetGamesHandler")
	}

	if o.LeagueGetLeagueByUUIDHandler == nil {
		unregistered = append(unregistered, "league.GetLeagueByUUIDHandler")
	}

	if o.SeasonGetLeagueSeasonByUUIDHandler == nil {
		unregistered = append(unregistered, "season.GetLeagueSeasonByUUIDHandler")
	}

	if o.SeasonGetLeagueSeasonsHandler == nil {
		unregistered = append(unregistered, "season.GetLeagueSeasonsHandler")
	}

	if o.LeagueGetLeaguesHandler == nil {
		unregistered = append(unregistered, "league.GetLeaguesHandler")
	}

	if o.GamerPostGamersHandler == nil {
		unregistered = append(unregistered, "gamer.PostGamersHandler")
	}

	if o.GamePostGamesHandler == nil {
		unregistered = append(unregistered, "game.PostGamesHandler")
	}

	if o.SeasonPostLeagueSeasonHandler == nil {
		unregistered = append(unregistered, "season.PostLeagueSeasonHandler")
	}

	if o.LeaguePostLeaguesHandler == nil {
		unregistered = append(unregistered, "league.PostLeaguesHandler")
	}

	if o.GamePutGameByUUIDHandler == nil {
		unregistered = append(unregistered, "game.PutGameByUUIDHandler")
	}

	if o.LeaguePutLeagueByUUIDHandler == nil {
		unregistered = append(unregistered, "league.PutLeagueByUUIDHandler")
	}

	if o.SeasonPutLeagueSeasonByUUIDHandler == nil {
		unregistered = append(unregistered, "season.PutLeagueSeasonByUUIDHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *JiliAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *JiliAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	return nil

}

// ConsumersFor gets the consumers for the specified media types
func (o *JiliAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/x-www-form-urlencoded":
			result["application/x-www-form-urlencoded"] = o.UrlformConsumer

		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *JiliAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *JiliAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

func (o *JiliAPI) initHandlerCache() {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["DELETE"] == nil {
		o.handlers[strings.ToUpper("DELETE")] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/games/{UUID}"] = game.NewDeleteGameByUUID(o.context, o.GameDeleteGameByUUIDHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers[strings.ToUpper("DELETE")] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/leagues/{UUID}"] = league.NewDeleteLeagueByUUID(o.context, o.LeagueDeleteLeagueByUUIDHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers[strings.ToUpper("DELETE")] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/seasons/{SUUID}"] = season.NewDeleteLeagueSeasonByUUID(o.context, o.SeasonDeleteLeagueSeasonByUUIDHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/games/{UUID}"] = game.NewGetGameByUUID(o.context, o.GameGetGameByUUIDHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/games"] = game.NewGetGames(o.context, o.GameGetGamesHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/leagues/{UUID}"] = league.NewGetLeagueByUUID(o.context, o.LeagueGetLeagueByUUIDHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/seasons/{SUUID}"] = season.NewGetLeagueSeasonByUUID(o.context, o.SeasonGetLeagueSeasonByUUIDHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/leagues/{{UUID}}/seasons"] = season.NewGetLeagueSeasons(o.context, o.SeasonGetLeagueSeasonsHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/leagues"] = league.NewGetLeagues(o.context, o.LeagueGetLeaguesHandler)

	if o.handlers["POST"] == nil {
		o.handlers[strings.ToUpper("POST")] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/gamers"] = gamer.NewPostGamers(o.context, o.GamerPostGamersHandler)

	if o.handlers["POST"] == nil {
		o.handlers[strings.ToUpper("POST")] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/games"] = game.NewPostGames(o.context, o.GamePostGamesHandler)

	if o.handlers["POST"] == nil {
		o.handlers[strings.ToUpper("POST")] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/leagues/{{UUID}}/seasons"] = season.NewPostLeagueSeason(o.context, o.SeasonPostLeagueSeasonHandler)

	if o.handlers["POST"] == nil {
		o.handlers[strings.ToUpper("POST")] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/leagues"] = league.NewPostLeagues(o.context, o.LeaguePostLeaguesHandler)

	if o.handlers["PUT"] == nil {
		o.handlers[strings.ToUpper("PUT")] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/games/{UUID}"] = game.NewPutGameByUUID(o.context, o.GamePutGameByUUIDHandler)

	if o.handlers["PUT"] == nil {
		o.handlers[strings.ToUpper("PUT")] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/leagues/{UUID}"] = league.NewPutLeagueByUUID(o.context, o.LeaguePutLeagueByUUIDHandler)

	if o.handlers["PUT"] == nil {
		o.handlers[strings.ToUpper("PUT")] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/seasons/{SUUID}"] = season.NewPutLeagueSeasonByUUID(o.context, o.SeasonPutLeagueSeasonByUUIDHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *JiliAPI) Serve(builder middleware.Builder) http.Handler {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}

	return o.context.APIHandler(builder)
}
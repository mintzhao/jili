package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/mintzhao/jili/golang/restapi/operations"
	"github.com/mintzhao/jili/golang/restapi/operations/game"
	"github.com/mintzhao/jili/golang/restapi/operations/gamer"
	"github.com/mintzhao/jili/golang/restapi/operations/league"
	"github.com/mintzhao/jili/golang/restapi/operations/season"
)

// This file is safe to edit. Once it exists it will not be overwritten

//go:generate swagger generate server --target ../golang --name  --spec ../rest_api.json

func configureFlags(api *operations.JiliAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.JiliAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// s.api.Logger = log.Printf

	api.UrlformConsumer = runtime.DiscardConsumer

	api.JSONProducer = runtime.JSONProducer()

	api.GameDeleteGameByUUIDHandler = game.DeleteGameByUUIDHandlerFunc(func(params game.DeleteGameByUUIDParams) middleware.Responder {
		return middleware.NotImplemented("operation game.DeleteGameByUUID has not yet been implemented")
	})
	api.LeagueDeleteLeagueByUUIDHandler = league.DeleteLeagueByUUIDHandlerFunc(func(params league.DeleteLeagueByUUIDParams) middleware.Responder {
		return middleware.NotImplemented("operation league.DeleteLeagueByUUID has not yet been implemented")
	})
	api.SeasonDeleteLeagueSeasonByUUIDHandler = season.DeleteLeagueSeasonByUUIDHandlerFunc(func(params season.DeleteLeagueSeasonByUUIDParams) middleware.Responder {
		return middleware.NotImplemented("operation season.DeleteLeagueSeasonByUUID has not yet been implemented")
	})
	api.GameGetGameByUUIDHandler = game.GetGameByUUIDHandlerFunc(func(params game.GetGameByUUIDParams) middleware.Responder {
		return middleware.NotImplemented("operation game.GetGameByUUID has not yet been implemented")
	})
	api.GameGetGamesHandler = game.GetGamesHandlerFunc(func(params game.GetGamesParams) middleware.Responder {
		return middleware.NotImplemented("operation game.GetGames has not yet been implemented")
	})
	api.LeagueGetLeagueByUUIDHandler = league.GetLeagueByUUIDHandlerFunc(func(params league.GetLeagueByUUIDParams) middleware.Responder {
		return middleware.NotImplemented("operation league.GetLeagueByUUID has not yet been implemented")
	})
	api.SeasonGetLeagueSeasonByUUIDHandler = season.GetLeagueSeasonByUUIDHandlerFunc(func(params season.GetLeagueSeasonByUUIDParams) middleware.Responder {
		return middleware.NotImplemented("operation season.GetLeagueSeasonByUUID has not yet been implemented")
	})
	api.SeasonGetLeagueSeasonsHandler = season.GetLeagueSeasonsHandlerFunc(func(params season.GetLeagueSeasonsParams) middleware.Responder {
		return middleware.NotImplemented("operation season.GetLeagueSeasons has not yet been implemented")
	})
	api.LeagueGetLeaguesHandler = league.GetLeaguesHandlerFunc(func(params league.GetLeaguesParams) middleware.Responder {
		return middleware.NotImplemented("operation league.GetLeagues has not yet been implemented")
	})
	api.GamerPostGamersHandler = gamer.PostGamersHandlerFunc(func(params gamer.PostGamersParams) middleware.Responder {
		return middleware.NotImplemented("operation gamer.PostGamers has not yet been implemented")
	})
	api.GamePostGamesHandler = game.PostGamesHandlerFunc(func(params game.PostGamesParams) middleware.Responder {
		return middleware.NotImplemented("operation game.PostGames has not yet been implemented")
	})
	api.SeasonPostLeagueSeasonHandler = season.PostLeagueSeasonHandlerFunc(func(params season.PostLeagueSeasonParams) middleware.Responder {
		return middleware.NotImplemented("operation season.PostLeagueSeason has not yet been implemented")
	})
	api.LeaguePostLeaguesHandler = league.PostLeaguesHandlerFunc(func(params league.PostLeaguesParams) middleware.Responder {
		return middleware.NotImplemented("operation league.PostLeagues has not yet been implemented")
	})
	api.GamePutGameByUUIDHandler = game.PutGameByUUIDHandlerFunc(func(params game.PutGameByUUIDParams) middleware.Responder {
		return middleware.NotImplemented("operation game.PutGameByUUID has not yet been implemented")
	})
	api.LeaguePutLeagueByUUIDHandler = league.PutLeagueByUUIDHandlerFunc(func(params league.PutLeagueByUUIDParams) middleware.Responder {
		return middleware.NotImplemented("operation league.PutLeagueByUUID has not yet been implemented")
	})
	api.SeasonPutLeagueSeasonByUUIDHandler = season.PutLeagueSeasonByUUIDHandlerFunc(func(params season.PutLeagueSeasonByUUIDParams) middleware.Responder {
		return middleware.NotImplemented("operation season.PutLeagueSeasonByUUID has not yet been implemented")
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}

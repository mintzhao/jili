package game

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/mintzhao/jili/golang/models"
)

// NewPostGamesParams creates a new PostGamesParams object
// with the default values initialized.
func NewPostGamesParams() PostGamesParams {
	var ()
	return PostGamesParams{}
}

// PostGamesParams contains all the bound params for the post games operation
// typically these are obtained from a http.Request
//
// swagger:parameters postGames
type PostGamesParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*新增比赛的信息
	  Required: true
	  In: body
	*/
	Game []*models.Game
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *PostGamesParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body []*models.Game
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("game", "body"))
			} else {
				res = append(res, errors.NewParseError("game", "body", "", err))
			}

		} else {
			for _, io := range o.Game {
				if err := io.Validate(route.Formats); err != nil {
					res = append(res, err)
					break
				}
			}

			if len(res) == 0 {
				o.Game = body
			}
		}

	} else {
		res = append(res, errors.Required("game", "body"))
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

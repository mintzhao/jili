package season

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetLeagueSeasonByUUIDParams creates a new GetLeagueSeasonByUUIDParams object
// with the default values initialized.
func NewGetLeagueSeasonByUUIDParams() GetLeagueSeasonByUUIDParams {
	var ()
	return GetLeagueSeasonByUUIDParams{}
}

// GetLeagueSeasonByUUIDParams contains all the bound params for the get league season by UUID operation
// typically these are obtained from a http.Request
//
// swagger:parameters getLeagueSeasonByUUID
type GetLeagueSeasonByUUIDParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*League season identity
	  Required: true
	  In: path
	*/
	SUUID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *GetLeagueSeasonByUUIDParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	rSUUID, rhkSUUID, _ := route.Params.GetOK("SUUID")
	if err := o.bindSUUID(rSUUID, rhkSUUID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetLeagueSeasonByUUIDParams) bindSUUID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.SUUID = raw

	return nil
}
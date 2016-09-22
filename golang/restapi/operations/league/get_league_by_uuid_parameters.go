package league

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetLeagueByUUIDParams creates a new GetLeagueByUUIDParams object
// with the default values initialized.
func NewGetLeagueByUUIDParams() GetLeagueByUUIDParams {
	var ()
	return GetLeagueByUUIDParams{}
}

// GetLeagueByUUIDParams contains all the bound params for the get league by UUID operation
// typically these are obtained from a http.Request
//
// swagger:parameters getLeagueByUUID
type GetLeagueByUUIDParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*League identity
	  Required: true
	  In: path
	*/
	UUID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *GetLeagueByUUIDParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	rUUID, rhkUUID, _ := route.Params.GetOK("UUID")
	if err := o.bindUUID(rUUID, rhkUUID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetLeagueByUUIDParams) bindUUID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.UUID = raw

	return nil
}

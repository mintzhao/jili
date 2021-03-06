package league

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteLeagueByUUIDParams creates a new DeleteLeagueByUUIDParams object
// with the default values initialized.
func NewDeleteLeagueByUUIDParams() DeleteLeagueByUUIDParams {
	var ()
	return DeleteLeagueByUUIDParams{}
}

// DeleteLeagueByUUIDParams contains all the bound params for the delete league by UUID operation
// typically these are obtained from a http.Request
//
// swagger:parameters deleteLeagueByUUID
type DeleteLeagueByUUIDParams struct {

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
func (o *DeleteLeagueByUUIDParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
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

func (o *DeleteLeagueByUUIDParams) bindUUID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.UUID = raw

	return nil
}

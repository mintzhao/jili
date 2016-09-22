package season

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/mintzhao/jili/golang/models"
)

/*GetLeagueSeasonByUUIDOK 联赛赛季

swagger:response getLeagueSeasonByUuidOK
*/
type GetLeagueSeasonByUUIDOK struct {

	// In: body
	Payload *models.Season `json:"body,omitempty"`
}

// NewGetLeagueSeasonByUUIDOK creates GetLeagueSeasonByUUIDOK with default headers values
func NewGetLeagueSeasonByUUIDOK() *GetLeagueSeasonByUUIDOK {
	return &GetLeagueSeasonByUUIDOK{}
}

// WithPayload adds the payload to the get league season by Uuid o k response
func (o *GetLeagueSeasonByUUIDOK) WithPayload(payload *models.Season) *GetLeagueSeasonByUUIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get league season by Uuid o k response
func (o *GetLeagueSeasonByUUIDOK) SetPayload(payload *models.Season) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetLeagueSeasonByUUIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetLeagueSeasonByUUIDDefault Unexpected error

swagger:response getLeagueSeasonByUuidDefault
*/
type GetLeagueSeasonByUUIDDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetLeagueSeasonByUUIDDefault creates GetLeagueSeasonByUUIDDefault with default headers values
func NewGetLeagueSeasonByUUIDDefault(code int) *GetLeagueSeasonByUUIDDefault {
	if code <= 0 {
		code = 500
	}

	return &GetLeagueSeasonByUUIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get league season by UUID default response
func (o *GetLeagueSeasonByUUIDDefault) WithStatusCode(code int) *GetLeagueSeasonByUUIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get league season by UUID default response
func (o *GetLeagueSeasonByUUIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get league season by UUID default response
func (o *GetLeagueSeasonByUUIDDefault) WithPayload(payload *models.Error) *GetLeagueSeasonByUUIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get league season by UUID default response
func (o *GetLeagueSeasonByUUIDDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetLeagueSeasonByUUIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
package season

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/mintzhao/jili/golang/models"
)

/*GetLeagueSeasonsOK 联赛赛季列表

swagger:response getLeagueSeasonsOK
*/
type GetLeagueSeasonsOK struct {

	// In: body
	Payload []*models.Season `json:"body,omitempty"`
}

// NewGetLeagueSeasonsOK creates GetLeagueSeasonsOK with default headers values
func NewGetLeagueSeasonsOK() *GetLeagueSeasonsOK {
	return &GetLeagueSeasonsOK{}
}

// WithPayload adds the payload to the get league seasons o k response
func (o *GetLeagueSeasonsOK) WithPayload(payload []*models.Season) *GetLeagueSeasonsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get league seasons o k response
func (o *GetLeagueSeasonsOK) SetPayload(payload []*models.Season) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetLeagueSeasonsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*GetLeagueSeasonsDefault Unexpected error

swagger:response getLeagueSeasonsDefault
*/
type GetLeagueSeasonsDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetLeagueSeasonsDefault creates GetLeagueSeasonsDefault with default headers values
func NewGetLeagueSeasonsDefault(code int) *GetLeagueSeasonsDefault {
	if code <= 0 {
		code = 500
	}

	return &GetLeagueSeasonsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get league seasons default response
func (o *GetLeagueSeasonsDefault) WithStatusCode(code int) *GetLeagueSeasonsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get league seasons default response
func (o *GetLeagueSeasonsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get league seasons default response
func (o *GetLeagueSeasonsDefault) WithPayload(payload *models.Error) *GetLeagueSeasonsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get league seasons default response
func (o *GetLeagueSeasonsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetLeagueSeasonsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

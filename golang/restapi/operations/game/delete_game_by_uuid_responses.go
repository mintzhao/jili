package game

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/mintzhao/jili/golang/models"
)

/*DeleteGameByUUIDOK 删除比赛成功

swagger:response deleteGameByUuidOK
*/
type DeleteGameByUUIDOK struct {
}

// NewDeleteGameByUUIDOK creates DeleteGameByUUIDOK with default headers values
func NewDeleteGameByUUIDOK() *DeleteGameByUUIDOK {
	return &DeleteGameByUUIDOK{}
}

// WriteResponse to the client
func (o *DeleteGameByUUIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
}

/*DeleteGameByUUIDDefault Unexpected error

swagger:response deleteGameByUuidDefault
*/
type DeleteGameByUUIDDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteGameByUUIDDefault creates DeleteGameByUUIDDefault with default headers values
func NewDeleteGameByUUIDDefault(code int) *DeleteGameByUUIDDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteGameByUUIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete game by UUID default response
func (o *DeleteGameByUUIDDefault) WithStatusCode(code int) *DeleteGameByUUIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete game by UUID default response
func (o *DeleteGameByUUIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete game by UUID default response
func (o *DeleteGameByUUIDDefault) WithPayload(payload *models.Error) *DeleteGameByUUIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete game by UUID default response
func (o *DeleteGameByUUIDDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteGameByUUIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

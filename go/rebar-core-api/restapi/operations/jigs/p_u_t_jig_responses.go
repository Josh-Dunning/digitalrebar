package jigs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/digitalrebar/digitalrebar/go/rebar-core-api/models"
)

/*PUTJigOK p u t jig o k

swagger:response pUTJigOK
*/
type PUTJigOK struct {

	// In: body
	Payload *models.JigOutput `json:"body,omitempty"`
}

// NewPUTJigOK creates PUTJigOK with default headers values
func NewPUTJigOK() *PUTJigOK {
	return &PUTJigOK{}
}

// WithPayload adds the payload to the p u t jig o k response
func (o *PUTJigOK) WithPayload(payload *models.JigOutput) *PUTJigOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the p u t jig o k response
func (o *PUTJigOK) SetPayload(payload *models.JigOutput) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PUTJigOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
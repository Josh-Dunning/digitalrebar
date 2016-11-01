package eventselectors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/digitalrebar/digitalrebar/go/rebar-core-api/models"
)

/*GETEventselectorOK g e t eventselector o k

swagger:response gETEventselectorOK
*/
type GETEventselectorOK struct {

	// In: body
	Payload *models.EventselectorOutput `json:"body,omitempty"`
}

// NewGETEventselectorOK creates GETEventselectorOK with default headers values
func NewGETEventselectorOK() *GETEventselectorOK {
	return &GETEventselectorOK{}
}

// WithPayload adds the payload to the g e t eventselector o k response
func (o *GETEventselectorOK) WithPayload(payload *models.EventselectorOutput) *GETEventselectorOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the g e t eventselector o k response
func (o *GETEventselectorOK) SetPayload(payload *models.EventselectorOutput) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GETEventselectorOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
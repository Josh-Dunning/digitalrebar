package networkallocations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/digitalrebar/digitalrebar/go/rebar-core-api/models"
)

/*GETNetworkallocationOK g e t networkallocation o k

swagger:response gETNetworkallocationOK
*/
type GETNetworkallocationOK struct {

	// In: body
	Payload *models.NetworkallocationOutput `json:"body,omitempty"`
}

// NewGETNetworkallocationOK creates GETNetworkallocationOK with default headers values
func NewGETNetworkallocationOK() *GETNetworkallocationOK {
	return &GETNetworkallocationOK{}
}

// WithPayload adds the payload to the g e t networkallocation o k response
func (o *GETNetworkallocationOK) WithPayload(payload *models.NetworkallocationOutput) *GETNetworkallocationOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the g e t networkallocation o k response
func (o *GETNetworkallocationOK) SetPayload(payload *models.NetworkallocationOutput) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GETNetworkallocationOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
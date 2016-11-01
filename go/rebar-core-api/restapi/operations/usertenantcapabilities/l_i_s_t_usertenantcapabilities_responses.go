package usertenantcapabilities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/digitalrebar/digitalrebar/go/rebar-core-api/models"
)

/*LISTUsertenantcapabilitiesOK l i s t usertenantcapabilities o k

swagger:response lISTUsertenantcapabilitiesOK
*/
type LISTUsertenantcapabilitiesOK struct {

	// In: body
	Payload []*models.UsertenantcapabilityOutput `json:"body,omitempty"`
}

// NewLISTUsertenantcapabilitiesOK creates LISTUsertenantcapabilitiesOK with default headers values
func NewLISTUsertenantcapabilitiesOK() *LISTUsertenantcapabilitiesOK {
	return &LISTUsertenantcapabilitiesOK{}
}

// WithPayload adds the payload to the l i s t usertenantcapabilities o k response
func (o *LISTUsertenantcapabilitiesOK) WithPayload(payload []*models.UsertenantcapabilityOutput) *LISTUsertenantcapabilitiesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the l i s t usertenantcapabilities o k response
func (o *LISTUsertenantcapabilitiesOK) SetPayload(payload []*models.UsertenantcapabilityOutput) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *LISTUsertenantcapabilitiesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}
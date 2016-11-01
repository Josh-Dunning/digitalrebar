package networkallocations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/digitalrebar/digitalrebar/go/rebar-core-api/models"
)

// GETNetworkallocationReader is a Reader for the GETNetworkallocation structure.
type GETNetworkallocationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETNetworkallocationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGETNetworkallocationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGETNetworkallocationOK creates a GETNetworkallocationOK with default headers values
func NewGETNetworkallocationOK() *GETNetworkallocationOK {
	return &GETNetworkallocationOK{}
}

/*GETNetworkallocationOK handles this case with default header values.

GETNetworkallocationOK g e t networkallocation o k
*/
type GETNetworkallocationOK struct {
	Payload *models.NetworkallocationOutput
}

func (o *GETNetworkallocationOK) Error() string {
	return fmt.Sprintf("[GET /networkallocations/{id}][%d] gETNetworkallocationOK  %+v", 200, o.Payload)
}

func (o *GETNetworkallocationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NetworkallocationOutput)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
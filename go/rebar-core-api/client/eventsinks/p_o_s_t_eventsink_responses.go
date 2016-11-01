package eventsinks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/digitalrebar/digitalrebar/go/rebar-core-api/models"
)

// POSTEventsinkReader is a Reader for the POSTEventsink structure.
type POSTEventsinkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *POSTEventsinkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPOSTEventsinkCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPOSTEventsinkCreated creates a POSTEventsinkCreated with default headers values
func NewPOSTEventsinkCreated() *POSTEventsinkCreated {
	return &POSTEventsinkCreated{}
}

/*POSTEventsinkCreated handles this case with default header values.

POSTEventsinkCreated p o s t eventsink created
*/
type POSTEventsinkCreated struct {
	Payload *models.EventsinkOutput
}

func (o *POSTEventsinkCreated) Error() string {
	return fmt.Sprintf("[POST /eventsinks][%d] pOSTEventsinkCreated  %+v", 201, o.Payload)
}

func (o *POSTEventsinkCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EventsinkOutput)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
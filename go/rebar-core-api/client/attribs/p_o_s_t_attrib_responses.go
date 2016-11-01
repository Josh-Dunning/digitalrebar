package attribs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/digitalrebar/digitalrebar/go/rebar-core-api/models"
)

// POSTAttribReader is a Reader for the POSTAttrib structure.
type POSTAttribReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *POSTAttribReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPOSTAttribCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPOSTAttribCreated creates a POSTAttribCreated with default headers values
func NewPOSTAttribCreated() *POSTAttribCreated {
	return &POSTAttribCreated{}
}

/*POSTAttribCreated handles this case with default header values.

POSTAttribCreated p o s t attrib created
*/
type POSTAttribCreated struct {
	Payload *models.AttribOutput
}

func (o *POSTAttribCreated) Error() string {
	return fmt.Sprintf("[POST /attribs][%d] pOSTAttribCreated  %+v", 201, o.Payload)
}

func (o *POSTAttribCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AttribOutput)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
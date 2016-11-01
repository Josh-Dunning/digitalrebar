package tenants

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/digitalrebar/digitalrebar/go/rebar-core-api/models"
)

// PUTTenantReader is a Reader for the PUTTenant structure.
type PUTTenantReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PUTTenantReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPUTTenantOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPUTTenantOK creates a PUTTenantOK with default headers values
func NewPUTTenantOK() *PUTTenantOK {
	return &PUTTenantOK{}
}

/*PUTTenantOK handles this case with default header values.

PUTTenantOK p u t tenant o k
*/
type PUTTenantOK struct {
	Payload *models.TenantOutput
}

func (o *PUTTenantOK) Error() string {
	return fmt.Sprintf("[PUT /tenants/{id}][%d] pUTTenantOK  %+v", 200, o.Payload)
}

func (o *PUTTenantOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TenantOutput)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
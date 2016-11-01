package barclamps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DELETEBarclampReader is a Reader for the DELETEBarclamp structure.
type DELETEBarclampReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DELETEBarclampReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDELETEBarclampNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDELETEBarclampNoContent creates a DELETEBarclampNoContent with default headers values
func NewDELETEBarclampNoContent() *DELETEBarclampNoContent {
	return &DELETEBarclampNoContent{}
}

/*DELETEBarclampNoContent handles this case with default header values.

DELETEBarclampNoContent d e l e t e barclamp no content
*/
type DELETEBarclampNoContent struct {
}

func (o *DELETEBarclampNoContent) Error() string {
	return fmt.Sprintf("[DELETE /barclamps/{id}][%d] dELETEBarclampNoContent ", 204)
}

func (o *DELETEBarclampNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
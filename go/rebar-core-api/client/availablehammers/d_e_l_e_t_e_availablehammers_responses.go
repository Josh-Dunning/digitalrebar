package availablehammers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DELETEAvailablehammersReader is a Reader for the DELETEAvailablehammers structure.
type DELETEAvailablehammersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DELETEAvailablehammersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDELETEAvailablehammersNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDELETEAvailablehammersNoContent creates a DELETEAvailablehammersNoContent with default headers values
func NewDELETEAvailablehammersNoContent() *DELETEAvailablehammersNoContent {
	return &DELETEAvailablehammersNoContent{}
}

/*DELETEAvailablehammersNoContent handles this case with default header values.

DELETEAvailablehammersNoContent d e l e t e availablehammers no content
*/
type DELETEAvailablehammersNoContent struct {
}

func (o *DELETEAvailablehammersNoContent) Error() string {
	return fmt.Sprintf("[DELETE /availablehammers/{id}][%d] dELETEAvailablehammersNoContent ", 204)
}

func (o *DELETEAvailablehammersNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
package dnsnamefilters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

/*DELETEDnsnamefiltersNoContent d e l e t e dnsnamefilters no content

swagger:response dELETEDnsnamefiltersNoContent
*/
type DELETEDnsnamefiltersNoContent struct {
}

// NewDELETEDnsnamefiltersNoContent creates DELETEDnsnamefiltersNoContent with default headers values
func NewDELETEDnsnamefiltersNoContent() *DELETEDnsnamefiltersNoContent {
	return &DELETEDnsnamefiltersNoContent{}
}

// WriteResponse to the client
func (o *DELETEDnsnamefiltersNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(204)
}
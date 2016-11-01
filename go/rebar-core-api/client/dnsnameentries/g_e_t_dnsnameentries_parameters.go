package dnsnameentries

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGETDnsnameentriesParams creates a new GETDnsnameentriesParams object
// with the default values initialized.
func NewGETDnsnameentriesParams() *GETDnsnameentriesParams {
	var ()
	return &GETDnsnameentriesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGETDnsnameentriesParamsWithTimeout creates a new GETDnsnameentriesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGETDnsnameentriesParamsWithTimeout(timeout time.Duration) *GETDnsnameentriesParams {
	var ()
	return &GETDnsnameentriesParams{

		timeout: timeout,
	}
}

// NewGETDnsnameentriesParamsWithContext creates a new GETDnsnameentriesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGETDnsnameentriesParamsWithContext(ctx context.Context) *GETDnsnameentriesParams {
	var ()
	return &GETDnsnameentriesParams{

		Context: ctx,
	}
}

/*GETDnsnameentriesParams contains all the parameters to send to the API endpoint
for the g e t dnsnameentries operation typically these are written to a http.Request
*/
type GETDnsnameentriesParams struct {

	/*ID*/
	ID string

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the g e t dnsnameentries params
func (o *GETDnsnameentriesParams) WithTimeout(timeout time.Duration) *GETDnsnameentriesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the g e t dnsnameentries params
func (o *GETDnsnameentriesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the g e t dnsnameentries params
func (o *GETDnsnameentriesParams) WithContext(ctx context.Context) *GETDnsnameentriesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the g e t dnsnameentries params
func (o *GETDnsnameentriesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithID adds the id to the g e t dnsnameentries params
func (o *GETDnsnameentriesParams) WithID(id string) *GETDnsnameentriesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the g e t dnsnameentries params
func (o *GETDnsnameentriesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GETDnsnameentriesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
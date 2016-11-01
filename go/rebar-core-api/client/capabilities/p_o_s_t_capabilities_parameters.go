package capabilities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/digitalrebar/digitalrebar/go/rebar-core-api/models"
)

// NewPOSTCapabilitiesParams creates a new POSTCapabilitiesParams object
// with the default values initialized.
func NewPOSTCapabilitiesParams() *POSTCapabilitiesParams {
	var ()
	return &POSTCapabilitiesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPOSTCapabilitiesParamsWithTimeout creates a new POSTCapabilitiesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPOSTCapabilitiesParamsWithTimeout(timeout time.Duration) *POSTCapabilitiesParams {
	var ()
	return &POSTCapabilitiesParams{

		timeout: timeout,
	}
}

// NewPOSTCapabilitiesParamsWithContext creates a new POSTCapabilitiesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPOSTCapabilitiesParamsWithContext(ctx context.Context) *POSTCapabilitiesParams {
	var ()
	return &POSTCapabilitiesParams{

		Context: ctx,
	}
}

/*POSTCapabilitiesParams contains all the parameters to send to the API endpoint
for the p o s t capabilities operation typically these are written to a http.Request
*/
type POSTCapabilitiesParams struct {

	/*Body*/
	Body *models.CapabilitiesInput

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the p o s t capabilities params
func (o *POSTCapabilitiesParams) WithTimeout(timeout time.Duration) *POSTCapabilitiesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the p o s t capabilities params
func (o *POSTCapabilitiesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the p o s t capabilities params
func (o *POSTCapabilitiesParams) WithContext(ctx context.Context) *POSTCapabilitiesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the p o s t capabilities params
func (o *POSTCapabilitiesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the p o s t capabilities params
func (o *POSTCapabilitiesParams) WithBody(body *models.CapabilitiesInput) *POSTCapabilitiesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the p o s t capabilities params
func (o *POSTCapabilitiesParams) SetBody(body *models.CapabilitiesInput) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *POSTCapabilitiesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Body == nil {
		o.Body = new(models.CapabilitiesInput)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
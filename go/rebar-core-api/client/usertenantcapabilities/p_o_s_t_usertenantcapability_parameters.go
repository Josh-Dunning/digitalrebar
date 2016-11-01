package usertenantcapabilities

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

// NewPOSTUsertenantcapabilityParams creates a new POSTUsertenantcapabilityParams object
// with the default values initialized.
func NewPOSTUsertenantcapabilityParams() *POSTUsertenantcapabilityParams {
	var ()
	return &POSTUsertenantcapabilityParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPOSTUsertenantcapabilityParamsWithTimeout creates a new POSTUsertenantcapabilityParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPOSTUsertenantcapabilityParamsWithTimeout(timeout time.Duration) *POSTUsertenantcapabilityParams {
	var ()
	return &POSTUsertenantcapabilityParams{

		timeout: timeout,
	}
}

// NewPOSTUsertenantcapabilityParamsWithContext creates a new POSTUsertenantcapabilityParams object
// with the default values initialized, and the ability to set a context for a request
func NewPOSTUsertenantcapabilityParamsWithContext(ctx context.Context) *POSTUsertenantcapabilityParams {
	var ()
	return &POSTUsertenantcapabilityParams{

		Context: ctx,
	}
}

/*POSTUsertenantcapabilityParams contains all the parameters to send to the API endpoint
for the p o s t usertenantcapability operation typically these are written to a http.Request
*/
type POSTUsertenantcapabilityParams struct {

	/*Body*/
	Body *models.UsertenantcapabilityInput

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the p o s t usertenantcapability params
func (o *POSTUsertenantcapabilityParams) WithTimeout(timeout time.Duration) *POSTUsertenantcapabilityParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the p o s t usertenantcapability params
func (o *POSTUsertenantcapabilityParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the p o s t usertenantcapability params
func (o *POSTUsertenantcapabilityParams) WithContext(ctx context.Context) *POSTUsertenantcapabilityParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the p o s t usertenantcapability params
func (o *POSTUsertenantcapabilityParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the p o s t usertenantcapability params
func (o *POSTUsertenantcapabilityParams) WithBody(body *models.UsertenantcapabilityInput) *POSTUsertenantcapabilityParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the p o s t usertenantcapability params
func (o *POSTUsertenantcapabilityParams) SetBody(body *models.UsertenantcapabilityInput) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *POSTUsertenantcapabilityParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Body == nil {
		o.Body = new(models.UsertenantcapabilityInput)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
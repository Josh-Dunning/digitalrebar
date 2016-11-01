package tenants

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

// NewPUTTenantParams creates a new PUTTenantParams object
// with the default values initialized.
func NewPUTTenantParams() *PUTTenantParams {
	var ()
	return &PUTTenantParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPUTTenantParamsWithTimeout creates a new PUTTenantParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPUTTenantParamsWithTimeout(timeout time.Duration) *PUTTenantParams {
	var ()
	return &PUTTenantParams{

		timeout: timeout,
	}
}

// NewPUTTenantParamsWithContext creates a new PUTTenantParams object
// with the default values initialized, and the ability to set a context for a request
func NewPUTTenantParamsWithContext(ctx context.Context) *PUTTenantParams {
	var ()
	return &PUTTenantParams{

		Context: ctx,
	}
}

/*PUTTenantParams contains all the parameters to send to the API endpoint
for the p u t tenant operation typically these are written to a http.Request
*/
type PUTTenantParams struct {

	/*Body*/
	Body *models.TenantInput
	/*ID*/
	ID string

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the p u t tenant params
func (o *PUTTenantParams) WithTimeout(timeout time.Duration) *PUTTenantParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the p u t tenant params
func (o *PUTTenantParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the p u t tenant params
func (o *PUTTenantParams) WithContext(ctx context.Context) *PUTTenantParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the p u t tenant params
func (o *PUTTenantParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the p u t tenant params
func (o *PUTTenantParams) WithBody(body *models.TenantInput) *PUTTenantParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the p u t tenant params
func (o *PUTTenantParams) SetBody(body *models.TenantInput) {
	o.Body = body
}

// WithID adds the id to the p u t tenant params
func (o *PUTTenantParams) WithID(id string) *PUTTenantParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the p u t tenant params
func (o *PUTTenantParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PUTTenantParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Body == nil {
		o.Body = new(models.TenantInput)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
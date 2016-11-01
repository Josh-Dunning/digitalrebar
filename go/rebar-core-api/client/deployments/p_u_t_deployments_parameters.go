package deployments

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

// NewPUTDeploymentsParams creates a new PUTDeploymentsParams object
// with the default values initialized.
func NewPUTDeploymentsParams() *PUTDeploymentsParams {
	var ()
	return &PUTDeploymentsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPUTDeploymentsParamsWithTimeout creates a new PUTDeploymentsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPUTDeploymentsParamsWithTimeout(timeout time.Duration) *PUTDeploymentsParams {
	var ()
	return &PUTDeploymentsParams{

		timeout: timeout,
	}
}

// NewPUTDeploymentsParamsWithContext creates a new PUTDeploymentsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPUTDeploymentsParamsWithContext(ctx context.Context) *PUTDeploymentsParams {
	var ()
	return &PUTDeploymentsParams{

		Context: ctx,
	}
}

/*PUTDeploymentsParams contains all the parameters to send to the API endpoint
for the p u t deployments operation typically these are written to a http.Request
*/
type PUTDeploymentsParams struct {

	/*Body*/
	Body *models.DeploymentsInput
	/*ID*/
	ID string

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the p u t deployments params
func (o *PUTDeploymentsParams) WithTimeout(timeout time.Duration) *PUTDeploymentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the p u t deployments params
func (o *PUTDeploymentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the p u t deployments params
func (o *PUTDeploymentsParams) WithContext(ctx context.Context) *PUTDeploymentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the p u t deployments params
func (o *PUTDeploymentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the p u t deployments params
func (o *PUTDeploymentsParams) WithBody(body *models.DeploymentsInput) *PUTDeploymentsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the p u t deployments params
func (o *PUTDeploymentsParams) SetBody(body *models.DeploymentsInput) {
	o.Body = body
}

// WithID adds the id to the p u t deployments params
func (o *PUTDeploymentsParams) WithID(id string) *PUTDeploymentsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the p u t deployments params
func (o *PUTDeploymentsParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PUTDeploymentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Body == nil {
		o.Body = new(models.DeploymentsInput)
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
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
)

// NewGETDeploymentsParams creates a new GETDeploymentsParams object
// with the default values initialized.
func NewGETDeploymentsParams() *GETDeploymentsParams {
	var ()
	return &GETDeploymentsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGETDeploymentsParamsWithTimeout creates a new GETDeploymentsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGETDeploymentsParamsWithTimeout(timeout time.Duration) *GETDeploymentsParams {
	var ()
	return &GETDeploymentsParams{

		timeout: timeout,
	}
}

// NewGETDeploymentsParamsWithContext creates a new GETDeploymentsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGETDeploymentsParamsWithContext(ctx context.Context) *GETDeploymentsParams {
	var ()
	return &GETDeploymentsParams{

		Context: ctx,
	}
}

/*GETDeploymentsParams contains all the parameters to send to the API endpoint
for the g e t deployments operation typically these are written to a http.Request
*/
type GETDeploymentsParams struct {

	/*ID*/
	ID string

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the g e t deployments params
func (o *GETDeploymentsParams) WithTimeout(timeout time.Duration) *GETDeploymentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the g e t deployments params
func (o *GETDeploymentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the g e t deployments params
func (o *GETDeploymentsParams) WithContext(ctx context.Context) *GETDeploymentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the g e t deployments params
func (o *GETDeploymentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithID adds the id to the g e t deployments params
func (o *GETDeploymentsParams) WithID(id string) *GETDeploymentsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the g e t deployments params
func (o *GETDeploymentsParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GETDeploymentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
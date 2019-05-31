// Code generated by go-swagger; DO NOT EDIT.

package compute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetInstancesInstanceIDVolumeAttachmentsParams creates a new GetInstancesInstanceIDVolumeAttachmentsParams object
// with the default values initialized.
func NewGetInstancesInstanceIDVolumeAttachmentsParams() *GetInstancesInstanceIDVolumeAttachmentsParams {
	var ()
	return &GetInstancesInstanceIDVolumeAttachmentsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetInstancesInstanceIDVolumeAttachmentsParamsWithTimeout creates a new GetInstancesInstanceIDVolumeAttachmentsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetInstancesInstanceIDVolumeAttachmentsParamsWithTimeout(timeout time.Duration) *GetInstancesInstanceIDVolumeAttachmentsParams {
	var ()
	return &GetInstancesInstanceIDVolumeAttachmentsParams{

		timeout: timeout,
	}
}

// NewGetInstancesInstanceIDVolumeAttachmentsParamsWithContext creates a new GetInstancesInstanceIDVolumeAttachmentsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetInstancesInstanceIDVolumeAttachmentsParamsWithContext(ctx context.Context) *GetInstancesInstanceIDVolumeAttachmentsParams {
	var ()
	return &GetInstancesInstanceIDVolumeAttachmentsParams{

		Context: ctx,
	}
}

// NewGetInstancesInstanceIDVolumeAttachmentsParamsWithHTTPClient creates a new GetInstancesInstanceIDVolumeAttachmentsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetInstancesInstanceIDVolumeAttachmentsParamsWithHTTPClient(client *http.Client) *GetInstancesInstanceIDVolumeAttachmentsParams {
	var ()
	return &GetInstancesInstanceIDVolumeAttachmentsParams{
		HTTPClient: client,
	}
}

/*GetInstancesInstanceIDVolumeAttachmentsParams contains all the parameters to send to the API endpoint
for the get instances instance ID volume attachments operation typically these are written to a http.Request
*/
type GetInstancesInstanceIDVolumeAttachmentsParams struct {

	/*InstanceID
	  The instance identifier

	*/
	InstanceID string
	/*ResourceGroupID
	  Filters the collection to resources within the resource group of the specified identifier

	*/
	ResourceGroupID *string
	/*Tag
	  Filters the collection to resources containing the specified tag

	*/
	Tag *string
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get instances instance ID volume attachments params
func (o *GetInstancesInstanceIDVolumeAttachmentsParams) WithTimeout(timeout time.Duration) *GetInstancesInstanceIDVolumeAttachmentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get instances instance ID volume attachments params
func (o *GetInstancesInstanceIDVolumeAttachmentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get instances instance ID volume attachments params
func (o *GetInstancesInstanceIDVolumeAttachmentsParams) WithContext(ctx context.Context) *GetInstancesInstanceIDVolumeAttachmentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get instances instance ID volume attachments params
func (o *GetInstancesInstanceIDVolumeAttachmentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get instances instance ID volume attachments params
func (o *GetInstancesInstanceIDVolumeAttachmentsParams) WithHTTPClient(client *http.Client) *GetInstancesInstanceIDVolumeAttachmentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get instances instance ID volume attachments params
func (o *GetInstancesInstanceIDVolumeAttachmentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInstanceID adds the instanceID to the get instances instance ID volume attachments params
func (o *GetInstancesInstanceIDVolumeAttachmentsParams) WithInstanceID(instanceID string) *GetInstancesInstanceIDVolumeAttachmentsParams {
	o.SetInstanceID(instanceID)
	return o
}

// SetInstanceID adds the instanceId to the get instances instance ID volume attachments params
func (o *GetInstancesInstanceIDVolumeAttachmentsParams) SetInstanceID(instanceID string) {
	o.InstanceID = instanceID
}

// WithResourceGroupID adds the resourceGroupID to the get instances instance ID volume attachments params
func (o *GetInstancesInstanceIDVolumeAttachmentsParams) WithResourceGroupID(resourceGroupID *string) *GetInstancesInstanceIDVolumeAttachmentsParams {
	o.SetResourceGroupID(resourceGroupID)
	return o
}

// SetResourceGroupID adds the resourceGroupId to the get instances instance ID volume attachments params
func (o *GetInstancesInstanceIDVolumeAttachmentsParams) SetResourceGroupID(resourceGroupID *string) {
	o.ResourceGroupID = resourceGroupID
}

// WithTag adds the tag to the get instances instance ID volume attachments params
func (o *GetInstancesInstanceIDVolumeAttachmentsParams) WithTag(tag *string) *GetInstancesInstanceIDVolumeAttachmentsParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the get instances instance ID volume attachments params
func (o *GetInstancesInstanceIDVolumeAttachmentsParams) SetTag(tag *string) {
	o.Tag = tag
}

// WithVersion adds the version to the get instances instance ID volume attachments params
func (o *GetInstancesInstanceIDVolumeAttachmentsParams) WithVersion(version string) *GetInstancesInstanceIDVolumeAttachmentsParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get instances instance ID volume attachments params
func (o *GetInstancesInstanceIDVolumeAttachmentsParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *GetInstancesInstanceIDVolumeAttachmentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param instance_id
	if err := r.SetPathParam("instance_id", o.InstanceID); err != nil {
		return err
	}

	if o.ResourceGroupID != nil {

		// query param resource_group.id
		var qrResourceGroupID string
		if o.ResourceGroupID != nil {
			qrResourceGroupID = *o.ResourceGroupID
		}
		qResourceGroupID := qrResourceGroupID
		if qResourceGroupID != "" {
			if err := r.SetQueryParam("resource_group.id", qResourceGroupID); err != nil {
				return err
			}
		}

	}

	if o.Tag != nil {

		// query param tag
		var qrTag string
		if o.Tag != nil {
			qrTag = *o.Tag
		}
		qTag := qrTag
		if qTag != "" {
			if err := r.SetQueryParam("tag", qTag); err != nil {
				return err
			}
		}

	}

	// query param version
	qrVersion := o.Version
	qVersion := qrVersion
	if qVersion != "" {
		if err := r.SetQueryParam("version", qVersion); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
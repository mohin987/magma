// Code generated by go-swagger; DO NOT EDIT.

package networks

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

	models "magma/orc8r/cloud/api/v1/go/models"
)

// NewPostNetworksNetworkIDDNSRecordsDomainParams creates a new PostNetworksNetworkIDDNSRecordsDomainParams object
// with the default values initialized.
func NewPostNetworksNetworkIDDNSRecordsDomainParams() *PostNetworksNetworkIDDNSRecordsDomainParams {
	var ()
	return &PostNetworksNetworkIDDNSRecordsDomainParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostNetworksNetworkIDDNSRecordsDomainParamsWithTimeout creates a new PostNetworksNetworkIDDNSRecordsDomainParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostNetworksNetworkIDDNSRecordsDomainParamsWithTimeout(timeout time.Duration) *PostNetworksNetworkIDDNSRecordsDomainParams {
	var ()
	return &PostNetworksNetworkIDDNSRecordsDomainParams{

		timeout: timeout,
	}
}

// NewPostNetworksNetworkIDDNSRecordsDomainParamsWithContext creates a new PostNetworksNetworkIDDNSRecordsDomainParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostNetworksNetworkIDDNSRecordsDomainParamsWithContext(ctx context.Context) *PostNetworksNetworkIDDNSRecordsDomainParams {
	var ()
	return &PostNetworksNetworkIDDNSRecordsDomainParams{

		Context: ctx,
	}
}

// NewPostNetworksNetworkIDDNSRecordsDomainParamsWithHTTPClient creates a new PostNetworksNetworkIDDNSRecordsDomainParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostNetworksNetworkIDDNSRecordsDomainParamsWithHTTPClient(client *http.Client) *PostNetworksNetworkIDDNSRecordsDomainParams {
	var ()
	return &PostNetworksNetworkIDDNSRecordsDomainParams{
		HTTPClient: client,
	}
}

/*PostNetworksNetworkIDDNSRecordsDomainParams contains all the parameters to send to the API endpoint
for the post networks network ID DNS records domain operation typically these are written to a http.Request
*/
type PostNetworksNetworkIDDNSRecordsDomainParams struct {

	/*Domain
	  DNS record domain

	*/
	Domain string
	/*NetworkID
	  Network ID

	*/
	NetworkID string
	/*Record
	  Custom DNS record for the domain

	*/
	Record *models.DNSConfigRecord

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post networks network ID DNS records domain params
func (o *PostNetworksNetworkIDDNSRecordsDomainParams) WithTimeout(timeout time.Duration) *PostNetworksNetworkIDDNSRecordsDomainParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post networks network ID DNS records domain params
func (o *PostNetworksNetworkIDDNSRecordsDomainParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post networks network ID DNS records domain params
func (o *PostNetworksNetworkIDDNSRecordsDomainParams) WithContext(ctx context.Context) *PostNetworksNetworkIDDNSRecordsDomainParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post networks network ID DNS records domain params
func (o *PostNetworksNetworkIDDNSRecordsDomainParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post networks network ID DNS records domain params
func (o *PostNetworksNetworkIDDNSRecordsDomainParams) WithHTTPClient(client *http.Client) *PostNetworksNetworkIDDNSRecordsDomainParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post networks network ID DNS records domain params
func (o *PostNetworksNetworkIDDNSRecordsDomainParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDomain adds the domain to the post networks network ID DNS records domain params
func (o *PostNetworksNetworkIDDNSRecordsDomainParams) WithDomain(domain string) *PostNetworksNetworkIDDNSRecordsDomainParams {
	o.SetDomain(domain)
	return o
}

// SetDomain adds the domain to the post networks network ID DNS records domain params
func (o *PostNetworksNetworkIDDNSRecordsDomainParams) SetDomain(domain string) {
	o.Domain = domain
}

// WithNetworkID adds the networkID to the post networks network ID DNS records domain params
func (o *PostNetworksNetworkIDDNSRecordsDomainParams) WithNetworkID(networkID string) *PostNetworksNetworkIDDNSRecordsDomainParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the post networks network ID DNS records domain params
func (o *PostNetworksNetworkIDDNSRecordsDomainParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithRecord adds the record to the post networks network ID DNS records domain params
func (o *PostNetworksNetworkIDDNSRecordsDomainParams) WithRecord(record *models.DNSConfigRecord) *PostNetworksNetworkIDDNSRecordsDomainParams {
	o.SetRecord(record)
	return o
}

// SetRecord adds the record to the post networks network ID DNS records domain params
func (o *PostNetworksNetworkIDDNSRecordsDomainParams) SetRecord(record *models.DNSConfigRecord) {
	o.Record = record
}

// WriteToRequest writes these params to a swagger request
func (o *PostNetworksNetworkIDDNSRecordsDomainParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param domain
	if err := r.SetPathParam("domain", o.Domain); err != nil {
		return err
	}

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	if o.Record != nil {
		if err := r.SetBodyParam(o.Record); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
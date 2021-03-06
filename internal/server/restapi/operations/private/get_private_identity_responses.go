// Code generated by go-swagger; DO NOT EDIT.

package private

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"identityserver/pkg/models/rest"
)

// GetPrivateIdentityOKCode is the HTTP code returned for type GetPrivateIdentityOK
const GetPrivateIdentityOKCode int = 200

/*GetPrivateIdentityOK View private Identity

swagger:response getPrivateIdentityOK
*/
type GetPrivateIdentityOK struct {

	/*
	  In: Body
	*/
	Payload *rest.IdentityResponse `json:"body,omitempty"`
}

// NewGetPrivateIdentityOK creates GetPrivateIdentityOK with default headers values
func NewGetPrivateIdentityOK() *GetPrivateIdentityOK {

	return &GetPrivateIdentityOK{}
}

// WithPayload adds the payload to the get private identity o k response
func (o *GetPrivateIdentityOK) WithPayload(payload *rest.IdentityResponse) *GetPrivateIdentityOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get private identity o k response
func (o *GetPrivateIdentityOK) SetPayload(payload *rest.IdentityResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPrivateIdentityOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetPrivateIdentityUnauthorizedCode is the HTTP code returned for type GetPrivateIdentityUnauthorized
const GetPrivateIdentityUnauthorizedCode int = 401

/*GetPrivateIdentityUnauthorized Authentication information is missing or invalid

swagger:response getPrivateIdentityUnauthorized
*/
type GetPrivateIdentityUnauthorized struct {
	/*

	 */
	WWWAuthenticate string `json:"WWW_Authenticate"`
}

// NewGetPrivateIdentityUnauthorized creates GetPrivateIdentityUnauthorized with default headers values
func NewGetPrivateIdentityUnauthorized() *GetPrivateIdentityUnauthorized {

	return &GetPrivateIdentityUnauthorized{}
}

// WithWWWAuthenticate adds the wWWAuthenticate to the get private identity unauthorized response
func (o *GetPrivateIdentityUnauthorized) WithWWWAuthenticate(wWWAuthenticate string) *GetPrivateIdentityUnauthorized {
	o.WWWAuthenticate = wWWAuthenticate
	return o
}

// SetWWWAuthenticate sets the wWWAuthenticate to the get private identity unauthorized response
func (o *GetPrivateIdentityUnauthorized) SetWWWAuthenticate(wWWAuthenticate string) {
	o.WWWAuthenticate = wWWAuthenticate
}

// WriteResponse to the client
func (o *GetPrivateIdentityUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header WWW_Authenticate

	wWWAuthenticate := o.WWWAuthenticate
	if wWWAuthenticate != "" {
		rw.Header().Set("WWW_Authenticate", wWWAuthenticate)
	}

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// GetPrivateIdentityPreconditionFailedCode is the HTTP code returned for type GetPrivateIdentityPreconditionFailed
const GetPrivateIdentityPreconditionFailedCode int = 412

/*GetPrivateIdentityPreconditionFailed Identity is not set

swagger:response getPrivateIdentityPreconditionFailed
*/
type GetPrivateIdentityPreconditionFailed struct {

	/*
	  In: Body
	*/
	Payload *rest.Error `json:"body,omitempty"`
}

// NewGetPrivateIdentityPreconditionFailed creates GetPrivateIdentityPreconditionFailed with default headers values
func NewGetPrivateIdentityPreconditionFailed() *GetPrivateIdentityPreconditionFailed {

	return &GetPrivateIdentityPreconditionFailed{}
}

// WithPayload adds the payload to the get private identity precondition failed response
func (o *GetPrivateIdentityPreconditionFailed) WithPayload(payload *rest.Error) *GetPrivateIdentityPreconditionFailed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get private identity precondition failed response
func (o *GetPrivateIdentityPreconditionFailed) SetPayload(payload *rest.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPrivateIdentityPreconditionFailed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(412)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetPrivateIdentityDefault Internal server error

swagger:response getPrivateIdentityDefault
*/
type GetPrivateIdentityDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *rest.Error `json:"body,omitempty"`
}

// NewGetPrivateIdentityDefault creates GetPrivateIdentityDefault with default headers values
func NewGetPrivateIdentityDefault(code int) *GetPrivateIdentityDefault {
	if code <= 0 {
		code = 500
	}

	return &GetPrivateIdentityDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get private identity default response
func (o *GetPrivateIdentityDefault) WithStatusCode(code int) *GetPrivateIdentityDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get private identity default response
func (o *GetPrivateIdentityDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get private identity default response
func (o *GetPrivateIdentityDefault) WithPayload(payload *rest.Error) *GetPrivateIdentityDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get private identity default response
func (o *GetPrivateIdentityDefault) SetPayload(payload *rest.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPrivateIdentityDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/runtime"

	models "webservice/Swagger/simpleswagger/models"
)

// GetIDOKCode is the HTTP code returned for type GetIDOK
const GetIDOKCode int = 200

/*GetIDOK get file from jenkins

swagger:response getIdOK
*/
type GetIDOK struct {
	/*

	 */
	ContentDisposition string `json:"Content-Disposition"`
	/*

	 */
	ContentType string `json:"Content-Type"`

	/*
	  In: Body
	*/
	Payload io.ReadCloser `json:"body,omitempty"`
}

// NewGetIDOK creates GetIDOK with default headers values
func NewGetIDOK() *GetIDOK {

	return &GetIDOK{}
}

// WithContentDisposition adds the contentDisposition to the get Id o k response
func (o *GetIDOK) WithContentDisposition(contentDisposition string) *GetIDOK {
	o.ContentDisposition = contentDisposition
	return o
}

// SetContentDisposition sets the contentDisposition to the get Id o k response
func (o *GetIDOK) SetContentDisposition(contentDisposition string) {
	o.ContentDisposition = contentDisposition
}

// WithContentType adds the contentType to the get Id o k response
func (o *GetIDOK) WithContentType(contentType string) *GetIDOK {
	o.ContentType = contentType
	return o
}

// SetContentType sets the contentType to the get Id o k response
func (o *GetIDOK) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithPayload adds the payload to the get Id o k response
func (o *GetIDOK) WithPayload(payload io.ReadCloser) *GetIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get Id o k response
func (o *GetIDOK) SetPayload(payload io.ReadCloser) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Content-Disposition

	contentDisposition := o.ContentDisposition
	if contentDisposition != "" {
		rw.Header().Set("Content-Disposition", contentDisposition)
	}

	// response header Content-Type

	contentType := o.ContentType
	if contentType != "" {
		rw.Header().Set("Content-Type", contentType)
	}

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*GetIDDefault generic error response

swagger:response getIdDefault
*/
type GetIDDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetIDDefault creates GetIDDefault with default headers values
func NewGetIDDefault(code int) *GetIDDefault {
	if code <= 0 {
		code = 500
	}

	return &GetIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get ID default response
func (o *GetIDDefault) WithStatusCode(code int) *GetIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get ID default response
func (o *GetIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get ID default response
func (o *GetIDDefault) WithPayload(payload *models.Error) *GetIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get ID default response
func (o *GetIDDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

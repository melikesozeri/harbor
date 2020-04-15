// Code generated by go-swagger; DO NOT EDIT.

package repository

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/goharbor/harbor/src/server/v2.0/models"
)

// UpdateRepositoryOKCode is the HTTP code returned for type UpdateRepositoryOK
const UpdateRepositoryOKCode int = 200

/*UpdateRepositoryOK Success

swagger:response updateRepositoryOK
*/
type UpdateRepositoryOK struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`
}

// NewUpdateRepositoryOK creates UpdateRepositoryOK with default headers values
func NewUpdateRepositoryOK() *UpdateRepositoryOK {

	return &UpdateRepositoryOK{}
}

// WithXRequestID adds the xRequestId to the update repository o k response
func (o *UpdateRepositoryOK) WithXRequestID(xRequestID string) *UpdateRepositoryOK {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the update repository o k response
func (o *UpdateRepositoryOK) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WriteResponse to the client
func (o *UpdateRepositoryOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header X-Request-Id

	xRequestID := o.XRequestID
	if xRequestID != "" {
		rw.Header().Set("X-Request-Id", xRequestID)
	}

	rw.Header().Del(runtime.HeaderContentType) // Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// UpdateRepositoryBadRequestCode is the HTTP code returned for type UpdateRepositoryBadRequest
const UpdateRepositoryBadRequestCode int = 400

/*UpdateRepositoryBadRequest Bad request

swagger:response updateRepositoryBadRequest
*/
type UpdateRepositoryBadRequest struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload models.Errors `json:"body,omitempty"`
}

// NewUpdateRepositoryBadRequest creates UpdateRepositoryBadRequest with default headers values
func NewUpdateRepositoryBadRequest() *UpdateRepositoryBadRequest {

	return &UpdateRepositoryBadRequest{}
}

// WithXRequestID adds the xRequestId to the update repository bad request response
func (o *UpdateRepositoryBadRequest) WithXRequestID(xRequestID string) *UpdateRepositoryBadRequest {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the update repository bad request response
func (o *UpdateRepositoryBadRequest) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the update repository bad request response
func (o *UpdateRepositoryBadRequest) WithPayload(payload models.Errors) *UpdateRepositoryBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update repository bad request response
func (o *UpdateRepositoryBadRequest) SetPayload(payload models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateRepositoryBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header X-Request-Id

	xRequestID := o.XRequestID
	if xRequestID != "" {
		rw.Header().Set("X-Request-Id", xRequestID)
	}

	rw.WriteHeader(400)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.Errors{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// UpdateRepositoryUnauthorizedCode is the HTTP code returned for type UpdateRepositoryUnauthorized
const UpdateRepositoryUnauthorizedCode int = 401

/*UpdateRepositoryUnauthorized Unauthorized

swagger:response updateRepositoryUnauthorized
*/
type UpdateRepositoryUnauthorized struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload models.Errors `json:"body,omitempty"`
}

// NewUpdateRepositoryUnauthorized creates UpdateRepositoryUnauthorized with default headers values
func NewUpdateRepositoryUnauthorized() *UpdateRepositoryUnauthorized {

	return &UpdateRepositoryUnauthorized{}
}

// WithXRequestID adds the xRequestId to the update repository unauthorized response
func (o *UpdateRepositoryUnauthorized) WithXRequestID(xRequestID string) *UpdateRepositoryUnauthorized {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the update repository unauthorized response
func (o *UpdateRepositoryUnauthorized) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the update repository unauthorized response
func (o *UpdateRepositoryUnauthorized) WithPayload(payload models.Errors) *UpdateRepositoryUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update repository unauthorized response
func (o *UpdateRepositoryUnauthorized) SetPayload(payload models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateRepositoryUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header X-Request-Id

	xRequestID := o.XRequestID
	if xRequestID != "" {
		rw.Header().Set("X-Request-Id", xRequestID)
	}

	rw.WriteHeader(401)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.Errors{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// UpdateRepositoryForbiddenCode is the HTTP code returned for type UpdateRepositoryForbidden
const UpdateRepositoryForbiddenCode int = 403

/*UpdateRepositoryForbidden Forbidden

swagger:response updateRepositoryForbidden
*/
type UpdateRepositoryForbidden struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload models.Errors `json:"body,omitempty"`
}

// NewUpdateRepositoryForbidden creates UpdateRepositoryForbidden with default headers values
func NewUpdateRepositoryForbidden() *UpdateRepositoryForbidden {

	return &UpdateRepositoryForbidden{}
}

// WithXRequestID adds the xRequestId to the update repository forbidden response
func (o *UpdateRepositoryForbidden) WithXRequestID(xRequestID string) *UpdateRepositoryForbidden {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the update repository forbidden response
func (o *UpdateRepositoryForbidden) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the update repository forbidden response
func (o *UpdateRepositoryForbidden) WithPayload(payload models.Errors) *UpdateRepositoryForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update repository forbidden response
func (o *UpdateRepositoryForbidden) SetPayload(payload models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateRepositoryForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header X-Request-Id

	xRequestID := o.XRequestID
	if xRequestID != "" {
		rw.Header().Set("X-Request-Id", xRequestID)
	}

	rw.WriteHeader(403)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.Errors{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// UpdateRepositoryNotFoundCode is the HTTP code returned for type UpdateRepositoryNotFound
const UpdateRepositoryNotFoundCode int = 404

/*UpdateRepositoryNotFound Not found

swagger:response updateRepositoryNotFound
*/
type UpdateRepositoryNotFound struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload models.Errors `json:"body,omitempty"`
}

// NewUpdateRepositoryNotFound creates UpdateRepositoryNotFound with default headers values
func NewUpdateRepositoryNotFound() *UpdateRepositoryNotFound {

	return &UpdateRepositoryNotFound{}
}

// WithXRequestID adds the xRequestId to the update repository not found response
func (o *UpdateRepositoryNotFound) WithXRequestID(xRequestID string) *UpdateRepositoryNotFound {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the update repository not found response
func (o *UpdateRepositoryNotFound) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the update repository not found response
func (o *UpdateRepositoryNotFound) WithPayload(payload models.Errors) *UpdateRepositoryNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update repository not found response
func (o *UpdateRepositoryNotFound) SetPayload(payload models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateRepositoryNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header X-Request-Id

	xRequestID := o.XRequestID
	if xRequestID != "" {
		rw.Header().Set("X-Request-Id", xRequestID)
	}

	rw.WriteHeader(404)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.Errors{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// UpdateRepositoryInternalServerErrorCode is the HTTP code returned for type UpdateRepositoryInternalServerError
const UpdateRepositoryInternalServerErrorCode int = 500

/*UpdateRepositoryInternalServerError Internal server error

swagger:response updateRepositoryInternalServerError
*/
type UpdateRepositoryInternalServerError struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload models.Errors `json:"body,omitempty"`
}

// NewUpdateRepositoryInternalServerError creates UpdateRepositoryInternalServerError with default headers values
func NewUpdateRepositoryInternalServerError() *UpdateRepositoryInternalServerError {

	return &UpdateRepositoryInternalServerError{}
}

// WithXRequestID adds the xRequestId to the update repository internal server error response
func (o *UpdateRepositoryInternalServerError) WithXRequestID(xRequestID string) *UpdateRepositoryInternalServerError {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the update repository internal server error response
func (o *UpdateRepositoryInternalServerError) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the update repository internal server error response
func (o *UpdateRepositoryInternalServerError) WithPayload(payload models.Errors) *UpdateRepositoryInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update repository internal server error response
func (o *UpdateRepositoryInternalServerError) SetPayload(payload models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateRepositoryInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header X-Request-Id

	xRequestID := o.XRequestID
	if xRequestID != "" {
		rw.Header().Set("X-Request-Id", xRequestID)
	}

	rw.WriteHeader(500)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.Errors{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

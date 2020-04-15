// Code generated by go-swagger; DO NOT EDIT.

package artifact

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/goharbor/harbor/src/server/v2.0/models"
)

// CopyArtifactCreatedCode is the HTTP code returned for type CopyArtifactCreated
const CopyArtifactCreatedCode int = 201

/*CopyArtifactCreated Created

swagger:response copyArtifactCreated
*/
type CopyArtifactCreated struct {
	/*The location of the resource

	 */
	Location string `json:"Location"`
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`
}

// NewCopyArtifactCreated creates CopyArtifactCreated with default headers values
func NewCopyArtifactCreated() *CopyArtifactCreated {

	return &CopyArtifactCreated{}
}

// WithLocation adds the location to the copy artifact created response
func (o *CopyArtifactCreated) WithLocation(location string) *CopyArtifactCreated {
	o.Location = location
	return o
}

// SetLocation sets the location to the copy artifact created response
func (o *CopyArtifactCreated) SetLocation(location string) {
	o.Location = location
}

// WithXRequestID adds the xRequestId to the copy artifact created response
func (o *CopyArtifactCreated) WithXRequestID(xRequestID string) *CopyArtifactCreated {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the copy artifact created response
func (o *CopyArtifactCreated) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WriteResponse to the client
func (o *CopyArtifactCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Location

	location := o.Location
	if location != "" {
		rw.Header().Set("Location", location)
	}

	// response header X-Request-Id

	xRequestID := o.XRequestID
	if xRequestID != "" {
		rw.Header().Set("X-Request-Id", xRequestID)
	}

	rw.Header().Del(runtime.HeaderContentType) // Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

// CopyArtifactBadRequestCode is the HTTP code returned for type CopyArtifactBadRequest
const CopyArtifactBadRequestCode int = 400

/*CopyArtifactBadRequest Bad request

swagger:response copyArtifactBadRequest
*/
type CopyArtifactBadRequest struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload models.Errors `json:"body,omitempty"`
}

// NewCopyArtifactBadRequest creates CopyArtifactBadRequest with default headers values
func NewCopyArtifactBadRequest() *CopyArtifactBadRequest {

	return &CopyArtifactBadRequest{}
}

// WithXRequestID adds the xRequestId to the copy artifact bad request response
func (o *CopyArtifactBadRequest) WithXRequestID(xRequestID string) *CopyArtifactBadRequest {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the copy artifact bad request response
func (o *CopyArtifactBadRequest) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the copy artifact bad request response
func (o *CopyArtifactBadRequest) WithPayload(payload models.Errors) *CopyArtifactBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the copy artifact bad request response
func (o *CopyArtifactBadRequest) SetPayload(payload models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CopyArtifactBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// CopyArtifactUnauthorizedCode is the HTTP code returned for type CopyArtifactUnauthorized
const CopyArtifactUnauthorizedCode int = 401

/*CopyArtifactUnauthorized Unauthorized

swagger:response copyArtifactUnauthorized
*/
type CopyArtifactUnauthorized struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload models.Errors `json:"body,omitempty"`
}

// NewCopyArtifactUnauthorized creates CopyArtifactUnauthorized with default headers values
func NewCopyArtifactUnauthorized() *CopyArtifactUnauthorized {

	return &CopyArtifactUnauthorized{}
}

// WithXRequestID adds the xRequestId to the copy artifact unauthorized response
func (o *CopyArtifactUnauthorized) WithXRequestID(xRequestID string) *CopyArtifactUnauthorized {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the copy artifact unauthorized response
func (o *CopyArtifactUnauthorized) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the copy artifact unauthorized response
func (o *CopyArtifactUnauthorized) WithPayload(payload models.Errors) *CopyArtifactUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the copy artifact unauthorized response
func (o *CopyArtifactUnauthorized) SetPayload(payload models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CopyArtifactUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// CopyArtifactForbiddenCode is the HTTP code returned for type CopyArtifactForbidden
const CopyArtifactForbiddenCode int = 403

/*CopyArtifactForbidden Forbidden

swagger:response copyArtifactForbidden
*/
type CopyArtifactForbidden struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload models.Errors `json:"body,omitempty"`
}

// NewCopyArtifactForbidden creates CopyArtifactForbidden with default headers values
func NewCopyArtifactForbidden() *CopyArtifactForbidden {

	return &CopyArtifactForbidden{}
}

// WithXRequestID adds the xRequestId to the copy artifact forbidden response
func (o *CopyArtifactForbidden) WithXRequestID(xRequestID string) *CopyArtifactForbidden {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the copy artifact forbidden response
func (o *CopyArtifactForbidden) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the copy artifact forbidden response
func (o *CopyArtifactForbidden) WithPayload(payload models.Errors) *CopyArtifactForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the copy artifact forbidden response
func (o *CopyArtifactForbidden) SetPayload(payload models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CopyArtifactForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// CopyArtifactNotFoundCode is the HTTP code returned for type CopyArtifactNotFound
const CopyArtifactNotFoundCode int = 404

/*CopyArtifactNotFound Not found

swagger:response copyArtifactNotFound
*/
type CopyArtifactNotFound struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload models.Errors `json:"body,omitempty"`
}

// NewCopyArtifactNotFound creates CopyArtifactNotFound with default headers values
func NewCopyArtifactNotFound() *CopyArtifactNotFound {

	return &CopyArtifactNotFound{}
}

// WithXRequestID adds the xRequestId to the copy artifact not found response
func (o *CopyArtifactNotFound) WithXRequestID(xRequestID string) *CopyArtifactNotFound {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the copy artifact not found response
func (o *CopyArtifactNotFound) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the copy artifact not found response
func (o *CopyArtifactNotFound) WithPayload(payload models.Errors) *CopyArtifactNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the copy artifact not found response
func (o *CopyArtifactNotFound) SetPayload(payload models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CopyArtifactNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// CopyArtifactInternalServerErrorCode is the HTTP code returned for type CopyArtifactInternalServerError
const CopyArtifactInternalServerErrorCode int = 500

/*CopyArtifactInternalServerError Internal server error

swagger:response copyArtifactInternalServerError
*/
type CopyArtifactInternalServerError struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload models.Errors `json:"body,omitempty"`
}

// NewCopyArtifactInternalServerError creates CopyArtifactInternalServerError with default headers values
func NewCopyArtifactInternalServerError() *CopyArtifactInternalServerError {

	return &CopyArtifactInternalServerError{}
}

// WithXRequestID adds the xRequestId to the copy artifact internal server error response
func (o *CopyArtifactInternalServerError) WithXRequestID(xRequestID string) *CopyArtifactInternalServerError {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the copy artifact internal server error response
func (o *CopyArtifactInternalServerError) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the copy artifact internal server error response
func (o *CopyArtifactInternalServerError) WithPayload(payload models.Errors) *CopyArtifactInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the copy artifact internal server error response
func (o *CopyArtifactInternalServerError) SetPayload(payload models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CopyArtifactInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

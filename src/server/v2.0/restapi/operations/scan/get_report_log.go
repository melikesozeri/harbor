// Code generated by go-swagger; DO NOT EDIT.

package scan

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetReportLogHandlerFunc turns a function with the right signature into a get report log handler
type GetReportLogHandlerFunc func(GetReportLogParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetReportLogHandlerFunc) Handle(params GetReportLogParams) middleware.Responder {
	return fn(params)
}

// GetReportLogHandler interface for that can handle valid get report log params
type GetReportLogHandler interface {
	Handle(GetReportLogParams) middleware.Responder
}

// NewGetReportLog creates a new http.Handler for the get report log operation
func NewGetReportLog(ctx *middleware.Context, handler GetReportLogHandler) *GetReportLog {
	return &GetReportLog{Context: ctx, Handler: handler}
}

/*GetReportLog swagger:route GET /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/scan/{report_id}/log scan getReportLog

Get the log of the scan report

Get the log of the scan report

*/
type GetReportLog struct {
	Context *middleware.Context
	Handler GetReportLogHandler
}

func (o *GetReportLog) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetReportLogParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

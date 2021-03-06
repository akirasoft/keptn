// Code generated by go-swagger; DO NOT EDIT.

package event

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	errors "github.com/go-openapi/errors"
	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"
	validate "github.com/go-openapi/validate"
)

// SaveEventHandlerFunc turns a function with the right signature into a save event handler
type SaveEventHandlerFunc func(SaveEventParams) middleware.Responder

// Handle executing the request and returning a response
func (fn SaveEventHandlerFunc) Handle(params SaveEventParams) middleware.Responder {
	return fn(params)
}

// SaveEventHandler interface for that can handle valid save event params
type SaveEventHandler interface {
	Handle(SaveEventParams) middleware.Responder
}

// NewSaveEvent creates a new http.Handler for the save event operation
func NewSaveEvent(ctx *middleware.Context, handler SaveEventHandler) *SaveEvent {
	return &SaveEvent{Context: ctx, Handler: handler}
}

/*SaveEvent swagger:route POST /event event saveEvent

Saves an event to the datastore

*/
type SaveEvent struct {
	Context *middleware.Context
	Handler SaveEventHandler
}

func (o *SaveEvent) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewSaveEventParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// SaveEventBody save event body
// swagger:model SaveEventBody
type SaveEventBody struct {

	// contenttype
	Contenttype string `json:"contenttype,omitempty"`

	// data
	Data interface{} `json:"data,omitempty"`

	// extensions
	Extensions interface{} `json:"extensions,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// source
	// Required: true
	Source *string `json:"source"`

	// specversion
	// Required: true
	Specversion *string `json:"specversion"`

	// time
	// Format: date-time
	Time strfmt.DateTime `json:"time,omitempty"`

	// type
	// Required: true
	Type *string `json:"type"`

	// shkeptncontext
	Shkeptncontext string `json:"shkeptncontext,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *SaveEventBody) UnmarshalJSON(raw []byte) error {
	// SaveEventParamsBodyAO0
	var dataSaveEventParamsBodyAO0 struct {
		Contenttype string `json:"contenttype,omitempty"`

		Data interface{} `json:"data,omitempty"`

		Extensions interface{} `json:"extensions,omitempty"`

		ID *string `json:"id"`

		Source *string `json:"source"`

		Specversion *string `json:"specversion"`

		Time strfmt.DateTime `json:"time,omitempty"`

		Type *string `json:"type"`
	}
	if err := swag.ReadJSON(raw, &dataSaveEventParamsBodyAO0); err != nil {
		return err
	}

	o.Contenttype = dataSaveEventParamsBodyAO0.Contenttype

	o.Data = dataSaveEventParamsBodyAO0.Data

	o.Extensions = dataSaveEventParamsBodyAO0.Extensions

	o.ID = dataSaveEventParamsBodyAO0.ID

	o.Source = dataSaveEventParamsBodyAO0.Source

	o.Specversion = dataSaveEventParamsBodyAO0.Specversion

	o.Time = dataSaveEventParamsBodyAO0.Time

	o.Type = dataSaveEventParamsBodyAO0.Type

	// SaveEventParamsBodyAO1
	var dataSaveEventParamsBodyAO1 struct {
		Shkeptncontext string `json:"shkeptncontext,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataSaveEventParamsBodyAO1); err != nil {
		return err
	}

	o.Shkeptncontext = dataSaveEventParamsBodyAO1.Shkeptncontext

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o SaveEventBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	var dataSaveEventParamsBodyAO0 struct {
		Contenttype string `json:"contenttype,omitempty"`

		Data interface{} `json:"data,omitempty"`

		Extensions interface{} `json:"extensions,omitempty"`

		ID *string `json:"id"`

		Source *string `json:"source"`

		Specversion *string `json:"specversion"`

		Time strfmt.DateTime `json:"time,omitempty"`

		Type *string `json:"type"`
	}

	dataSaveEventParamsBodyAO0.Contenttype = o.Contenttype

	dataSaveEventParamsBodyAO0.Data = o.Data

	dataSaveEventParamsBodyAO0.Extensions = o.Extensions

	dataSaveEventParamsBodyAO0.ID = o.ID

	dataSaveEventParamsBodyAO0.Source = o.Source

	dataSaveEventParamsBodyAO0.Specversion = o.Specversion

	dataSaveEventParamsBodyAO0.Time = o.Time

	dataSaveEventParamsBodyAO0.Type = o.Type

	jsonDataSaveEventParamsBodyAO0, errSaveEventParamsBodyAO0 := swag.WriteJSON(dataSaveEventParamsBodyAO0)
	if errSaveEventParamsBodyAO0 != nil {
		return nil, errSaveEventParamsBodyAO0
	}
	_parts = append(_parts, jsonDataSaveEventParamsBodyAO0)

	var dataSaveEventParamsBodyAO1 struct {
		Shkeptncontext string `json:"shkeptncontext,omitempty"`
	}

	dataSaveEventParamsBodyAO1.Shkeptncontext = o.Shkeptncontext

	jsonDataSaveEventParamsBodyAO1, errSaveEventParamsBodyAO1 := swag.WriteJSON(dataSaveEventParamsBodyAO1)
	if errSaveEventParamsBodyAO1 != nil {
		return nil, errSaveEventParamsBodyAO1
	}
	_parts = append(_parts, jsonDataSaveEventParamsBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this save event body
func (o *SaveEventBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSpecversion(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SaveEventBody) validateID(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"id", "body", o.ID); err != nil {
		return err
	}

	return nil
}

func (o *SaveEventBody) validateSource(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"source", "body", o.Source); err != nil {
		return err
	}

	return nil
}

func (o *SaveEventBody) validateSpecversion(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"specversion", "body", o.Specversion); err != nil {
		return err
	}

	return nil
}

func (o *SaveEventBody) validateTime(formats strfmt.Registry) error {

	if swag.IsZero(o.Time) { // not required
		return nil
	}

	if err := validate.FormatOf("body"+"."+"time", "body", "date-time", o.Time.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *SaveEventBody) validateType(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"type", "body", o.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SaveEventBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SaveEventBody) UnmarshalBinary(b []byte) error {
	var res SaveEventBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// SaveEventDefaultBody save event default body
// swagger:model SaveEventDefaultBody
type SaveEventDefaultBody struct {

	// code
	Code int64 `json:"code,omitempty"`

	// fields
	Fields string `json:"fields,omitempty"`

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this save event default body
func (o *SaveEventDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SaveEventDefaultBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("saveEvent default"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SaveEventDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SaveEventDefaultBody) UnmarshalBinary(b []byte) error {
	var res SaveEventDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

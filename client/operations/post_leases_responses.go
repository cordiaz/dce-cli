// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// PostLeasesReader is a Reader for the PostLeases structure.
type PostLeasesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLeasesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostLeasesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostLeasesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostLeasesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostLeasesConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostLeasesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLeasesCreated creates a PostLeasesCreated with default headers values
func NewPostLeasesCreated() *PostLeasesCreated {
	return &PostLeasesCreated{}
}

/*PostLeasesCreated handles this case with default header values.

PostLeasesCreated post leases created
*/
type PostLeasesCreated struct {
	AccessControlAllowHeaders string

	AccessControlAllowMethods string

	AccessControlAllowOrigin string

	Payload *PostLeasesCreatedBody
}

func (o *PostLeasesCreated) Error() string {
	return fmt.Sprintf("[POST /leases][%d] postLeasesCreated  %+v", 201, o.Payload)
}

func (o *PostLeasesCreated) GetPayload() *PostLeasesCreatedBody {
	return o.Payload
}

func (o *PostLeasesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Access-Control-Allow-Headers
	o.AccessControlAllowHeaders = response.GetHeader("Access-Control-Allow-Headers")

	// response header Access-Control-Allow-Methods
	o.AccessControlAllowMethods = response.GetHeader("Access-Control-Allow-Methods")

	// response header Access-Control-Allow-Origin
	o.AccessControlAllowOrigin = response.GetHeader("Access-Control-Allow-Origin")

	o.Payload = new(PostLeasesCreatedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLeasesBadRequest creates a PostLeasesBadRequest with default headers values
func NewPostLeasesBadRequest() *PostLeasesBadRequest {
	return &PostLeasesBadRequest{}
}

/*PostLeasesBadRequest handles this case with default header values.

If the "expiresOn" date specified is non-zero but less than the current epoch date,  "Requested lease has a desired expiry date less than today: <date>" or "Failed to Parse Request Body" if the request body is blank or incorrectly formatted.

*/
type PostLeasesBadRequest struct {
}

func (o *PostLeasesBadRequest) Error() string {
	return fmt.Sprintf("[POST /leases][%d] postLeasesBadRequest ", 400)
}

func (o *PostLeasesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostLeasesForbidden creates a PostLeasesForbidden with default headers values
func NewPostLeasesForbidden() *PostLeasesForbidden {
	return &PostLeasesForbidden{}
}

/*PostLeasesForbidden handles this case with default header values.

Failed to authenticate request
*/
type PostLeasesForbidden struct {
}

func (o *PostLeasesForbidden) Error() string {
	return fmt.Sprintf("[POST /leases][%d] postLeasesForbidden ", 403)
}

func (o *PostLeasesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostLeasesConflict creates a PostLeasesConflict with default headers values
func NewPostLeasesConflict() *PostLeasesConflict {
	return &PostLeasesConflict{}
}

/*PostLeasesConflict handles this case with default header values.

Conflict if there is an existing lease already active with the provided principal and account.
*/
type PostLeasesConflict struct {
}

func (o *PostLeasesConflict) Error() string {
	return fmt.Sprintf("[POST /leases][%d] postLeasesConflict ", 409)
}

func (o *PostLeasesConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostLeasesInternalServerError creates a PostLeasesInternalServerError with default headers values
func NewPostLeasesInternalServerError() *PostLeasesInternalServerError {
	return &PostLeasesInternalServerError{}
}

/*PostLeasesInternalServerError handles this case with default header values.

Server errors if the database cannot be reached.
*/
type PostLeasesInternalServerError struct {
}

func (o *PostLeasesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /leases][%d] postLeasesInternalServerError ", 500)
}

func (o *PostLeasesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*PostLeasesBody post leases body
swagger:model PostLeasesBody
*/
type PostLeasesBody struct {

	// budget amount
	// Required: true
	BudgetAmount *float64 `json:"budgetAmount"`

	// budget currency
	// Required: true
	BudgetCurrency *string `json:"budgetCurrency"`

	// budget notification emails
	// Required: true
	BudgetNotificationEmails []string `json:"budgetNotificationEmails"`

	// expires on
	ExpiresOn float64 `json:"expiresOn,omitempty"`

	// principal Id
	// Required: true
	PrincipalID *string `json:"principalId"`
}

// Validate validates this post leases body
func (o *PostLeasesBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateBudgetAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateBudgetCurrency(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateBudgetNotificationEmails(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePrincipalID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostLeasesBody) validateBudgetAmount(formats strfmt.Registry) error {

	if err := validate.Required("lease"+"."+"budgetAmount", "body", o.BudgetAmount); err != nil {
		return err
	}

	return nil
}

func (o *PostLeasesBody) validateBudgetCurrency(formats strfmt.Registry) error {

	if err := validate.Required("lease"+"."+"budgetCurrency", "body", o.BudgetCurrency); err != nil {
		return err
	}

	return nil
}

func (o *PostLeasesBody) validateBudgetNotificationEmails(formats strfmt.Registry) error {

	if err := validate.Required("lease"+"."+"budgetNotificationEmails", "body", o.BudgetNotificationEmails); err != nil {
		return err
	}

	return nil
}

func (o *PostLeasesBody) validatePrincipalID(formats strfmt.Registry) error {

	if err := validate.Required("lease"+"."+"principalId", "body", o.PrincipalID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostLeasesBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostLeasesBody) UnmarshalBinary(b []byte) error {
	var res PostLeasesBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostLeasesCreatedBody Lease Details
swagger:model PostLeasesCreatedBody
*/
type PostLeasesCreatedBody struct {

	// accountId of the AWS account
	AccountID string `json:"accountId,omitempty"`

	// budget amount
	BudgetAmount float64 `json:"budgetAmount,omitempty"`

	// budget currency
	BudgetCurrency string `json:"budgetCurrency,omitempty"`

	// budget notification emails
	BudgetNotificationEmails []string `json:"budgetNotificationEmails"`

	// creation date in epoch seconds
	CreatedOn float64 `json:"createdOn,omitempty"`

	// date lease should expire in epoch seconds
	ExpiresOn float64 `json:"expiresOn,omitempty"`

	// Lease ID
	ID string `json:"id,omitempty"`

	// date last modified in epoch seconds
	LastModifiedOn float64 `json:"lastModifiedOn,omitempty"`

	// Status of the Lease.
	// "Active": The principal is leased and has access to the account
	// "Inactive": The lease has become inactive, either through expiring, exceeding budget, or by request.
	//
	// Enum: [Active Inactive]
	LeaseStatus string `json:"leaseStatus,omitempty"`

	// date lease status was last modified in epoch seconds
	LeaseStatusModifiedOn float64 `json:"leaseStatusModifiedOn,omitempty"`

	// A reason behind the lease status.
	// "LeaseExpired": The lease exceeded its expiration time ("expiresOn") and
	// the associated account was reset and returned to the account pool.
	// "LeaseOverBudget": The lease exceeded its budgeted amount and the
	// associated account was reset and returned to the account pool.
	// "LeaseDestroyed": The lease was adminstratively ended, which can be done
	// via the leases API.
	// "LeaseActive": The lease is active.
	// "LeaseRolledBack": A system error occurred while provisioning the lease.
	// and it was rolled back.
	//
	// Enum: [LeaseExpired LeaseOverBudget LeaseDestroyed LeaseActive LeaseRolledBack]
	LeaseStatusReason string `json:"leaseStatusReason,omitempty"`

	// principalId of the lease to get
	PrincipalID string `json:"principalId,omitempty"`
}

// Validate validates this post leases created body
func (o *PostLeasesCreatedBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLeaseStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLeaseStatusReason(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var postLeasesCreatedBodyTypeLeaseStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Active","Inactive"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postLeasesCreatedBodyTypeLeaseStatusPropEnum = append(postLeasesCreatedBodyTypeLeaseStatusPropEnum, v)
	}
}

const (

	// PostLeasesCreatedBodyLeaseStatusActive captures enum value "Active"
	PostLeasesCreatedBodyLeaseStatusActive string = "Active"

	// PostLeasesCreatedBodyLeaseStatusInactive captures enum value "Inactive"
	PostLeasesCreatedBodyLeaseStatusInactive string = "Inactive"
)

// prop value enum
func (o *PostLeasesCreatedBody) validateLeaseStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, postLeasesCreatedBodyTypeLeaseStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *PostLeasesCreatedBody) validateLeaseStatus(formats strfmt.Registry) error {

	if swag.IsZero(o.LeaseStatus) { // not required
		return nil
	}

	// value enum
	if err := o.validateLeaseStatusEnum("postLeasesCreated"+"."+"leaseStatus", "body", o.LeaseStatus); err != nil {
		return err
	}

	return nil
}

var postLeasesCreatedBodyTypeLeaseStatusReasonPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["LeaseExpired","LeaseOverBudget","LeaseDestroyed","LeaseActive","LeaseRolledBack"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postLeasesCreatedBodyTypeLeaseStatusReasonPropEnum = append(postLeasesCreatedBodyTypeLeaseStatusReasonPropEnum, v)
	}
}

const (

	// PostLeasesCreatedBodyLeaseStatusReasonLeaseExpired captures enum value "LeaseExpired"
	PostLeasesCreatedBodyLeaseStatusReasonLeaseExpired string = "LeaseExpired"

	// PostLeasesCreatedBodyLeaseStatusReasonLeaseOverBudget captures enum value "LeaseOverBudget"
	PostLeasesCreatedBodyLeaseStatusReasonLeaseOverBudget string = "LeaseOverBudget"

	// PostLeasesCreatedBodyLeaseStatusReasonLeaseDestroyed captures enum value "LeaseDestroyed"
	PostLeasesCreatedBodyLeaseStatusReasonLeaseDestroyed string = "LeaseDestroyed"

	// PostLeasesCreatedBodyLeaseStatusReasonLeaseActive captures enum value "LeaseActive"
	PostLeasesCreatedBodyLeaseStatusReasonLeaseActive string = "LeaseActive"

	// PostLeasesCreatedBodyLeaseStatusReasonLeaseRolledBack captures enum value "LeaseRolledBack"
	PostLeasesCreatedBodyLeaseStatusReasonLeaseRolledBack string = "LeaseRolledBack"
)

// prop value enum
func (o *PostLeasesCreatedBody) validateLeaseStatusReasonEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, postLeasesCreatedBodyTypeLeaseStatusReasonPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *PostLeasesCreatedBody) validateLeaseStatusReason(formats strfmt.Registry) error {

	if swag.IsZero(o.LeaseStatusReason) { // not required
		return nil
	}

	// value enum
	if err := o.validateLeaseStatusReasonEnum("postLeasesCreated"+"."+"leaseStatusReason", "body", o.LeaseStatusReason); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostLeasesCreatedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostLeasesCreatedBody) UnmarshalBinary(b []byte) error {
	var res PostLeasesCreatedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

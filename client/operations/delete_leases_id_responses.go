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

// DeleteLeasesIDReader is a Reader for the DeleteLeasesID structure.
type DeleteLeasesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteLeasesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteLeasesIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteLeasesIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteLeasesIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteLeasesIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteLeasesIDOK creates a DeleteLeasesIDOK with default headers values
func NewDeleteLeasesIDOK() *DeleteLeasesIDOK {
	return &DeleteLeasesIDOK{}
}

/*DeleteLeasesIDOK handles this case with default header values.

DeleteLeasesIDOK delete leases Id o k
*/
type DeleteLeasesIDOK struct {
	Payload *DeleteLeasesIDOKBody
}

func (o *DeleteLeasesIDOK) Error() string {
	/*
		#nosec CWE-89: false positive. No sql here.
	*/
	return fmt.Sprintf("[DELETE /leases/{id}][%d] deleteLeasesIdOK  %+v", 200, o.Payload)
}

func (o *DeleteLeasesIDOK) GetPayload() *DeleteLeasesIDOKBody {
	return o.Payload
}

func (o *DeleteLeasesIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DeleteLeasesIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLeasesIDBadRequest creates a DeleteLeasesIDBadRequest with default headers values
func NewDeleteLeasesIDBadRequest() *DeleteLeasesIDBadRequest {
	return &DeleteLeasesIDBadRequest{}
}

/*DeleteLeasesIDBadRequest handles this case with default header values.

"Failed to Parse Request Body" if the request body is blank or incorrectly formatted. or if there are no account leases found for the specified accountId or if the account specified is not already Active.

*/
type DeleteLeasesIDBadRequest struct {
}

func (o *DeleteLeasesIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /leases/{id}][%d] deleteLeasesIdBadRequest ", 400)
}

func (o *DeleteLeasesIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteLeasesIDForbidden creates a DeleteLeasesIDForbidden with default headers values
func NewDeleteLeasesIDForbidden() *DeleteLeasesIDForbidden {
	return &DeleteLeasesIDForbidden{}
}

/*DeleteLeasesIDForbidden handles this case with default header values.

Failed to authenticate request
*/
type DeleteLeasesIDForbidden struct {
}

func (o *DeleteLeasesIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /leases/{id}][%d] deleteLeasesIdForbidden ", 403)
}

func (o *DeleteLeasesIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteLeasesIDInternalServerError creates a DeleteLeasesIDInternalServerError with default headers values
func NewDeleteLeasesIDInternalServerError() *DeleteLeasesIDInternalServerError {
	return &DeleteLeasesIDInternalServerError{}
}

/*DeleteLeasesIDInternalServerError handles this case with default header values.

Server errors if the database cannot be reached.
*/
type DeleteLeasesIDInternalServerError struct {
}

func (o *DeleteLeasesIDInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /leases/{id}][%d] deleteLeasesIdInternalServerError ", 500)
}

func (o *DeleteLeasesIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*DeleteLeasesIDOKBody Lease Details
swagger:model DeleteLeasesIDOKBody
*/
type DeleteLeasesIDOKBody struct {

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

// Validate validates this delete leases ID o k body
func (o *DeleteLeasesIDOKBody) Validate(formats strfmt.Registry) error {
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

var deleteLeasesIdOKBodyTypeLeaseStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Active","Inactive"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		deleteLeasesIdOKBodyTypeLeaseStatusPropEnum = append(deleteLeasesIdOKBodyTypeLeaseStatusPropEnum, v)
	}
}

const (

	// DeleteLeasesIDOKBodyLeaseStatusActive captures enum value "Active"
	DeleteLeasesIDOKBodyLeaseStatusActive string = "Active"

	// DeleteLeasesIDOKBodyLeaseStatusInactive captures enum value "Inactive"
	DeleteLeasesIDOKBodyLeaseStatusInactive string = "Inactive"
)

// prop value enum
func (o *DeleteLeasesIDOKBody) validateLeaseStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, deleteLeasesIdOKBodyTypeLeaseStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *DeleteLeasesIDOKBody) validateLeaseStatus(formats strfmt.Registry) error {

	if swag.IsZero(o.LeaseStatus) { // not required
		return nil
	}

	// value enum
	if err := o.validateLeaseStatusEnum("deleteLeasesIdOK"+"."+"leaseStatus", "body", o.LeaseStatus); err != nil {
		return err
	}

	return nil
}

var deleteLeasesIdOKBodyTypeLeaseStatusReasonPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["LeaseExpired","LeaseOverBudget","LeaseDestroyed","LeaseActive","LeaseRolledBack"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		deleteLeasesIdOKBodyTypeLeaseStatusReasonPropEnum = append(deleteLeasesIdOKBodyTypeLeaseStatusReasonPropEnum, v)
	}
}

const (

	// DeleteLeasesIDOKBodyLeaseStatusReasonLeaseExpired captures enum value "LeaseExpired"
	DeleteLeasesIDOKBodyLeaseStatusReasonLeaseExpired string = "LeaseExpired"

	// DeleteLeasesIDOKBodyLeaseStatusReasonLeaseOverBudget captures enum value "LeaseOverBudget"
	DeleteLeasesIDOKBodyLeaseStatusReasonLeaseOverBudget string = "LeaseOverBudget"

	// DeleteLeasesIDOKBodyLeaseStatusReasonLeaseDestroyed captures enum value "LeaseDestroyed"
	DeleteLeasesIDOKBodyLeaseStatusReasonLeaseDestroyed string = "LeaseDestroyed"

	// DeleteLeasesIDOKBodyLeaseStatusReasonLeaseActive captures enum value "LeaseActive"
	DeleteLeasesIDOKBodyLeaseStatusReasonLeaseActive string = "LeaseActive"

	// DeleteLeasesIDOKBodyLeaseStatusReasonLeaseRolledBack captures enum value "LeaseRolledBack"
	DeleteLeasesIDOKBodyLeaseStatusReasonLeaseRolledBack string = "LeaseRolledBack"
)

// prop value enum
func (o *DeleteLeasesIDOKBody) validateLeaseStatusReasonEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, deleteLeasesIdOKBodyTypeLeaseStatusReasonPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *DeleteLeasesIDOKBody) validateLeaseStatusReason(formats strfmt.Registry) error {

	if swag.IsZero(o.LeaseStatusReason) { // not required
		return nil
	}

	// value enum
	if err := o.validateLeaseStatusReasonEnum("deleteLeasesIdOK"+"."+"leaseStatusReason", "body", o.LeaseStatusReason); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteLeasesIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteLeasesIDOKBody) UnmarshalBinary(b []byte) error {
	var res DeleteLeasesIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

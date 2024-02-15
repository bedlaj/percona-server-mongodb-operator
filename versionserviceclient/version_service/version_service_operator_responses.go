// Code generated by go-swagger; DO NOT EDIT.

package version_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/percona/percona-server-mongodb-operator/versionserviceclient/models"
)

// VersionServiceOperatorReader is a Reader for the VersionServiceOperator structure.
type VersionServiceOperatorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VersionServiceOperatorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVersionServiceOperatorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVersionServiceOperatorDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVersionServiceOperatorOK creates a VersionServiceOperatorOK with default headers values
func NewVersionServiceOperatorOK() *VersionServiceOperatorOK {
	return &VersionServiceOperatorOK{}
}

/*
VersionServiceOperatorOK handles this case with default header values.

A successful response.
*/
type VersionServiceOperatorOK struct {
	Payload *models.VersionOperatorResponse
}

func (o *VersionServiceOperatorOK) Error() string {
	return fmt.Sprintf("[GET /versions/v1/{product}/{operatorVersion}][%d] versionServiceOperatorOK  %+v", 200, o.Payload)
}

func (o *VersionServiceOperatorOK) GetPayload() *models.VersionOperatorResponse {
	return o.Payload
}

func (o *VersionServiceOperatorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VersionOperatorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVersionServiceOperatorDefault creates a VersionServiceOperatorDefault with default headers values
func NewVersionServiceOperatorDefault(code int) *VersionServiceOperatorDefault {
	return &VersionServiceOperatorDefault{
		_statusCode: code,
	}
}

/*
VersionServiceOperatorDefault handles this case with default header values.

An unexpected error response
*/
type VersionServiceOperatorDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// Code gets the status code for the version service operator default response
func (o *VersionServiceOperatorDefault) Code() int {
	return o._statusCode
}

func (o *VersionServiceOperatorDefault) Error() string {
	return fmt.Sprintf("[GET /versions/v1/{product}/{operatorVersion}][%d] VersionService_Operator default  %+v", o._statusCode, o.Payload)
}

func (o *VersionServiceOperatorDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *VersionServiceOperatorDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

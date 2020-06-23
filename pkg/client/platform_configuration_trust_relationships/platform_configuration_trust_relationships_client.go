// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by go-swagger; DO NOT EDIT.

package platform_configuration_trust_relationships

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new platform configuration trust relationships API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for platform configuration trust relationships API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateTrustRelationship(params *CreateTrustRelationshipParams, authInfo runtime.ClientAuthInfoWriter) (*CreateTrustRelationshipCreated, error)

	DeleteTrustRelationship(params *DeleteTrustRelationshipParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteTrustRelationshipOK, error)

	GetTrustRelationship(params *GetTrustRelationshipParams, authInfo runtime.ClientAuthInfoWriter) (*GetTrustRelationshipOK, error)

	GetTrustRelationships(params *GetTrustRelationshipsParams, authInfo runtime.ClientAuthInfoWriter) (*GetTrustRelationshipsOK, error)

	UpdateTrustRelationship(params *UpdateTrustRelationshipParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateTrustRelationshipOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateTrustRelationship creates trust relationship

  Creates a trust relationship.
*/
func (a *Client) CreateTrustRelationship(params *CreateTrustRelationshipParams, authInfo runtime.ClientAuthInfoWriter) (*CreateTrustRelationshipCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateTrustRelationshipParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "create-trust-relationship",
		Method:             "POST",
		PathPattern:        "/platform/configuration/trust-relationships",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateTrustRelationshipReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateTrustRelationshipCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for create-trust-relationship: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteTrustRelationship deletes trust relationship

  Deletes a trust relationship.
*/
func (a *Client) DeleteTrustRelationship(params *DeleteTrustRelationshipParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteTrustRelationshipOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteTrustRelationshipParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "delete-trust-relationship",
		Method:             "DELETE",
		PathPattern:        "/platform/configuration/trust-relationships/{trust_relationship_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteTrustRelationshipReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteTrustRelationshipOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for delete-trust-relationship: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetTrustRelationship gets trust relationship

  Retrieves information about a trust relationship. `local` can be used as the ID to obtain the local trust relationship.
*/
func (a *Client) GetTrustRelationship(params *GetTrustRelationshipParams, authInfo runtime.ClientAuthInfoWriter) (*GetTrustRelationshipOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTrustRelationshipParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get-trust-relationship",
		Method:             "GET",
		PathPattern:        "/platform/configuration/trust-relationships/{trust_relationship_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTrustRelationshipReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetTrustRelationshipOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-trust-relationship: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetTrustRelationships lists trust relationships

  List trust relationships.
*/
func (a *Client) GetTrustRelationships(params *GetTrustRelationshipsParams, authInfo runtime.ClientAuthInfoWriter) (*GetTrustRelationshipsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTrustRelationshipsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get-trust-relationships",
		Method:             "GET",
		PathPattern:        "/platform/configuration/trust-relationships",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTrustRelationshipsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetTrustRelationshipsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-trust-relationships: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateTrustRelationship updates trust relationship

  Updates a trust relationship.
*/
func (a *Client) UpdateTrustRelationship(params *UpdateTrustRelationshipParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateTrustRelationshipOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateTrustRelationshipParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "update-trust-relationship",
		Method:             "PUT",
		PathPattern:        "/platform/configuration/trust-relationships/{trust_relationship_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateTrustRelationshipReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateTrustRelationshipOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for update-trust-relationship: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
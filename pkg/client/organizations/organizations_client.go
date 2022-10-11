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

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new organizations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for organizations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AcceptOrganizationInvitation(params *AcceptOrganizationInvitationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AcceptOrganizationInvitationOK, error)

	CreateOrganizationInvitations(params *CreateOrganizationInvitationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateOrganizationInvitationsCreated, error)

	DeleteOrganizationInvitations(params *DeleteOrganizationInvitationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteOrganizationInvitationsOK, error)

	DeleteOrganizationMemberships(params *DeleteOrganizationMembershipsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteOrganizationMembershipsOK, error)

	GetOrganization(params *GetOrganizationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOrganizationOK, error)

	GetOrganizationInvitation(params *GetOrganizationInvitationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOrganizationInvitationOK, error)

	ListOrganizationInvitations(params *ListOrganizationInvitationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListOrganizationInvitationsOK, error)

	ListOrganizationMembers(params *ListOrganizationMembersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListOrganizationMembersOK, error)

	ListOrganizations(params *ListOrganizationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListOrganizationsOK, error)

	UpdateOrganization(params *UpdateOrganizationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateOrganizationOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AcceptOrganizationInvitation accepts an organization invitation

  Accepts an organization invitation. Currently unavailable in self-hosted ECE.
*/
func (a *Client) AcceptOrganizationInvitation(params *AcceptOrganizationInvitationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AcceptOrganizationInvitationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAcceptOrganizationInvitationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "accept-organization-invitation",
		Method:             "POST",
		PathPattern:        "/organizations/invitations/{invitation_token}/_accept",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AcceptOrganizationInvitationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AcceptOrganizationInvitationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for accept-organization-invitation: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateOrganizationInvitations creates organization invitations

  Creates or refreshes organization invitations. Currently unavailable in self-hosted ECE.
*/
func (a *Client) CreateOrganizationInvitations(params *CreateOrganizationInvitationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateOrganizationInvitationsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateOrganizationInvitationsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "create-organization-invitations",
		Method:             "POST",
		PathPattern:        "/organizations/{organization_id}/invitations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateOrganizationInvitationsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateOrganizationInvitationsCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for create-organization-invitations: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteOrganizationInvitations deletes organization invitations

  Deletes one or more organization invitations. Currently unavailable in self-hosted ECE.
*/
func (a *Client) DeleteOrganizationInvitations(params *DeleteOrganizationInvitationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteOrganizationInvitationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteOrganizationInvitationsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "delete-organization-invitations",
		Method:             "DELETE",
		PathPattern:        "/organizations/{organization_id}/invitations/{invitation_tokens}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteOrganizationInvitationsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteOrganizationInvitationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for delete-organization-invitations: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteOrganizationMemberships deletes organization memberships

  Deletes one or more organization memberships. Currently unavailable in self-hosted ECE.
*/
func (a *Client) DeleteOrganizationMemberships(params *DeleteOrganizationMembershipsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteOrganizationMembershipsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteOrganizationMembershipsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "delete-organization-memberships",
		Method:             "DELETE",
		PathPattern:        "/organizations/{organization_id}/members/{user_ids}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteOrganizationMembershipsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteOrganizationMembershipsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for delete-organization-memberships: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetOrganization fetches organization information

  Fetch a single organization by id. Currently unavailable in self-hosted ECE.
*/
func (a *Client) GetOrganization(params *GetOrganizationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOrganizationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrganizationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get-organization",
		Method:             "GET",
		PathPattern:        "/organizations/{organization_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOrganizationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetOrganizationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-organization: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetOrganizationInvitation gets organization invitation

  Gets a single invitation to an organization by token. Currently unavailable in self-hosted ECE.
*/
func (a *Client) GetOrganizationInvitation(params *GetOrganizationInvitationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOrganizationInvitationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrganizationInvitationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get-organization-invitation",
		Method:             "GET",
		PathPattern:        "/organizations/invitations/{invitation_token}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOrganizationInvitationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetOrganizationInvitationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-organization-invitation: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListOrganizationInvitations lists organization invitations

  Fetch open invitations to the selected organization. Currently unavailable in self-hosted ECE.
*/
func (a *Client) ListOrganizationInvitations(params *ListOrganizationInvitationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListOrganizationInvitationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListOrganizationInvitationsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "list-organization-invitations",
		Method:             "GET",
		PathPattern:        "/organizations/{organization_id}/invitations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListOrganizationInvitationsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListOrganizationInvitationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for list-organization-invitations: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListOrganizationMembers lists organization members

  Fetch users belonging to the selected organization. Currently unavailable in self-hosted ECE.
*/
func (a *Client) ListOrganizationMembers(params *ListOrganizationMembersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListOrganizationMembersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListOrganizationMembersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "list-organization-members",
		Method:             "GET",
		PathPattern:        "/organizations/{organization_id}/members",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListOrganizationMembersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListOrganizationMembersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for list-organization-members: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListOrganizations lists organizations

  Fetch organizations available to the current user. Currently unavailable in self-hosted ECE.
*/
func (a *Client) ListOrganizations(params *ListOrganizationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListOrganizationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListOrganizationsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "list-organizations",
		Method:             "GET",
		PathPattern:        "/organizations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListOrganizationsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListOrganizationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for list-organizations: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateOrganization updates organization

  Updates an existing organization. Currently unavailable in self-hosted ECE.
*/
func (a *Client) UpdateOrganization(params *UpdateOrganizationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateOrganizationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateOrganizationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "update-organization",
		Method:             "PUT",
		PathPattern:        "/organizations/{organization_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateOrganizationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateOrganizationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for update-organization: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
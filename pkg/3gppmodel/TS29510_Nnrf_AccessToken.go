// Package model provides primitives for NRF OAuth2 Access Token.
//
// Reference: 3GPP TS 29.510 V18.5.0, Nnrf_AccessToken Service
//
// This file contains spec-accurate Go structs for the NRF OAuth2
// Access Token request/response types.
package gppmodel

// ============================================================================
// AccessTokenReq — OAuth2 Access Token Request
// ============================================================================

// AccessTokenReq contains information related to the access token request.
// Reference: TS 29.510 Section 6.2.6.2.2 (AccessTokenReq)
type AccessTokenReq struct {
	// GrantType is the OAuth2 grant type. Must be "client_credentials". (Required)
	GrantType AccessTokenGrantType `json:"grant_type"`

	// NfInstanceId is the NF instance ID of the service consumer. (Required)
	NfInstanceId NfInstanceId `json:"nfInstanceId"`

	// NfType is the type of the NF service consumer (e.g. "NWDAF").
	NfType string `json:"nfType,omitempty"`

	// TargetNfType is the type of the target NF service producer (e.g. "AMF").
	TargetNfType string `json:"targetNfType,omitempty"`

	// Scope is the space-separated list of requested scopes. (Required)
	Scope AccessTokenScope `json:"scope"`

	// TargetNfInstanceId is the NF instance ID of the target NF.
	TargetNfInstanceId NfInstanceId `json:"targetNfInstanceId,omitempty"`

	// RequesterPlmn is the PLMN ID of the requester NF.
	RequesterPlmn *PlmnId `json:"requesterPlmn,omitempty"`

	// RequesterPlmnList is the list of PLMN IDs of the requester NF.
	RequesterPlmnList []PlmnId `json:"requesterPlmnList,omitempty"`

	// RequesterSnssaiList is the list of S-NSSAIs of the requester NF.
	RequesterSnssaiList []Snssai `json:"requesterSnssaiList,omitempty"`

	// RequesterFqdn is the FQDN of the requester NF.
	RequesterFqdn Fqdn `json:"requesterFqdn,omitempty"`

	// RequesterSnpnList is the list of SNPN identities of the requester NF.
	RequesterSnpnList []PlmnIdNid `json:"requesterSnpnList,omitempty"`

	// TargetPlmn is the PLMN ID of the target NF.
	TargetPlmn *PlmnId `json:"targetPlmn,omitempty"`

	// TargetSnpn is the SNPN identity of the target NF.
	TargetSnpn *PlmnIdNid `json:"targetSnpn,omitempty"`

	// TargetSnssaiList is the list of S-NSSAIs of the target NF.
	TargetSnssaiList []Snssai `json:"targetSnssaiList,omitempty"`

	// TargetNsiList is the list of NSI IDs of the target NF.
	TargetNsiList []string `json:"targetNsiList,omitempty"`

	// TargetNfSetId is the NF Set ID of the target NF.
	TargetNfSetId NfSetId `json:"targetNfSetId,omitempty"`

	// TargetNfServiceSetId is the NF Service Set ID of the target NF service.
	TargetNfServiceSetId NfServiceSetId `json:"targetNfServiceSetId,omitempty"`

	// HnrfAccessTokenUri is the URI of the hNRF to obtain an access token.
	HnrfAccessTokenUri Uri `json:"hnrfAccessTokenUri,omitempty"`

	// SourceNfInstanceId is the NF instance ID of the source NF.
	SourceNfInstanceId NfInstanceId `json:"sourceNfInstanceId,omitempty"`
}

// ============================================================================
// AccessTokenRsp — OAuth2 Access Token Response
// ============================================================================

// AccessTokenRsp contains information related to the access token response.
// Reference: TS 29.510 Section 6.2.6.2.3 (AccessTokenRsp)
type AccessTokenRsp struct {
	// AccessToken is the JWS Compact Serialized representation of JWS signed
	// JSON object (AccessTokenClaims). (Required)
	AccessToken string `json:"access_token"`

	// TokenType is the type of the token. (Required)
	TokenType AccessTokenType `json:"token_type"`

	// ExpiresIn is the lifetime in seconds of the access token.
	ExpiresIn int `json:"expires_in,omitempty"`

	// Scope is the granted scope (space-separated).
	Scope AccessTokenScope `json:"scope,omitempty"`
}

// ============================================================================
// AccessTokenClaims — JWT Claims inside the access token
// ============================================================================

// AccessTokenClaims represents the claims data structure for the access token.
// Reference: TS 29.510 Section 6.2.6.2.4 (AccessTokenClaims)
type AccessTokenClaims struct {
	// Iss is the NF instance ID of the NRF that issued the token. (Required)
	Iss NfInstanceId `json:"iss"`

	// Sub is the NF instance ID of the NF service consumer. (Required)
	Sub NfInstanceId `json:"sub"`

	// Aud is the intended audience. Can be a single NFType or a list of NF instance IDs. (Required)
	Aud interface{} `json:"aud"`

	// Scope is the granted scope. (Required)
	Scope AccessTokenScope `json:"scope"`

	// Exp is the expiration time as Unix timestamp in seconds. (Required)
	Exp int `json:"exp"`

	// ConsumerPlmnId is the PLMN ID of the NF service consumer.
	ConsumerPlmnId *PlmnId `json:"consumerPlmnId,omitempty"`

	// ConsumerSnpnId is the SNPN ID of the NF service consumer.
	ConsumerSnpnId *PlmnIdNid `json:"consumerSnpnId,omitempty"`

	// ProducerPlmnId is the PLMN ID of the NF service producer.
	ProducerPlmnId *PlmnId `json:"producerPlmnId,omitempty"`

	// ProducerSnpnId is the SNPN ID of the NF service producer.
	ProducerSnpnId *PlmnIdNid `json:"producerSnpnId,omitempty"`

	// ProducerSnssaiList is the list of S-NSSAIs of the NF service producer.
	ProducerSnssaiList []Snssai `json:"producerSnssaiList,omitempty"`

	// ProducerNsiList is the list of NSI IDs of the NF service producer.
	ProducerNsiList []string `json:"producerNsiList,omitempty"`

	// ProducerNfSetId is the NF Set ID of the NF service producer.
	ProducerNfSetId NfSetId `json:"producerNfSetId,omitempty"`

	// ProducerNfServiceSetId is the NF Service Set ID of the NF service producer.
	ProducerNfServiceSetId NfServiceSetId `json:"producerNfServiceSetId,omitempty"`

	// SourceNfInstanceId is the NF instance ID of the source NF.
	SourceNfInstanceId NfInstanceId `json:"sourceNfInstanceId,omitempty"`
}

// ============================================================================
// AccessTokenErr — OAuth2 Error Response
// ============================================================================

// AccessTokenErr represents an error returned in the access token response.
// Reference: TS 29.510 Section 6.2.6.2.5 (AccessTokenErr)
type AccessTokenErr struct {
	// Error is the OAuth2 error code. (Required)
	Error AccessTokenErrorCode `json:"error"`

	// ErrorDescription is a human-readable description of the error.
	ErrorDescription string `json:"error_description,omitempty"`

	// ErrorUri is a URI identifying a human-readable web page with error info.
	ErrorUri string `json:"error_uri,omitempty"`
}

// ============================================================================
// Simple Types & Enums
// ============================================================================

// AccessTokenGrantType defines valid OAuth2 grant types for NRF access token.
type AccessTokenGrantType string

const (
	// ClientCredentialsGrant is the only supported grant type for NRF OAuth2.
	ClientCredentialsGrant AccessTokenGrantType = "client_credentials"
)

// AccessTokenType defines the token type returned by NRF.
type AccessTokenType string

const (
	// BearerToken is the Bearer token type.
	BearerToken AccessTokenType = "Bearer"
)

// AccessTokenErrorCode defines OAuth2 error codes for access token errors.
type AccessTokenErrorCode string

const (
	AccessTokenErrInvalidRequest       AccessTokenErrorCode = "invalid_request"
	AccessTokenErrInvalidClient        AccessTokenErrorCode = "invalid_client"
	AccessTokenErrInvalidGrant         AccessTokenErrorCode = "invalid_grant"
	AccessTokenErrUnauthorizedClient   AccessTokenErrorCode = "unauthorized_client"
	AccessTokenErrUnsupportedGrantType AccessTokenErrorCode = "unsupported_grant_type"
	AccessTokenErrInvalidScope         AccessTokenErrorCode = "invalid_scope"
)

// AccessTokenScope is a space-separated list of OAuth2 scopes.
// Pattern: ^([a-zA-Z0-9_:-]+)( [a-zA-Z0-9_:-]+)*$
type AccessTokenScope string

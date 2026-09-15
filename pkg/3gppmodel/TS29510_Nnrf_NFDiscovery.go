// Package model provides primitives for NRF NFDiscovery Service.
//
// Reference: 3GPP TS 29.510 V18.6.0, Nnrf_NFDiscovery Service
//
// This file contains spec-accurate Go structs for the NRF NF Discovery
// request/response types. Types from TS29510_Nnrf_NFManagement.yaml that
// are not yet generated use interface{} stubs.
package gppmodel

import (
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// ============================================================================
// SearchResult — Top-level NRF Discovery response
// ============================================================================

// SearchResult contains the list of NF Profiles returned in a Discovery response.
// Reference: TS 29.510 Section 6.2.6.3.2 (SearchResult)
type SearchResult struct {
	// ValidityPeriod is the validity period for the search results (in seconds). (Required)
	ValidityPeriod int `json:"validityPeriod"`

	// NfInstances contains the list of NF Profiles matching the query. (Required)
	NfInstances []NFProfile `json:"nfInstances"`

	// CompleteNfInstances contains NF Profiles with complete data (full profile query).
	CompleteNfInstances []NFProfile `json:"completeNfInstances,omitempty"`

	// SearchId is the identifier of the stored search result.
	SearchId string `json:"searchId,omitempty"`

	// NumNfInstComplete is the number of complete NF instance profiles returned.
	NumNfInstComplete *Uint32 `json:"numNfInstComplete,omitempty"`

	// PreferredSearch indicates which preferred query parameters matched.
	PreferredSearch *PreferredSearch `json:"preferredSearch,omitempty"`

	// NrfSupportedFeatures lists features supported by the NRF.
	NrfSupportedFeatures SupportedFeatures `json:"nrfSupportedFeatures,omitempty"`

	// NfInstanceList is a map of matching NF instances keyed by NF instance ID.
	NfInstanceList map[string]NfInstanceInfo `json:"nfInstanceList,omitempty"`

	// SearchResultInfo contains additional search result information.
	SearchResultInfo *SearchResultInfo `json:"searchResultInfo,omitempty"`

	// AlteredPriorityInd indicates whether the NRF altered the priorities.
	AlteredPriorityInd bool `json:"alteredPriorityInd,omitempty"`

	// NoProfileMatchInfo provides the reason for no matching profile.
	NoProfileMatchInfo *NoProfileMatchInfo `json:"noProfileMatchInfo,omitempty"`

	// IgnoredQueryParams lists query parameters that were ignored.
	IgnoredQueryParams []string `json:"ignoredQueryParams,omitempty"`
}

// ============================================================================
// StoredSearchResult
// ============================================================================

// StoredSearchResult contains a complete search result stored by the NRF.
// Reference: TS 29.510 Section 6.2.6.3.3 (StoredSearchResult)
type StoredSearchResult struct {
	// NfInstances contains the list of NF Profiles. (Required)
	NfInstances []NFProfile `json:"nfInstances"`

	// CompleteNfInstances contains NF Profiles with complete data.
	CompleteNfInstances []NFProfile `json:"completeNfInstances,omitempty"`
}

// ============================================================================
// NFProfile — NF Instance profile discovered by the NRF
// ============================================================================

// NFProfile contains information of an NF Instance discovered by the NRF.
// Reference: TS 29.510 Section 6.2.6.3.4 (NFProfile)
type NFProfile struct {
	// NfInstanceId uniquely identifies the NF instance (UUID v4). (Required)
	NfInstanceId openapi_types.UUID `json:"nfInstanceId"`

	// NfInstanceName is the human-readable name of the NF instance.
	NfInstanceName string `json:"nfInstanceName,omitempty"`

	// NfType is the type of Network Function. (Required)
	NfType NFType `json:"nfType"`

	// NfStatus is the status of the NF instance. (Required)
	NfStatus NFStatus `json:"nfStatus"`

	// CollocatedNfInstances lists collocated NF instances.
	CollocatedNfInstances []interface{} `json:"collocatedNfInstances,omitempty"`

	// PlmnList contains PLMN IDs of the NF instance.
	PlmnList []PlmnId `json:"plmnList,omitempty"`

	// SNssais contains S-NSSAIs supported by the NF instance.
	SNssais []ExtSnssai `json:"sNssais,omitempty"`

	// PerPlmnSnssaiList lists S-NSSAIs per PLMN.
	PerPlmnSnssaiList []interface{} `json:"perPlmnSnssaiList,omitempty"`

	// NsiList contains NSI identifiers.
	NsiList []string `json:"nsiList,omitempty"`

	// Fqdn is the FQDN of the NF instance.
	Fqdn Fqdn `json:"fqdn,omitempty"`

	// InterPlmnFqdn is the FQDN for inter-PLMN communication.
	InterPlmnFqdn Fqdn `json:"interPlmnFqdn,omitempty"`

	// Ipv4Addresses contains IPv4 addresses of the NF instance.
	Ipv4Addresses []Ipv4Addr `json:"ipv4Addresses,omitempty"`

	// Ipv6Addresses contains IPv6 addresses of the NF instance.
	Ipv6Addresses []Ipv6Addr `json:"ipv6Addresses,omitempty"`

	// AllowedPlmns lists PLMNs allowed to access the NF.
	AllowedPlmns []PlmnId `json:"allowedPlmns,omitempty"`

	// AllowedSnpns lists SNPNs allowed to access the NF.
	AllowedSnpns []PlmnIdNid `json:"allowedSnpns,omitempty"`

	// AllowedNfTypes lists NF types allowed to access.
	AllowedNfTypes []NFType `json:"allowedNfTypes,omitempty"`

	// AllowedNfDomains lists NF domains allowed to access.
	AllowedNfDomains []string `json:"allowedNfDomains,omitempty"`

	// AllowedNssais lists S-NSSAIs allowed for access.
	AllowedNssais []ExtSnssai `json:"allowedNssais,omitempty"`

	// AllowedRuleSet maps JSON pointer IDs to rule sets.
	AllowedRuleSet map[string]interface{} `json:"allowedRuleSet,omitempty"`

	// Capacity is the static capacity of the NF instance (0-65535).
	Capacity int `json:"capacity,omitempty"`

	// Load is the current load percentage (0-100).
	Load int `json:"load,omitempty"`

	// LoadTimeStamp is the timestamp when the load was measured.
	LoadTimeStamp *DateTime `json:"loadTimeStamp,omitempty"`

	// Locality describes the location of the NF instance.
	Locality string `json:"locality,omitempty"`

	// ExtLocality maps locality types to locality descriptions.
	ExtLocality map[string]string `json:"extLocality,omitempty"`

	// Priority is the priority of the NF instance (0-65535, lower = higher priority).
	Priority int `json:"priority,omitempty"`

	// UdrInfo contains UDR-specific information.
	UdrInfo interface{} `json:"udrInfo,omitempty"`

	// UdrInfoList maps keys to UdrInfo entries.
	UdrInfoList map[string]interface{} `json:"udrInfoList,omitempty"`

	// UdmInfo contains UDM-specific information.
	UdmInfo interface{} `json:"udmInfo,omitempty"`

	// UdmInfoList maps keys to UdmInfo entries.
	UdmInfoList map[string]interface{} `json:"udmInfoList,omitempty"`

	// AusfInfo contains AUSF-specific information.
	AusfInfo interface{} `json:"ausfInfo,omitempty"`

	// AusfInfoList maps keys to AusfInfo entries.
	AusfInfoList map[string]interface{} `json:"ausfInfoList,omitempty"`

	// AmfInfo contains AMF-specific information.
	AmfInfo interface{} `json:"amfInfo,omitempty"`

	// AmfInfoList maps keys to AmfInfo entries.
	AmfInfoList map[string]interface{} `json:"amfInfoList,omitempty"`

	// SmfInfo contains SMF-specific information.
	SmfInfo interface{} `json:"smfInfo,omitempty"`

	// SmfInfoList maps keys to SmfInfo entries.
	SmfInfoList map[string]interface{} `json:"smfInfoList,omitempty"`

	// UpfInfo contains UPF-specific information.
	UpfInfo interface{} `json:"upfInfo,omitempty"`

	// UpfInfoList maps keys to UpfInfo entries.
	UpfInfoList map[string]interface{} `json:"upfInfoList,omitempty"`

	// PcfInfo contains PCF-specific information.
	PcfInfo interface{} `json:"pcfInfo,omitempty"`

	// PcfInfoList maps keys to PcfInfo entries.
	PcfInfoList map[string]interface{} `json:"pcfInfoList,omitempty"`

	// BsfInfo contains BSF-specific information.
	BsfInfo interface{} `json:"bsfInfo,omitempty"`

	// BsfInfoList maps keys to BsfInfo entries.
	BsfInfoList map[string]interface{} `json:"bsfInfoList,omitempty"`

	// ChfInfo contains CHF-specific information.
	ChfInfo interface{} `json:"chfInfo,omitempty"`

	// ChfInfoList maps keys to ChfInfo entries.
	ChfInfoList map[string]interface{} `json:"chfInfoList,omitempty"`

	// UdsfInfo contains UDSF-specific information.
	UdsfInfo interface{} `json:"udsfInfo,omitempty"`

	// UdsfInfoList maps keys to UdsfInfo entries.
	UdsfInfoList map[string]interface{} `json:"udsfInfoList,omitempty"`

	// NwdafInfo contains NWDAF-specific information.
	NwdafInfo interface{} `json:"nwdafInfo,omitempty"`

	// NwdafInfoList maps keys to NwdafInfo entries.
	NwdafInfoList map[string]interface{} `json:"nwdafInfoList,omitempty"`

	// NefInfo contains NEF-specific information.
	NefInfo interface{} `json:"nefInfo,omitempty"`

	// PcscfInfoList maps keys to PcscfInfo entries.
	PcscfInfoList map[string]interface{} `json:"pcscfInfoList,omitempty"`

	// HssInfoList maps keys to HssInfo entries.
	HssInfoList map[string]interface{} `json:"hssInfoList,omitempty"`

	// CustomInfo contains custom/vendor-specific information.
	CustomInfo interface{} `json:"customInfo,omitempty"`

	// RecoveryTime is the timestamp when the NF recovered.
	RecoveryTime *DateTime `json:"recoveryTime,omitempty"`

	// NfServicePersistence indicates whether NF service instances are persistent.
	NfServicePersistence bool `json:"nfServicePersistence,omitempty"`

	// NfServices is the deprecated list of NF services.
	NfServices []NFService `json:"nfServices,omitempty"`

	// NfServiceList maps serviceInstanceId to NFService entries.
	NfServiceList map[string]NFService `json:"nfServiceList,omitempty"`

	// DefaultNotificationSubscriptions lists default notification subscriptions.
	DefaultNotificationSubscriptions []interface{} `json:"defaultNotificationSubscriptions,omitempty"`

	// LmfInfo contains LMF-specific information.
	LmfInfo interface{} `json:"lmfInfo,omitempty"`

	// GmlcInfo contains GMLC-specific information.
	GmlcInfo interface{} `json:"gmlcInfo,omitempty"`

	// SnpnList lists SNPN IDs.
	SnpnList []PlmnIdNid `json:"snpnList,omitempty"`

	// NfSetIdList lists NF Set IDs.
	NfSetIdList []NfSetId `json:"nfSetIdList,omitempty"`

	// ServingScope lists areas that can be served.
	ServingScope []string `json:"servingScope,omitempty"`

	// LcHSupportInd indicates support for load control.
	LcHSupportInd bool `json:"lcHSupportInd,omitempty"`

	// OlcHSupportInd indicates support for overload control.
	OlcHSupportInd bool `json:"olcHSupportInd,omitempty"`

	// NfSetRecoveryTimeList maps NfSetId to recovery time.
	NfSetRecoveryTimeList map[string]DateTime `json:"nfSetRecoveryTimeList,omitempty"`

	// ServiceSetRecoveryTimeList maps NfServiceSetId to recovery time.
	ServiceSetRecoveryTimeList map[string]DateTime `json:"serviceSetRecoveryTimeList,omitempty"`

	// ScpDomains lists SCP domains the NF belongs to.
	ScpDomains []string `json:"scpDomains,omitempty"`

	// ScpInfo contains SCP-specific information.
	ScpInfo interface{} `json:"scpInfo,omitempty"`

	// SeppInfo contains SEPP-specific information.
	SeppInfo interface{} `json:"seppInfo,omitempty"`

	// VendorId is the vendor identifier.
	VendorId VendorId `json:"vendorId,omitempty"`

	// SupportedVendorSpecificFeatures maps enterprise codes to features.
	SupportedVendorSpecificFeatures map[string][]interface{} `json:"supportedVendorSpecificFeatures,omitempty"`

	// AanfInfoList maps keys to AanfInfo entries.
	AanfInfoList map[string]interface{} `json:"aanfInfoList,omitempty"`

	// MfafInfo contains MFAF-specific information.
	MfafInfo interface{} `json:"mfafInfo,omitempty"`

	// EasdfInfoList maps keys to EasdfInfo entries.
	EasdfInfoList map[string]interface{} `json:"easdfInfoList,omitempty"`

	// DccfInfo contains DCCF-specific information.
	DccfInfo interface{} `json:"dccfInfo,omitempty"`

	// NsacfInfoList maps keys to NsacfInfo entries.
	NsacfInfoList map[string]interface{} `json:"nsacfInfoList,omitempty"`

	// MbSmfInfoList maps keys to MbSmfInfo entries.
	MbSmfInfoList map[string]interface{} `json:"mbSmfInfoList,omitempty"`

	// TsctsfInfoList maps keys to TsctsfInfo entries.
	TsctsfInfoList map[string]interface{} `json:"tsctsfInfoList,omitempty"`

	// MbUpfInfoList maps keys to MbUpfInfo entries.
	MbUpfInfoList map[string]interface{} `json:"mbUpfInfoList,omitempty"`

	// TrustAfInfo contains Trusted AF-specific information.
	TrustAfInfo interface{} `json:"trustAfInfo,omitempty"`

	// NssaafInfo contains NSSAAF-specific information.
	NssaafInfo interface{} `json:"nssaafInfo,omitempty"`

	// HniList lists Home Network Identifiers.
	HniList []Fqdn `json:"hniList,omitempty"`

	// IwmscInfo contains IWMSC-specific information.
	IwmscInfo interface{} `json:"iwmscInfo,omitempty"`

	// MnpfInfo contains MNPF-specific information.
	MnpfInfo interface{} `json:"mnpfInfo,omitempty"`

	// SmsfInfo contains SMSF-specific information.
	SmsfInfo interface{} `json:"smsfInfo,omitempty"`

	// DcsfInfoList maps keys to DcsfInfo entries.
	DcsfInfoList map[string]interface{} `json:"dcsfInfoList,omitempty"`

	// MrfInfoList maps keys to MrfInfo entries.
	MrfInfoList map[string]interface{} `json:"mrfInfoList,omitempty"`

	// MrfpInfoList maps keys to MrfpInfo entries.
	MrfpInfoList map[string]interface{} `json:"mrfpInfoList,omitempty"`

	// MfInfoList maps keys to MfInfo entries.
	MfInfoList map[string]interface{} `json:"mfInfoList,omitempty"`

	// AdrfInfoList maps keys to AdrfInfo entries.
	AdrfInfoList map[string]interface{} `json:"adrfInfoList,omitempty"`

	// SelectionConditions contains selection condition parameters.
	SelectionConditions interface{} `json:"selectionConditions,omitempty"`

	// CanaryRelease indicates canary release participation.
	CanaryRelease bool `json:"canaryRelease,omitempty"`

	// ExclusiveCanaryReleaseSelection indicates exclusive canary selection.
	ExclusiveCanaryReleaseSelection bool `json:"exclusiveCanaryReleaseSelection,omitempty"`

	// SharedProfileDataId is the shared profile data identifier (UUID).
	SharedProfileDataId string `json:"sharedProfileDataId,omitempty"`
}

// ============================================================================
// NFService — NF Service Instance information
// ============================================================================

// NFService contains information of a given NF Service Instance.
// Reference: TS 29.510 Section 6.2.6.3.15 (NFService)
type NFService struct {
	// ServiceInstanceId is the service instance identifier. (Required)
	ServiceInstanceId string `json:"serviceInstanceId"`

	// ServiceName is the name of the service. (Required)
	ServiceName ServiceName `json:"serviceName"`

	// Versions lists the API versions supported. (Required)
	Versions []NFServiceVersion `json:"versions"`

	// Scheme is the URI scheme (http or https). (Required)
	Scheme UriScheme `json:"scheme"`

	// NfServiceStatus is the status of the NF service. (Required)
	NfServiceStatus NFServiceStatus `json:"nfServiceStatus"`

	// Fqdn is the FQDN of the NF service.
	Fqdn Fqdn `json:"fqdn,omitempty"`

	// InterPlmnFqdn is the FQDN for inter-PLMN communication.
	InterPlmnFqdn Fqdn `json:"interPlmnFqdn,omitempty"`

	// IpEndPoints lists the IP endpoints for this service.
	IpEndPoints []IpEndPoint `json:"ipEndPoints,omitempty"`

	// ApiPrefix is the API prefix for constructing service URIs.
	ApiPrefix string `json:"apiPrefix,omitempty"`

	// CallbackUriPrefixList lists callback URI prefix items.
	CallbackUriPrefixList []interface{} `json:"callbackUriPrefixList,omitempty"`

	// DefaultNotificationSubscriptions lists default notification subscriptions.
	DefaultNotificationSubscriptions []interface{} `json:"defaultNotificationSubscriptions,omitempty"`

	// AllowedPlmns lists PLMNs allowed to access this service.
	AllowedPlmns []PlmnId `json:"allowedPlmns,omitempty"`

	// AllowedSnpns lists SNPNs allowed to access this service.
	AllowedSnpns []PlmnIdNid `json:"allowedSnpns,omitempty"`

	// AllowedNfTypes lists NF types allowed to access this service.
	AllowedNfTypes []NFType `json:"allowedNfTypes,omitempty"`

	// AllowedNfDomains lists NF domains allowed to access this service.
	AllowedNfDomains []string `json:"allowedNfDomains,omitempty"`

	// AllowedNssais lists S-NSSAIs allowed for this service.
	AllowedNssais []ExtSnssai `json:"allowedNssais,omitempty"`

	// Capacity is the static capacity (0-65535).
	Capacity int `json:"capacity,omitempty"`

	// Load is the current load percentage (0-100).
	Load int `json:"load,omitempty"`

	// LoadTimeStamp is the timestamp when load was measured.
	LoadTimeStamp *DateTime `json:"loadTimeStamp,omitempty"`

	// Priority is the priority (0-65535, lower = higher).
	Priority int `json:"priority,omitempty"`

	// RecoveryTime is the timestamp when the service recovered.
	RecoveryTime *DateTime `json:"recoveryTime,omitempty"`

	// SupportedFeatures lists features supported by the service.
	SupportedFeatures SupportedFeatures `json:"supportedFeatures,omitempty"`

	// NfServiceSetIdList lists NF Service Set IDs.
	NfServiceSetIdList []NfServiceSetId `json:"nfServiceSetIdList,omitempty"`

	// SNssais lists S-NSSAIs for this service.
	SNssais []ExtSnssai `json:"sNssais,omitempty"`

	// PerPlmnSnssaiList lists S-NSSAIs per PLMN for this service.
	PerPlmnSnssaiList []interface{} `json:"perPlmnSnssaiList,omitempty"`

	// VendorId is the vendor identifier.
	VendorId VendorId `json:"vendorId,omitempty"`

	// SupportedVendorSpecificFeatures maps enterprise codes to features.
	SupportedVendorSpecificFeatures map[string][]interface{} `json:"supportedVendorSpecificFeatures,omitempty"`

	// Oauth2Required indicates if OAuth2 is required.
	Oauth2Required bool `json:"oauth2Required,omitempty"`

	// AllowedOperationsPerNfType maps NF type to allowed operations.
	AllowedOperationsPerNfType map[string][]string `json:"allowedOperationsPerNfType,omitempty"`

	// AllowedOperationsPerNfInstance maps NF instance ID to allowed operations.
	AllowedOperationsPerNfInstance map[string][]string `json:"allowedOperationsPerNfInstance,omitempty"`

	// AllowedOperationsPerNfInstanceOverrides indicates overrides.
	AllowedOperationsPerNfInstanceOverrides bool `json:"allowedOperationsPerNfInstanceOverrides,omitempty"`

	// AllowedScopesRuleSet maps JSON pointer IDs to rule sets.
	AllowedScopesRuleSet map[string]interface{} `json:"allowedScopesRuleSet,omitempty"`

	// SelectionConditions contains selection condition parameters.
	SelectionConditions interface{} `json:"selectionConditions,omitempty"`

	// CanaryRelease indicates canary release participation.
	CanaryRelease bool `json:"canaryRelease,omitempty"`

	// ExclusiveCanaryReleaseSelection indicates exclusive canary selection.
	ExclusiveCanaryReleaseSelection bool `json:"exclusiveCanaryReleaseSelection,omitempty"`

	// SharedServiceDataId is the shared service data identifier (UUID).
	SharedServiceDataId string `json:"sharedServiceDataId,omitempty"`
}

// ============================================================================
// Supporting Types
// ============================================================================

// PreferredSearch contains information on whether returned NFProfiles match
// the preferred query parameters.
// Reference: TS 29.510 Section 6.2.6.3.18 (PreferredSearch)
type PreferredSearch struct {
	PreferredTaiMatchInd               bool `json:"preferredTaiMatchInd,omitempty"`
	PreferredFullPlmnMatchInd          bool `json:"preferredFullPlmnMatchInd,omitempty"`
	PreferredApiVersionsMatchInd       bool `json:"preferredApiVersionsMatchInd,omitempty"`
	OtherApiVersionsInd                bool `json:"otherApiVersionsInd,omitempty"`
	PreferredLocalityMatchInd          bool `json:"preferredLocalityMatchInd,omitempty"`
	OtherLocalityInd                   bool `json:"otherLocalityInd,omitempty"`
	PreferredVendorSpecificFeaturesInd bool `json:"preferredVendorSpecificFeaturesInd,omitempty"`
	PreferredCollocatedNfTypeInd       bool `json:"preferredCollocatedNfTypeInd,omitempty"`
	PreferredPgwMatchInd               bool `json:"preferredPgwMatchInd,omitempty"`
	PreferredAnalyticsDelaysInd        bool `json:"preferredAnalyticsDelaysInd,omitempty"`
	PreferredFeaturesMatchInd          bool `json:"preferredFeaturesMatchInd,omitempty"`
	NoPreferredFeaturesInd             bool `json:"noPreferredFeaturesInd,omitempty"`
}

// NfInstanceInfo contains information on an NF profile matching a discovery request.
// Reference: TS 29.510 Section 6.2.6.3.19 (NfInstanceInfo)
type NfInstanceInfo struct {
	NrfDiscApiUri        Uri               `json:"nrfDiscApiUri,omitempty"`
	PreferredSearch      *PreferredSearch  `json:"preferredSearch,omitempty"`
	NrfAlteredPriorities map[string]int    `json:"nrfAlteredPriorities,omitempty"`
	NrfSupportedFeatures SupportedFeatures `json:"nrfSupportedFeatures,omitempty"`
}

// SearchResultInfo contains additional information to the search result.
// Reference: TS 29.510 Section 6.2.6.3.25 (SearchResultInfo)
type SearchResultInfo struct {
	UnsatisfiedTaiList []Tai `json:"unsatisfiedTaiList,omitempty"`
}

// NfServiceInstance identifies an NF service instance for exclusion queries.
// Reference: TS 29.510 Section 6.2.6.3.23 (NfServiceInstance)
type NfServiceInstance struct {
	ServiceInstanceId string          `json:"serviceInstanceId,omitempty"`
	NfInstanceId      *NfInstanceId   `json:"nfInstanceId,omitempty"`
	NfServiceSetId    *NfServiceSetId `json:"nfServiceSetId,omitempty"`
}

// ============================================================================
// Type Aliases for TS 29.510 NFManagement types not yet generated
// ============================================================================

// NFStatus represents the status of an NF instance.
// Reference: TS 29.510 Nnrf_NFManagement (NFStatus)
type NFStatus string

const (
	NFStatusRegistered     NFStatus = "REGISTERED"
	NFStatusSuspended      NFStatus = "SUSPENDED"
	NFStatusUndiscoverable NFStatus = "UNDISCOVERABLE"
)

// ServiceName represents a service name as defined in 3GPP TS 29.510.
type ServiceName string

// NFServiceStatus represents the status of an NF service.
type NFServiceStatus string

const (
	NFServiceStatusRegistered     NFServiceStatus = "REGISTERED"
	NFServiceStatusSuspended      NFServiceStatus = "SUSPENDED"
	NFServiceStatusUndiscoverable NFServiceStatus = "UNDISCOVERABLE"
)

// NFServiceVersion represents the API version of an NF service.
type NFServiceVersion struct {
	ApiVersionInUri  string        `json:"apiVersionInUri"`
	ApiFullVersion   string        `json:"apiFullVersion"`
	Expiry           *DateTime     `json:"expiry,omitempty"`
	Resources        []interface{} `json:"resources,omitempty"`
	CustomOperations []interface{} `json:"customOperations,omitempty"`
}

// IpEndPoint represents an IP endpoint of an NF service.
type IpEndPoint struct {
	Ipv4Address Ipv4Addr    `json:"ipv4Address,omitempty"`
	Ipv6Address Ipv6Addr    `json:"ipv6Address,omitempty"`
	Transport   interface{} `json:"transport,omitempty"`
	Port        int         `json:"port,omitempty"`
}

// NoProfileMatchReason represents the reason for no profile match.
// Values: REQUESTER_PLMN_NOT_ALLOWED, TARGET_NF_SUSPENDED,
// TARGET_NF_UNDISCOVERABLE, QUERY_PARAMS_COMBINATION_NO_MATCH,
// TARGET_NF_TYPE_NOT_SUPPORTED, UNSPECIFIED
type NoProfileMatchReason string

const (
	NoProfileMatchRequesterPlmnNotAllowed       NoProfileMatchReason = "REQUESTER_PLMN_NOT_ALLOWED"
	NoProfileMatchTargetNfSuspended             NoProfileMatchReason = "TARGET_NF_SUSPENDED"
	NoProfileMatchTargetNfUndiscoverable        NoProfileMatchReason = "TARGET_NF_UNDISCOVERABLE"
	NoProfileMatchQueryParamsCombinationNoMatch NoProfileMatchReason = "QUERY_PARAMS_COMBINATION_NO_MATCH"
	NoProfileMatchTargetNfTypeNotSupported      NoProfileMatchReason = "TARGET_NF_TYPE_NOT_SUPPORTED"
	NoProfileMatchUnspecified                   NoProfileMatchReason = "UNSPECIFIED"
)

// ============================================================================
// NrfDiscoveryParams — NRF NF Discovery query parameters
// Reference: 3GPP TS 29.510, Nnrf_NFDiscovery Service, GET /nf-instances
// ============================================================================

// NrfDiscoveryParams holds the parameters for an NRF NF Discovery query.
type NrfDiscoveryParams struct {
	// TargetNfType is the type of the target NF to discover (e.g., "AMF", "UDM"). (Required)
	TargetNfType string

	// RequesterNfType is the type of the requesting NF (e.g., "NWDAF"). (Required)
	RequesterNfType string

	// ServiceNames lists the service names the target NF must support.
	ServiceNames []ServiceName

	// Supi is the SUPI of the UE (for UE-specific discovery).
	Supi Supi

	// NfInstanceId is the specific NF instance ID to discover.
	NfInstanceId NfInstanceId

	// InternalGroupIdentity is an internal group identifier for group-based discovery.
	InternalGroupIdentity string

	// TargetNfInstanceIdList lists multiple NF instance IDs to discover.
	TargetNfInstanceIdList []NfInstanceId

	// Dnn is the Data Network Name for SMF/UPF discovery.
	Dnn Dnn

	// CompleteProfile requests complete NF profiles in the response.
	CompleteProfile bool

	// Limit limits the number of NF profiles returned.
	Limit int

	// PreferredLocality filters NFs by locality.
	PreferredLocality string

	// RequesterNfInstanceId is the NF instance ID of the requester.
	RequesterNfInstanceId NfInstanceId

	// PreferredCollocatedNfTypes lists preferred collocated NF types.
	PreferredCollocatedNfTypes []string

	// Guami is the GUAMI for AMF-specific discovery.
	Guami *Guami

	// SupportedFeatures lists supported features by the requester.
	SupportedFeatures SupportedFeatures

	// ExcludeNfInstList lists NF instances to exclude from results.
	ExcludeNfInstList []NfInstanceId

	// TargetNfSetId is the target NF Set ID.
	TargetNfSetId NfSetId

	// RequesterNfInstanceFqdn is the FQDN of the requester NF.
	RequesterNfInstanceFqdn Fqdn

	//
	Snssais []Snssai

	TaiList []Tai
}

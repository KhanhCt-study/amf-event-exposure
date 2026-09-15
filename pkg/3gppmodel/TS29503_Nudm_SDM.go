// Package model provides additional Nudm_SDM types not in the existing TS29503_Nudm.go.
//
// Reference: 3GPP TS 29.503 V18.5.0, Nudm_SDM Service
//
// This file covers:
//   - SDM Subscription request/response types (SdmSubscription, SdmSubsModification)
//   - User Consent Subscription Data types (UcSubscriptionData, UserConsent)
//   - Shared types (SubscriptionDataSets, DataSetName)
package gppmodel

// ============================================================================
// SDM Subscription Types
// ============================================================================

// SdmSubscription represents a subscription request to UDM for data change notifications.
// Reference: TS 29.503 Section 6.2.6.2.2 (SdmSubscription)
type SdmSubscription struct {
	// NfInstanceId is the NF instance ID of the service consumer. (Required)
	NfInstanceId NfInstanceId `json:"nfInstanceId"`

	// ImplicitUnsubscribe indicates implicit unsubscribe is supported.
	ImplicitUnsubscribe bool `json:"implicitUnsubscribe,omitempty"`

	// Expires is the subscription expiration time (ISO 8601).
	Expires DateTime `json:"expires,omitempty"`

	// CallbackReference is the URI where notifications are sent. (Required)
	CallbackReference Uri `json:"callbackReference"`

	// AmfServiceName is the service name of the AMF (for AMF subscriptions).
	AmfServiceName ServiceName `json:"amfServiceName,omitempty"`

	// MonitoredResourceUris lists the monitored resource URIs. (Required)
	MonitoredResourceUris []Uri `json:"monitoredResourceUris"`

	// SingleNssai is the S-NSSAI for slice-specific subscriptions.
	SingleNssai *Snssai `json:"singleNssai,omitempty"`

	// Dnn is the DNN for DNN-specific subscriptions.
	Dnn Dnn `json:"dnn,omitempty"`

	// SubscriptionId is the subscription identifier (set by UDM in response).
	SubscriptionId string `json:"subscriptionId,omitempty"`

	// PlmnId is the PLMN ID of the NF service consumer.
	PlmnId *PlmnId `json:"plmnId,omitempty"`

	// ImmediateReport indicates whether immediate reporting is requested.
	ImmediateReport bool `json:"immediateReport,omitempty"`

	// Report contains immediate report data if immediateReport is true.
	Report *ImmediateReport `json:"report,omitempty"`

	// SupportedFeatures lists supported features.
	SupportedFeatures SupportedFeatures `json:"supportedFeatures,omitempty"`

	// ContextInfo contains original request headers for context.
	ContextInfo *SdmContextInfo `json:"contextInfo,omitempty"`

	// NfChangeFilter indicates whether NF change filtering is requested.
	NfChangeFilter bool `json:"nfChangeFilter,omitempty"`

	// UniqueSubscription indicates this is a unique subscription.
	UniqueSubscription bool `json:"uniqueSubscription,omitempty"`

	// ResetIds lists reset identifiers for data restoration detection.
	ResetIds []string `json:"resetIds,omitempty"`

	// UeConSmfDataSubFilter contains SMF data subscription filter.
	UeConSmfDataSubFilter *UeContextInSmfDataSubFilter `json:"ueConSmfDataSubFilter,omitempty"`

	// AdjacentPlmns lists adjacent PLMN IDs.
	AdjacentPlmns []PlmnId `json:"adjacentPlmns,omitempty"`

	// DisasterRoamingInd indicates disaster roaming detection.
	DisasterRoamingInd bool `json:"disasterRoamingInd,omitempty"`

	// DataRestorationCallbackUri is the URI for data restoration notifications.
	DataRestorationCallbackUri Uri `json:"dataRestorationCallbackUri,omitempty"`

	// UdrRestartInd indicates UDR restart detection is requested.
	UdrRestartInd bool `json:"udrRestartInd,omitempty"`

	// ExpectedUeBehaviourThresholds maps JSON pointer → threshold data for UE behaviour.
	ExpectedUeBehaviourThresholds map[string]ExpectedUeBehaviourThreshold `json:"expectedUeBehaviourThresholds,omitempty"`
}

// SdmSubsModification represents a modification request for an SDM subscription.
// Reference: TS 29.503 Section 6.2.6.2.5 (SdmSubsModification)
type SdmSubsModification struct {
	// Expires is the new expiration time.
	Expires DateTime `json:"expires,omitempty"`

	// MonitoredResourceUris lists updated monitored resource URIs.
	MonitoredResourceUris []Uri `json:"monitoredResourceUris,omitempty"`

	// ExpectedUeBehaviourThresholds maps JSON pointer → threshold data.
	ExpectedUeBehaviourThresholds map[string]ExpectedUeBehaviourThreshold `json:"expectedUeBehaviourThresholds,omitempty"`
}

// ============================================================================
// User Consent Types
// ============================================================================

// UcSubscriptionData represents User Consent Subscription Data.
// Reference: TS 29.503 Section 5.4.2.2 (UcSubscriptionData)
type UcSubscriptionData struct {
	// UserConsentPerPurposeList maps consent purpose → user consent.
	UserConsentPerPurposeList map[string]UserConsent `json:"userConsentPerPurposeList"`
}

// ============================================================================
// SDM Context Info (separate from EE ContextInfo to avoid conflicts)
// ============================================================================

// SdmContextInfo contains original request headers for SDM context.
// Reference: TS 29.503 (ContextInfo)
type SdmContextInfo struct {
	// OrigHeaders contains the original request headers.
	OrigHeaders []string `json:"origHeaders,omitempty"`

	// RequestHeaders contains the request headers.
	RequestHeaders []string `json:"requestHeaders,omitempty"`
}

// ============================================================================
// UE Context In SMF Data Subscription Filter
// ============================================================================

// UeContextInSmfDataSubFilter is the subscription filter for UE Context In SMF Data.
type UeContextInSmfDataSubFilter struct {
	// DnnList is the list of DNNs to filter by.
	DnnList []Dnn `json:"dnnList,omitempty"`

	// SnssaiList is the list of S-NSSAIs to filter by.
	SnssaiList []Snssai `json:"snssaiList,omitempty"`

	// EmergencyInd indicates whether emergency sessions are included.
	EmergencyInd bool `json:"emergencyInd,omitempty"`
}

// ============================================================================
// Expected UE Behaviour Threshold
// ============================================================================

// ExpectedUeBehaviourThreshold contains thresholds for UE behaviour monitoring.
type ExpectedUeBehaviourThreshold struct {
	// ExpecedUeBehaviourDatasets lists the expected UE behaviour dataset names.
	ExpecedUeBehaviourDatasets []ExpecedUeBehaviourDataset `json:"expecedUeBehaviourDatasets,omitempty"`

	// SingleNssais is the list of S-NSSAIs to filter by.
	SingleNssais []Snssai `json:"singleNssais,omitempty"`

	// Dnns is the list of DNNs to filter by.
	Dnns []Dnn `json:"dnns,omitempty"`

	// ConfidenceLevel is the confidence level threshold (pattern: 0.00-1.00).
	ConfidenceLevel string `json:"confidenceLevel,omitempty"`

	// AccuracyLevel is the accuracy level threshold (pattern: 0.00-1.00).
	AccuracyLevel string `json:"accuracyLevel,omitempty"`
}

// ============================================================================
// Immediate Report (for SDM immediate reporting)
// ============================================================================

// ImmediateReport represents immediate report data returned in a subscription response.
type ImmediateReport struct {
	// SubscriptionDataSets contains the subscription data.
	SubscriptionDataSets *SubscriptionDataSetsSummary `json:"subscriptionDataSets,omitempty"`

	// Items contains additional items.
	Items []SubscriptionDataSetsSummary `json:"items,omitempty"`
}

// SubscriptionDataSetsSummary is a summary view of subscription data sets.
type SubscriptionDataSetsSummary struct {
	// AmData contains access and mobility subscription data.
	AmData *AccessAndMobilitySubscriptionDataSDM `json:"amData,omitempty"`

	// SmfSelData contains SMF selection subscription data.
	SmfSelData *SmfSelectionSubscriptionDataSDM `json:"smfSelData,omitempty"`

	// UecAmfData contains UE Context In AMF Data.
	UecAmfData *UeContextInAmfDataSDM `json:"uecAmfData,omitempty"`

	// UecSmfData contains UE Context In SMF Data.
	UecSmfData *UeContextInSmfDataSDM `json:"uecSmfData,omitempty"`

	// SmData contains Session Management Subscription Data.
	SmData *SmSubsDataSDM `json:"smData,omitempty"`

	// UcData contains User Consent Subscription Data.
	UcData *UcSubscriptionData `json:"ucData,omitempty"`
}

// ============================================================================
// SDM-specific Subset Types (different from full versions in YAML spec)
// ============================================================================

// AccessAndMobilitySubscriptionDataSDM is the SDM view of AM subscription data.
type AccessAndMobilitySubscriptionDataSDM struct {
	SupportedFeatures SupportedFeatures `json:"supportedFeatures,omitempty"`
	Gpsis             []Gpsi            `json:"gpsis,omitempty"`
	Nssai             *NssaiSDM         `json:"nssai,omitempty"`
	// Simplified: only fields needed for consent/context checks
}

// NssaiSDM is a simplified NSSAI for SDM responses.
type NssaiSDM struct {
	DefaultSingleNssais []Snssai `json:"defaultSingleNssais"`
	SingleNssais        []Snssai `json:"singleNssais,omitempty"`
}

// SmfSelectionSubscriptionDataSDM is the SDM view of SMF selection data.
type SmfSelectionSubscriptionDataSDM struct {
	SupportedFeatures SupportedFeatures `json:"supportedFeatures,omitempty"`
}

// UeContextInAmfDataSDM is the SDM view of UE Context In AMF Data.
type UeContextInAmfDataSDM struct {
	AmfInfo []AmfInfoSDM `json:"amfInfo,omitempty"`
}

// AmfInfoSDM is the SDM view of AMF info.
type AmfInfoSDM struct {
	AmfInstanceId NfInstanceId `json:"amfInstanceId"`
}

// UeContextInSmfDataSDM is the SDM view of UE Context In SMF Data.
type UeContextInSmfDataSDM struct {
	PduSessions map[string]PduSessionSDM `json:"pduSessions,omitempty"`
}

// PduSessionSDM is the SDM view of a PDU session.
type PduSessionSDM struct {
	Dnn           Dnn          `json:"dnn"`
	SmfInstanceId NfInstanceId `json:"smfInstanceId"`
	PlmnId        *PlmnId      `json:"plmnId,omitempty"`
}

// SmSubsDataSDM is the SDM view of SM subscription data.
type SmSubsDataSDM struct {
	DnnConfigurations map[string]interface{} `json:"dnnConfigurations,omitempty"`
}

// ============================================================================
// Group Identifiers (Nudm_SDM /group-data)
// Reference: TS 29.503 Nudm_SDM Section 6.2.6.2.31 (GroupIdentifiers)
// ============================================================================

// ExtGroupId represents an External Group Identifier.
// Pattern: "extgroupid-<local-part>@<domain-part>"
// Reference: TS 29.503 (ExtGroupId)
type ExtGroupId string

// UeId represents a UE identity within a group identifier response.
// Reference: TS 29.503 (UeId)
type UeId struct {
	// Supi is the Subscription Permanent Identifier. (Required)
	Supi Supi `json:"supi"`

	// GpsiList lists Generic Public Subscription Identifiers for the UE.
	GpsiList []Gpsi `json:"gpsiList,omitempty"`
}

// GroupIdentifiers represents the response from GET /group-data/group-identifiers.
// It provides the mapping between external/internal group IDs and UE identities.
// Reference: TS 29.503 Section 6.2.6.2.31 (GroupIdentifiers)
type GroupIdentifiers struct {
	// ExtGroupId is the External Group Identifier.
	ExtGroupId ExtGroupId `json:"extGroupId,omitempty"`

	// IntGroupId is the Internal Group Identifier.
	IntGroupId GroupId `json:"intGroupId,omitempty"`

	// UeIdList contains the list of UE identities belonging to the group.
	UeIdList []UeId `json:"ueIdList,omitempty"`
}

// GetGroupIdentifiersParams holds the query parameters for
// GET /group-data/group-identifiers on the UDM SDM service.
// Reference: TS 29.503 Nudm_SDM (GetGroupIdentifiers)
type GetGroupIdentifiersParams struct {
	// ExtGroupId is the External Group Identifier.
	// Mutually exclusive with IntGroupId.
	ExtGroupId ExtGroupId

	// IntGroupId is the Internal Group Identifier.
	// Mutually exclusive with ExtGroupId.
	IntGroupId GroupId

	// UeIdInd, if true, requests UE identities (SUPI/GPSI) in the response.
	UeIdInd bool

	// SupportedFeatures optionally lists supported features.
	SupportedFeatures SupportedFeatures

	// AfId is an optional AF identifier.
	AfId string
}

// ============================================================================
// Enums
// ============================================================================

// UserConsent indicates user's consent status.
// Reference: TS 29.503 (UserConsent)
type UserConsent string

const (
	// ConsentGiven indicates the user has given consent.
	ConsentGiven UserConsent = "CONSENT_GIVEN"

	// ConsentNotGiven indicates the user has not given consent.
	ConsentNotGiven UserConsent = "CONSENT_NOT_GIVEN"
)

// DataSetName identifies the requested data set for SDM operations.
// Reference: TS 29.503 Table 6.2.6.2.1-1
type DataSetName string

const (
	DataSetNameAm              DataSetName = "AM"
	DataSetNameSmfSel          DataSetName = "SMF_SEL"
	DataSetNameUecAmf          DataSetName = "UEC_AMF"
	DataSetNameUecSmf          DataSetName = "UEC_SMF"
	DataSetNameUecSmsf         DataSetName = "UEC_SMSF"
	DataSetNameSmsSubs         DataSetName = "SMS_SUBS"
	DataSetNameSm              DataSetName = "SM"
	DataSetNameTrace           DataSetName = "TRACE"
	DataSetNameSmsMng          DataSetName = "SMS_MNG"
	DataSetNameLcsPrivacy      DataSetName = "LCS_PRIVACY"
	DataSetNameLcsMo           DataSetName = "LCS_MO"
	DataSetNameLcsSubscription DataSetName = "LCS_SUBSCRIPTION"
	DataSetNameV2x             DataSetName = "V2X"
	DataSetNameLcsBca          DataSetName = "LCS_BCA"
	DataSetNameProse           DataSetName = "PROSE"
	DataSetNameMbs             DataSetName = "MBS"
	DataSetNameUc              DataSetName = "UC"
	DataSetNameA2x             DataSetName = "A2X"
)

// UcPurpose defines the purpose of user consent.
// Reference: TS 29.503 (UcPurpose)
type UcPurpose string

const (
	UcPurposeAnalytics         UcPurpose = "ANALYTICS"
	UcPurposeModelTraining     UcPurpose = "MODEL_TRAINING"
	UcPurposeNwCapExposure     UcPurpose = "NW_CAP_EXPOSURE"
	UcPurposeEdgeappUeLocation UcPurpose = "EDGEAPP_UE_LOCATION"
)

// ExpecedUeBehaviourDataset identifies a behavioural dataset.
type ExpecedUeBehaviourDataset string

const (
	ExpecedUeBehaviourStationaryIndication       ExpecedUeBehaviourDataset = "STATIONARY_INDICATION"
	ExpecedUeBehaviourCommunicationDurationTime  ExpecedUeBehaviourDataset = "COMMUNICATION_DURATION_TIME"
	ExpecedUeBehaviourPeriodicTime               ExpecedUeBehaviourDataset = "PERIODIC_TIME"
	ExpecedUeBehaviourScheduledCommunicationTime ExpecedUeBehaviourDataset = "SCHEDULED_COMMUNICATION_TIME"
	ExpecedUeBehaviourTrafficProfile             ExpecedUeBehaviourDataset = "TRAFFIC_PROFILE"
	ExpecedUeBehaviourBatteryIndication          ExpecedUeBehaviourDataset = "BATTERY_INDICATION"
	ExpecedUeBehaviourExpectedUmts               ExpecedUeBehaviourDataset = "EXPECTED_UMTS"
	ExpecedUeBehaviourExpectedInactivityTime     ExpecedUeBehaviourDataset = "EXPECTED_INACTIVITY_TIME"
)

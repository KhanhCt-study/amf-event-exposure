// Package model provides primitives for Nudm_EE Event Exposure Service.
//
// Reference: 3GPP TS 29.503 V18.4.0, Nudm_EventExposure Service
//
// This file contains spec-accurate Go structs for:
//   - EE Subscription request/response types
//   - Monitoring configuration and report types
//   - Event type enums and supporting simple types
//
// Shared types reused: PlmnId, Snssai, NfInstanceId, Uri, DateTime, etc.
package gppmodel

// ============================================================================
// Core Subscription Types
// ============================================================================

// EeSubscription represents the request body for creating an EE subscription.
// Reference: TS 29.503 Section 6.2.6.2.2 (EeSubscription)
type EeSubscription struct {
	// CallbackReference is the URI where event notifications are sent. (Required)
	CallbackReference Uri `json:"callbackReference"`

	// MonitoringConfigurations is a map of ReferenceId → MonitoringConfiguration. (Required)
	MonitoringConfigurations map[string]MonitoringConfiguration `json:"monitoringConfigurations"`

	// ReportingOptions contains global reporting options for all monitoring configs.
	ReportingOptions *ReportingOptions `json:"reportingOptions,omitempty"`

	// SupportedFeatures lists features supported by the NF service consumer.
	SupportedFeatures SupportedFeatures `json:"supportedFeatures,omitempty"`

	// SubscriptionId is the subscription identifier (set by UDM in response).
	SubscriptionId string `json:"subscriptionId,omitempty"`

	// ContextInfo contains original request headers for context.
	ContextInfo *ContextInfo `json:"contextInfo,omitempty"`

	// EpcAppliedInd indicates whether the subscription is applied in EPC.
	EpcAppliedInd bool `json:"epcAppliedInd,omitempty"`

	// ScefDiamHost is the SCEF Diameter host identity.
	ScefDiamHost DiameterIdentity `json:"scefDiamHost,omitempty"`

	// ScefDiamRealm is the SCEF Diameter realm identity.
	ScefDiamRealm DiameterIdentity `json:"scefDiamRealm,omitempty"`

	// NotifyCorrelationId is the notification correlation identifier.
	NotifyCorrelationId string `json:"notifyCorrelationId,omitempty"`

	// SecondCallbackRef is the second callback URI for monitoring revocation.
	SecondCallbackRef Uri `json:"secondCallbackRef,omitempty"`

	// Gpsi is the GPSI of the target UE.
	Gpsi Gpsi `json:"gpsi,omitempty"`

	// ExcludeGpsiList lists GPSIs to exclude from group monitoring.
	ExcludeGpsiList []Gpsi `json:"excludeGpsiList,omitempty"`

	// IncludeGpsiList lists GPSIs to include in group monitoring.
	IncludeGpsiList []Gpsi `json:"includeGpsiList,omitempty"`

	// DataRestorationCallbackUri is the URI for data restoration notifications.
	DataRestorationCallbackUri Uri `json:"dataRestorationCallbackUri,omitempty"`

	// UdrRestartInd indicates whether UDR restart detection is requested.
	UdrRestartInd bool `json:"udrRestartInd,omitempty"`
}

// CreatedEeSubscription represents the response body for a created EE subscription.
// Reference: TS 29.503 Section 6.2.6.2.3 (CreatedEeSubscription)
type CreatedEeSubscription struct {
	// EeSubscription is the created subscription data. (Required)
	EeSubscription *EeSubscription `json:"eeSubscription"`

	// NumberOfUes is the number of UEs affected by the subscription.
	NumberOfUes Uinteger `json:"numberOfUes,omitempty"`

	// EventReports contains immediate event reports if immediateFlag was set.
	EventReports []MonitoringReport `json:"eventReports,omitempty"`

	// EpcStatusInd indicates EPC status.
	EpcStatusInd bool `json:"epcStatusInd,omitempty"`

	// FivegOnlyInd indicates 5G-only status.
	FivegOnlyInd bool `json:"5gOnlyInd,omitempty"`

	// FailedMonitoringConfigs maps referenceId → FailedMonitoringConfiguration.
	FailedMonitoringConfigs map[string]FailedMonitoringConfiguration `json:"failedMonitoringConfigs,omitempty"`

	// FailedMoniConfigsEPC maps referenceId → FailedMonitoringConfiguration for EPC.
	FailedMoniConfigsEPC map[string]FailedMonitoringConfiguration `json:"failedMoniConfigsEPC,omitempty"`

	// ResetIds lists reset identifiers.
	ResetIds []string `json:"resetIds,omitempty"`

	// CurrentStatusNotAvailableList lists event types whose current status is unavailable.
	CurrentStatusNotAvailableList []EventType `json:"currentStatusNotAvailableList,omitempty"`
}

// ============================================================================
// Monitoring Configuration Types
// ============================================================================

// MonitoringConfiguration contains the configuration for a single monitoring event.
// Reference: TS 29.503 Section 6.2.6.2.5 (MonitoringConfiguration)
type MonitoringConfiguration struct {
	// EventType identifies the type of event to monitor. (Required)
	EventType EventType `json:"eventType"`

	// ImmediateFlag requests an immediate report.
	ImmediateFlag bool `json:"immediateFlag,omitempty"`

	// LocationReportingConfiguration specifies location reporting parameters.
	LocationReportingConfiguration *LocationReportingConfiguration `json:"locationReportingConfiguration,omitempty"`

	// AssociationType indicates IMEI or IMEISV change association.
	AssociationType AssociationType `json:"associationType,omitempty"`

	// DatalinkReportCfg configures data link delivery status reporting.
	DatalinkReportCfg *DatalinkReportingConfiguration `json:"datalinkReportCfg,omitempty"`

	// LossConnectivityCfg configures loss of connectivity detection.
	LossConnectivityCfg *LossConnectivityCfg `json:"lossConnectivityCfg,omitempty"`

	// MaximumLatency is the maximum latency in seconds.
	MaximumLatency DurationSec `json:"maximumLatency,omitempty"`

	// MaximumResponseTime is the maximum response time in seconds.
	MaximumResponseTime DurationSec `json:"maximumResponseTime,omitempty"`

	// SuggestedPacketNumDl is the suggested number of downlink packets.
	SuggestedPacketNumDl int `json:"suggestedPacketNumDl,omitempty"`

	// Dnn is the Data Network Name for PDU session monitoring.
	Dnn Dnn `json:"dnn,omitempty"`

	// SingleNssai is the S-NSSAI for slice-specific monitoring.
	SingleNssai *Snssai `json:"singleNssai,omitempty"`

	// AppId is the Application ID for app-specific monitoring.
	AppId ApplicationId `json:"appId,omitempty"`

	// PduSessionStatusCfg configures PDU session status reporting.
	PduSessionStatusCfg *PduSessionStatusCfg `json:"pduSessionStatusCfg,omitempty"`

	// ReachabilityForSmsCfg configures SMS reachability reporting.
	ReachabilityForSmsCfg *ReachabilityForSmsConfiguration `json:"reachabilityForSmsCfg,omitempty"`

	// MtcProviderInformation is the MTC provider identity.
	MtcProviderInformation MtcProviderInformation `json:"mtcProviderInformation,omitempty"`

	// AfId is the Application Function identifier.
	AfId string `json:"afId,omitempty"`

	// ReachabilityForDataCfg configures data reachability reporting.
	ReachabilityForDataCfg *ReachabilityForDataConfiguration `json:"reachabilityForDataCfg,omitempty"`

	// IdleStatusInd requests idle status indication.
	IdleStatusInd bool `json:"idleStatusInd,omitempty"`

	// MonitoringSuspension contains monitoring suspension configuration.
	MonitoringSuspension *MonitoringSuspension `json:"monitoringSuspension,omitempty"`
}

// FailedMonitoringConfiguration contains the event type and failed cause.
// Reference: TS 29.503 Section 6.2.6.2.6 (FailedMonitoringConfiguration)
type FailedMonitoringConfiguration struct {
	// EventType is the event type that failed. (Required)
	EventType EventType `json:"eventType"`

	// FailedCause is the reason for failure. (Required)
	FailedCause FailedCause `json:"failedCause"`
}

// ============================================================================
// Report Types
// ============================================================================

// MonitoringReport represents a monitoring event notification.
// Reference: TS 29.503 Section 6.2.6.2.23 (MonitoringReport)
type MonitoringReport struct {
	// ReferenceId is the reference identifier for the monitoring event. (Required)
	ReferenceId ReferenceId `json:"referenceId"`

	// EventType identifies the event type. (Required)
	EventType EventType `json:"eventType"`

	// Report contains the event-specific report data.
	Report *Report `json:"report,omitempty"`

	// ReachabilityForSmsReport contains SMS reachability report data.
	ReachabilityForSmsReport *ReachabilityForSmsReport `json:"reachabilityForSmsReport,omitempty"`

	// Gpsi is the GPSI of the affected UE.
	Gpsi Gpsi `json:"gpsi,omitempty"`

	// TimeStamp is the timestamp when the event was detected. (Required)
	TimeStamp DateTime `json:"timeStamp"`

	// ReachabilityReport contains data reachability report data.
	ReachabilityReport *ReachabilityReport `json:"reachabilityReport,omitempty"`
}

// Report is a union of all possible event report types.
// Reference: TS 29.503 (Report)
type Report struct {
	ChangeOfSupiPeiAssociation *ChangeOfSupiPeiAssociationReport `json:"changeOfSupiPeiAssociationReport,omitempty"`
	RoamingStatus              *RoamingStatusReport              `json:"roamingStatusReport,omitempty"`
	CnTypeChange               *CnTypeChangeReport               `json:"cnTypeChangeReport,omitempty"`
	CmInfo                     *CmInfoReport                     `json:"cmInfoReport,omitempty"`
	LossConnectivity           *LossConnectivityReport           `json:"lossConnectivityReport,omitempty"`
	Location                   *LocationReport                   `json:"locationReport,omitempty"`
	PdnConnectivityStat        *PdnConnectivityStatReport        `json:"pdnConnectivityStatReport,omitempty"`
	GroupMembListChanges       *GroupMembListChanges             `json:"groupMembListChanges,omitempty"`
}

// ChangeOfSupiPeiAssociationReport reports a SUPI-PEI association change.
type ChangeOfSupiPeiAssociationReport struct {
	NewPei Pei `json:"newPei"`
}

// RoamingStatusReport reports roaming status changes.
type RoamingStatusReport struct {
	Roaming         bool    `json:"roaming"`
	NewServingPlmn  *PlmnId `json:"newServingPlmn"`
	AccessType      AccessType `json:"accessType,omitempty"`
	Purged          *bool   `json:"purged,omitempty"`
}

// CnTypeChangeReport reports core network type changes.
type CnTypeChangeReport struct {
	NewCnType CnType `json:"newCnType"`
	OldCnType CnType `json:"oldCnType,omitempty"`
}

// CmInfoReport reports connection management state changes.
type CmInfoReport struct {
	OldCmInfoList []CmInfo `json:"oldCmInfoList,omitempty"`
	NewCmInfoList []CmInfo `json:"newCmInfoList"`
}

// LossConnectivityReport reports loss of connectivity.
type LossConnectivityReport struct {
	LossOfConnectReason LossOfConnectivityReason `json:"lossOfConnectReason"`
}

// LocationReport reports UE location information.
type LocationReport struct {
	Location *UserLocation `json:"location"`
}

// PdnConnectivityStatReport reports PDN connectivity status.
type PdnConnectivityStatReport struct {
	PdnConnStat  PdnConnectivityStatus `json:"pdnConnStat"`
	Dnn          Dnn                   `json:"dnn,omitempty"`
	PduSeId      PduSessionId          `json:"pduSeId,omitempty"`
	Ipv4Addr     Ipv4Addr              `json:"ipv4Addr,omitempty"`
	Ipv6Prefixes []Ipv6Prefix          `json:"ipv6Prefixes,omitempty"`
	Ipv6Addrs    []Ipv6Addr            `json:"ipv6Addrs,omitempty"`
	PduSessType  PduSessionType        `json:"pduSessType,omitempty"`
}

// GroupMembListChanges reports changes to a group's member list.
type GroupMembListChanges struct {
	AddedUEs   []Gpsi `json:"addedUEs,omitempty"`
	RemovedUEs []Gpsi `json:"removedUEs,omitempty"`
}

// ReachabilityForSmsReport reports SMS reachability status.
type ReachabilityForSmsReport struct {
	SmsfAccessType      AccessType `json:"smsfAccessType"`
	MaxAvailabilityTime DateTime   `json:"maxAvailabilityTime,omitempty"`
}

// ReachabilityReport reports data reachability status.
type ReachabilityReport struct {
	AmfInstanceId        NfInstanceId     `json:"amfInstanceId,omitempty"`
	AccessTypeList       []AccessType     `json:"accessTypeList,omitempty"`
	Reachability         UeReachability   `json:"reachability,omitempty"`
	MaxAvailabilityTime  DateTime         `json:"maxAvailabilityTime,omitempty"`
	IdleStatusIndication *IdleStatusIndication `json:"idleStatusIndication,omitempty"`
}

// ============================================================================
// Configuration Sub-Types
// ============================================================================

// ReportingOptions contains global reporting options for monitoring events.
// Reference: TS 29.503 (ReportingOptions)
type ReportingOptions struct {
	ReportMode          EventReportMode                 `json:"reportMode,omitempty"`
	MaxNumOfReports     MaxNumOfReports                 `json:"maxNumOfReports,omitempty"`
	Expiry              DateTime                        `json:"expiry,omitempty"`
	SamplingRatio       SamplingRatio                   `json:"samplingRatio,omitempty"`
	GuardTime           DurationSec                     `json:"guardTime,omitempty"`
	ReportPeriod        DurationSec                     `json:"reportPeriod,omitempty"`
	NotifFlag           NotificationFlag                `json:"notifFlag,omitempty"`
	MutingExcInstructions *MutingExceptionInstructions  `json:"mutingExcInstructions,omitempty"`
	MutingNotSettings   *MutingNotificationsSettings    `json:"mutingNotSettings,omitempty"`
	VarRepPeriodInfo    []VarRepPeriod                  `json:"varRepPeriodInfo,omitempty"`
}

// LocationReportingConfiguration specifies location reporting parameters.
type LocationReportingConfiguration struct {
	CurrentLocation bool             `json:"currentLocation"`
	OneTime         bool             `json:"oneTime,omitempty"`
	Accuracy        LocationAccuracy `json:"accuracy,omitempty"`
	N3gppAccuracy   LocationAccuracy `json:"n3gppAccuracy,omitempty"`
}

// DatalinkReportingConfiguration configures data link delivery status reporting.
type DatalinkReportingConfiguration struct {
	DddTrafficDes []DddTrafficDescriptor `json:"dddTrafficDes,omitempty"`
	Dnn           Dnn                    `json:"dnn,omitempty"`
	Slice         *Snssai                `json:"slice,omitempty"`
	DddStatusList []DlDataDeliveryStatus `json:"dddStatusList,omitempty"`
}

// LossConnectivityCfg configures loss of connectivity detection.
type LossConnectivityCfg struct {
	MaxDetectionTime DurationSec `json:"maxDetectionTime,omitempty"`
}

// PduSessionStatusCfg configures PDU session status reporting.
type PduSessionStatusCfg struct {
	Dnn Dnn `json:"dnn,omitempty"`
}

// ReachabilityForDataConfiguration configures data reachability reporting.
type ReachabilityForDataConfiguration struct {
	ReportCfg   ReachabilityForDataReportConfig `json:"reportCfg"`
	MinInterval DurationSec                     `json:"minInterval,omitempty"`
}

// MonitoringSuspension contains monitoring suspension configuration.
type MonitoringSuspension struct {
	SuspendedInsidePlmnList  []PlmnIdNid `json:"suspendedInsidePlmnList,omitempty"`
	SuspendedOutsidePlmnList []PlmnIdNid `json:"suspendedOutsidePlmnList,omitempty"`
}

// ============================================================================
// Revocation Types
// ============================================================================

// EeMonitoringRevoked represents a monitoring revocation notification.
// Reference: TS 29.503 (EeMonitoringRevoked)
type EeMonitoringRevoked struct {
	RevokedMonitoringEventList map[string]MonitoringEvent `json:"revokedMonitoringEventList"`
	RemovedGpsi                Gpsi                       `json:"removedGpsi,omitempty"`
	ExcludeGpsiList            []Gpsi                     `json:"excludeGpsiList,omitempty"`
}

// MonitoringEvent contains a revoked monitoring event type and cause.
type MonitoringEvent struct {
	EventType    EventType    `json:"eventType"`
	RevokedCause RevokedCause `json:"revokedCause,omitempty"`
}

// ============================================================================
// Error Types
// ============================================================================

// EeSubscriptionError represents an EE subscription error response.
// Reference: TS 29.503 (EeSubscriptionError)
type EeSubscriptionError struct {
	ProblemDetails
	EeSubscriptionErrorAddInfo
}

// EeSubscriptionErrorAddInfo contains additional info for EE subscription errors.
type EeSubscriptionErrorAddInfo struct {
	SubType                   SubscriptionType                          `json:"subType,omitempty"`
	FailedMonitoringConfigs   map[string]FailedMonitoringConfiguration  `json:"failedMonitoringConfigs,omitempty"`
	FailedMoniConfigsEPC      map[string]FailedMonitoringConfiguration  `json:"failedMoniConfigsEPC,omitempty"`
}

// ContextInfo is reused from SDM. Defined here for EE usage.
// Reference: TS 29.503 (ContextInfo)
type ContextInfo struct {
	OrigHeaders    []string `json:"origHeaders,omitempty"`
	RequestHeaders []string `json:"requestHeaders,omitempty"`
}

// ============================================================================
// Simple Types
// ============================================================================

// ReferenceId is a uint64 reference identifier for monitoring configuration entries.
type ReferenceId = Uint64

// MaxNumOfReports is the maximum number of reports.
type MaxNumOfReports = int

// ============================================================================
// Enums
// ============================================================================

// EventType identifies the type of monitoring event in Nudm_EE.
// Reference: TS 29.503 Table 6.2.6.2.5-1
type EventType string

const (
	EventTypeLossOfConnectivity           EventType = "LOSS_OF_CONNECTIVITY"
	EventTypeUeReachabilityForData        EventType = "UE_REACHABILITY_FOR_DATA"
	EventTypeUeReachabilityForSms         EventType = "UE_REACHABILITY_FOR_SMS"
	EventTypeLocationReporting            EventType = "LOCATION_REPORTING"
	EventTypeChangeOfSupiPeiAssociation   EventType = "CHANGE_OF_SUPI_PEI_ASSOCIATION"
	EventTypeRoamingStatus                EventType = "ROAMING_STATUS"
	EventTypeCommunicationFailure         EventType = "COMMUNICATION_FAILURE"
	EventTypeAvailabilityAfterDdnFailure  EventType = "AVAILABILITY_AFTER_DDN_FAILURE"
	EventTypeCnTypeChange                 EventType = "CN_TYPE_CHANGE"
	EventTypeDlDataDeliveryStatus         EventType = "DL_DATA_DELIVERY_STATUS"
	EventTypePdnConnectivityStatus        EventType = "PDN_CONNECTIVITY_STATUS"
	EventTypeUeConnectionManagementState  EventType = "UE_CONNECTION_MANAGEMENT_STATE"
	EventTypeAccessTypeReport             EventType = "ACCESS_TYPE_REPORT"
	EventTypeRegistrationStateReport      EventType = "REGISTRATION_STATE_REPORT"
	EventTypeConnectivityStateReport      EventType = "CONNECTIVITY_STATE_REPORT"
	EventTypeTypeAllocationCodeReport     EventType = "TYPE_ALLOCATION_CODE_REPORT"
	EventTypeFrequentMobilityRegistration EventType = "FREQUENT_MOBILITY_REGISTRATION_REPORT"
	EventTypePduSesRel                    EventType = "PDU_SES_REL"
	EventTypePduSesEst                    EventType = "PDU_SES_EST"
	EventTypeUeMemoryAvailableForSms      EventType = "UE_MEMORY_AVAILABLE_FOR_SMS"
	EventTypeGroupMemberListChange        EventType = "GROUP_MEMBER_LIST_CHANGE"
	EventTypeQosMon                       EventType = "QOS_MON"
)

// LocationAccuracy indicates the location accuracy level.
type LocationAccuracy string

const (
	LocationAccuracyCellLevel  LocationAccuracy = "CELL_LEVEL"
	LocationAccuracyRanNodeLevel LocationAccuracy = "RAN_NODE_LEVEL"
	LocationAccuracyTaLevel    LocationAccuracy = "TA_LEVEL"
	LocationAccuracyN3iwfLevel LocationAccuracy = "N3IWF_LEVEL"
	LocationAccuracyUeIp       LocationAccuracy = "UE_IP"
	LocationAccuracyUePort     LocationAccuracy = "UE_PORT"
)

// CnType indicates the core network type.
type CnType string

const (
	CnTypeSingle4g  CnType = "SINGLE_4G"
	CnTypeSingle5g  CnType = "SINGLE_5G"
	CnTypeDual4g5g  CnType = "DUAL_4G5G"
)

// AssociationType indicates whether IMEI or IMEISV change is monitored.
type AssociationType string

const (
	AssociationTypeImeiChange   AssociationType = "IMEI_CHANGE"
	AssociationTypeImeisvChange AssociationType = "IMEISV_CHANGE"
)

// EventReportMode indicates the event report mode.
type EventReportMode string

const (
	EventReportModePeriodic         EventReportMode = "PERIODIC"
	EventReportModeOnEventDetection EventReportMode = "ON_EVENT_DETECTION"
)

// ReachabilityForDataReportConfig indicates the report configuration for data reachability.
type ReachabilityForDataReportConfig string

const (
	ReachabilityForDataReportDirect   ReachabilityForDataReportConfig = "DIRECT_REPORT"
	ReachabilityForDataReportIndirect ReachabilityForDataReportConfig = "INDIRECT_REPORT"
)

// ReachabilityForSmsConfiguration indicates SMS reachability configuration.
type ReachabilityForSmsConfiguration string

const (
	ReachabilityForSmsOverNas ReachabilityForSmsConfiguration = "REACHABILITY_FOR_SMS_OVER_NAS"
	ReachabilityForSmsOverIp  ReachabilityForSmsConfiguration = "REACHABILITY_FOR_SMS_OVER_IP"
)

// RevokedCause indicates the revocation cause for a monitoring event.
type RevokedCause string

const (
	RevokedCauseNotAllowed       RevokedCause = "NOT_ALLOWED"
	RevokedCauseExcludedFromGroup RevokedCause = "EXCLUDED_FROM_GROUP"
	RevokedCauseGpsiRemoved      RevokedCause = "GPSI_REMOVED"
)

// FailedCause indicates the failure cause for a monitoring configuration.
type FailedCause string

const (
	FailedCauseAfNotAllowed                 FailedCause = "AF_NOT_ALLOWED"
	FailedCauseMtcProviderNotAllowed        FailedCause = "MTC_PROVIDER_NOT_ALLOWED"
	FailedCauseMonitoringNotAllowed         FailedCause = "MONITORING_NOT_ALLOWED"
	FailedCauseUnsupportedMonitoringEventType FailedCause = "UNSUPPORTED_MONITORING_EVENT_TYPE"
	FailedCauseUnsupportedMonitoringReportOptions FailedCause = "UNSUPPORTED_MONITORING_REPORT_OPTIONS"
	FailedCauseUnspecified                  FailedCause = "UNSPECIFIED"
)

// SubscriptionType indicates the type of EE subscription.
type SubscriptionType string

// PdnConnectivityStatus indicates PDN connectivity status.
type PdnConnectivityStatus string

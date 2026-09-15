package gppmodel

// ============================================================================
// TS 29.518 Namf_EventExposure — Notification Delivery & Subscription Request Structs
// Reference: 3GPP TS 29.518 V18.5.0, Namf_EventExposure Service
//
// This file contains spec-accurate Go structs for:
//   - Subscription request payloads (Consumer → AMF)
//   - Notification payloads (AMF → Consumer)
//
// Types use the "Namf" prefix to avoid conflicts with simplified types in event.go.
//
// Shared types reused from event.go: Tai, Ncgi, PlmnID, Snssai, UserLocation.
// ============================================================================

// ---------------------------------------------------------------------------
// Notification Payload Types
// ---------------------------------------------------------------------------

// NamfEventNotification represents the request body of the Event Notification
// callback (POST to eventNotifyUri). This is the top-level payload sent from
// AMF to the subscribed Consumer.
// Reference: TS 29.518 Section 6.2.6.2.4 (AmfEventNotification)
type NamfEventNotification struct {
	// NotifyCorrelationId is the notification correlation identifier.
	NotifyCorrelationId string `json:"notifyCorrelationId,omitempty"`

	// SubsChangeNotifyCorrelationId is the correlation ID for subscription
	// change notifications.
	SubsChangeNotifyCorrelationId string `json:"subsChangeNotifyCorrelationId,omitempty"`

	// ReportList contains one or more event reports in this notification.
	ReportList []NamfEventReport `json:"reportList,omitempty"`

	// EventSubsSyncInfo contains subscription synchronization information.
	EventSubsSyncInfo *NamfEventSubsSyncInfo `json:"eventSubsSyncInfo,omitempty"`
}

// NamfEventReport represents a report triggered by a subscribed event type.
// Reference: TS 29.518 Section 6.2.6.2.5 (AmfEventReport)
type NamfEventReport struct {
	// Type identifies the AMF event type. (Required)
	Type AmfEventType `json:"type"`

	// State represents the state of the subscribed event. (Required)
	State NamfEventState `json:"state"`

	// TimeStamp is the ISO 8601 timestamp when the event was detected. (Required)
	TimeStamp string `json:"timeStamp"`

	// SubscriptionId links back to the subscription that triggered this report.
	SubscriptionId string `json:"subscriptionId,omitempty"`

	// AnyUe indicates if the report applies to any UE.
	AnyUe *bool `json:"anyUe,omitempty"`

	// Supi is the Subscription Permanent Identifier of the affected UE.
	Supi string `json:"supi,omitempty"`

	// AreaList contains areas related to the event.
	AreaList []NamfEventArea `json:"areaList,omitempty"`

	// RefId is the reference identifier from the UDM subscription.
	RefId *int32 `json:"refId,omitempty"`

	// Gpsi is the Generic Public Subscription Identifier.
	Gpsi string `json:"gpsi,omitempty"`

	// Pei is the Permanent Equipment Identifier.
	Pei string `json:"pei,omitempty"`

	// Location is the current user location reported by the AMF.
	// Reuses UserLocation from event.go.
	Location *UserLocation `json:"location,omitempty"`

	// AdditionalLocation is an additional user location.
	AdditionalLocation *UserLocation `json:"additionalLocation,omitempty"`

	// Timezone is the UE's timezone (e.g., "+09:00").
	Timezone string `json:"timezone,omitempty"`

	// AccessTypeList lists access types related to the event.
	// Values: "3GPP_ACCESS", "NON_3GPP_ACCESS".
	AccessTypeList []AccessType `json:"accessTypeList,omitempty"`

	// RmInfoList contains registration management state information.
	RmInfoList []RmInfo `json:"rmInfoList,omitempty"`

	// CmInfoList contains connection management state information.
	CmInfoList []CmInfo `json:"cmInfoList,omitempty"`

	// Reachability indicates UE reachability status.
	Reachability UeReachability `json:"reachability,omitempty"`

	// CommFailure describes a communication failure detected by AMF.
	CommFailure *CommunicationFailure `json:"commFailure,omitempty"`

	// LossOfConnectReason describes the reason for loss of connectivity.
	LossOfConnectReason LossOfConnectivityReason `json:"lossOfConnectReason,omitempty"`

	// NumberOfUes is the number of UEs in area.
	NumberOfUes *int `json:"numberOfUes,omitempty"`

	// FiveGsUserStateList contains 5GS User State information per access type.
	FiveGsUserStateList []FiveGsUserStateInfo `json:"5gsUserStateList,omitempty"`

	// TypeCode is the IMEI TAC (Type Allocation Code), pattern: "imeitac-XXXXXXXX".
	TypeCode string `json:"typeCode,omitempty"`

	// RegistrationNumber is the frequent mobility registration count.
	RegistrationNumber *int `json:"registrationNumber,omitempty"`

	// MaxAvailabilityTime is the maximum availability time of the UE (ISO 8601).
	MaxAvailabilityTime string `json:"maxAvailabilityTime,omitempty"`

	// UeIdExt contains extended UE identity information.
	UeIdExt []UEIdExt `json:"ueIdExt,omitempty"`

	// SnssaiTaiList contains S-NSSAI to TAI mapping information.
	SnssaiTaiList []SnssaiTaiMapping `json:"snssaiTaiList,omitempty"`

	// IdleStatusIndication represents the idle status indication.
	IdleStatusIndication *IdleStatusIndication `json:"idleStatusIndication,omitempty"`

	// UeAccessBehaviorTrends contains UE access behavior trend reports.
	UeAccessBehaviorTrends []UeAccessBehaviorReportItem `json:"ueAccessBehaviorTrends,omitempty"`

	// UeLocationTrends contains UE location trend reports.
	UeLocationTrends []UeLocationTrendsReportItem `json:"ueLocationTrends,omitempty"`

	// MmTransLocationReportList contains MM transaction location reports.
	MmTransLocationReportList []MmTransactionLocationReportItem `json:"mmTransLocationReportList,omitempty"`

	// MmTransSliceReportList contains MM transaction slice reports.
	MmTransSliceReportList []MmTransactionSliceReportItem `json:"mmTransSliceReportList,omitempty"`

	// TermReason is the subscription termination reason.
	TermReason SubTerminationReason `json:"termReason,omitempty"`

	// UnavailabilityPeriod is the unavailability period in seconds.
	UnavailabilityPeriod *int `json:"unavailabilityPeriod,omitempty"`
}

// NamfEventState represents the state of a subscribed event.
// Reference: TS 29.518 (AmfEventState)
type NamfEventState struct {
	// Active indicates if the event subscription is active. (Required)
	Active bool `json:"active"`

	// RemainReports is the remaining number of reports to be generated.
	RemainReports *int `json:"remainReports,omitempty"`

	// RemainDuration is the remaining duration in seconds.
	RemainDuration *int `json:"remainDuration,omitempty"`
}

// NamfEventArea represents an area to be monitored by an AMF event.
// Reference: TS 29.518 (AmfEventArea)
type NamfEventArea struct {
	// PresenceInfo contains presence information for the area.
	PresenceInfo *PresenceInfo `json:"presenceInfo,omitempty"`

	// LadnInfo contains LADN information.
	LadnInfo *LadnInfo `json:"ladnInfo,omitempty"`

	// SNssai is the S-NSSAI associated with the area. Reuses Snssai from event.go.
	SNssai *Snssai `json:"sNssai,omitempty"`

	// NsiId is the Network Slice Instance Identifier.
	NsiId string `json:"nsiId,omitempty"`
}

// NamfEventSubsSyncInfo contains AMF Event Subscriptions synchronization information.
// Reference: TS 29.518 (AmfEventSubsSyncInfo)
type NamfEventSubsSyncInfo struct {
	// SubscriptionList contains the list of subscription information. (Required)
	SubscriptionList []NamfEventSubscriptionInfo `json:"subscriptionList"`
}

// NamfEventSubscriptionInfo contains individual AMF Event Subscription information.
// Reference: TS 29.518 (AmfEventSubscriptionInfo)
type NamfEventSubscriptionInfo struct {
	// SubId is the subscription identifier URI. (Required)
	SubId string `json:"subId"`

	// NotifyCorrelationId is the notification correlation identifier.
	NotifyCorrelationId string `json:"notifyCorrelationId,omitempty"`

	// RefIdList is the list of reference identifiers. (Required)
	RefIdList []int32 `json:"refIdList"`

	// OldSubId is the old subscription identifier URI.
	OldSubId string `json:"oldSubId,omitempty"`
}

// ---------------------------------------------------------------------------
// Supporting Types
// ---------------------------------------------------------------------------

// RmInfo represents the registration state of a UE for an access type.
// Reference: TS 29.518 (RmInfo)
type RmInfo struct {
	// RmState is the registration management state. (Required)
	RmState RmState `json:"rmState"`

	// AccessType is the access type. (Required)
	AccessType AccessType `json:"accessType"`
}

// CmInfo represents the connection management state of a UE for an access type.
// Reference: TS 29.518 (CmInfo)
type CmInfo struct {
	// CmState is the connection management state. (Required)
	CmState CmState `json:"cmState"`

	// AccessType is the access type. (Required)
	AccessType AccessType `json:"accessType"`
}

// LadnInfo represents LADN (Local Area Data Network) Information.
// Reference: TS 29.518 (LadnInfo)
type LadnInfo struct {
	// Ladn is the LADN DNN. (Required)
	Ladn string `json:"ladn"`

	// Presence is the presence state of the LADN.
	Presence string `json:"presence,omitempty"`
}

// FiveGsUserStateInfo represents the 5GS User state of the UE for an access type.
// Reference: TS 29.518 (5GsUserStateInfo)
type FiveGsUserStateInfo struct {
	// FiveGsUserState is the 5GS User State. (Required)
	FiveGsUserState FiveGsUserState `json:"5gsUserState"`

	// AccessType is the access type. (Required)
	AccessType AccessType `json:"accessType"`
}

// UEIdExt represents an extended UE Identity.
// Reference: TS 29.518 (UEIdExt)
type UEIdExt struct {
	// Supi is the Subscription Permanent Identifier.
	Supi string `json:"supi,omitempty"`

	// Gpsi is the Generic Public Subscription Identifier.
	Gpsi string `json:"gpsi,omitempty"`
}

// SnssaiTaiMapping represents restricted or unrestricted S-NSSAIs per TAI(s).
// Reference: TS 29.518 (SnssaiTaiMapping)
type SnssaiTaiMapping struct {
	// ReportingArea is the target area for the mapping. (Required)
	ReportingArea TargetArea `json:"reportingArea"`

	// AccessTypeList lists access types for this mapping.
	AccessTypeList []AccessType `json:"accessTypeList,omitempty"`

	// SupportedSnssaiList lists supported S-NSSAIs.
	SupportedSnssaiList []SupportedSnssai `json:"supportedSnssaiList,omitempty"`
}

// TargetArea represents a TA list, TAI range list, or any TA.
// Reference: TS 29.518 (TargetArea)
type TargetArea struct {
	// TaList is a list of Tracking Area Identities. Reuses Tai from event.go.
	TaList []Tai `json:"taList,omitempty"`

	// TaiRangeList is a list of TAI ranges.
	TaiRangeList []TaiRange `json:"taiRangeList,omitempty"`

	// AnyTa indicates any Tracking Area. Default: false.
	AnyTa *bool `json:"anyTa,omitempty"`
}

// SupportedSnssai represents a supported S-NSSAI with optional restriction.
// Reference: TS 29.518 (SupportedSnssai)
type SupportedSnssai struct {
	// SNssai is the S-NSSAI. (Required) Reuses Snssai from event.go.
	SNssai Snssai `json:"sNssai"`

	// RestrictionInd indicates whether the S-NSSAI is restricted. Default: false.
	RestrictionInd *bool `json:"restrictionInd,omitempty"`
}

// IdleStatusIndication represents the idle status indication.
// Reference: TS 29.518 (IdleStatusIndication)
type IdleStatusIndication struct {
	// TimeStamp is the time when the UE transitioned to idle (ISO 8601).
	TimeStamp string `json:"timeStamp,omitempty"`

	// ActiveTime is the active time in seconds.
	ActiveTime *int `json:"activeTime,omitempty"`

	// SubsRegTimer is the subscribed periodic registration timer in seconds.
	SubsRegTimer *int `json:"subsRegTimer,omitempty"`

	// EdrxCycleLength is the eDRX cycle length.
	EdrxCycleLength *int `json:"edrxCycleLength,omitempty"`

	// SuggestedNumOfDlPackets is the suggested number of downlink packets.
	SuggestedNumOfDlPackets *int `json:"suggestedNumOfDlPackets,omitempty"`
}

// UeAccessBehaviorReportItem is a report item for UE Access Behavior Trends event.
// Reference: TS 29.518 (UeAccessBehaviorReportItem)
type UeAccessBehaviorReportItem struct {
	// StateTransitionType is the access state transition type. (Required)
	StateTransitionType AccessStateTransitionType `json:"stateTransitionType"`

	// Spacing is the time spacing in seconds. (Required)
	Spacing int `json:"spacing"`

	// Duration is the duration in seconds. (Required)
	Duration int `json:"duration"`
}

// UeLocationTrendsReportItem is a report item for UE Location Trends event.
// Reference: TS 29.518 (UeLocationTrendsReportItem)
type UeLocationTrendsReportItem struct {
	// Tai is the Tracking Area Identity. Reuses Tai from event.go.
	Tai *Tai `json:"tai,omitempty"`

	// Ncgi is the NR Cell Global Identity. Reuses Ncgi from event.go.
	Ncgi *Ncgi `json:"ncgi,omitempty"`

	// Ecgi is the E-UTRA Cell Global Identity.
	Ecgi *Ecgi `json:"ecgi,omitempty"`

	// N3gaLocation is the Non-3GPP access location.
	N3gaLocation *N3gaLocation `json:"n3gaLocation,omitempty"`

	// Spacing is the time spacing in seconds. (Required)
	Spacing int `json:"spacing"`

	// Duration is the duration in seconds. (Required)
	Duration int `json:"duration"`

	// Timestamp is the event timestamp (ISO 8601). (Required)
	Timestamp string `json:"timestamp"`
}

// MmTransactionLocationReportItem is a UE MM Transaction Report Item per Location.
// Reference: TS 29.518 (MmTransactionLocationReportItem)
type MmTransactionLocationReportItem struct {
	// Tai is the Tracking Area Identity.
	Tai *Tai `json:"tai,omitempty"`

	// Ncgi is the NR Cell Global Identity.
	Ncgi *Ncgi `json:"ncgi,omitempty"`

	// Ecgi is the E-UTRA Cell Global Identity.
	Ecgi *Ecgi `json:"ecgi,omitempty"`

	// N3gaLocation is the Non-3GPP access location.
	N3gaLocation *N3gaLocation `json:"n3gaLocation,omitempty"`

	// Timestamp is the event timestamp (ISO 8601). (Required)
	Timestamp string `json:"timestamp"`

	// Transactions is the number of MM transactions. (Required)
	Transactions int `json:"transactions"`
}

// MmTransactionSliceReportItem is a UE MM Transaction Report Item per Slice.
// Reference: TS 29.518 (MmTransactionSliceReportItem)
type MmTransactionSliceReportItem struct {
	// Snssai is the S-NSSAI. Reuses Snssai from event.go.
	Snssai *Snssai `json:"snssai,omitempty"`

	// Timestamp is the event timestamp (ISO 8601). (Required)
	Timestamp string `json:"timestamp"`

	// Transactions is the number of MM transactions. (Required)
	Transactions int `json:"transactions"`
}

// ---------------------------------------------------------------------------
// Enum String Types
// ---------------------------------------------------------------------------

// AmfEventType describes the supported event types of Namf_EventExposure Service.
type AmfEventType string

const (
	AmfEventTypeLocationReport               AmfEventType = "LOCATION_REPORT"
	AmfEventTypePresenceInAoiReport          AmfEventType = "PRESENCE_IN_AOI_REPORT"
	AmfEventTypeTimezoneReport               AmfEventType = "TIMEZONE_REPORT"
	AmfEventTypeAccessTypeReport             AmfEventType = "ACCESS_TYPE_REPORT"
	AmfEventTypeRegistrationStateReport      AmfEventType = "REGISTRATION_STATE_REPORT"
	AmfEventTypeConnectivityStateReport      AmfEventType = "CONNECTIVITY_STATE_REPORT"
	AmfEventTypeReachabilityReport           AmfEventType = "REACHABILITY_REPORT"
	AmfEventTypeCommunicationFailureReport   AmfEventType = "COMMUNICATION_FAILURE_REPORT"
	AmfEventTypeUesInAreaReport              AmfEventType = "UES_IN_AREA_REPORT"
	AmfEventTypeSubscriptionIdChange         AmfEventType = "SUBSCRIPTION_ID_CHANGE"
	AmfEventTypeSubscriptionIdAddition       AmfEventType = "SUBSCRIPTION_ID_ADDITION"
	AmfEventTypeSubscriptionTermination      AmfEventType = "SUBSCRIPTION_TERMINATION"
	AmfEventTypeLossOfConnectivity           AmfEventType = "LOSS_OF_CONNECTIVITY"
	AmfEventTypeFiveGsUserStateReport        AmfEventType = "5GS_USER_STATE_REPORT"
	AmfEventTypeAvailabilityAfterDdnFailure  AmfEventType = "AVAILABILITY_AFTER_DDN_FAILURE"
	AmfEventTypeTypeAllocationCodeReport     AmfEventType = "TYPE_ALLOCATION_CODE_REPORT"
	AmfEventTypeFrequentMobilityRegistration AmfEventType = "FREQUENT_MOBILITY_REGISTRATION_REPORT"
	AmfEventTypeSnssaiTaMappingReport        AmfEventType = "SNSSAI_TA_MAPPING_REPORT"
	AmfEventTypeUeLocationTrends             AmfEventType = "UE_LOCATION_TRENDS"
	AmfEventTypeUeAccessBehaviorTrends       AmfEventType = "UE_ACCESS_BEHAVIOR_TRENDS"
	AmfEventTypeUeMmTransactionReport        AmfEventType = "UE_MM_TRANSACTION_REPORT"
)

// UeReachability describes the reachability of the UE.
type UeReachability string

const (
	UeReachabilityUnreachable    UeReachability = "UNREACHABLE"
	UeReachabilityReachable      UeReachability = "REACHABLE"
	UeReachabilityRegulatoryOnly UeReachability = "REGULATORY_ONLY"
)

// RmState describes the registration management state of a UE.
type RmState string

const (
	RmStateRegistered   RmState = "REGISTERED"
	RmStateDeregistered RmState = "DEREGISTERED"
)

// CmState describes the connection management state of a UE.
type CmState string

const (
	CmStateIdle      CmState = "IDLE"
	CmStateConnected CmState = "CONNECTED"
)

// FiveGsUserState describes the 5GS User State of a UE.
type FiveGsUserState string

const (
	FiveGsUserStateDeregistered                   FiveGsUserState = "DEREGISTERED"
	FiveGsUserStateConnectedNotReachableForPaging FiveGsUserState = "CONNECTED_NOT_REACHABLE_FOR_PAGING"
	FiveGsUserStateConnectedReachableForPaging    FiveGsUserState = "CONNECTED_REACHABLE_FOR_PAGING"
	FiveGsUserStateNotProvidedFromAmf             FiveGsUserState = "NOT_PROVIDED_FROM_AMF"
)

// LossOfConnectivityReason describes the reason for loss of connectivity.
type LossOfConnectivityReason string

const (
	LossOfConnectivityReasonDeregistered            LossOfConnectivityReason = "DEREGISTERED"
	LossOfConnectivityReasonMaxDetectionTimeExpired LossOfConnectivityReason = "MAX_DETECTION_TIME_EXPIRED"
	LossOfConnectivityReasonPurged                  LossOfConnectivityReason = "PURGED"
	LossOfConnectivityReasonUnavailablePeriod       LossOfConnectivityReason = "UNAVAILABLE_PERIOD"
)

// AccessStateTransitionType describes the Access State Transition Type.
type AccessStateTransitionType string

const (
	AccessStateTransitionAccessTypeChange3GPP       AccessStateTransitionType = "ACCESS_TYPE_CHANGE_3GPP"
	AccessStateTransitionAccessTypeChangeN3GPP      AccessStateTransitionType = "ACCESS_TYPE_CHANGE_N3GPP"
	AccessStateTransitionRmStateChangeDeregistered  AccessStateTransitionType = "RM_STATE_CHANGE_DEREGISTERED"
	AccessStateTransitionRmStateChangeRegistered    AccessStateTransitionType = "RM_STATE_CHANGE_REGISTERED"
	AccessStateTransitionCmStateChangeIdle          AccessStateTransitionType = "CM_STATE_CHANGE_IDLE"
	AccessStateTransitionCmStateChangeConnected     AccessStateTransitionType = "CM_STATE_CHANGE_CONNECTED"
	AccessStateTransitionHandover                   AccessStateTransitionType = "HANDOVER"
	AccessStateTransitionMobilityRegistrationUpdate AccessStateTransitionType = "MOBILITY_REGISTRATION_UPDATE"
)

// SubTerminationReason describes the Subscription Termination Reason.
type SubTerminationReason string

const (
	SubTerminationReasonInvalidSubscription       SubTerminationReason = "INVALID_SUBSCRIPTION"
	SubTerminationReasonSubscriptionNotAuthorized SubTerminationReason = "SUBSCRIPTION_NOT_AUTHORIZED"
)

// ============================================================================
// Subscription Request Types (Consumer → AMF)
// ============================================================================

// AmfCreateEventSubscription is the request body for creating an AMF event
// subscription (POST /subscriptions).
// Reference: TS 29.518 Section 6.2.6.2.2 (AmfCreateEventSubscription)
type AmfCreateEventSubscription struct {
	// Subscription contains the event subscription details. (Required)
	Subscription AmfEventSubscription `json:"subscription"`

	// SupportedFeatures lists features supported by the consumer.
	SupportedFeatures SupportedFeatures `json:"supportedFeatures,omitempty"`

	// OldGuami is the old GUAMI for AMF relocation scenarios.
	OldGuami *Guami `json:"oldGuami,omitempty"`
}

// AmfCreatedEventSubscription is the response body for a successful AMF event
// subscription creation.
// Reference: TS 29.518 Section 6.2.6.2.3 (AmfCreatedEventSubscription)
type AmfCreatedEventSubscription struct {
	// Subscription contains the event subscription details. (Required)
	Subscription AmfEventSubscription `json:"subscription"`

	// SubscriptionId is the URI of the created subscription resource. (Required)
	SubscriptionId Uri `json:"subscriptionId"`

	// ReportList contains initial event reports if immediate reporting was requested.
	ReportList []NamfEventReport `json:"reportList,omitempty"`

	// SupportedFeatures lists features supported by the AMF.
	SupportedFeatures SupportedFeatures `json:"supportedFeatures,omitempty"`
}

// AmfUpdateEventSubscriptionItem describes a single modification to an AMF Event
// Subscription via JSON Patch (RFC 6902).
// Reference: TS 29.518 Section 6.2.6.2.6 (AmfUpdateEventSubscriptionItem)
type AmfUpdateEventSubscriptionItem struct {
	// Op is the JSON Patch operation: "add", "remove", or "replace". (Required)
	Op string `json:"op"`

	// Path is the JSON Pointer path to the property to modify. (Required)
	Path string `json:"path"`

	// Value is the new AmfEvent value (for "add" or "replace" ops).
	Value *AmfEvent `json:"value,omitempty"`

	// PresenceInfo is the PresenceInfo value for path targeting presenceInfoList.
	PresenceInfo *PresenceInfo `json:"presenceInfo,omitempty"`

	// ExcludeSupiList is for path "/excludeSupiList".
	ExcludeSupiList []Supi `json:"excludeSupiList,omitempty"`

	// ExcludeGpsiList is for path "/excludeGpsiList".
	ExcludeGpsiList []Gpsi `json:"excludeGpsiList,omitempty"`

	// IncludeSupiList is for path "/includeSupiList".
	IncludeSupiList []Supi `json:"includeSupiList,omitempty"`

	// IncludeGpsiList is for path "/includeGpsiList".
	IncludeGpsiList []Gpsi `json:"includeGpsiList,omitempty"`

	// NotifyForSupiList is for path targeting notifyForSupiList.
	NotifyForSupiList []Supi `json:"notifyForSupiList,omitempty"`

	// NotifyForGroupList is for path targeting notifyForGroupList.
	NotifyForGroupList []GroupId `json:"notifyForGroupList,omitempty"`

	// NotifyForSnssaiDnnList is for path targeting notifyForSnssaiDnnList.
	NotifyForSnssaiDnnList []SnssaiDnnItem `json:"notifyForSnssaiDnnList,omitempty"`
}

// AmfUpdatedEventSubscription represents a successful update on an AMF Event Subscription.
// Reference: TS 29.518 Section 6.2.6.2.8 (AmfUpdatedEventSubscription)
type AmfUpdatedEventSubscription struct {
	// Subscription contains the updated event subscription details. (Required)
	Subscription AmfEventSubscription `json:"subscription"`

	// ReportList contains event reports if triggered by the update.
	ReportList []NamfEventReport `json:"reportList,omitempty"`
}

// AmfEventSubscription represents an individual event subscription resource on AMF.
// Reference: TS 29.518 Section 6.2.6.2.4 (AmfEventSubscription)
type AmfEventSubscription struct {
	// EventList contains the list of events to subscribe to. (Required)
	EventList []AmfEvent `json:"eventList"`

	// EventNotifyUri is the URI where AMF sends event notifications. (Required)
	EventNotifyUri Uri `json:"eventNotifyUri"`

	// NotifyCorrelationId is the notification correlation identifier. (Required)
	NotifyCorrelationId string `json:"notifyCorrelationId"`

	// NfId is the NF Instance ID of the consumer. (Required)
	NfId NfInstanceId `json:"nfId"`

	// SubsChangeNotifyUri is the URI for subscription change notifications.
	SubsChangeNotifyUri Uri `json:"subsChangeNotifyUri,omitempty"`

	// SubsChangeNotifyCorrelationId is the correlation ID for subscription changes.
	SubsChangeNotifyCorrelationId string `json:"subsChangeNotifyCorrelationId,omitempty"`

	// Supi is the Subscription Permanent Identifier of the targeted UE.
	Supi Supi `json:"supi,omitempty"`

	// GroupId is the group identifier for group-based subscriptions.
	GroupId GroupId `json:"groupId,omitempty"`

	// ExcludeSupiList lists SUPIs to exclude.
	ExcludeSupiList []Supi `json:"excludeSupiList,omitempty"`

	// ExcludeGpsiList lists GPSIs to exclude.
	ExcludeGpsiList []Gpsi `json:"excludeGpsiList,omitempty"`

	// IncludeSupiList lists SUPIs to include.
	IncludeSupiList []Supi `json:"includeSupiList,omitempty"`

	// IncludeGpsiList lists GPSIs to include.
	IncludeGpsiList []Gpsi `json:"includeGpsiList,omitempty"`

	// Gpsi is the Generic Public Subscription Identifier.
	Gpsi Gpsi `json:"gpsi,omitempty"`

	// Pei is the Permanent Equipment Identifier.
	Pei Pei `json:"pei,omitempty"`

	// AnyUE indicates the subscription applies to any UE.
	AnyUE bool `json:"anyUE,omitempty"`

	// Options specifies the event reporting mode.
	Options *AmfEventMode `json:"options,omitempty"`

	// SourceNfType is the type of the NF that is the source of the event.
	SourceNfType NFType `json:"sourceNfType,omitempty"`

	// TermNotifyInd indicates whether termination notification is required.
	TermNotifyInd bool `json:"termNotifyInd,omitempty"`
}

// AmfEvent describes an event to be subscribed.
// Reference: TS 29.518 Section 6.2.6.2.5 (AmfEvent)
type AmfEvent struct {
	// Type identifies the AMF event type. (Required)
	Type AmfEventType `json:"type"`

	// ImmediateFlag requests immediate reporting.
	ImmediateFlag bool `json:"immediateFlag,omitempty"`

	// AreaList contains the areas to monitor.
	AreaList []NamfEventArea `json:"areaList,omitempty"`

	// LocationFilterList contains location-based filters.
	LocationFilterList []LocationFilter `json:"locationFilterList,omitempty"`

	// RefId is a reference ID from UDM subscription.
	RefId int32 `json:"refId,omitempty"`

	// TrafficDescriptorList contains traffic descriptor filters.
	TrafficDescriptorList []TrafficDescriptor `json:"trafficDescriptorList,omitempty"`

	// ReportUeReachable requests reporting of UE reachability.
	ReportUeReachable bool `json:"reportUeReachable,omitempty"`

	// ReachabilityFilter specifies the reachability filter.
	ReachabilityFilter *ReachabilityFilter `json:"reachabilityFilter,omitempty"`

	// UdmDetectInd indicates UDM detection.
	UdmDetectInd bool `json:"udmDetectInd,omitempty"`

	// MaxReports is the maximum number of reports.
	MaxReports int `json:"maxReports,omitempty"`

	// PresenceInfoList maps praId to PresenceInfo.
	PresenceInfoList map[string]PresenceInfo `json:"presenceInfoList,omitempty"`

	// MaxResponseTime is the maximum response time in seconds.
	MaxResponseTime DurationSec `json:"maxResponseTime,omitempty"`

	// TargetArea specifies the target area.
	TargetArea *TargetArea `json:"targetArea,omitempty"`

	// SnssaiFilter contains S-NSSAI based filters.
	SnssaiFilter []ExtSnssai `json:"snssaiFilter,omitempty"`

	// UeInAreaFilter specifies UE in-area filtering.
	UeInAreaFilter *UeInAreaFilter `json:"ueInAreaFilter,omitempty"`

	// MinInterval is the minimum reporting interval in seconds.
	MinInterval DurationSec `json:"minInterval,omitempty"`

	// NextReport is the timestamp for the next report.
	NextReport *DateTime `json:"nextReport,omitempty"`

	// IdleStatusInd indicates idle status reporting.
	IdleStatusInd bool `json:"idleStatusInd,omitempty"`
}

// AmfEventMode describes how reports shall be generated by a subscribed event.
// Reference: TS 29.518 Section 6.2.6.2.6 (AmfEventMode)
type AmfEventMode struct {
	// Trigger specifies the event trigger type. (Required)
	Trigger AmfEventTrigger `json:"trigger"`

	// MaxReports is the maximum number of reports.
	MaxReports int `json:"maxReports,omitempty"`

	// Expiry is the subscription expiry time.
	Expiry *DateTime `json:"expiry,omitempty"`

	// RepPeriod is the repetition period in seconds for periodic reporting.
	RepPeriod DurationSec `json:"repPeriod,omitempty"`

	// SampRatio is the sampling ratio.
	SampRatio SamplingRatio `json:"sampRatio,omitempty"`

	// PartitioningCriteria lists criteria for partitioning.
	PartitioningCriteria []PartitioningCriteria `json:"partitioningCriteria,omitempty"`

	// NotifFlag is the notification flag.
	NotifFlag NotificationFlag `json:"notifFlag,omitempty"`

	// MutingExcInstructions specifies muting exception instructions.
	MutingExcInstructions *MutingExceptionInstructions `json:"mutingExcInstructions,omitempty"`

	// MutingNotSettings specifies muting notification settings (read-only).
	MutingNotSettings *MutingNotificationsSettings `json:"mutingNotSettings,omitempty"`

	// VarRepPeriodInfo contains variable reporting period information.
	VarRepPeriodInfo []VarRepPeriod `json:"varRepPeriodInfo,omitempty"`
}

// ============================================================================
// Subscription Request Supporting Types
// ============================================================================

// AmfEventTrigger describes how the event notification should be triggered.
// Reference: TS 29.518 Section 6.2.6.2.7 (AmfEventTrigger)
type AmfEventTrigger string

const (
	AmfEventTriggerOneTime    AmfEventTrigger = "ONE_TIME"
	AmfEventTriggerContinuous AmfEventTrigger = "CONTINUOUS"
	AmfEventTriggerPeriodic   AmfEventTrigger = "PERIODIC"
)

// LocationFilter contains location-based filtering criteria for events.
// Reference: TS 29.518 (LocationFilter)
type LocationFilter struct {
	// Tai is the Tracking Area Identity.
	Tai *Tai `json:"tai,omitempty"`

	// Ncgi is the NR Cell Global Identity.
	Ncgi *Ncgi `json:"ncgi,omitempty"`

	// Ecgi is the E-UTRA Cell Global Identity.
	Ecgi *Ecgi `json:"ecgi,omitempty"`

	// N3gaLocation is the Non-3GPP access location.
	N3gaLocation *N3gaLocation `json:"n3gaLocation,omitempty"`
}

// TrafficDescriptor describes a traffic descriptor for event filtering.
// Reference: TS 29.518 (TrafficDescriptor)
type TrafficDescriptor struct {
	// Dnn is the Data Network Name.
	Dnn Dnn `json:"dnn,omitempty"`

	// SNssai is the S-NSSAI.
	SNssai *Snssai `json:"sNssai,omitempty"`
}

// ReachabilityFilter specifies the reachability filter type.
// Reference: TS 29.518 (ReachabilityFilter)
type ReachabilityFilter string

const (
	ReachabilityFilterUeReachableInd      ReachabilityFilter = "UE_REACHABLE_IND"
	ReachabilityFilterMaxAvailabilityTime ReachabilityFilter = "MAX_AVAILABILITY_TIME"
)

// UeInAreaFilter specifies UE in-area filtering criteria.
// Reference: TS 29.518 (UeInAreaFilter)
type UeInAreaFilter struct {
	// UeInAreaType is the type of in-area filter.
	UeInAreaType UeInAreaType `json:"ueInAreaType,omitempty"`

	// PresenceInfo contains presence information.
	PresenceInfo *PresenceInfo `json:"presenceInfo,omitempty"`
}

// UeInAreaType describes how UE presence in area is determined.
type UeInAreaType string

const (
	UeInAreaTypeTrue  UeInAreaType = "TRUE"
	UeInAreaTypeFalse UeInAreaType = "FALSE"
)

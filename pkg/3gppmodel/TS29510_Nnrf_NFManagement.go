// Package model provides primitives for NRF NFManagement Service.
//
// Reference: 3GPP TS 29.510 V18.6.0, Nnrf_NFManagement Service
//
// This file contains spec-accurate Go structs for the NRF NF Management
// subscription types — used to receive notifications when NF instances
// register, deregister, or change their profiles.
package gppmodel

// ============================================================================
// Subscription Types
// ============================================================================

// NrfSubscriptionData represents the request/response body for NRF NF status
// subscriptions (POST /subscriptions).
// Reference: TS 29.510 Section 6.2.6.2.3 (SubscriptionData)
type NrfSubscriptionData struct {
	// NfStatusNotificationUri is the callback URI where NRF sends notifications. (Required)
	NfStatusNotificationUri Uri `json:"nfStatusNotificationUri"`

	// ReqNfInstanceId is the NF instance ID of the subscribing NF.
	ReqNfInstanceId NfInstanceId `json:"reqNfInstanceId,omitempty"`

	// SharedDataIds lists shared data IDs to monitor.
	SharedDataIds []string `json:"sharedDataIds,omitempty"`

	// SubscrCond is the condition that determines which NFs to monitor.
	// One of: NfInstanceIdCond, NfInstanceIdListCond, NfTypeCond, ServiceNameCond, etc.
	SubscrCond *NrfSubscrCond `json:"subscrCond,omitempty"`

	// SubscriptionId is the subscription identifier (set by NRF in response, read-only).
	SubscriptionId string `json:"subscriptionId,omitempty"`

	// ValidityTime is the subscription expiration time (ISO 8601).
	ValidityTime DateTime `json:"validityTime,omitempty"`

	// ReqNotifEvents lists the notification event types to subscribe to.
	// Values: NF_REGISTERED, NF_DEREGISTERED, NF_PROFILE_CHANGED, SHARED_DATA_CHANGED.
	ReqNotifEvents []NotificationEventType `json:"reqNotifEvents,omitempty"`

	// PlmnId is the PLMN ID of the subscribing NF.
	PlmnId PlmnId `json:"plmnId,omitempty"`

	// Nid is the NID of the subscribing NF (for SNPN).
	Nid Nid `json:"nid,omitempty"`

	// NotifCondition filters which changes trigger notifications.
	NotifCondition *NotifCondition `json:"notifCondition,omitempty"`

	// ReqNfType is the NF type of the requester.
	ReqNfType *NFType `json:"reqNfType,omitempty"`

	// ReqNfFqdn is the FQDN of the requester NF.
	ReqNfFqdn Fqdn `json:"reqNfFqdn,omitempty"`

	// ReqSnssais lists S-NSSAIs of the requester.
	ReqSnssais []ExtSnssai `json:"reqSnssais,omitempty"`

	// ReqPlmnList lists PLMNs of the requester.
	ReqPlmnList []PlmnId `json:"reqPlmnList,omitempty"`

	// ReqSnpnList lists SNPNs of the requester.
	ReqSnpnList []PlmnIdNid `json:"reqSnpnList,omitempty"`

	// ServingScope lists the serving scope of the requester.
	ServingScope []string `json:"servingScope,omitempty"`

	// RequesterFeatures lists features supported by the requester (write-only).
	RequesterFeatures SupportedFeatures `json:"requesterFeatures,omitempty"`

	// NrfSupportedFeatures lists features supported by the NRF (read-only).
	NrfSupportedFeatures SupportedFeatures `json:"nrfSupportedFeatures,omitempty"`

	// HnrfUri is the URI of the home NRF.
	HnrfUri Uri `json:"hnrfUri,omitempty"`

	// OnboardingCapability indicates onboarding capability.
	OnboardingCapability bool `json:"onboardingCapability,omitempty"`

	// TargetHni is the target home network identifier.
	TargetHni Fqdn `json:"targetHni,omitempty"`

	// PreferredLocality filters NFs by locality.
	PreferredLocality string `json:"preferredLocality,omitempty"`

	// CompleteProfileSubscription requests complete NF profiles in notifications.
	CompleteProfileSubscription bool `json:"completeProfileSubscription,omitempty"`
}

// ============================================================================
// Subscription Condition Types
// ============================================================================

// NrfSubscrCond is the condition determining which NFs to monitor.
// It is a oneOf: the first non-nil field indicates the condition type.
// Reference: TS 29.510 Section 6.2.6.2.4 (SubscrCond)
type NrfSubscrCond struct {
	// NfInstanceIdCond monitors a specific NF instance.
	NfInstanceIdCond *NfInstanceIdCond `json:"nfInstanceIdCond,omitempty"`

	// NfTypeCond monitors all NFs of a given type.
	NfTypeCond *NfTypeCond `json:"nfTypeCond,omitempty"`

	// ServiceNameCond monitors NFs supporting a given service name.
	ServiceNameCond *ServiceNameCond `json:"serviceNameCond,omitempty"`

	// AmfCond monitors AMFs by AMF Set ID / Region ID.
	AmfCond *AmfCond `json:"amfCond,omitempty"`

	// NfSetCond monitors NFs in a given NF Set.
	NfSetCond *NfSetCond `json:"nfSetCond,omitempty"`
}

// NfInstanceIdCond subscribes to notifications for a specific NF Instance ID.
// Reference: TS 29.510 (NfInstanceIdCond)
type NfInstanceIdCond struct {
	NfInstanceId NfInstanceId `json:"nfInstanceId"`
}

// NfTypeCond subscribes to notifications for all NFs of a given NF type.
// Reference: TS 29.510 (NfTypeCond)
type NfTypeCond struct {
	NfType NFType `json:"nfType"`
}

// ServiceNameCond subscribes to notifications for NFs supporting a given service name.
// Reference: TS 29.510 (ServiceNameCond)
type ServiceNameCond struct {
	ServiceName ServiceName `json:"serviceName"`
}

// AmfCond subscribes to notifications for AMFs based on Set ID / Region ID.
// Reference: TS 29.510 (AmfCond)
type AmfCond struct {
	AmfSetId    AmfSetId    `json:"amfSetId,omitempty"`
	AmfRegionId AmfRegionId `json:"amfRegionId,omitempty"`
}

// NfSetCond subscribes to notifications for NFs in a given NF Set.
// Reference: TS 29.510 (NfSetCond)
type NfSetCond struct {
	NfSetId NfSetId `json:"nfSetId"`
}

// NotifCondition filters which NF profile attribute changes trigger notifications.
// Reference: TS 29.510 (NotifCondition)
type NotifCondition struct {
	// MonitoredAttributes lists attributes whose change triggers a notification.
	MonitoredAttributes []string `json:"monitoredAttributes,omitempty"`

	// UnmonitoredAttributes lists attributes whose change does NOT trigger a notification.
	UnmonitoredAttributes []string `json:"unmonitoredAttributes,omitempty"`
}

// ============================================================================
// Notification Types (NRF → Consumer)
// ============================================================================

// NrfNotificationData is the notification payload sent by NRF to the subscribed
// NF when an NF status event occurs.
// Reference: TS 29.510 Section 6.2.6.2.5 (NotificationData)
type NrfNotificationData struct {
	// Event is the type of notification event. (Required)
	Event NotificationEventType `json:"event"`

	// NfInstanceUri is the URI of the NF instance that triggered the event. (Required)
	NfInstanceUri Uri `json:"nfInstanceUri"`

	// NfProfile contains the NF profile (for NF_REGISTERED or NF_PROFILE_CHANGED).
	NfProfile *NFProfile `json:"nfProfile,omitempty"`

	// ProfileChanges lists individual profile changes (for NF_PROFILE_CHANGED).
	ProfileChanges []ChangeItem `json:"profileChanges,omitempty"`

	// CompleteNfProfile contains the complete NF profile.
	CompleteNfProfile *NFProfile `json:"completeNfProfile,omitempty"`

	// SharedDataChanges lists shared data changes (for SHARED_DATA_CHANGED).
	SharedDataChanges []ChangeItem `json:"sharedDataChanges,omitempty"`

	// ConditionEvent indicates whether the NF was added or removed from the condition.
	ConditionEvent ConditionEventType `json:"conditionEvent,omitempty"`

	// SubscriptionContext contains subscription context information.
	SubscriptionContext *SubscriptionContext `json:"subscriptionContext,omitempty"`
}

// SubscriptionContext contains context data related to a created subscription.
// Reference: TS 29.510 (SubscriptionContext)
type SubscriptionContext struct {
	// SubscriptionId is the subscription identifier. (Required)
	SubscriptionId string `json:"subscriptionId"`

	// SubscrCond is the subscription condition that triggered the notification.
	SubscrCond *NrfSubscrCond `json:"subscrCond,omitempty"`
}

// ============================================================================
// Enums
// ============================================================================

// NotificationEventType describes the types of events sent in NRF notifications.
// Reference: TS 29.510 (NotificationEventType)
type NotificationEventType string

const (
	NotificationEventNfRegistered      NotificationEventType = "NF_REGISTERED"
	NotificationEventNfDeregistered    NotificationEventType = "NF_DEREGISTERED"
	NotificationEventNfProfileChanged  NotificationEventType = "NF_PROFILE_CHANGED"
	NotificationEventSharedDataChanged NotificationEventType = "SHARED_DATA_CHANGED"
)

// ConditionEventType indicates whether an NF was added or removed from a
// subscription condition.
// Reference: TS 29.510 (ConditionEventType)
type ConditionEventType string

const (
	ConditionEventNfAdded   ConditionEventType = "NF_ADDED"
	ConditionEventNfRemoved ConditionEventType = "NF_REMOVED"
)

// ============================================================================
// ChangeItem — reused from CommonData; alias for clarity
// ============================================================================

// ChangeItem is already defined in TS29571_CommonData.go.
// It represents an individual change in an NF profile.

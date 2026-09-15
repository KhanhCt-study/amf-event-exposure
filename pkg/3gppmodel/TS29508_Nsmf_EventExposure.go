package gppmodel

// ============================================================================
// TS 29.508 Nsmf_EventExposure — Notification Delivery Structs
// Reference: 3GPP TS 29.508 V18.5.0, Session Management Event Exposure Service
//
// This file contains spec-accurate Go structs for the SMF Event Exposure
// notification payloads:
//   - NsmfEventExposureNotification: event notification batch from SMF to Consumer
//   - AckOfNotify: consumer acknowledgement of a notification
//
// Shared types reused from:
//   - event.go: PlmnID, Snssai
//   - TS29518_Namf_EventExposure.go: CommunicationFailure, AccessType
// ============================================================================

// ---------------------------------------------------------------------------
// Notification Payload Types
// ---------------------------------------------------------------------------

// NsmfEventExposureNotification represents notifications on events that occurred.
// This is the top-level payload sent from SMF to the subscribed Consumer via
// HTTP POST to notifUri.
// Reference: TS 29.508 (NsmfEventExposureNotification)
type NsmfEventExposureNotification struct {
	// NotifId is the notification correlation ID. (Required)
	NotifId string `json:"notifId"`

	// EventNotifs contains one or more individual event notifications. (Required)
	EventNotifs []NsmfEventNotification `json:"eventNotifs"`

	// AckUri is the URI for the consumer to send an acknowledgement.
	AckUri string `json:"ackUri,omitempty"`
}

// NsmfEventNotification represents a notification related to a single event.
// Reference: TS 29.508 (EventNotification)
type NsmfEventNotification struct {
	// Event identifies the SMF event type. (Required)
	Event SmfEvent `json:"event"`

	// TimeStamp is the ISO 8601 timestamp when the event was detected. (Required)
	TimeStamp string `json:"timeStamp"`

	// Supi is the Subscription Permanent Identifier.
	Supi string `json:"supi,omitempty"`

	// Gpsi is the Generic Public Subscription Identifier.
	Gpsi string `json:"gpsi,omitempty"`

	// UeIpAddr is the IP address allocated to the UE.
	UeIpAddr *IpAddr `json:"ueIpAddr,omitempty"`

	// TransacInfos contains SMF transaction information.
	TransacInfos []TransactionInfo `json:"transacInfos,omitempty"`

	// SourceDnai is the source DNAI (before path change).
	SourceDnai string `json:"sourceDnai,omitempty"`

	// TargetDnai is the target DNAI (after path change).
	TargetDnai string `json:"targetDnai,omitempty"`

	// DnaiChgType indicates the DNAI change type.
	// Values: "EARLY", "EARLY_LATE", "LATE".
	DnaiChgType string `json:"dnaiChgType,omitempty"`

	// CandidateDnais lists candidate DNAI(s) for the PDU Session.
	CandidateDnais []string `json:"candidateDnais,omitempty"`

	// CandDnaisPrioInd indicates if candidate DNAIs are in priority order.
	CandDnaisPrioInd *bool `json:"candDnaisPrioInd,omitempty"`

	// EasRediscoverInd indicates EAS re-discovery was performed.
	EasRediscoverInd *bool `json:"easRediscoverInd,omitempty"`

	// SourceUeIpv4Addr is the source UE IPv4 address (before DNAI change).
	SourceUeIpv4Addr string `json:"sourceUeIpv4Addr,omitempty"`

	// SourceUeIpv6Prefix is the source UE IPv6 prefix.
	SourceUeIpv6Prefix string `json:"sourceUeIpv6Prefix,omitempty"`

	// TargetUeIpv4Addr is the target UE IPv4 address (after DNAI change).
	TargetUeIpv4Addr string `json:"targetUeIpv4Addr,omitempty"`

	// TargetUeIpv6Prefix is the target UE IPv6 prefix.
	TargetUeIpv6Prefix string `json:"targetUeIpv6Prefix,omitempty"`

	// UeMac is the UE MAC address.
	UeMac string `json:"ueMac,omitempty"`

	// AdIpv4Addr is the Application Detection IPv4 address.
	AdIpv4Addr string `json:"adIpv4Addr,omitempty"`

	// AdIpv6Prefix is the Application Detection IPv6 prefix.
	AdIpv6Prefix string `json:"adIpv6Prefix,omitempty"`

	// ReIpv4Addr is the Redundant IPv4 address.
	ReIpv4Addr string `json:"reIpv4Addr,omitempty"`

	// ReIpv6Prefix is the Redundant IPv6 prefix.
	ReIpv6Prefix string `json:"reIpv6Prefix,omitempty"`

	// PlmnId is the PLMN identity. Reuses PlmnID from event.go.
	PlmnId *PlmnID `json:"plmnId,omitempty"`

	// AccType is the access type.
	AccType AccessType `json:"accType,omitempty"`

	// PduAccTypes lists PDU session access types.
	PduAccTypes []AccessType `json:"pduAccTypes,omitempty"`

	// PduSeId is the PDU Session ID.
	PduSeId *int `json:"pduSeId,omitempty"`

	// RatType is the RAT type.
	RatType string `json:"ratType,omitempty"`

	// DddStatus is the Downlink Data Delivery status.
	DddStatus string `json:"dddStatus,omitempty"`

	// MaxWaitTime is the maximum waiting time (ISO 8601).
	MaxWaitTime string `json:"maxWaitTime,omitempty"`

	// CommFailure describes a communication failure.
	// Reuses CommunicationFailure from TS29518_Namf_EventExposure.go.
	CommFailure *CommunicationFailure `json:"commFailure,omitempty"`

	// Ipv4Addr is the UE's IPv4 address.
	Ipv4Addr string `json:"ipv4Addr,omitempty"`

	// Ipv6Prefixes is a list of UE's IPv6 prefixes.
	Ipv6Prefixes []string `json:"ipv6Prefixes,omitempty"`

	// Ipv6Addrs is a list of UE's IPv6 addresses.
	Ipv6Addrs []string `json:"ipv6Addrs,omitempty"`

	// PduSessType is the PDU Session type.
	// Values: "IPV4", "IPV6", "IPV4V6", "UNSTRUCTURED", "ETHERNET".
	PduSessType string `json:"pduSessType,omitempty"`

	// SscMode is the SSC mode.
	SscMode string `json:"sscMode,omitempty"`

	// Qfi is the QoS Flow Identifier (0-63).
	Qfi *int `json:"qfi,omitempty"`

	// AppId is the Application Identifier.
	AppId string `json:"appId,omitempty"`

	// Dnn is the Data Network Name.
	Dnn string `json:"dnn,omitempty"`

	// Snssai is the S-NSSAI. Reuses Snssai from event.go.
	Snssai *Snssai `json:"snssai,omitempty"`

	// UlDelays contains uplink packet delay measurements (ms).
	UlDelays []int `json:"ulDelays,omitempty"`

	// DlDelays contains downlink packet delay measurements (ms).
	DlDelays []int `json:"dlDelays,omitempty"`

	// RtDelays contains round-trip packet delay measurements (ms).
	RtDelays []int `json:"rtDelays,omitempty"`

	// UlCongInfo is the uplink congestion information.
	UlCongInfo *int `json:"ulCongInfo,omitempty"`

	// DlCongInfo is the downlink congestion information.
	DlCongInfo *int `json:"dlCongInfo,omitempty"`

	// UlDataRate is the uplink data rate (bps).
	UlDataRate string `json:"ulDataRate,omitempty"`

	// DlDataRate is the downlink data rate (bps).
	DlDataRate string `json:"dlDataRate,omitempty"`

	// SmNasFromUe contains SM NAS messages received from UE.
	SmNasFromUe *SmNasFromUe `json:"smNasFromUe,omitempty"`

	// SmNasFromSmf contains SM NAS messages sent by SMF to UE.
	SmNasFromSmf *SmNasFromSmf `json:"smNasFromSmf,omitempty"`

	// UpRedTrans indicates whether redundant transmission is setup (true) or terminated (false).
	UpRedTrans *bool `json:"upRedTrans,omitempty"`

	// SsId is the SSID that the PDU session is related to.
	SsId string `json:"ssId,omitempty"`

	// BssId is the BSSID that the PDU session is related to.
	BssId string `json:"bssId,omitempty"`

	// StartWlan is the WLAN connection start time (ISO 8601).
	StartWlan string `json:"startWlan,omitempty"`

	// EndWlan is the WLAN connection end time (ISO 8601).
	EndWlan string `json:"endWlan,omitempty"`

	// PduSessInfos contains PDU session related information.
	PduSessInfos []PduSessionInformation `json:"pduSessInfos,omitempty"`

	// UpfInfo contains UPF information.
	UpfInfo *UpfInformation `json:"upfInfo,omitempty"`

	// Pdmf indicates packet delay measurement failure.
	Pdmf *bool `json:"pdmf,omitempty"`

	// SatBackhaulCat is the Satellite Backhaul Category.
	SatBackhaulCat string `json:"satBackhaulCat,omitempty"`

	// SupportedFeatures is the list of supported features.
	SupportedFeatures string `json:"supportedFeatures,omitempty"`

	// FiveQi is the 5G QoS Identifier (0-255).
	FiveQi *int `json:"5qi,omitempty"`
}

// NsmfAckOfNotify represents an acknowledgement of an event notification.
// Sent by the Consumer back to the SMF via the ackUri callback.
// Reference: TS 29.508 (AckOfNotify)
type NsmfAckOfNotify struct {
	// NotifId is the notification correlation ID. (Required)
	NotifId string `json:"notifId"`

	// AckResult contains the application-layer handling result. (Required)
	AckResult AfResultInfo `json:"ackResult"`

	// Supi is the Subscription Permanent Identifier.
	Supi string `json:"supi,omitempty"`

	// Gpsi is the Generic Public Subscription Identifier.
	Gpsi string `json:"gpsi,omitempty"`
}

// ---------------------------------------------------------------------------
// Supporting Types
// ---------------------------------------------------------------------------

// TransactionInfo represents SMF Transaction Information.
// Reference: TS 29.508 (TransactionInfo)
type TransactionInfo struct {
	// Transaction is the transaction count. (Required)
	Transaction int `json:"transaction"`

	// Snssai is the S-NSSAI. Reuses Snssai from event.go.
	Snssai *Snssai `json:"snssai,omitempty"`

	// AppIds contains application identifiers.
	AppIds []string `json:"appIds,omitempty"`

	// TransacMetrics indicates SM transaction metrics.
	TransacMetrics []TransactionMetric `json:"transacMetrics,omitempty"`
}

// SmNasFromUe represents SM NAS messages received by SMF from UE.
// Reference: TS 29.508 (SmNasFromUe)
type SmNasFromUe struct {
	// SmNasType is the SM NAS message type. (Required)
	SmNasType string `json:"smNasType"`

	// TimeStamp is the timestamp (ISO 8601). (Required)
	TimeStamp string `json:"timeStamp"`
}

// SmNasFromSmf represents SM NAS messages sent by SMF to UE with
// congestion control information.
// Reference: TS 29.508 (SmNasFromSmf)
type SmNasFromSmf struct {
	// SmNasType is the SM NAS message type. (Required)
	SmNasType string `json:"smNasType"`

	// TimeStamp is the timestamp (ISO 8601). (Required)
	TimeStamp string `json:"timeStamp"`

	// BackoffTimer is the back-off timer duration in seconds. (Required)
	BackoffTimer int `json:"backoffTimer"`

	// AppliedSmccType is the applied SM congestion control type. (Required)
	// Values: "DNN_CC", "SNSSAI_CC".
	AppliedSmccType AppliedSmccType `json:"appliedSmccType"`
}

// PduSessionInformation represents PDU session related information.
// Reference: TS 29.508 (PduSessionInformation)
type PduSessionInformation struct {
	// PduSessId is the PDU Session ID.
	PduSessId *int `json:"pduSessId,omitempty"`

	// SessInfo contains session information.
	SessInfo *NsmfPduSessionInfo `json:"sessInfo,omitempty"`
}

// NsmfPduSessionInfo represents session information details.
// Reference: TS 29.508 (PduSessionInfo)
type NsmfPduSessionInfo struct {
	// N4SessId is the identifier of the N4 session.
	N4SessId string `json:"n4SessId,omitempty"`

	// SessInactiveTimer is the session inactivity timer in seconds.
	SessInactiveTimer *int `json:"sessInactiveTimer,omitempty"`

	// PduSessStatus is the PDU Session status.
	// Values: "ACTIVATED", "DEACTIVATED".
	PduSessStatus PduSessionStatus `json:"pduSessStatus,omitempty"`
}

// UpfInformation represents the ID/address/FQDN of the UPF.
// Reference: TS 29.508 (UpfInformation)
type UpfInformation struct {
	// UpfId is the UPF identifier.
	UpfId string `json:"upfId,omitempty"`

	// UpfAddr contains the UPF address or FQDN.
	UpfAddr *AddrFqdn `json:"upfAddr,omitempty"`
}

// AfResultInfo identifies the result of application layer handling.
// Reference: TS 29.522 TrafficInfluence (AfResultInfo)
type AfResultInfo struct {
	// AfStatus is the application-layer result status. (Required)
	// Values: "SUCCESS", "TEMPORARY_CONGESTION", "RELOC_NO_ALLOWED", "OTHER".
	AfStatus AfResultStatus `json:"afStatus"`

	// UpBuffInd indicates if uplink traffic buffering is needed.
	UpBuffInd *bool `json:"upBuffInd,omitempty"`
}

// ---------------------------------------------------------------------------
// Enum String Types
// ---------------------------------------------------------------------------

// SmfEvent represents the types of events that can be subscribed on SMF.
type SmfEvent string

const (
	SmfEventAcTyCh             SmfEvent = "AC_TY_CH"
	SmfEventUpPathCh           SmfEvent = "UP_PATH_CH"
	SmfEventPduSesRel          SmfEvent = "PDU_SES_REL"
	SmfEventPlmnCh             SmfEvent = "PLMN_CH"
	SmfEventUeIpCh             SmfEvent = "UE_IP_CH"
	SmfEventRatTyCh            SmfEvent = "RAT_TY_CH"
	SmfEventDdds               SmfEvent = "DDDS"
	SmfEventCommFail           SmfEvent = "COMM_FAIL"
	SmfEventPduSesEst          SmfEvent = "PDU_SES_EST"
	SmfEventQfiAlloc           SmfEvent = "QFI_ALLOC"
	SmfEventQosMon             SmfEvent = "QOS_MON"
	SmfEventSmccExp            SmfEvent = "SMCC_EXP"
	SmfEventDispersion         SmfEvent = "DISPERSION"
	SmfEventRedTransExp        SmfEvent = "RED_TRANS_EXP"
	SmfEventWlanInfo           SmfEvent = "WLAN_INFO"
	SmfEventUpfInfo            SmfEvent = "UPF_INFO"
	SmfEventUpStatusInfo       SmfEvent = "UP_STATUS_INFO"
	SmfEventSatbCh             SmfEvent = "SATB_CH"
	SmfEventTrafficCorrelation SmfEvent = "TRAFFIC_CORRELATION"
)

// AppliedSmccType indicates the type of applied SM congestion control.
type AppliedSmccType string

const (
	AppliedSmccTypeDnnCC    AppliedSmccType = "DNN_CC"
	AppliedSmccTypeSnssaiCC AppliedSmccType = "SNSSAI_CC"
)

// TransactionMetric represents the metric on UE SM transactions.
type TransactionMetric string

const (
	TransactionMetricPduSesEst   TransactionMetric = "PDU_SES_EST"
	TransactionMetricPduSesAuth  TransactionMetric = "PDU_SES_AUTH"
	TransactionMetricPduSesModif TransactionMetric = "PDU_SES_MODIF"
	TransactionMetricPduSesRel   TransactionMetric = "PDU_SES_REL"
)

// PduSessionStatus represents the status of the PDU Session.
type PduSessionStatus string

const (
	PduSessionStatusActivated   PduSessionStatus = "ACTIVATED"
	PduSessionStatusDeactivated PduSessionStatus = "DEACTIVATED"
)

// AfResultStatus represents the status of application handling result.
type AfResultStatus string

const (
	AfResultStatusSuccess             AfResultStatus = "SUCCESS"
	AfResultStatusTemporaryCongestion AfResultStatus = "TEMPORARY_CONGESTION"
	AfResultStatusRelocNoAllowed      AfResultStatus = "RELOC_NO_ALLOWED"
	AfResultStatusOther               AfResultStatus = "OTHER"
)

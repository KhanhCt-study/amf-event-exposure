package gppmodel

// ============================================================================
// TS 29.503 Nudm_SDM / Nudm_UECM — Notification Delivery Structs
// Reference: 3GPP TS 29.503 V18.5.0, Unified Data Management Services
//
// This file contains spec-accurate Go structs for UDM notification payloads:
//   - ModificationNotification (Nudm_SDM): subscriber data change notification
//   - DataRestorationNotification (Nudm_UECM): UDR data-loss recovery notification
//
// Shared types reused from event.go: PlmnID, Snssai.
// ============================================================================

// ---------------------------------------------------------------------------
// ModificationNotification (Nudm_SDM)
// ---------------------------------------------------------------------------

// NudmModificationNotification represents the notification payload sent by UDM
// to subscribed NFs when subscriber data is modified.
// Reference: TS 29.503 Nudm_SDM (ModificationNotification)
type NudmModificationNotification struct {
	// NotifyItems contains the list of changed data items. (Required)
	NotifyItems []NotifyItem `json:"notifyItems"`

	// SubscriptionId is the subscription identifier.
	SubscriptionId string `json:"subscriptionId,omitempty"`
}

// ---------------------------------------------------------------------------
// DataRestorationNotification (Nudm_UECM)
// ---------------------------------------------------------------------------

// NudmDataRestorationNotification represents the notification payload sent by
// UDM to registered NFs when a data-loss event occurs at the UDR, identifying
// potentially affected UEs.
// Reference: TS 29.503 Nudm_UECM (DataRestorationNotification)
type NudmDataRestorationNotification struct {
	// LastReplicationTime is the timestamp of the last successful replication
	// before the data-loss event (ISO 8601).
	LastReplicationTime string `json:"lastReplicationTime,omitempty"`

	// RecoveryTime is the timestamp when the data was recovered (ISO 8601).
	RecoveryTime string `json:"recoveryTime,omitempty"`

	// PlmnId identifies the PLMN whose data was affected.
	PlmnId *PlmnID `json:"plmnId,omitempty"`

	// SupiRanges contains ranges of affected SUPIs.
	SupiRanges []SupiRange `json:"supiRanges,omitempty"`

	// GpsiRanges contains ranges of affected GPSIs.
	GpsiRanges []IdentityRange `json:"gpsiRanges,omitempty"`

	// ResetIds is a list of reset identifiers.
	ResetIds []string `json:"resetIds,omitempty"`

	// SNssaiList contains affected S-NSSAIs. Reuses Snssai from event.go.
	SNssaiList []Snssai `json:"sNssaiList,omitempty"`

	// DnnList contains affected DNNs.
	DnnList []string `json:"dnnList,omitempty"`

	// UdmGroupId is the UDM group identifier.
	UdmGroupId string `json:"udmGroupId,omitempty"`
}

// ============================================================================
// TS 29.503 Nudm_UECM — Registration Data Types
// ============================================================================

// Amf3GppAccessRegistration contains the set of information relevant to the
// AMF where the UE has registered via 3GPP access.
// Reference: TS 29.503 Section 6.2.6.2.3 (Amf3GppAccessRegistration)
type Amf3GppAccessRegistration struct {
	// AmfInstanceId is the NF Instance ID of the serving AMF. (Required)
	AmfInstanceId NfInstanceId `json:"amfInstanceId"`

	// SupportedFeatures lists features supported by the AMF.
	SupportedFeatures SupportedFeatures `json:"supportedFeatures,omitempty"`

	// PurgeFlag indicates if the UE is to be purged.
	PurgeFlag *PurgeFlag `json:"purgeFlag,omitempty"`

	// Pei is the Permanent Equipment Identifier.
	Pei Pei `json:"pei,omitempty"`

	// ImsVoPs indicates the IMS Voice over PS session support.
	ImsVoPs *ImsVoPs `json:"imsVoPs,omitempty"`

	// DeregCallbackUri is the URI for deregistration callback. (Required)
	DeregCallbackUri Uri `json:"deregCallbackUri"`

	// AmfServiceNameDereg is the service name for deregistration.
	AmfServiceNameDereg ServiceName `json:"amfServiceNameDereg,omitempty"`

	// PcscfRestorationCallbackUri is the URI for P-CSCF restoration callback.
	PcscfRestorationCallbackUri Uri `json:"pcscfRestorationCallbackUri,omitempty"`

	// AmfServiceNamePcscfRest is the service name for P-CSCF restoration.
	AmfServiceNamePcscfRest ServiceName `json:"amfServiceNamePcscfRest,omitempty"`

	// InitialRegistrationInd indicates if this is the initial registration.
	InitialRegistrationInd bool `json:"initialRegistrationInd,omitempty"`

	// EmergencyRegistrationInd indicates if this is an emergency registration.
	EmergencyRegistrationInd bool `json:"emergencyRegistrationInd,omitempty"`

	// Guami is the Globally Unique AMF Identifier. (Required)
	Guami Guami `json:"guami"`

	// BackupAmfInfo lists backup AMF information.
	BackupAmfInfo []BackupAmfInfo `json:"backupAmfInfo,omitempty"`

	// DrFlag indicates dual registration support.
	DrFlag *DualRegistrationFlag `json:"drFlag,omitempty"`

	// RatType is the Radio Access Technology type. (Required)
	RatType RatType `json:"ratType"`

	// UrrpIndicator indicates URRP-AMF support.
	UrrpIndicator bool `json:"urrpIndicator,omitempty"`

	// AmfEeSubscriptionId is the URI for AMF Event Exposure subscription.
	AmfEeSubscriptionId Uri `json:"amfEeSubscriptionId,omitempty"`

	// EpsInterworkingInfo contains EPS interworking information.
	EpsInterworkingInfo *EpsInterworkingInfo `json:"epsInterworkingInfo,omitempty"`

	// UeSrvccCapability indicates UE SRVCC capability.
	UeSrvccCapability bool `json:"ueSrvccCapability,omitempty"`

	// RegistrationTime is the timestamp of registration.
	RegistrationTime *DateTime `json:"registrationTime,omitempty"`

	// VgmlcAddress contains the VGMLC address.
	VgmlcAddress *VgmlcAddress `json:"vgmlcAddress,omitempty"`

	// ContextInfo contains additional context information.
	ContextInfo interface{} `json:"contextInfo,omitempty"`

	// NoEeSubscriptionInd indicates no EE subscription.
	NoEeSubscriptionInd bool `json:"noEeSubscriptionInd,omitempty"`

	// Supi is the Subscription Permanent Identifier.
	Supi Supi `json:"supi,omitempty"`

	// UeReachableInd indicates UE reachability.
	UeReachableInd *UeReachableInd `json:"ueReachableInd,omitempty"`

	// ReRegistrationRequired indicates if re-registration is required.
	ReRegistrationRequired bool `json:"reRegistrationRequired,omitempty"`

	// AdminDeregSubWithdrawn indicates admin deregistration withdrawn.
	AdminDeregSubWithdrawn bool `json:"adminDeregSubWithdrawn,omitempty"`

	// DataRestorationCallbackUri is the URI for data restoration callback.
	DataRestorationCallbackUri Uri `json:"dataRestorationCallbackUri,omitempty"`

	// ResetIds lists reset identifiers.
	ResetIds []string `json:"resetIds,omitempty"`

	// DisasterRoamingInd indicates disaster roaming support.
	DisasterRoamingInd bool `json:"disasterRoamingInd,omitempty"`

	// UeMINTCapability indicates UE MINT capability.
	UeMINTCapability bool `json:"ueMINTCapability,omitempty"`

	// SorSnpnSiSupported indicates SOR-SNPN-SI support.
	SorSnpnSiSupported bool `json:"sorSnpnSiSupported,omitempty"`

	// UdrRestartInd indicates UDR restart.
	UdrRestartInd bool `json:"udrRestartInd,omitempty"`

	// LastSynchronizationTime is the last synchronization timestamp.
	LastSynchronizationTime *DateTime `json:"lastSynchronizationTime,omitempty"`
}

// RegistrationDataSets wraps the response from GetRegistrations (Nudm_UECM).
// Reference: TS 29.503 (RegistrationDataSets)
type RegistrationDataSets struct {
	// Amf3Gpp contains the AMF 3GPP access registration.
	Amf3Gpp *Amf3GppAccessRegistration `json:"amf3Gpp,omitempty"`

	// AmfNon3Gpp contains the AMF non-3GPP access registration.
	// Note: per TS 29.503, the non-3GPP type (AmfNon3GppAccessRegistration)
	// differs slightly from 3GPP (imsVoPs is required). For pragmatic
	// reasons, both are stored as Amf3GppAccessRegistration pointers.
	AmfNon3Gpp *Amf3GppAccessRegistration `json:"amfNon3Gpp,omitempty"`

	// SmfRegistration contains the SMF registration information.
	SmfRegistration *SmfRegistrationInfo `json:"smfRegistration,omitempty"`

	// Smsf3Gpp contains the SMSF 3GPP registration.
	Smsf3Gpp *SmsfRegistration `json:"smsf3Gpp,omitempty"`

	// SmsfNon3Gpp contains the SMSF non-3GPP registration.
	SmsfNon3Gpp *SmsfRegistration `json:"smsfNon3Gpp,omitempty"`

	// IpSmGw contains the IP-SM-GW registration.
	IpSmGw *IpSmGwRegistration `json:"ipSmGw,omitempty"`

	// NwdafRegistration contains the NWDAF registration information.
	NwdafRegistration *NwdafRegistrationInfo `json:"nwdafRegistration,omitempty"`
}

// ============================================================================
// UDM UECM Supporting Types
// ============================================================================

// PurgeFlag indicates whether the UE is to be purged from the AMF.
type PurgeFlag string

const (
	PurgeFlagPurgeAllowed    PurgeFlag = "PURGE_ALLOWED"
	PurgeFlagPurgeNotAllowed PurgeFlag = "PURGE_NOT_ALLOWED"
)

// ImsVoPs indicates the IMS Voice over PS session support.
type ImsVoPs string

const (
	ImsVoPsHOMOLOGOUS    ImsVoPs = "HOMOLOGOUS"
	ImsVoPsNONHOMOLOGOUS ImsVoPs = "NON_HOMOLOGOUS"
)

// DualRegistrationFlag indicates whether dual registration is supported.
type DualRegistrationFlag string

const (
	DualRegistrationFlagYes DualRegistrationFlag = "YES"
	DualRegistrationFlagNo  DualRegistrationFlag = "NO"
)

// EpsInterworkingInfo contains EPS interworking information.
type EpsInterworkingInfo struct {
	// EpsIwkInd indicates EPS interworking support.
	EpsIwkInd bool `json:"epsIwkInd,omitempty"`
}

// VgmlcAddress contains the VGMLC address information.
type VgmlcAddress struct {
	// VgmlcAddress is the VGMLC address.
	VgmlcAddress string `json:"vgmlcAddress,omitempty"`

	// VgmlcNumber is the VGMLC number.
	VgmlcNumber string `json:"vgmlcNumber,omitempty"`
}

// UeReachableInd indicates whether the UE is reachable.
type UeReachableInd string

const (
	UeReachableIndReachable    UeReachableInd = "REACHABLE"
	UeReachableIndNotReachable UeReachableInd = "NOT_REACHABLE"
	UeReachableIndUnknown      UeReachableInd = "UNKNOWN"
)

// SmfRegistrationInfo contains the list of SMF registrations for a UE.
// Reference: TS 29.503 Nudm_UECM (SmfRegistrationInfo)
type SmfRegistrationInfo struct {
	// SmfRegistrationList contains the list of SMF registrations. (Required)
	SmfRegistrationList []SmfRegistration `json:"smfRegistrationList"`
}

// SmfRegistration contains SMF registration information for a PDU session.
// Reference: TS 29.503 Nudm_UECM (SmfRegistration)
type SmfRegistration struct {
	// SmfInstanceId is the NF Instance ID of the SMF. (Required)
	SmfInstanceId NfInstanceId `json:"smfInstanceId"`

	// SmfSetId is the SMF Set ID.
	SmfSetId NfSetId `json:"smfSetId,omitempty"`

	// SupportedFeatures lists supported features.
	SupportedFeatures SupportedFeatures `json:"supportedFeatures,omitempty"`

	// PduSessionId is the PDU Session ID. (Required)
	PduSessionId PduSessionId `json:"pduSessionId"`

	// SingleNssai is the S-NSSAI of the PDU session. (Required)
	SingleNssai Snssai `json:"singleNssai"`

	// Dnn is the Data Network Name.
	Dnn Dnn `json:"dnn,omitempty"`

	// PlmnId is the PLMN ID of the SMF. (Required)
	PlmnId PlmnId `json:"plmnId"`

	// RegistrationTime is the timestamp of registration.
	RegistrationTime *DateTime `json:"registrationTime,omitempty"`

	// ContextInfo contains context information.
	ContextInfo interface{} `json:"contextInfo,omitempty"`
}

// SmsfRegistration contains SMSF registration information.
// Reference: TS 29.503 Nudm_UECM (SmsfRegistration)
type SmsfRegistration struct {
	// SmsfInstanceId is the NF Instance ID of the SMSF. (Required)
	SmsfInstanceId NfInstanceId `json:"smsfInstanceId"`

	// SmsfSetId is the SMSF Set ID.
	SmsfSetId NfSetId `json:"smsfSetId,omitempty"`

	// PlmnId is the PLMN ID. (Required)
	PlmnId PlmnId `json:"plmnId"`

	// RegistrationTime is the timestamp of registration.
	RegistrationTime *DateTime `json:"registrationTime,omitempty"`
}

// IpSmGwRegistration contains IP-SM-GW registration information.
// Reference: TS 29.503 Nudm_UECM (IpSmGwRegistration)
type IpSmGwRegistration struct {
	// NfInstanceId is the NF Instance ID of the IP-SM-GW.
	NfInstanceId NfInstanceId `json:"nfInstanceId,omitempty"`

	// IpsmgwIpv4 is the IPv4 address of the IP-SM-GW.
	IpsmgwIpv4 Ipv4Addr `json:"ipsmgwIpv4,omitempty"`

	// IpsmgwIpv6 is the IPv6 address of the IP-SM-GW.
	IpsmgwIpv6 Ipv6Addr `json:"ipsmgwIpv6,omitempty"`

	// IpsmgwFqdn is the FQDN of the IP-SM-GW.
	IpsmgwFqdn Fqdn `json:"ipsmgwFqdn,omitempty"`
}

// NwdafRegistrationInfo contains the list of NWDAF registrations for a UE.
// Reference: TS 29.503 Nudm_UECM (NwdafRegistrationInfo)
type NwdafRegistrationInfo struct {
	// NwdafRegistrationList contains the list of NWDAF registrations. (Required)
	NwdafRegistrationList []NwdafRegistration `json:"nwdafRegistrationList"`
}

// NwdafRegistration contains NWDAF registration information for a UE.
// Reference: TS 29.503 Nudm_UECM (NwdafRegistration)
type NwdafRegistration struct {
	// NwdafInstanceId is the NF Instance ID of the NWDAF. (Required)
	NwdafInstanceId NfInstanceId `json:"nwdafInstanceId"`

	// AnalyticsIds lists the Analytics IDs provided by this NWDAF. (Required)
	// TODO: Replace interface{} with proper AnalyticsId type from TS29520_Nnwdaf_AnalyticsInfo
	AnalyticsIds []interface{} `json:"analyticsIds"`

	// NwdafSetId is the NWDAF Set ID.
	NwdafSetId NfSetId `json:"nwdafSetId,omitempty"`

	// RegistrationTime is the timestamp of registration.
	RegistrationTime *DateTime `json:"registrationTime,omitempty"`

	// ContextInfo contains context information.
	ContextInfo interface{} `json:"contextInfo,omitempty"`
}

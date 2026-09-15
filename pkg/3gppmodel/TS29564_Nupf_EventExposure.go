package gppmodel

// ============================================================================
// TS 29.564 Nupf_EventExposure — Notification Delivery Structs
// Reference: 3GPP TS 29.564 V18.5.0, UPF Event Exposure Service
//
// This file contains spec-accurate Go structs for the UPF Event Exposure
// notification payload (NF Source → Consumer).
//
// Shared types reused from event.go: Snssai
// ============================================================================

// ---------------------------------------------------------------------------
// Notification Payload Types
// ---------------------------------------------------------------------------

// NupfNotificationData represents the list of NotificationItems.
// Reference: TS 29.564 Section 6.1.6.2.2 (NotificationData)
type NupfNotificationData struct {
	// NotificationItems contains one or more event reports. (Required)
	NotificationItems []NupfNotificationItem `json:"notificationItems"`

	// CorrelationId is the correlation identifier.
	CorrelationId string `json:"correlationId,omitempty"`

	// AchievedSampRatio represents the achieved sampling ratio.
	AchievedSampRatio *int `json:"achievedSampRatio,omitempty"`
}

// NupfNotificationItem represents a report on one subscribed event.
// Reference: TS 29.564 (NotificationItem)
type NupfNotificationItem struct {
	// EventType identifies the UPF event type. (Required)
	EventType NupfEventType `json:"eventType"`

	// UeIpv4Addr is the UE IPv4 address.
	UeIpv4Addr string `json:"ueIpv4Addr,omitempty"`

	// UeIpv6Prefix is the UE IPv6 prefix.
	UeIpv6Prefix string `json:"ueIpv6Prefix,omitempty"`

	// UeMacAddr is the UE MAC address.
	UeMacAddr string `json:"ueMacAddr,omitempty"`

	// Dnn is the Data Network Name.
	Dnn string `json:"dnn,omitempty"`

	// Snssai is the S-NSSAI. Reuses Snssai from event.go.
	Snssai *Snssai `json:"snssai,omitempty"`

	// Gpsi is the Generic Public Subscription Identifier.
	Gpsi string `json:"gpsi,omitempty"`

	// Supi is the Subscription Permanent Identifier.
	Supi string `json:"supi,omitempty"`

	// TimeStamp is the ISO 8601 timestamp when the event was detected. (Required)
	TimeStamp string `json:"timeStamp"`

	// StartTime is the ISO 8601 timestamp for the start of the measurement.
	StartTime string `json:"startTime,omitempty"`

	// QosMonitoringMeasurement contains QoS monitoring data.
	QosMonitoringMeasurement *QosMonitoringMeasurement `json:"qosMonitoringMeasurement,omitempty"`

	// TscMngtInfo contains TSC management info.
	TscMngtInfo *TscManagementInfo `json:"tscMngtInfo,omitempty"`

	// UserDataUsageMeasurements contains user data usage measurements.
	UserDataUsageMeasurements []UserDataUsageMeasurements `json:"userDataUsageMeasurements,omitempty"`
}

// ---------------------------------------------------------------------------
// Supporting Types
// ---------------------------------------------------------------------------

// QosMonitoringMeasurement contains QoS Monitoring Measurement information.
// Reference: TS 29.564 (QosMonitoringMeasurement)
type QosMonitoringMeasurement struct {
	DlPacketDelay     *int   `json:"dlPacketDelay,omitempty"`
	UlPacketDelay     *int   `json:"ulPacketDelay,omitempty"`
	RtrPacketDelay    *int   `json:"rtrPacketDelay,omitempty"`
	MeasureFailure    *bool  `json:"measureFailure,omitempty"`
	DlAveThroughput   string `json:"dlAveThroughput,omitempty"`
	UlAveThroughput   string `json:"ulAveThroughput,omitempty"`
	DlCongestion      *int   `json:"dlCongestion,omitempty"` // 0-10000
	UlCongestion      *int   `json:"ulCongestion,omitempty"` // 0-10000
	DefaultQosFlowInd *bool  `json:"defaultQosFlowInd,omitempty"`
}

// TscManagementInfo contains TSC Management Information.
// Reference: TS 29.564 (TscManagementInfo)
type TscManagementInfo struct {
	// Pmics represents Port Management Container (from TS 29.512)
	Pmics []string `json:"pmics,omitempty"`
	// Umic represents Bridge Management Container (from TS 29.512)
	Umic string `json:"umic,omitempty"`
}

// UserDataUsageMeasurements contains data usage measurements.
// Reference: TS 29.564 (UserDataUsageMeasurements)
type UserDataUsageMeasurements struct {
	AppId                           string                           `json:"appId,omitempty"`
	FlowInfo                        interface{}                      `json:"flowInfo,omitempty"` // From TS 29.512 FlowInformation
	VolumeMeasurement               *VolumeMeasurement               `json:"volumeMeasurement,omitempty"`
	ThroughputMeasurement           *ThroughputMeasurement           `json:"throughputMeasurement,omitempty"`
	ApplicationRelatedInformation   *ApplicationRelatedInformation   `json:"applicationRelatedInformation,omitempty"`
	ThroughputStatisticsMeasurement *ThroughputStatisticsMeasurement `json:"throughputStatisticsMeasurement,omitempty"`
}

// VolumeMeasurement contains Volume Measurement information.
// Reference: TS 29.564 (VolumeMeasurement)
type VolumeMeasurement struct {
	TotalVolume      *int64  `json:"totalVolume,omitempty"`
	UlVolume         *int64  `json:"ulVolume,omitempty"`
	DlVolume         *int64  `json:"dlVolume,omitempty"`
	TotalNbOfPackets *uint64 `json:"totalNbOfPackets,omitempty"`
	UlNbOfPackets    *uint64 `json:"ulNbOfPackets,omitempty"`
	DlNbOfPackets    *uint64 `json:"dlNbOfPackets,omitempty"`
}

// ThroughputMeasurement contains Throughput Measurement information.
// Reference: TS 29.564 (ThroughputMeasurement)
type ThroughputMeasurement struct {
	UlThroughput       string `json:"ulThroughput,omitempty"`
	DlThroughput       string `json:"dlThroughput,omitempty"`
	UlPacketThroughput string `json:"ulPacketThroughput,omitempty"`
	DlPacketThroughput string `json:"dlPacketThroughput,omitempty"`
}

// ApplicationRelatedInformation contains Application Related Information.
// Reference: TS 29.564 (ApplicationRelatedInformation)
type ApplicationRelatedInformation struct {
	Urls           []string            `json:"urls,omitempty"`
	DomainInfoList []DomainInformation `json:"domainInfoList,omitempty"`
}

// ThroughputStatisticsMeasurement contains Throughput Statistics Measurement.
// Reference: TS 29.564 (ThroughputStatisticsMeasurement)
type ThroughputStatisticsMeasurement struct {
	UlAverageThroughput       string `json:"ulAverageThroughput,omitempty"`
	DlAverageThroughput       string `json:"dlAverageThroughput,omitempty"`
	UlPeakThroughput          string `json:"ulPeakThroughput,omitempty"`
	DlPeakThroughPut          string `json:"dlPeakThroughPut,omitempty"`
	UlAveragePacketThroughput string `json:"ulAveragePacketThroughput,omitempty"`
	DlAveragePacketThroughput string `json:"dlAveragePacketThroughput,omitempty"`
	UlPeakPacketThroughput    string `json:"ulPeakPacketThroughput,omitempty"`
	DlPeakPacketThroughput    string `json:"dlPeakPacketThroughput,omitempty"`
}

// DomainInformation contains Domain Information.
// Reference: TS 29.564 (DomainInformation)
type DomainInformation struct {
	DomainName         string     `json:"domainName"` // Required
	DomainNameProtocol DnProtocol `json:"domainNameProtocol,omitempty"`
}

// ---------------------------------------------------------------------------
// Enum String Types
// ---------------------------------------------------------------------------

// NupfEventType represents the UPF Event Type.
type NupfEventType string

const (
	NupfEventTypeQosMonitoring         NupfEventType = "QOS_MONITORING"
	NupfEventTypeUserDataUsageMeasures NupfEventType = "USER_DATA_USAGE_MEASURES"
	NupfEventTypeUserDataUsageTrends   NupfEventType = "USER_DATA_USAGE_TRENDS"
	NupfEventTypeTscMngtInfo           NupfEventType = "TSC_MNGT_INFO"
)

// DnProtocol represents the Domain Name Protocol.
type DnProtocol string

const (
	DnProtocolDnsQname DnProtocol = "DNS_QNAME"
	DnProtocolTlsSni   DnProtocol = "TLS_SNI"
	DnProtocolTlsSan   DnProtocol = "TLS_SAN"
	DnProtocolTlsScn   DnProtocol = "TLS_SCN"
)

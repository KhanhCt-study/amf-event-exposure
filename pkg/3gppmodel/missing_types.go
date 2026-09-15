// Package model provides stubs for types referenced by generated code but not yet
// generated from their respective YAML specs. These will be replaced when the
// corresponding oapi-codegen generation is run against the full spec set.
package gppmodel

import "encoding/json"

// ReportingInformation Represents the reporting information for event subscriptions.
// TODO: Generate from TS29523_Npcf_EventExposure.yaml or equivalent spec.
type ReportingInformation struct {
	// ImmRep Indicates whether immediate reporting is required.
	ImmRep *bool `json:"immRep,omitempty"`

	// NotifMethod Indicates the notification method (PERIODIC, ONE_TIME, ON_EVENT_DETECTION).
	NotifMethod *NotificationMethod `json:"notifMethod,omitempty"`

	// MaxReportNbr Maximum number of reports.
	MaxReportNbr *int `json:"maxReportNbr,omitempty"`

	// MonDur Represents the monitoring duration in seconds.
	MonDur *int `json:"monDur,omitempty"`

	// RepPeriod Repetition period in seconds.
	RepPeriod *int `json:"repPeriod,omitempty"`

	// SampRatio Sampling ratio for data collection.
	SampRatio *int `json:"sampRatio,omitempty"`

	// GroupRepTime Maximum time for grouping reports in seconds.
	GroupRepTime *int `json:"groupRepTime,omitempty"`

	// NotifFlag Notification flag (ACTIVATE, DEACTIVATE, RETRIEVAL).
	NotifFlag *NotificationFlag `json:"notifFlag,omitempty"`

	// PartitionCriteria Partitioning criteria for the impacted UEs before performing sampling.
	// Each element is a PartitioningCriteria enum value (e.g., TAC, SUBPLMN, SNSSAI, DNN, GEOAREA).
	PartitionCriteria *[]PartitioningCriteria `json:"partitionCriteria,omitempty"`
}

// DataNotification Represents a data subscription notification body.
// TODO: Generate from full Nnwdaf_DataManagement YAML spec.
type DataNotification struct {
	// DataAccessProfile Information about the data source.
	DataAccessProfile interface{} `json:"dataAccessProfile,omitempty"`

	// DataItem The actual data collected.
	DataItem interface{} `json:"dataItem,omitempty"`

	// TimeStamp Timestamp when the data was collected.
	TimeStamp *DateTime `json:"timeStamp,omitempty"`
}

// DataSubscription Represents a data subscription configuration.
// TODO: Generate from full Nnwdaf_DataManagement YAML spec.
type DataSubscription struct {
	// DataSources List of data source identifiers.
	DataSources interface{} `json:"dataSources,omitempty"`

	// DataTypes Types of data to collect.
	DataTypes interface{} `json:"dataTypes,omitempty"`

	// ReportingInfo Reporting configuration for the data collection.
	ReportingInfo *ReportingInformation `json:"reportingInfo,omitempty"`
}

// ============================================================================
// Stubs for types referenced in CommonData files but not yet generated
// from their respective YAML specs.
// ============================================================================

// SponsoringStatus Represents the status of a sponsoring agreement.
// TODO: Generate from appropriate 3GPP YAML spec.
type SponsoringStatus struct {
	// AspId Application Service Provider identifier.
	AspId *string `json:"aspId,omitempty"`

	// SponsorEnabled Indicates whether sponsoring is enabled.
	SponsorEnabled *bool `json:"sponsorEnabled,omitempty"`
}

// EthFlowDescription Represents an Ethernet flow description.
// TODO: Generate from appropriate 3GPP YAML spec.
type EthFlowDescription struct {
	// DestMacAddr Destination MAC address.
	DestMacAddr *string `json:"destMacAddr,omitempty"`

	// EthType Ethernet type.
	EthType *string `json:"ethType,omitempty"`

	// VLanTags VLAN tags.
	VLanTags *[]string `json:"vLanTags,omitempty"`
}

// TosTrafficClass Represents the Type of Service (TOS) traffic class.
// TODO: Generate from appropriate 3GPP YAML spec.
type TosTrafficClass struct {
	// TosValue Type of Service value.
	TosValue *int `json:"tosValue,omitempty"`
}

// CivicAddress Represents a civic address.
// TODO: Generate from appropriate 3GPP YAML spec.
type CivicAddress struct {
	// Country Country name.
	Country *string `json:"country,omitempty"`

	// City City name.
	City *string `json:"city,omitempty"`

	// Street Street name.
	Street *string `json:"street,omitempty"`
}

// GeographicArea Represents a geographic area.
// TODO: Generate from appropriate 3GPP YAML spec.
type GeographicArea struct {
	// Shape Shape of the geographic area.
	Shape interface{} `json:"shape,omitempty"`
}

// NetworkAreaInfo Represents network area information.
// TODO: Generate from appropriate 3GPP YAML spec.
type NetworkAreaInfo struct {
	// Tai Tracking Area Identities.
	Tais *[]Tai `json:"tais,omitempty"`

	// Ncgi NR Cell Global Identities.
	Ncgis *[]Ncgi `json:"ncgis,omitempty"`

	// Ecgi E-UTRAN Cell Global Identities.
	Ecgis *[]Ecgi `json:"ecgis,omitempty"`
}

// PlmnID represents the PLMN identifier (MCC + MNC).
// TODO: Generate from appropriate 3GPP YAML spec.
type PlmnID struct {
	// Mcc Mobile Country Code.
	Mcc *string `json:"mcc,omitempty"`

	// Mnc Mobile Network Code.
	Mnc *string `json:"mnc,omitempty"`
}

// SupiRange Represents a range of SUPIs.
// TODO: Generate from appropriate 3GPP YAML spec.
type SupiRange struct {
	// Start Starting SUPI.
	Start *Supi `json:"start,omitempty"`

	// End Ending SUPI.
	End *Supi `json:"end,omitempty"`
}

// IdentityRange Represents a range of identities.
// TODO: Generate from appropriate 3GPP YAML spec.
type IdentityRange struct {
	// Start Starting identity.
	Start *string `json:"start,omitempty"`

	// End Ending identity.
	End *string `json:"end,omitempty"`
}

// CommunicationFailure Represents a communication failure event.
// TODO: Generate from appropriate 3GPP YAML spec.
type CommunicationFailure struct {
	// NasReleaseCode NAS release code.
	NasReleaseCode *string `json:"nasReleaseCode,omitempty"`

	// RanReleaseCode RAN release code.
	RanReleaseCode *string `json:"ranReleaseCode,omitempty"`
}

// AddrFqdn Represents an address (IP or FQDN).
// TODO: Generate from appropriate 3GPP YAML spec.
type AddrFqdn struct {
	// IpAddr IP address.
	IpAddr *string `json:"ipAddr,omitempty"`

	// Fqdn FQDN.
	Fqdn *string `json:"fqdn,omitempty"`
}

// TaiRange Represents a range of Tracking Area Identities.
// TODO: Generate from appropriate 3GPP YAML spec.
type TaiRange struct {
	// PlmnId PLMN Identifier.
	PlmnId *PlmnID `json:"plmnId,omitempty"`

	// TacRangeList Range of TACs.
	TacRangeList *[]TacRange `json:"tacRangeList,omitempty"`
}

// TacRange Represents a range of Tracking Area Codes.
// TODO: Generate from appropriate 3GPP YAML spec.
type TacRange struct {
	// Start Starting TAC.
	Start *string `json:"start,omitempty"`

	// End Ending TAC.
	End *string `json:"end,omitempty"`
}

// FlowDescription Represents a flow description.
// TODO: Generate from appropriate 3GPP YAML spec.
type FlowDescription struct {
	// IpTrafficFilter IP traffic filter.
	IpTrafficFilter *string `json:"ipTrafficFilter,omitempty"`

	// EthTrafficFilter Ethernet traffic filter.
	EthTrafficFilter *EthFlowDescription `json:"ethTrafficFilter,omitempty"`
}

// PositioningMethod Represents a positioning method.
// TODO: Generate from appropriate 3GPP YAML spec.
type PositioningMethod struct {
	union json.RawMessage
}

// GeographicalArea Represents a geographical area.
// TODO: Generate from appropriate 3GPP YAML spec.
type GeographicalArea struct {
	// TaiList List of TAIs.
	TaiList *[]Tai `json:"taiList,omitempty"`

	// NcgiList List of NCGIs.
	NcgiList *[]Ncgi `json:"ncgiList,omitempty"`
}

// GeographicalCoordinates Represents geographical coordinates.
// TODO: Generate from appropriate 3GPP YAML spec.
type GeographicalCoordinates struct {
	// Lon Longitude.
	Lon *float64 `json:"lon,omitempty"`

	// Lat Latitude.
	Lat *float64 `json:"lat,omitempty"`
}

// MLModelAddr Represents the address of an ML model.
// TODO: Generate from appropriate 3GPP YAML spec.
type MLModelAddr struct {
	// MLModelAddrStr ML model address as string.
	MLModelAddrStr *string `json:"mlModelAddrStr,omitempty"`
}

// ModelProvisionParamsExt Represents extended parameters for ML model provisioning.
// TODO: Generate from appropriate 3GPP YAML spec.
type ModelProvisionParamsExt struct {
	// AdrfId ADRF identifier.
	AdrfId *NfInstanceId `json:"adrfId,omitempty"`
}

// NsiId Represents a Network Slice Instance identifier.
// TODO: Generate from appropriate 3GPP YAML spec.
type NsiId = string

// SvcExperience Represents service experience information.
// TODO: Generate from appropriate 3GPP YAML spec.
type SvcExperience struct {
	// Mos Mean Opinion Score.
	Mos *float32 `json:"mos,omitempty"`
}

// DomainNameProtocol Represents a domain name protocol.
// TODO: Generate from appropriate 3GPP YAML spec.
type DomainNameProtocol struct {
	union json.RawMessage
}

// VelocityEstimate Represents a velocity estimate.
// TODO: Generate from appropriate 3GPP YAML spec.
type VelocityEstimate struct {
	// HSpeed Horizontal speed.
	HSpeed *float32 `json:"hSpeed,omitempty"`
}

// ExpectedUeBehaviourData Represents expected UE behaviour data.
// TODO: Generate from appropriate 3GPP YAML spec.
type ExpectedUeBehaviourData struct {
	// StationaryIndication Stationary indication.
	StationaryIndication *bool `json:"stationaryIndication,omitempty"`
}

// NFType Represents the type of a Network Function.
// TODO: Generate from appropriate 3GPP YAML spec.
type NFType struct {
	union json.RawMessage
}

// NoProfileMatchInfo Represents information about no profile match.
// TODO: Generate from appropriate 3GPP YAML spec.
type NoProfileMatchInfo struct {
	// Reason Reason for no match.
	Reason *string `json:"reason,omitempty"`
}

// VendorId Represents a vendor identifier.
// TODO: Generate from appropriate 3GPP YAML spec.
type VendorId = string

// RelativeCartesianLocation Represents a relative cartesian location.
// TODO: Generate from appropriate 3GPP YAML spec.
type RelativeCartesianLocation struct {
	// X X coordinate.
	X *float32 `json:"x,omitempty"`

	// Y Y coordinate.
	Y *float32 `json:"y,omitempty"`
}

// Point Represents a geographic point.
// TODO: Generate from appropriate 3GPP YAML spec.
type Point struct {
	// GeographicalCoordinates Geographical coordinates.
	GeographicalCoordinates *GeographicalCoordinates `json:"geographicalCoordinates,omitempty"`
}

// PointAltitude Represents a point with altitude.
// TODO: Generate from appropriate 3GPP YAML spec.
type PointAltitude struct {
	// Point Geographic point.
	Point *Point `json:"point,omitempty"`

	// Altitude Altitude value.
	Altitude *float64 `json:"altitude,omitempty"`
}

// LocalOrigin Represents a local origin for coordinate systems.
// TODO: Generate from appropriate 3GPP YAML spec.
type LocalOrigin struct {
	// CoordinateId Coordinate identifier.
	CoordinateId *string `json:"coordinateId,omitempty"`
}

// CodecData Represents codec data information.
// TODO: Generate from appropriate 3GPP YAML spec.
type CodecData = string

// DataCollectionPurpose Represents the purpose of data collection.
// TODO: Generate from appropriate 3GPP YAML spec.
type DataCollectionPurpose struct {
	union json.RawMessage
}

// DeletionAlert Represents a deletion alert for data management.
// TODO: Generate from appropriate 3GPP YAML spec.
type DeletionAlert struct {
	// SubIds List of subscription IDs to be deleted.
	SubIds *[]string `json:"subIds,omitempty"`
}

// FetchInstruction Represents a fetch instruction for data retrieval.
// TODO: Generate from appropriate 3GPP YAML spec.
type FetchInstruction struct {
	// SubIds List of subscription IDs for fetching.
	SubIds *[]string `json:"subIds,omitempty"`

	// StartTime Start time for data fetch.
	StartTime *DateTime `json:"startTime,omitempty"`

	// EndTime End time for data fetch.
	EndTime *DateTime `json:"endTime,omitempty"`
}

// FormattingInstruction Represents formatting instructions for data.
// TODO: Generate from appropriate 3GPP YAML spec.
type FormattingInstruction struct {
	union json.RawMessage
}

// NotifSummaryReport Represents a notification summary report.
// TODO: Generate from appropriate 3GPP YAML spec.
type NotifSummaryReport struct {
	// SubId Subscription ID.
	SubId *string `json:"subId,omitempty"`

	// NotifCount Number of notifications.
	NotifCount *int `json:"notifCount,omitempty"`
}

// NotifyEndpoint Represents a notification endpoint.
// TODO: Generate from appropriate 3GPP YAML spec.
type NotifyEndpoint struct {
	// NotifyUri Notification URI.
	NotifyUri *Uri `json:"notifyUri,omitempty"`

	// NotifId Notification correlation ID.
	NotifId *string `json:"notifId,omitempty"`
}

// ProcessingInstruction Represents processing instructions for data.
// TODO: Generate from appropriate 3GPP YAML spec.
type ProcessingInstruction struct {
	union json.RawMessage
}

// StorageHandlingInformation Represents storage handling information.
// TODO: Generate from appropriate 3GPP YAML spec.
type StorageHandlingInformation struct {
	union json.RawMessage
}

// ReservPriority Represents a reservation priority.
// TODO: Generate from appropriate 3GPP YAML spec.
type ReservPriority struct {
	union json.RawMessage
}

// MediaType Represents a media type.
// TODO: Generate from appropriate 3GPP YAML spec.
type MediaType = string

// AfAppId Represents an AF application identifier.
// TODO: Generate from appropriate 3GPP YAML spec.
type AfAppId = string
type CustomUeMobility struct {
	// Ts: Timestamp of the analytics (Optional)
	Ts *DateTime `json:"ts,omitempty"`

	// RecurringTime: Time period for the analytics (Optional)
	RecurringTime *string `json:"recurringTime,omitempty"`

	// Duration: Duration of the UE presence in the location in seconds
	Duration *int32 `json:"duration,omitempty"`

	// DurationVariance: Variance of the duration
	DurationVariance *float32 `json:"durationVariance,omitempty"`

	// LocInfos: Danh sách các vị trí của thiết bị (Bắt buộc phải có)
	LocInfos []LocationInfo `json:"locInfos"`
}

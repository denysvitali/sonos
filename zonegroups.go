package sonos

import "encoding/xml"

type VanishedDevice struct {
	XMLName                 xml.Name `xml:"VanishedDevice"`
	UUID                    string   `xml:"UUID"`
	Location                string   `xml:"Location"`
	ZoneName                string   `xml:"ZoneName"`
	Icon                    string   `xml:"Icon"`
	Configuration           string   `xml:"Configuration"`
	SoftwareVersion         string   `xml:"SoftwareVersion"`
	SWGen                   string   `xml:"SWGen"`
	MinCompatibleVersion    string   `xml:"MinCompatibleVersion"`
	LegacyCompatibleVersion string   `xml:"LegacyCompatibleVersion"`
	BootSeq                 string   `xml:"BootSeq"`
	TVConfigurationError    string   `xml:"TVConfigurationError"`
	HdmiCecAvailable        string   `xml:"HdmiCecAvailable"`
	WirelessMode            string   `xml:"WirelessMode"`
	WirelessLeafOnly        string   `xml:"WirelessLeafOnly"`
	HasConfiguredSSID       string   `xml:"HasConfiguredSSID"`
	ChannelFreq             string   `xml:"ChannelFreq"`
	BehindWifiExtender      string   `xml:"BehindWifiExtender"`
	WifiEnabled             string   `xml:"WifiEnabled"`
	Orientation             string   `xml:"Orientation"`
	RoomCalibrationState    string   `xml:"RoomCalibrationState"`
	SecureRegState          string   `xml:"SecureRegState"`
	VoiceConfigState        string   `xml:"VoiceConfigState"`
	MicEnabled              string   `xml:"MicEnabled"`
	AirPlayEnabled          string   `xml:"AirPlayEnabled"`
	IdleState               string   `xml:"IdleState"`
	MoreInfo                string   `xml:"MoreInfo"`
}

type Satellite struct {
	XMLName                 xml.Name `xml:"Satellite"`
	UUID                    string   `xml:"UUID"`
	Location                string   `xml:"Location"`
	ZoneName                string   `xml:"ZoneName"`
	Icon                    string   `xml:"Icon"`
	Configuration           string   `xml:"Configuration"`
	SoftwareVersion         string   `xml:"SoftwareVersion"`
	SWGen                   string   `xml:"SWGen"`
	MinCompatibleVersion    string   `xml:"MinCompatibleVersion"`
	LegacyCompatibleVersion string   `xml:"LegacyCompatibleVersion"`
	BootSeq                 string   `xml:"BootSeq"`
	TVConfigurationError    string   `xml:"TVConfigurationError"`
	HdmiCecAvailable        string   `xml:"HdmiCecAvailable"`
	WirelessMode            string   `xml:"WirelessMode"`
	WirelessLeafOnly        string   `xml:"WirelessLeafOnly"`
	HasConfiguredSSID       string   `xml:"HasConfiguredSSID"`
	ChannelFreq             string   `xml:"ChannelFreq"`
	BehindWifiExtender      string   `xml:"BehindWifiExtender"`
	WifiEnabled             string   `xml:"WifiEnabled"`
	Orientation             string   `xml:"Orientation"`
	RoomCalibrationState    string   `xml:"RoomCalibrationState"`
	SecureRegState          string   `xml:"SecureRegState"`
	VoiceConfigState        string   `xml:"VoiceConfigState"`
	MicEnabled              string   `xml:"MicEnabled"`
	AirPlayEnabled          string   `xml:"AirPlayEnabled"`
	IdleState               string   `xml:"IdleState"`
	MoreInfo                string   `xml:"MoreInfo"`
}

type ZoneGroupMember struct {
	XMLName                 xml.Name         `xml:"ZoneGroupMember"`
	UUID                    string           `xml:"UUID,attr"`
	Location                string           `xml:"Location,attr"`
	ZoneName                string           `xml:"ZoneName,attr"`
	Icon                    string           `xml:"Icon,attr"`
	Configuration           string           `xml:"Configuration,attr"`
	SoftwareVersion         string           `xml:"SoftwareVersion,attr"`
	SWGen                   string           `xml:"SWGen,attr"`
	MinCompatibleVersion    string           `xml:"MinCompatibleVersion,attr"`
	LegacyCompatibleVersion string           `xml:"LegacyCompatibleVersion,attr"`
	HTForwardEnabled        string           `xml:"HTForwardEnabled,attr"`
	BootSeq                 string           `xml:"BootSeq,attr"`
	TVConfigurationError    string           `xml:"TVConfigurationError,attr"`
	HdmiCecAvailable        string           `xml:"HdmiCecAvailable,attr"`
	WirelessMode            string           `xml:"WirelessMode,attr"`
	WirelessLeafOnly        string           `xml:"WirelessLeafOnly,attr"`
	HasConfiguredSSID       string           `xml:"HasConfiguredSSID,attr"`
	ChannelFreq             string           `xml:"ChannelFreq,attr"`
	BehindWifiExtender      string           `xml:"BehindWifiExtender,attr"`
	WifiEnabled             string           `xml:"WifiEnabled,attr"`
	Orientation             string           `xml:"Orientation,attr"`
	RoomCalibrationState    string           `xml:"RoomCalibrationState,attr"`
	SecureRegState          string           `xml:"SecureRegState,attr"`
	VoiceConfigState        string           `xml:"VoiceConfigState,attr"`
	MicEnabled              string           `xml:"MicEnabled,attr"`
	AirPlayEnabled          bool             `xml:"AirPlayEnabled,attr"`
	VirtualLineInSource     string           `xml:"VirtualLineInSource,attr"`
	IdleState               string           `xml:"IdleState,attr"`
	MoreInfo                string           `xml:"MoreInfo,attr"`
	Satellite               []Satellite      `xml:"Satellite>Satellite"`
	VanishedDevice          []VanishedDevice `xml:"VanishedDevices>VanishedDevice"`
}

type ZoneGroup struct {
	XMLName         xml.Name          `xml:"ZoneGroup"`
	Coordinator     string            `xml:"Coordinator,attr"`
	ID              string            `xml:"ID,attr"`
	ZoneGroupMember []ZoneGroupMember `xml:"ZoneGroupMember"`
}

type ZoneGroupState struct {
	XMLName    xml.Name    `xml:"ZoneGroupState"`
	ZoneGroups []ZoneGroup `xml:"ZoneGroups>ZoneGroup"`
}

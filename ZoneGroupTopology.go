package sonos

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ZoneGroupTopologyService struct {
	controlEndpoint *url.URL
	eventEndpoint   *url.URL
}

func NewZoneGroupTopologyService(deviceUrl *url.URL) *ZoneGroupTopologyService {
	c, _ := url.Parse("/ZoneGroupTopology/Control")
	e, _ := url.Parse("/ZoneGroupTopology/Event")
	return &ZoneGroupTopologyService{
		controlEndpoint: deviceUrl.ResolveReference(c),
		eventEndpoint:   deviceUrl.ResolveReference(e),
	}
}
func (s *ZoneGroupTopologyService) ControlEndpoint() *url.URL {
	return s.controlEndpoint
}
func (s *ZoneGroupTopologyService) EventEndpoint() *url.URL {
	return s.eventEndpoint
}

type ZoneGroupTopologyEnvelope struct {
	XMLName       xml.Name              `xml:"s:Envelope"`
	XMLNameSpace  string                `xml:"xmlns:s,attr"`
	EncodingStyle string                `xml:"s:encodingStyle,attr"`
	Body          ZoneGroupTopologyBody `xml:"s:Body"`
}
type ZoneGroupTopologyBody struct {
	XMLName                   xml.Name                                        `xml:"s:Body"`
	CheckForUpdate            *ZoneGroupTopologyCheckForUpdateArgs            `xml:"u:CheckForUpdate,omitempty"`
	BeginSoftwareUpdate       *ZoneGroupTopologyBeginSoftwareUpdateArgs       `xml:"u:BeginSoftwareUpdate,omitempty"`
	ReportUnresponsiveDevice  *ZoneGroupTopologyReportUnresponsiveDeviceArgs  `xml:"u:ReportUnresponsiveDevice,omitempty"`
	ReportAlarmStartedRunning *ZoneGroupTopologyReportAlarmStartedRunningArgs `xml:"u:ReportAlarmStartedRunning,omitempty"`
	SubmitDiagnostics         *ZoneGroupTopologySubmitDiagnosticsArgs         `xml:"u:SubmitDiagnostics,omitempty"`
	RegisterMobileDevice      *ZoneGroupTopologyRegisterMobileDeviceArgs      `xml:"u:RegisterMobileDevice,omitempty"`
	GetZoneGroupAttributes    *ZoneGroupTopologyGetZoneGroupAttributesArgs    `xml:"u:GetZoneGroupAttributes,omitempty"`
	GetZoneGroupState         *ZoneGroupTopologyGetZoneGroupStateArgs         `xml:"u:GetZoneGroupState,omitempty"`
}
type ZoneGroupTopologyEnvelopeResponse struct {
	XMLName       xml.Name                      `xml:"Envelope"`
	XMLNameSpace  string                        `xml:"xmlns:s,attr"`
	EncodingStyle string                        `xml:"encodingStyle,attr"`
	Body          ZoneGroupTopologyBodyResponse `xml:"Body"`
}
type ZoneGroupTopologyBodyResponse struct {
	XMLName                   xml.Name                                            `xml:"Body"`
	CheckForUpdate            *ZoneGroupTopologyCheckForUpdateResponse            `xml:"CheckForUpdateResponse"`
	BeginSoftwareUpdate       *ZoneGroupTopologyBeginSoftwareUpdateResponse       `xml:"BeginSoftwareUpdateResponse"`
	ReportUnresponsiveDevice  *ZoneGroupTopologyReportUnresponsiveDeviceResponse  `xml:"ReportUnresponsiveDeviceResponse"`
	ReportAlarmStartedRunning *ZoneGroupTopologyReportAlarmStartedRunningResponse `xml:"ReportAlarmStartedRunningResponse"`
	SubmitDiagnostics         *ZoneGroupTopologySubmitDiagnosticsResponse         `xml:"SubmitDiagnosticsResponse"`
	RegisterMobileDevice      *ZoneGroupTopologyRegisterMobileDeviceResponse      `xml:"RegisterMobileDeviceResponse"`
	GetZoneGroupAttributes    *ZoneGroupTopologyGetZoneGroupAttributesResponse    `xml:"GetZoneGroupAttributesResponse"`
	GetZoneGroupState         *ZoneGroupTopologyGetZoneGroupStateResponse         `xml:"GetZoneGroupStateResponse"`
}

func (s *ZoneGroupTopologyService) _ZoneGroupTopologyExec(soapAction string, httpClient *http.Client, envelope *ZoneGroupTopologyEnvelope) (*ZoneGroupTopologyEnvelopeResponse, error) {
	postBody, err := xml.Marshal(envelope)
	if err != nil {
		return nil, err
	}
	// fmt.Printf("soapAction %s: postBody %v\n", soapAction, string(postBody))
	req, err := http.NewRequest("POST", s.controlEndpoint.String(), bytes.NewBuffer(postBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "text/xml; charset=\"utf-8\"")
	req.Header.Set("SOAPAction", soapAction)
	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	responseBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	// fmt.Printf("responseBody %v\n", string(responseBody))
	var envelopeResponse ZoneGroupTopologyEnvelopeResponse
	err = xml.Unmarshal(responseBody, &envelopeResponse)
	if err != nil {
		return nil, err
	}
	return &envelopeResponse, nil
}

type ZoneGroupTopologyCheckForUpdateArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	// Allowed Value: All
	// Allowed Value: Software
	UpdateType string `xml:"UpdateType"`
	CachedOnly bool   `xml:"CachedOnly"`
	Version    string `xml:"Version"`
}
type ZoneGroupTopologyCheckForUpdateResponse struct {
	UpdateItem string `xml:"UpdateItem"`
}

func (s *ZoneGroupTopologyService) CheckForUpdate(httpClient *http.Client, args *ZoneGroupTopologyCheckForUpdateArgs) (*ZoneGroupTopologyCheckForUpdateResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:ZoneGroupTopology:1"
	r, err := s._ZoneGroupTopologyExec("urn:schemas-upnp-org:service:ZoneGroupTopology:1#CheckForUpdate", httpClient,
		&ZoneGroupTopologyEnvelope{
			Body:          ZoneGroupTopologyBody{CheckForUpdate: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.CheckForUpdate == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.CheckForUpdate, nil
}

type ZoneGroupTopologyBeginSoftwareUpdateArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	UpdateURL    string `xml:"UpdateURL"`
	Flags        uint32 `xml:"Flags"`
	ExtraOptions string `xml:"ExtraOptions"`
}
type ZoneGroupTopologyBeginSoftwareUpdateResponse struct {
}

func (s *ZoneGroupTopologyService) BeginSoftwareUpdate(httpClient *http.Client, args *ZoneGroupTopologyBeginSoftwareUpdateArgs) (*ZoneGroupTopologyBeginSoftwareUpdateResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:ZoneGroupTopology:1"
	r, err := s._ZoneGroupTopologyExec("urn:schemas-upnp-org:service:ZoneGroupTopology:1#BeginSoftwareUpdate", httpClient,
		&ZoneGroupTopologyEnvelope{
			Body:          ZoneGroupTopologyBody{BeginSoftwareUpdate: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.BeginSoftwareUpdate == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.BeginSoftwareUpdate, nil
}

type ZoneGroupTopologyReportUnresponsiveDeviceArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	DeviceUUID   string `xml:"DeviceUUID"`
	// Allowed Value: Remove
	// Allowed Value: TopologyMonitorProbe
	// Allowed Value: VerifyThenRemoveSystemwide
	DesiredAction string `xml:"DesiredAction"`
}
type ZoneGroupTopologyReportUnresponsiveDeviceResponse struct {
}

func (s *ZoneGroupTopologyService) ReportUnresponsiveDevice(httpClient *http.Client, args *ZoneGroupTopologyReportUnresponsiveDeviceArgs) (*ZoneGroupTopologyReportUnresponsiveDeviceResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:ZoneGroupTopology:1"
	r, err := s._ZoneGroupTopologyExec("urn:schemas-upnp-org:service:ZoneGroupTopology:1#ReportUnresponsiveDevice", httpClient,
		&ZoneGroupTopologyEnvelope{
			Body:          ZoneGroupTopologyBody{ReportUnresponsiveDevice: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.ReportUnresponsiveDevice == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.ReportUnresponsiveDevice, nil
}

type ZoneGroupTopologyReportAlarmStartedRunningArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
}
type ZoneGroupTopologyReportAlarmStartedRunningResponse struct {
}

func (s *ZoneGroupTopologyService) ReportAlarmStartedRunning(httpClient *http.Client, args *ZoneGroupTopologyReportAlarmStartedRunningArgs) (*ZoneGroupTopologyReportAlarmStartedRunningResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:ZoneGroupTopology:1"
	r, err := s._ZoneGroupTopologyExec("urn:schemas-upnp-org:service:ZoneGroupTopology:1#ReportAlarmStartedRunning", httpClient,
		&ZoneGroupTopologyEnvelope{
			Body:          ZoneGroupTopologyBody{ReportAlarmStartedRunning: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.ReportAlarmStartedRunning == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.ReportAlarmStartedRunning, nil
}

type ZoneGroupTopologySubmitDiagnosticsArgs struct {
	XMLNameSpace       string `xml:"xmlns:u,attr"`
	IncludeControllers bool   `xml:"IncludeControllers"`
	Type               string `xml:"Type"`
}
type ZoneGroupTopologySubmitDiagnosticsResponse struct {
	DiagnosticID uint32 `xml:"DiagnosticID"`
}

func (s *ZoneGroupTopologyService) SubmitDiagnostics(httpClient *http.Client, args *ZoneGroupTopologySubmitDiagnosticsArgs) (*ZoneGroupTopologySubmitDiagnosticsResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:ZoneGroupTopology:1"
	r, err := s._ZoneGroupTopologyExec("urn:schemas-upnp-org:service:ZoneGroupTopology:1#SubmitDiagnostics", httpClient,
		&ZoneGroupTopologyEnvelope{
			Body:          ZoneGroupTopologyBody{SubmitDiagnostics: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.SubmitDiagnostics == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.SubmitDiagnostics, nil
}

type ZoneGroupTopologyRegisterMobileDeviceArgs struct {
	XMLNameSpace     string `xml:"xmlns:u,attr"`
	MobileDeviceName string `xml:"MobileDeviceName"`
	MobileDeviceUDN  string `xml:"MobileDeviceUDN"`
	MobileIPAndPort  string `xml:"MobileIPAndPort"`
}
type ZoneGroupTopologyRegisterMobileDeviceResponse struct {
}

func (s *ZoneGroupTopologyService) RegisterMobileDevice(httpClient *http.Client, args *ZoneGroupTopologyRegisterMobileDeviceArgs) (*ZoneGroupTopologyRegisterMobileDeviceResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:ZoneGroupTopology:1"
	r, err := s._ZoneGroupTopologyExec("urn:schemas-upnp-org:service:ZoneGroupTopology:1#RegisterMobileDevice", httpClient,
		&ZoneGroupTopologyEnvelope{
			Body:          ZoneGroupTopologyBody{RegisterMobileDevice: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.RegisterMobileDevice == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.RegisterMobileDevice, nil
}

type ZoneGroupTopologyGetZoneGroupAttributesArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
}
type ZoneGroupTopologyGetZoneGroupAttributesResponse struct {
	CurrentZoneGroupName          string `xml:"CurrentZoneGroupName"`
	CurrentZoneGroupID            string `xml:"CurrentZoneGroupID"`
	CurrentZonePlayerUUIDsInGroup string `xml:"CurrentZonePlayerUUIDsInGroup"`
	CurrentMuseHouseholdId        string `xml:"CurrentMuseHouseholdId"`
}

func (s *ZoneGroupTopologyService) GetZoneGroupAttributes(httpClient *http.Client, args *ZoneGroupTopologyGetZoneGroupAttributesArgs) (*ZoneGroupTopologyGetZoneGroupAttributesResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:ZoneGroupTopology:1"
	r, err := s._ZoneGroupTopologyExec("urn:schemas-upnp-org:service:ZoneGroupTopology:1#GetZoneGroupAttributes", httpClient,
		&ZoneGroupTopologyEnvelope{
			Body:          ZoneGroupTopologyBody{GetZoneGroupAttributes: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetZoneGroupAttributes == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.GetZoneGroupAttributes, nil
}

type ZoneGroupTopologyGetZoneGroupStateArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
}
type ZoneGroupTopologyGetZoneGroupStateResponse struct {
	ZoneGroupState string `xml:"ZoneGroupState"`
}

func (s *ZoneGroupTopologyService) GetZoneGroupState(httpClient *http.Client, args *ZoneGroupTopologyGetZoneGroupStateArgs) (*ZoneGroupTopologyGetZoneGroupStateResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:ZoneGroupTopology:1"
	r, err := s._ZoneGroupTopologyExec("urn:schemas-upnp-org:service:ZoneGroupTopology:1#GetZoneGroupState", httpClient,
		&ZoneGroupTopologyEnvelope{
			Body:          ZoneGroupTopologyBody{GetZoneGroupState: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetZoneGroupState == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.GetZoneGroupState, nil
}

type ZoneGroupTopologyUpnpEvent struct {
	XMLName      xml.Name                    `xml:"propertyset"`
	XMLNameSpace string                      `xml:"xmlns:e,attr"`
	Properties   []ZoneGroupTopologyProperty `xml:"property"`
}
type ZoneGroupTopologyProperty struct {
	XMLName                 xml.Name `xml:"property"`
	AvailableSoftwareUpdate *string  `xml:"AvailableSoftwareUpdate"`
	ZoneGroupState          *string  `xml:"ZoneGroupState"`
	ThirdPartyMediaServersX *string  `xml:"ThirdPartyMediaServersX"`
	AlarmRunSequence        *string  `xml:"AlarmRunSequence"`
	MuseHouseholdId         *string  `xml:"MuseHouseholdId"`
	ZoneGroupName           *string  `xml:"ZoneGroupName"`
	ZoneGroupID             *string  `xml:"ZoneGroupID"`
	ZonePlayerUUIDsInGroup  *string  `xml:"ZonePlayerUUIDsInGroup"`
	AreasUpdateID           *string  `xml:"AreasUpdateID"`
	SourceAreasUpdateID     *string  `xml:"SourceAreasUpdateID"`
	NetsettingsUpdateID     *string  `xml:"NetsettingsUpdateID"`
}

func ZoneGroupTopologyDispatchEvent(zp *ZonePlayer, body []byte) {
	var evt ZoneGroupTopologyUpnpEvent
	err := xml.Unmarshal(body, &evt)
	if err != nil {
		return
	}
	for _, prop := range evt.Properties {
		switch {
		case prop.AvailableSoftwareUpdate != nil:
			dispatchZoneGroupTopologyAvailableSoftwareUpdate(zp, *prop.AvailableSoftwareUpdate) // string
		case prop.ZoneGroupState != nil:
			dispatchZoneGroupTopologyZoneGroupState(zp, *prop.ZoneGroupState) // string
		case prop.ThirdPartyMediaServersX != nil:
			dispatchZoneGroupTopologyThirdPartyMediaServersX(zp, *prop.ThirdPartyMediaServersX) // string
		case prop.AlarmRunSequence != nil:
			dispatchZoneGroupTopologyAlarmRunSequence(zp, *prop.AlarmRunSequence) // string
		case prop.MuseHouseholdId != nil:
			dispatchZoneGroupTopologyMuseHouseholdId(zp, *prop.MuseHouseholdId) // string
		case prop.ZoneGroupName != nil:
			dispatchZoneGroupTopologyZoneGroupName(zp, *prop.ZoneGroupName) // string
		case prop.ZoneGroupID != nil:
			dispatchZoneGroupTopologyZoneGroupID(zp, *prop.ZoneGroupID) // string
		case prop.ZonePlayerUUIDsInGroup != nil:
			dispatchZoneGroupTopologyZonePlayerUUIDsInGroup(zp, *prop.ZonePlayerUUIDsInGroup) // string
		case prop.AreasUpdateID != nil:
			dispatchZoneGroupTopologyAreasUpdateID(zp, *prop.AreasUpdateID) // string
		case prop.SourceAreasUpdateID != nil:
			dispatchZoneGroupTopologySourceAreasUpdateID(zp, *prop.SourceAreasUpdateID) // string
		case prop.NetsettingsUpdateID != nil:
			dispatchZoneGroupTopologyNetsettingsUpdateID(zp, *prop.NetsettingsUpdateID) // string
		}
	}
}

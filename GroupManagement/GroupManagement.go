// Code generated by makeservice. DO NOT EDIT.

// Package groupmanagement is a generated GroupManagement package.
package groupmanagement

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	_ServiceURN     = "urn:schemas-upnp-org:service:GroupManagement:1"
	_EncodingSchema = "http://schemas.xmlsoap.org/soap/encoding/"
	_EnvelopeSchema = "http://schemas.xmlsoap.org/soap/envelope/"
)

type ServiceOption func(*Service)

func WithClient(c *http.Client) ServiceOption {
	return func(s *Service) {
		s.client = c
	}
}

func WithLocation(u *url.URL) ServiceOption {
	return func(s *Service) {
		s.location = u
	}
}

type Service struct {
	controlEndpoint *url.URL
	eventEndpoint   *url.URL

	location *url.URL
	client   *http.Client
}

func NewService(opts ...ServiceOption) *Service {
	s := &Service{}

	c, err := url.Parse("/GroupManagement/Control")
	if nil != err {
		panic(err)
	}
	e, err := url.Parse("/GroupManagement/Event")
	if nil != err {
		panic(err)
	}

	for _, opt := range opts {
		opt(s)
	}

	if s.client == nil {
		panic("no client location")
	}
	if s.location == nil {
		panic("empty location")
	}

	s.controlEndpoint = s.location.ResolveReference(c)
	s.eventEndpoint = s.location.ResolveReference(e)

	return s
}

func (s *Service) ControlEndpoint() *url.URL {
	return s.controlEndpoint
}

func (s *Service) EventEndpoint() *url.URL {
	return s.eventEndpoint
}

func (s *Service) Location() *url.URL {
	return s.location
}

func (s *Service) Client() *http.Client {
	return s.client
}

type Envelope struct {
	XMLName       xml.Name `xml:"s:Envelope"`
	Xmlns         string   `xml:"xmlns:s,attr"`
	EncodingStyle string   `xml:"s:encodingStyle,attr"`
	Body          Body     `xml:"s:Body"`
}
type Body struct {
	XMLName                    xml.Name                        `xml:"s:Body"`
	AddMember                  *AddMemberArgs                  `xml:"u:AddMember,omitempty"`
	RemoveMember               *RemoveMemberArgs               `xml:"u:RemoveMember,omitempty"`
	ReportTrackBufferingResult *ReportTrackBufferingResultArgs `xml:"u:ReportTrackBufferingResult,omitempty"`
	SetSourceAreaIds           *SetSourceAreaIdsArgs           `xml:"u:SetSourceAreaIds,omitempty"`
}
type EnvelopeResponse struct {
	XMLName       xml.Name     `xml:"Envelope"`
	Xmlns         string       `xml:"xmlns:s,attr"`
	EncodingStyle string       `xml:"encodingStyle,attr"`
	Body          BodyResponse `xml:"Body"`
}
type BodyResponse struct {
	XMLName                    xml.Name                            `xml:"Body"`
	AddMember                  *AddMemberResponse                  `xml:"AddMemberResponse,omitempty"`
	RemoveMember               *RemoveMemberResponse               `xml:"RemoveMemberResponse,omitempty"`
	ReportTrackBufferingResult *ReportTrackBufferingResultResponse `xml:"ReportTrackBufferingResultResponse,omitempty"`
	SetSourceAreaIds           *SetSourceAreaIdsResponse           `xml:"SetSourceAreaIdsResponse,omitempty"`
}

func (s *Service) exec(actionName string, envelope *Envelope) (*EnvelopeResponse, error) {
	marshaled, err := xml.Marshal(envelope)
	if err != nil {
		return nil, err
	}
	postBody := []byte("<?xml version=\"1.0\"?>")
	postBody = append(postBody, marshaled...)
	req, err := http.NewRequest("POST", s.controlEndpoint.String(), bytes.NewBuffer(postBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "text/xml; charset=\"utf-8\"")
	req.Header.Set("SOAPAction", _ServiceURN+"#"+actionName)
	res, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	responseBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var envelopeResponse EnvelopeResponse
	err = xml.Unmarshal(responseBody, &envelopeResponse)
	if err != nil {
		return nil, err
	}
	return &envelopeResponse, nil
}

type AddMemberArgs struct {
	Xmlns    string `xml:"xmlns:u,attr"`
	MemberID string `xml:"MemberID"`
	BootSeq  uint32 `xml:"BootSeq"`
}
type AddMemberResponse struct {
	CurrentTransportSettings string `xml:"CurrentTransportSettings"`
	CurrentURI               string `xml:"CurrentURI"`
	GroupUUIDJoined          string `xml:"GroupUUIDJoined"`
	ResetVolumeAfter         bool   `xml:"ResetVolumeAfter"`
	VolumeAVTransportURI     string `xml:"VolumeAVTransportURI"`
}

func (s *Service) AddMember(args *AddMemberArgs) (*AddMemberResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`AddMember`,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{AddMember: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.AddMember == nil {
		return nil, errors.New(`unexpected response from service calling groupmanagement.AddMember()`)
	}

	return r.Body.AddMember, nil
}

type RemoveMemberArgs struct {
	Xmlns    string `xml:"xmlns:u,attr"`
	MemberID string `xml:"MemberID"`
}
type RemoveMemberResponse struct {
}

func (s *Service) RemoveMember(args *RemoveMemberArgs) (*RemoveMemberResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`RemoveMember`,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{RemoveMember: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.RemoveMember == nil {
		return nil, errors.New(`unexpected response from service calling groupmanagement.RemoveMember()`)
	}

	return r.Body.RemoveMember, nil
}

type ReportTrackBufferingResultArgs struct {
	Xmlns      string `xml:"xmlns:u,attr"`
	MemberID   string `xml:"MemberID"`
	ResultCode int32  `xml:"ResultCode"`
}
type ReportTrackBufferingResultResponse struct {
}

func (s *Service) ReportTrackBufferingResult(args *ReportTrackBufferingResultArgs) (*ReportTrackBufferingResultResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`ReportTrackBufferingResult`,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{ReportTrackBufferingResult: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.ReportTrackBufferingResult == nil {
		return nil, errors.New(`unexpected response from service calling groupmanagement.ReportTrackBufferingResult()`)
	}

	return r.Body.ReportTrackBufferingResult, nil
}

type SetSourceAreaIdsArgs struct {
	Xmlns                string `xml:"xmlns:u,attr"`
	DesiredSourceAreaIds string `xml:"DesiredSourceAreaIds"`
}
type SetSourceAreaIdsResponse struct {
}

func (s *Service) SetSourceAreaIds(args *SetSourceAreaIdsArgs) (*SetSourceAreaIdsResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`SetSourceAreaIds`,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{SetSourceAreaIds: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.SetSourceAreaIds == nil {
		return nil, errors.New(`unexpected response from service calling groupmanagement.SetSourceAreaIds()`)
	}

	return r.Body.SetSourceAreaIds, nil
}

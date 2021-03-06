// Code generated by makeservice. DO NOT EDIT.

// Package virtuallinein is a generated VirtualLineIn package.
package virtuallinein

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	_ServiceURN     = "urn:schemas-upnp-org:service:VirtualLineIn:1"
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

	c, err := url.Parse("/MediaRenderer/VirtualLineIn/Control")
	if nil != err {
		panic(err)
	}
	e, err := url.Parse("/MediaRenderer/VirtualLineIn/Event")
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
	XMLName           xml.Name               `xml:"s:Body"`
	StartTransmission *StartTransmissionArgs `xml:"u:StartTransmission,omitempty"`
	StopTransmission  *StopTransmissionArgs  `xml:"u:StopTransmission,omitempty"`
	Play              *PlayArgs              `xml:"u:Play,omitempty"`
	Pause             *PauseArgs             `xml:"u:Pause,omitempty"`
	Next              *NextArgs              `xml:"u:Next,omitempty"`
	Previous          *PreviousArgs          `xml:"u:Previous,omitempty"`
	Stop              *StopArgs              `xml:"u:Stop,omitempty"`
	SetVolume         *SetVolumeArgs         `xml:"u:SetVolume,omitempty"`
}
type EnvelopeResponse struct {
	XMLName       xml.Name     `xml:"Envelope"`
	Xmlns         string       `xml:"xmlns:s,attr"`
	EncodingStyle string       `xml:"encodingStyle,attr"`
	Body          BodyResponse `xml:"Body"`
}
type BodyResponse struct {
	XMLName           xml.Name                   `xml:"Body"`
	StartTransmission *StartTransmissionResponse `xml:"StartTransmissionResponse,omitempty"`
	StopTransmission  *StopTransmissionResponse  `xml:"StopTransmissionResponse,omitempty"`
	Play              *PlayResponse              `xml:"PlayResponse,omitempty"`
	Pause             *PauseResponse             `xml:"PauseResponse,omitempty"`
	Next              *NextResponse              `xml:"NextResponse,omitempty"`
	Previous          *PreviousResponse          `xml:"PreviousResponse,omitempty"`
	Stop              *StopResponse              `xml:"StopResponse,omitempty"`
	SetVolume         *SetVolumeResponse         `xml:"SetVolumeResponse,omitempty"`
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

type StartTransmissionArgs struct {
	Xmlns         string `xml:"xmlns:u,attr"`
	InstanceID    uint32 `xml:"InstanceID"`
	CoordinatorID string `xml:"CoordinatorID"`
}
type StartTransmissionResponse struct {
	CurrentTransportSettings string `xml:"CurrentTransportSettings"`
}

func (s *Service) StartTransmission(args *StartTransmissionArgs) (*StartTransmissionResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`StartTransmission`,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{StartTransmission: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.StartTransmission == nil {
		return nil, errors.New(`unexpected response from service calling virtuallinein.StartTransmission()`)
	}

	return r.Body.StartTransmission, nil
}

type StopTransmissionArgs struct {
	Xmlns         string `xml:"xmlns:u,attr"`
	InstanceID    uint32 `xml:"InstanceID"`
	CoordinatorID string `xml:"CoordinatorID"`
}
type StopTransmissionResponse struct {
}

func (s *Service) StopTransmission(args *StopTransmissionArgs) (*StopTransmissionResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`StopTransmission`,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{StopTransmission: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.StopTransmission == nil {
		return nil, errors.New(`unexpected response from service calling virtuallinein.StopTransmission()`)
	}

	return r.Body.StopTransmission, nil
}

type PlayArgs struct {
	Xmlns      string `xml:"xmlns:u,attr"`
	InstanceID uint32 `xml:"InstanceID"`
	Speed      string `xml:"Speed"`
}
type PlayResponse struct {
}

func (s *Service) Play(args *PlayArgs) (*PlayResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`Play`,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{Play: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.Play == nil {
		return nil, errors.New(`unexpected response from service calling virtuallinein.Play()`)
	}

	return r.Body.Play, nil
}

type PauseArgs struct {
	Xmlns      string `xml:"xmlns:u,attr"`
	InstanceID uint32 `xml:"InstanceID"`
}
type PauseResponse struct {
}

func (s *Service) Pause(args *PauseArgs) (*PauseResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`Pause`,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{Pause: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.Pause == nil {
		return nil, errors.New(`unexpected response from service calling virtuallinein.Pause()`)
	}

	return r.Body.Pause, nil
}

type NextArgs struct {
	Xmlns      string `xml:"xmlns:u,attr"`
	InstanceID uint32 `xml:"InstanceID"`
}
type NextResponse struct {
}

func (s *Service) Next(args *NextArgs) (*NextResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`Next`,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{Next: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.Next == nil {
		return nil, errors.New(`unexpected response from service calling virtuallinein.Next()`)
	}

	return r.Body.Next, nil
}

type PreviousArgs struct {
	Xmlns      string `xml:"xmlns:u,attr"`
	InstanceID uint32 `xml:"InstanceID"`
}
type PreviousResponse struct {
}

func (s *Service) Previous(args *PreviousArgs) (*PreviousResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`Previous`,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{Previous: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.Previous == nil {
		return nil, errors.New(`unexpected response from service calling virtuallinein.Previous()`)
	}

	return r.Body.Previous, nil
}

type StopArgs struct {
	Xmlns      string `xml:"xmlns:u,attr"`
	InstanceID uint32 `xml:"InstanceID"`
}
type StopResponse struct {
}

func (s *Service) Stop(args *StopArgs) (*StopResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`Stop`,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{Stop: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.Stop == nil {
		return nil, errors.New(`unexpected response from service calling virtuallinein.Stop()`)
	}

	return r.Body.Stop, nil
}

type SetVolumeArgs struct {
	Xmlns         string `xml:"xmlns:u,attr"`
	InstanceID    uint32 `xml:"InstanceID"`
	DesiredVolume uint16 `xml:"DesiredVolume"`
}
type SetVolumeResponse struct {
}

func (s *Service) SetVolume(args *SetVolumeArgs) (*SetVolumeResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`SetVolume`,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{SetVolume: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.SetVolume == nil {
		return nil, errors.New(`unexpected response from service calling virtuallinein.SetVolume()`)
	}

	return r.Body.SetVolume, nil
}

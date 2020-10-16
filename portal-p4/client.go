package portal_p4

import (
	"bytes"
	"context"
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"io"
	"net/http"
	base "team529.nl/go-portal-p4"
	"team529.nl/go-portal-p4/portal-p4/messages"
)

var envToEndpointBase = map[string]string{
	"test": "https://pp4-test.edsn.nl:443",
	"prod": "https://p4.edsn.nl:443",
}

type Client struct {
	EndpointBase string
	Certificate  *x509.Certificate
	PrivateKey   *rsa.PrivateKey
	HTTPClient   *http.Client
	ResponseHook func(*http.Response) *http.Response
	RequestHook  func(*http.Request) *http.Request
}

func NewClient(env string, certificate x509.Certificate, key rsa.PrivateKey) (*Client, error) {
	endpoint, ok := envToEndpointBase[env]
	if !ok {
		return nil, fmt.Errorf("env [%s] unknown", env)
	}

	client := Client{
		EndpointBase: endpoint,
		Certificate:  &certificate,
		PrivateKey:   &key,
	}

	return &client, nil
}

func (c *Client) do(ctx context.Context, method, uri, action string, in, out interface{}) error {

	if err := base.AssertPointer(in, "Non pointer passed to in"); err != nil {
		return err
	}

	if err := base.AssertPointer(out, "Non pointer passed to out"); err != nil {
		return err
	}

	var body io.Reader

	signer := base.Signer{
		Certificate: c.Certificate,
		PrivateKey:  c.PrivateKey,
	}

	base64Token, err := base.ToBase64(c.Certificate)
	if err != nil {
		return err
	}

	header := base.Header{Security: base.Security{
		MustUnderstand: "1",
		BinarySecurityToken: base.BinarySecurityToken{
			Id:           "x509-1",
			EncodingType: "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-soap-message-security-1.0#Base64Binary",
			ValueType:    "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-x509-token-profile-1.0#X509v3",
			Token:        base64Token,
		},
	}}
	var envelope base.SoapEnvelope
	envelope.Header = header
	envelope.Body.Id = "_0"
	if method == "POST" || method == "PUT" {
		envelope.Body.Message = in

		signature, err := signer.Sign(&envelope.Body)
		if err != nil {
			return err
		}

		envelope.Header.Security.Signature = *signature
		xml := base.ObjToXML(envelope).String()
		var b bytes.Buffer
		b.WriteString(xml)
		body = &b
	}
	req, err := http.NewRequest(method, uri, body)
	if err != nil {
		return err
	}
	req.Header.Set("SOAPAction", action)
	req = req.WithContext(ctx)
	if c.RequestHook != nil {
		req = c.RequestHook(req)
	}
	httpClient := c.HTTPClient
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	rsp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()
	if c.ResponseHook != nil {
		rsp = c.ResponseHook(rsp)
	}

	responseEnvelope := base.SoapEnvelope{}
	responseEnvelope.Body.Message = out

	xmlBuf := new(bytes.Buffer)
	xmlBuf.ReadFrom(rsp.Body)
	if err := base.ObjFromXMLString(xmlBuf.String(), &responseEnvelope); err != nil {
		return err
	}

	validator, err := base.NewValidator(c.EndpointBase)
	if err != nil {
		return err
	}

	if err := validator.ValidateXmlString(xmlBuf.String()); err != nil {
		return err
	}

	if envelope.Body.Fault != nil {
		return fmt.Errorf("%s: %s", envelope.Body.Fault.Code, envelope.Body.Fault.String)
	}
	return nil
}

func (c *Client) DatabatchRequest(ctx context.Context, part1 messages.P4CollectedDataBatchRequestEnvelope) (messages.P4CollectedDataBatchResponseEnvelope, error) {
	var response messages.P4CollectedDataBatchResponseEnvelope
	err := c.do(ctx, "POST", fmt.Sprintf("%s/P4BatchVerzoekMeterstand/P4Port", c.EndpointBase), "urn:P4CollectedDataBatchRequest", &part1, &response)
	return response, err
}

func (c *Client) DatabatchResult(ctx context.Context, part1 messages.P4CollectedDataBatchResultRequestEnvelope) (messages.P4CollectedDataBatchResultResponseEnvelope, error) {
	var response messages.P4CollectedDataBatchResultResponseEnvelope
	err := c.do(ctx, "POST", fmt.Sprintf("%s/P4BatchVerzoekMeterstand/P4Port", c.EndpointBase), "urn:P4CollectedDataBatchResult", &part1, &response)
	return response, err
}

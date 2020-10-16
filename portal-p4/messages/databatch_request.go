package messages

import "encoding/xml"

type P4CollectedDataBatchRequestEnvelope struct {
	XMLName                    xml.Name                                                      `xml:"urn:nedu:edsn:data:p4collecteddatabatchrequest:2:standard P4CollectedDataBatchRequestEnvelope"`
	EDSNBusinessDocumentHeader P4CollectedDataBatchRequestEnvelopeEDSNBusinessDocumentHeader `xml:"urn:nedu:edsn:data:p4collecteddatabatchrequest:2:standard EDSNBusinessDocumentHeader"`
	P4Content                  P4CollectedDataBatchRequestEnvelopeP4Content                  `xml:"urn:nedu:edsn:data:p4collecteddatabatchrequest:2:standard P4Content"`
}

type P4CollectedDataBatchRequestEnvelopeP4Content struct {
	P4MeteringPoint []P4CollectedDataBatchRequestEnvelopeP4ContentP4MeteringPoint `xml:"urn:nedu:edsn:data:p4collecteddatabatchrequest:2:standard P4MeteringPoint"`
}

type P4CollectedDataBatchRequestEnvelopeP4ContentP4MeteringPoint struct {
	EANID             GSRNEANCode         `xml:"urn:nedu:edsn:data:p4collecteddatabatchrequest:2:standard EANID"`
	ExternalReference TextType            `xml:"urn:nedu:edsn:data:p4collecteddatabatchrequest:2:standard ExternalReference"`
	QueryDate         Date                `xml:"urn:nedu:edsn:data:p4collecteddatabatchrequest:2:standard QueryDate,omitempty"`
	QueryReason       QueryReasonTypeCode `xml:"urn:nedu:edsn:data:p4collecteddatabatchrequest:2:standard QueryReason"`
}

type P4CollectedDataBatchRequestEnvelopeEDSNBusinessDocumentHeader struct {
	ContentHash       string                                                                   `xml:"urn:nedu:edsn:data:p4collecteddatabatchrequest:2:standard ContentHash,omitempty"`
	ConversationID    string                                                                   `xml:"urn:nedu:edsn:data:p4collecteddatabatchrequest:2:standard ConversationID,omitempty"`
	CorrelationID     string                                                                   `xml:"urn:nedu:edsn:data:p4collecteddatabatchrequest:2:standard CorrelationID,omitempty"`
	CreationTimestamp *DateTime                                                                `xml:"urn:nedu:edsn:data:p4collecteddatabatchrequest:2:standard CreationTimestamp"`
	DocumentID        string                                                                   `xml:"urn:nedu:edsn:data:p4collecteddatabatchrequest:2:standard DocumentID,omitempty"`
	ExpiresAt         *DateTime                                                                `xml:"urn:nedu:edsn:data:p4collecteddatabatchrequest:2:standard ExpiresAt,omitempty"`
	MessageID         string                                                                   `xml:"urn:nedu:edsn:data:p4collecteddatabatchrequest:2:standard MessageID"`
	ProcessTypeID     string                                                                   `xml:"urn:nedu:edsn:data:p4collecteddatabatchrequest:2:standard ProcessTypeID,omitempty"`
	RepeatedRequest   string                                                                   `xml:"urn:nedu:edsn:data:p4collecteddatabatchrequest:2:standard RepeatedRequest,omitempty"`
	TestRequest       string                                                                   `xml:"urn:nedu:edsn:data:p4collecteddatabatchrequest:2:standard TestRequest,omitempty"`
	Destination       P4CollectedDataBatchRequestEnvelopeEDSNBusinessDocumentHeaderDestination `xml:"urn:nedu:edsn:data:p4collecteddatabatchrequest:2:standard Destination"`
	Manifest          *P4CollectedDataBatchRequestEnvelopeEDSNBusinessDocumentHeaderManifest   `xml:"urn:nedu:edsn:data:p4collecteddatabatchrequest:2:standard Manifest,omitempty"`
	Source            P4CollectedDataBatchRequestEnvelopeEDSNBusinessDocumentHeaderSource      `xml:"urn:nedu:edsn:data:p4collecteddatabatchrequest:2:standard Source"`
}

type P4CollectedDataBatchRequestEnvelopeEDSNBusinessDocumentHeaderDestination struct {
	Receiver P4CollectedDataBatchRequestEnvelopeEDSNBusinessDocumentHeaderDestinationReceiver `xml:"urn:nedu:edsn:data:p4collecteddatabatchrequest:2:standard Receiver"`
	Service  *P4CollectedDataBatchRequestEnvelopeEDSNBusinessDocumentHeaderDestinationService `xml:"urn:nedu:edsn:data:p4collecteddatabatchrequest:2:standard Service,omitempty"`
}

type P4CollectedDataBatchRequestEnvelopeEDSNBusinessDocumentHeaderDestinationReceiver struct {
	Authority             string `xml:"urn:nedu:edsn:data:p4collecteddatabatchrequest:2:standard Authority,omitempty"`
	ContactTypeIdentifier string `xml:"urn:nedu:edsn:data:p4collecteddatabatchrequest:2:standard ContactTypeIdentifier,omitempty"`
	ReceiverID            string `xml:"urn:nedu:edsn:data:p4collecteddatabatchrequest:2:standard ReceiverID"`
}

type P4CollectedDataBatchRequestEnvelopeEDSNBusinessDocumentHeaderDestinationService struct {
	ServiceMethod string `xml:"urn:nedu:edsn:data:p4collecteddatabatchrequest:2:standard ServiceMethod,omitempty"`
	ServiceName   string `xml:"urn:nedu:edsn:data:p4collecteddatabatchrequest:2:standard ServiceName,omitempty"`
}

type P4CollectedDataBatchRequestEnvelopeEDSNBusinessDocumentHeaderManifest struct {
	NumberofItems float64                                                                             `xml:"urn:nedu:edsn:data:p4collecteddatabatchrequest:2:standard NumberofItems"`
	ManifestItem  []P4CollectedDataBatchRequestEnvelopeEDSNBusinessDocumentHeaderManifestManifestItem `xml:"urn:nedu:edsn:data:p4collecteddatabatchrequest:2:standard ManifestItem"`
}

type P4CollectedDataBatchRequestEnvelopeEDSNBusinessDocumentHeaderManifestManifestItem struct {
	Description               string `xml:"urn:nedu:edsn:data:p4collecteddatabatchrequest:2:standard Description,omitempty"`
	LanguageCode              string `xml:"urn:nedu:edsn:data:p4collecteddatabatchrequest:2:standard LanguageCode,omitempty"`
	MimeTypeQualifierCode     string `xml:"urn:nedu:edsn:data:p4collecteddatabatchrequest:2:standard MimeTypeQualifierCode"`
	UniformResourceIdentifier string `xml:"urn:nedu:edsn:data:p4collecteddatabatchrequest:2:standard UniformResourceIdentifier"`
}

type P4CollectedDataBatchRequestEnvelopeEDSNBusinessDocumentHeaderSource struct {
	Authority             string `xml:"urn:nedu:edsn:data:p4collecteddatabatchrequest:2:standard Authority,omitempty"`
	ContactTypeIdentifier string `xml:"urn:nedu:edsn:data:p4collecteddatabatchrequest:2:standard ContactTypeIdentifier,omitempty"`
	SenderID              string `xml:"urn:nedu:edsn:data:p4collecteddatabatchrequest:2:standard SenderID"`
}

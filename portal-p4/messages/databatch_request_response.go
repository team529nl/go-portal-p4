package messages

import "encoding/xml"

type P4CollectedDataBatchResponseEnvelope struct {
	XMLName                    xml.Name                                                       `xml:"urn:nedu:edsn:data:p4collecteddatabatchresponse:2:standard P4CollectedDataBatchResponseEnvelope"`
	EDSNBusinessDocumentHeader P4CollectedDataBatchResponseEnvelopeEDSNBusinessDocumentHeader `xml:"urn:nedu:edsn:data:p4collecteddatabatchresponse:2:standard EDSNBusinessDocumentHeader"`
	P4Content                  *P4CollectedDataBatchResponseEnvelopeP4Content                 `xml:"urn:nedu:edsn:data:p4collecteddatabatchresponse:2:standard P4Content"`
}

type P4CollectedDataBatchResponseEnvelopeP4Content struct {
	P4Rejection *P4CollectedDataBatchResponseEnvelopeP4ContentP4Rejection `xml:"urn:nedu:edsn:data:p4collecteddatabatchresponse:2:standard P4Rejection,omitempty"`
}

type P4CollectedDataBatchResponseEnvelopeP4ContentP4Rejection struct {
	Rejection P4CollectedDataBatchResponseEnvelopeRejectionP4Type `xml:"urn:nedu:edsn:data:p4collecteddatabatchresponse:2:standard Rejection"`
}

type P4CollectedDataBatchResponseEnvelopeRejectionP4Type struct {
	RejectionCode RejectionReasonType `xml:"urn:nedu:edsn:data:p4collecteddatabatchresponse:2:standard RejectionCode"`
	RejectionText string              `xml:"urn:nedu:edsn:data:p4collecteddatabatchresponse:2:standard RejectionText,omitempty"`
}

type P4CollectedDataBatchResponseEnvelopeEDSNBusinessDocumentHeader struct {
	ContentHash       string                                                                    `xml:"urn:nedu:edsn:data:p4collecteddatabatchresponse:2:standard ContentHash,omitempty"`
	ConversationID    string                                                                    `xml:"urn:nedu:edsn:data:p4collecteddatabatchresponse:2:standard ConversationID,omitempty"`
	CorrelationID     string                                                                    `xml:"urn:nedu:edsn:data:p4collecteddatabatchresponse:2:standard CorrelationID,omitempty"`
	CreationTimestamp *DateTime                                                                 `xml:"urn:nedu:edsn:data:p4collecteddatabatchresponse:2:standard CreationTimestamp"`
	DocumentID        string                                                                    `xml:"urn:nedu:edsn:data:p4collecteddatabatchresponse:2:standard DocumentID,omitempty"`
	ExpiresAt         *DateTime                                                                 `xml:"urn:nedu:edsn:data:p4collecteddatabatchresponse:2:standard ExpiresAt,omitempty"`
	MessageID         string                                                                    `xml:"urn:nedu:edsn:data:p4collecteddatabatchresponse:2:standard MessageID"`
	ProcessTypeID     string                                                                    `xml:"urn:nedu:edsn:data:p4collecteddatabatchresponse:2:standard ProcessTypeID,omitempty"`
	RepeatedRequest   string                                                                    `xml:"urn:nedu:edsn:data:p4collecteddatabatchresponse:2:standard RepeatedRequest,omitempty"`
	TestRequest       string                                                                    `xml:"urn:nedu:edsn:data:p4collecteddatabatchresponse:2:standard TestRequest,omitempty"`
	Destination       P4CollectedDataBatchResponseEnvelopeEDSNBusinessDocumentHeaderDestination `xml:"urn:nedu:edsn:data:p4collecteddatabatchresponse:2:standard Destination"`
	Manifest          *P4CollectedDataBatchResponseEnvelopeEDSNBusinessDocumentHeaderManifest   `xml:"urn:nedu:edsn:data:p4collecteddatabatchresponse:2:standard Manifest,omitempty"`
	Source            P4CollectedDataBatchResponseEnvelopeEDSNBusinessDocumentHeaderSource      `xml:"urn:nedu:edsn:data:p4collecteddatabatchresponse:2:standard Source"`
}

type P4CollectedDataBatchResponseEnvelopeEDSNBusinessDocumentHeaderDestination struct {
	Receiver P4CollectedDataBatchResponseEnvelopeEDSNBusinessDocumentHeaderDestinationReceiver `xml:"urn:nedu:edsn:data:p4collecteddatabatchresponse:2:standard Receiver"`
	Service  P4CollectedDataBatchResponseEnvelopeEDSNBusinessDocumentHeaderDestinationService  `xml:"urn:nedu:edsn:data:p4collecteddatabatchresponse:2:standard Service,omitempty"`
}

type P4CollectedDataBatchResponseEnvelopeEDSNBusinessDocumentHeaderDestinationReceiver struct {
	Authority             string `xml:"urn:nedu:edsn:data:p4collecteddatabatchresponse:2:standard Authority,omitempty"`
	ContactTypeIdentifier string `xml:"urn:nedu:edsn:data:p4collecteddatabatchresponse:2:standard ContactTypeIdentifier,omitempty"`
	ReceiverID            string `xml:"urn:nedu:edsn:data:p4collecteddatabatchresponse:2:standard ReceiverID"`
}

type P4CollectedDataBatchResponseEnvelopeEDSNBusinessDocumentHeaderDestinationService struct {
	ServiceMethod string `xml:"urn:nedu:edsn:data:p4collecteddatabatchresponse:2:standard ServiceMethod,omitempty"`
	ServiceName   string `xml:"urn:nedu:edsn:data:p4collecteddatabatchresponse:2:standard ServiceName,omitempty"`
}

type P4CollectedDataBatchResponseEnvelopeEDSNBusinessDocumentHeaderManifest struct {
	NumberofItems float64                                                                              `xml:"urn:nedu:edsn:data:p4collecteddatabatchresponse:2:standard NumberofItems"`
	ManifestItem  []P4CollectedDataBatchResponseEnvelopeEDSNBusinessDocumentHeaderManifestManifestItem `xml:"urn:nedu:edsn:data:p4collecteddatabatchresponse:2:standard ManifestItem"`
}

type P4CollectedDataBatchResponseEnvelopeEDSNBusinessDocumentHeaderManifestManifestItem struct {
	Description               string `xml:"urn:nedu:edsn:data:p4collecteddatabatchresponse:2:standard Description,omitempty"`
	LanguageCode              string `xml:"urn:nedu:edsn:data:p4collecteddatabatchresponse:2:standard LanguageCode,omitempty"`
	MimeTypeQualifierCode     string `xml:"urn:nedu:edsn:data:p4collecteddatabatchresponse:2:standard MimeTypeQualifierCode"`
	UniformResourceIdentifier string `xml:"urn:nedu:edsn:data:p4collecteddatabatchresponse:2:standard UniformResourceIdentifier"`
}

type P4CollectedDataBatchResponseEnvelopeEDSNBusinessDocumentHeaderSource struct {
	Authority             string `xml:"urn:nedu:edsn:data:p4collecteddatabatchresponse:2:standard Authority,omitempty"`
	ContactTypeIdentifier string `xml:"urn:nedu:edsn:data:p4collecteddatabatchresponse:2:standard ContactTypeIdentifier,omitempty"`
	SenderID              string `xml:"urn:nedu:edsn:data:p4collecteddatabatchresponse:2:standard SenderID"`
}

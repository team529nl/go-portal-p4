package messages

import "encoding/xml"

type P4CollectedDataBatchResultRequestEnvelope struct {
	XMLName                    xml.Name                                                            `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultrequest:2:standard P4CollectedDataBatchResultRequestEnvelope"`
	EDSNBusinessDocumentHeader P4CollectedDataBatchResultRequestEnvelopeEDSNBusinessDocumentHeader `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultrequest:2:standard EDSNBusinessDocumentHeader"`
	P4Content                  P4CollectedDataBatchResultRequestEnvelopeP4Content                  `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultrequest:2:standard P4Content"`
}

type P4CollectedDataBatchResultRequestEnvelopeP4Content struct {
}

type P4CollectedDataBatchResultRequestEnvelopeEDSNBusinessDocumentHeader struct {
	ContentHash       string                                                                         `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultrequest:2:standard ContentHash,omitempty"`
	ConversationID    string                                                                         `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultrequest:2:standard ConversationID,omitempty"`
	CorrelationID     string                                                                         `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultrequest:2:standard CorrelationID,omitempty"`
	CreationTimestamp *DateTime                                                                      `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultrequest:2:standard CreationTimestamp"`
	DocumentID        string                                                                         `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultrequest:2:standard DocumentID,omitempty"`
	ExpiresAt         *DateTime                                                                      `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultrequest:2:standard ExpiresAt,omitempty"`
	MessageID         string                                                                         `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultrequest:2:standard MessageID"`
	ProcessTypeID     string                                                                         `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultrequest:2:standard ProcessTypeID,omitempty"`
	RepeatedRequest   string                                                                         `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultrequest:2:standard RepeatedRequest,omitempty"`
	TestRequest       string                                                                         `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultrequest:2:standard TestRequest,omitempty"`
	Destination       P4CollectedDataBatchResultRequestEnvelopeEDSNBusinessDocumentHeaderDestination `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultrequest:2:standard Destination"`
	Manifest          *P4CollectedDataBatchResultRequestEnvelopeEDSNBusinessDocumentHeaderManifest   `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultrequest:2:standard Manifest,omitempty"`
	Source            P4CollectedDataBatchResultRequestEnvelopeEDSNBusinessDocumentHeaderSource      `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultrequest:2:standard Source"`
}

type P4CollectedDataBatchResultRequestEnvelopeEDSNBusinessDocumentHeaderDestination struct {
	Receiver P4CollectedDataBatchResultRequestEnvelopeEDSNBusinessDocumentHeaderDestinationReceiver `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultrequest:2:standard Receiver"`
	Service  P4CollectedDataBatchResultRequestEnvelopeEDSNBusinessDocumentHeaderDestinationService  `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultrequest:2:standard Service,omitempty"`
}

type P4CollectedDataBatchResultRequestEnvelopeEDSNBusinessDocumentHeaderDestinationReceiver struct {
	Authority             string `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultrequest:2:standard Authority,omitempty"`
	ContactTypeIdentifier string `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultrequest:2:standard ContactTypeIdentifier,omitempty"`
	ReceiverID            string `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultrequest:2:standard ReceiverID"`
}

type P4CollectedDataBatchResultRequestEnvelopeEDSNBusinessDocumentHeaderDestinationService struct {
	ServiceMethod string `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultrequest:2:standard ServiceMethod,omitempty"`
	ServiceName   string `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultrequest:2:standard ServiceName,omitempty"`
}

type P4CollectedDataBatchResultRequestEnvelopeEDSNBusinessDocumentHeaderManifest struct {
	NumberofItems float64                                                                                   `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultrequest:2:standard NumberofItems"`
	ManifestItem  []P4CollectedDataBatchResultRequestEnvelopeEDSNBusinessDocumentHeaderManifestManifestItem `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultrequest:2:standard ManifestItem"`
}

type P4CollectedDataBatchResultRequestEnvelopeEDSNBusinessDocumentHeaderManifestManifestItem struct {
	Description               string `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultrequest:2:standard Description,omitempty"`
	LanguageCode              string `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultrequest:2:standard LanguageCode,omitempty"`
	MimeTypeQualifierCode     string `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultrequest:2:standard MimeTypeQualifierCode"`
	UniformResourceIdentifier string `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultrequest:2:standard UniformResourceIdentifier"`
}

type P4CollectedDataBatchResultRequestEnvelopeEDSNBusinessDocumentHeaderSource struct {
	Authority             string `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultrequest:2:standard Authority,omitempty"`
	ContactTypeIdentifier string `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultrequest:2:standard ContactTypeIdentifier,omitempty"`
	SenderID              string `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultrequest:2:standard SenderID"`
}

package messages

import "encoding/xml"

type P4CollectedDataBatchResultResponseEnvelope struct {
	XMLName                    xml.Name                                                             `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultresponse:2:standard P4CollectedDataBatchResultResponseEnvelope"`
	EDSNBusinessDocumentHeader P4CollectedDataBatchResultResponseEnvelopeEDSNBusinessDocumentHeader `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultresponse:2:standard EDSNBusinessDocumentHeader"`
	P4Content                  P4CollectedDataBatchResultResponseEnvelopeP4Content                  `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultresponse:2:standard P4Content"`
}

type P4CollectedDataBatchResultResponseEnvelopeP4Content struct {
	P4MeteringPoint []P4CollectedDataBatchResultResponseEnvelopeP4ContentP4MeteringPoint `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultresponse:2:standard P4MeteringPoint,omitempty"`
}

type P4CollectedDataBatchResultResponseEnvelopeP4ContentP4MeteringPoint struct {
	EANID             GSRNEANCode                                                                        `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultresponse:2:standard EANID"`
	ExternalReference TextType                                                                           `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultresponse:2:standard ExternalReference"`
	QueryDate         Date                                                                               `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultresponse:2:standard QueryDate,omitempty"`
	QueryReason       QueryReasonTypeCode                                                                `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultresponse:2:standard QueryReason"`
	P4EnergyMeter     *[]P4CollectedDataBatchResultResponseEnvelopeP4ContentP4MeteringPointP4EnergyMeter `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultresponse:2:standard P4EnergyMeter"`
	P4Rejection       *[]P4CollectedDataBatchResultResponseEnvelopeP4ContentP4MeteringPointP4Rejection   `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultresponse:2:standard P4Rejection"`
}

type P4CollectedDataBatchResultResponseEnvelopeP4ContentP4MeteringPointP4Rejection struct {
	Rejection P4CollectedDataBatchResultResponseEnvelopeRejectionP4Type `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultresponse:2:standard Rejection"`
}

type P4CollectedDataBatchResultResponseEnvelopeRejectionP4Type struct {
	RejectionCode RejectionReasonType `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultresponse:2:standard RejectionCode"`
	RejectionText string              `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultresponse:2:standard RejectionText,omitempty"`
}

type P4CollectedDataBatchResultResponseEnvelopeP4ContentP4MeteringPointP4EnergyMeter struct {
	ID         EnergyMeterIDType                                                                           `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultresponse:2:standard ID"`
	P4Register []P4CollectedDataBatchResultResponseEnvelopeP4ContentP4MeteringPointP4EnergyMeterP4Register `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultresponse:2:standard P4Register"`
}

type P4CollectedDataBatchResultResponseEnvelopeP4ContentP4MeteringPointP4EnergyMeterP4Register struct {
	ID          EnergyRegisterIDType                                                                                 `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultresponse:2:standard ID"`
	MeasureUnit MeasureUnitCode                                                                                      `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultresponse:2:standard MeasureUnit"`
	P4Reading   []P4CollectedDataBatchResultResponseEnvelopeP4ContentP4MeteringPointP4EnergyMeterP4RegisterP4Reading `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultresponse:2:standard P4Reading"`
}

type P4CollectedDataBatchResultResponseEnvelopeP4ContentP4MeteringPointP4EnergyMeterP4RegisterP4Reading struct {
	Reading         float64  `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultresponse:2:standard Reading,omitempty"`
	ReadingDateTime DateTime `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultresponse:2:standard ReadingDateTime"`
}

type P4CollectedDataBatchResultResponseEnvelopeEDSNBusinessDocumentHeader struct {
	ContentHash       string                                                                          `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultresponse:2:standard ContentHash,omitempty"`
	ConversationID    string                                                                          `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultresponse:2:standard ConversationID,omitempty"`
	CorrelationID     string                                                                          `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultresponse:2:standard CorrelationID,omitempty"`
	CreationTimestamp *DateTime                                                                       `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultresponse:2:standard CreationTimestamp"`
	DocumentID        string                                                                          `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultresponse:2:standard DocumentID,omitempty"`
	ExpiresAt         *DateTime                                                                       `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultresponse:2:standard ExpiresAt,omitempty"`
	MessageID         string                                                                          `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultresponse:2:standard MessageID"`
	ProcessTypeID     string                                                                          `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultresponse:2:standard ProcessTypeID,omitempty"`
	RepeatedRequest   string                                                                          `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultresponse:2:standard RepeatedRequest,omitempty"`
	TestRequest       string                                                                          `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultresponse:2:standard TestRequest,omitempty"`
	Destination       P4CollectedDataBatchResultResponseEnvelopeEDSNBusinessDocumentHeaderDestination `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultresponse:2:standard Destination"`
	Manifest          *P4CollectedDataBatchResultResponseEnvelopeEDSNBusinessDocumentHeaderManifest   `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultresponse:2:standard Manifest,omitempty"`
	Source            P4CollectedDataBatchResultResponseEnvelopeEDSNBusinessDocumentHeaderSource      `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultresponse:2:standard Source"`
}

type P4CollectedDataBatchResultResponseEnvelopeEDSNBusinessDocumentHeaderDestination struct {
	Receiver P4CollectedDataBatchResultResponseEnvelopeEDSNBusinessDocumentHeaderDestinationReceiver `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultresponse:2:standard Receiver"`
	Service  P4CollectedDataBatchResultResponseEnvelopeEDSNBusinessDocumentHeaderDestinationService  `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultresponse:2:standard Service,omitempty"`
}

type P4CollectedDataBatchResultResponseEnvelopeEDSNBusinessDocumentHeaderDestinationReceiver struct {
	Authority             string `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultresponse:2:standard Authority,omitempty"`
	ContactTypeIdentifier string `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultresponse:2:standard ContactTypeIdentifier,omitempty"`
	ReceiverID            string `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultresponse:2:standard ReceiverID"`
}

type P4CollectedDataBatchResultResponseEnvelopeEDSNBusinessDocumentHeaderDestinationService struct {
	ServiceMethod string `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultresponse:2:standard ServiceMethod,omitempty"`
	ServiceName   string `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultresponse:2:standard ServiceName,omitempty"`
}

type P4CollectedDataBatchResultResponseEnvelopeEDSNBusinessDocumentHeaderManifest struct {
	NumberofItems float64                                                                                    `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultresponse:2:standard NumberofItems"`
	ManifestItem  []P4CollectedDataBatchResultResponseEnvelopeEDSNBusinessDocumentHeaderManifestManifestItem `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultresponse:2:standard ManifestItem"`
}

type P4CollectedDataBatchResultResponseEnvelopeEDSNBusinessDocumentHeaderManifestManifestItem struct {
	Description               string `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultresponse:2:standard Description,omitempty"`
	LanguageCode              string `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultresponse:2:standard LanguageCode,omitempty"`
	MimeTypeQualifierCode     string `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultresponse:2:standard MimeTypeQualifierCode"`
	UniformResourceIdentifier string `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultresponse:2:standard UniformResourceIdentifier"`
}

type P4CollectedDataBatchResultResponseEnvelopeEDSNBusinessDocumentHeaderSource struct {
	Authority             string `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultresponse:2:standard Authority,omitempty"`
	ContactTypeIdentifier string `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultresponse:2:standard ContactTypeIdentifier,omitempty"`
	SenderID              string `xml:"urn:nedu:edsn:data:p4collecteddatabatchresultresponse:2:standard SenderID"`
}

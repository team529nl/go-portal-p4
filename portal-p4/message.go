package portal_p4

import (
	"github.com/google/uuid"
	"team529.nl/go-portal-p4/portal-p4/messages"
	"time"
)

type Clock interface {
	Now() time.Time
}

type systemClock struct{}

func (c systemClock) Now() time.Time {
	return time.Now()
}

type IdGenerator interface {
	Generate() string
}

type uuidGenerator struct{}

func (g uuidGenerator) Generate() string {
	return uuid.New().String()
}

type constantId struct {
	Id string
}

func (g constantId) Generate() string {
	return g.Id
}

type Date struct {
	Year  int
	Month time.Month
	Day   int
}

//type QueryReason int
//
//const (
//	Dagstand QueryReason = iota
//	Intervalstand
//	Maandstand_recovery
//)
//
//func (qr QueryReason) String() string {
//	return [...]string{"DAY", "INT", "RCY"}[qr]
//}

type MessageFactory interface {
	DataRequest(eanCodes []messages.GSRNEANCode, dates []Date, reason messages.QueryReasonTypeCode) messages.P4CollectedDataBatchRequestEnvelope
	DataResultRequest() messages.P4CollectedDataBatchResultRequestEnvelope
}

type DefaultMessageFactory struct {
	sender      string
	receiver    string
	clock       Clock
	IdGenerator IdGenerator
	reference   string
}

func NewMessageFactory(sender string, receiver string, reference string) MessageFactory {
	return DefaultMessageFactory{
		sender:      sender,
		receiver:    receiver,
		clock:       new(systemClock),
		IdGenerator: new(uuidGenerator),
		reference:   reference,
	}
}

func NewConstantMessageIdMessageFactory(sender string, receiver string, reference string, messageId string) MessageFactory {
	return DefaultMessageFactory{
		sender:      sender,
		receiver:    receiver,
		clock:       new(systemClock),
		IdGenerator: constantId{Id: messageId},
		reference:   reference,
	}
}

func (f DefaultMessageFactory) dataRequestHeader() messages.P4CollectedDataBatchRequestEnvelopeEDSNBusinessDocumentHeader {
	ts := messages.DateTime(f.clock.Now())

	return messages.P4CollectedDataBatchRequestEnvelopeEDSNBusinessDocumentHeader{
		CreationTimestamp: &ts,
		MessageID:         f.IdGenerator.Generate(),
		Destination: messages.P4CollectedDataBatchRequestEnvelopeEDSNBusinessDocumentHeaderDestination{
			Receiver: messages.P4CollectedDataBatchRequestEnvelopeEDSNBusinessDocumentHeaderDestinationReceiver{
				ReceiverID: f.receiver,
			},
		},
		Source: messages.P4CollectedDataBatchRequestEnvelopeEDSNBusinessDocumentHeaderSource{
			SenderID: f.sender,
		},
	}
}

func (f DefaultMessageFactory) dataResultHeader() messages.P4CollectedDataBatchResultRequestEnvelopeEDSNBusinessDocumentHeader {
	ts := messages.DateTime(f.clock.Now())

	return messages.P4CollectedDataBatchResultRequestEnvelopeEDSNBusinessDocumentHeader{
		CreationTimestamp: &ts,
		MessageID:         f.IdGenerator.Generate(),
		Destination: messages.P4CollectedDataBatchResultRequestEnvelopeEDSNBusinessDocumentHeaderDestination{
			Receiver: messages.P4CollectedDataBatchResultRequestEnvelopeEDSNBusinessDocumentHeaderDestinationReceiver{
				ReceiverID: f.receiver,
			},
		},
		Source: messages.P4CollectedDataBatchResultRequestEnvelopeEDSNBusinessDocumentHeaderSource{
			SenderID: f.sender,
		},
	}
}

func (f DefaultMessageFactory) dataRequestMeteringPoint(ean messages.GSRNEANCode, date Date, reason messages.QueryReasonTypeCode) messages.P4CollectedDataBatchRequestEnvelopeP4ContentP4MeteringPoint {

	return messages.P4CollectedDataBatchRequestEnvelopeP4ContentP4MeteringPoint{
		EANID:             ean,
		ExternalReference: messages.TextType(f.reference),
		QueryDate:         messages.Date(time.Date(date.Year, date.Month, date.Day, 0, 0, 0, 0, time.Local)),
		QueryReason:       reason,
	}
}

func (f DefaultMessageFactory) DataRequest(eanCodes []messages.GSRNEANCode, dates []Date, reason messages.QueryReasonTypeCode) messages.P4CollectedDataBatchRequestEnvelope {
	var meteringPoints []messages.P4CollectedDataBatchRequestEnvelopeP4ContentP4MeteringPoint
	for _, ean := range eanCodes {
		for _, date := range dates {
			meteringPoints = append(meteringPoints, f.dataRequestMeteringPoint(ean, date, reason))
		}
	}

	return messages.P4CollectedDataBatchRequestEnvelope{
		EDSNBusinessDocumentHeader: f.dataRequestHeader(),
		P4Content:                  messages.P4CollectedDataBatchRequestEnvelopeP4Content{P4MeteringPoint: meteringPoints},
	}
}

func (f DefaultMessageFactory) DataResultRequest() messages.P4CollectedDataBatchResultRequestEnvelope {
	return messages.P4CollectedDataBatchResultRequestEnvelope{
		EDSNBusinessDocumentHeader: f.dataResultHeader(),
	}
}

package messages

import (
	"encoding/json"
	"time"
)

type DateTime time.Time

func (t *DateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t DateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

type Date time.Time

func (t *Date) UnmarshalText(text []byte) error {
	return (*xsdDate)(t).UnmarshalText(text)
}
func (t Date) MarshalText() ([]byte, error) {
	return xsdDate(t).MarshalText()
}

func (t Date) String() string {
	return (time.Time)(t).Format("2006-01-02")
}

// Must match the pattern [0-9]{18}
type GSRNEANCode string

// Must be at least 1 char long
type TextType string

type QueryReasonTypeCode int

const (
	Dagstand QueryReasonTypeCode = iota + 1
	Intervalstand
	Maandstand_recovery
)

var queryReasonTypeCodeToString = map[QueryReasonTypeCode]string{
	Dagstand:            "DAY",
	Intervalstand:       "INT",
	Maandstand_recovery: "RCY",
}

var queryReasonTypeCodeToDescription = map[QueryReasonTypeCode]string{
	Dagstand:            "Dagstand",
	Intervalstand:       "Intervalstand",
	Maandstand_recovery: "Maandstand",
}

var stringToQueryReasonTypeCode = map[string]QueryReasonTypeCode{
	"DAY": Dagstand,
	"INT": Intervalstand,
	"RCY": Maandstand_recovery,
}

func (t *QueryReasonTypeCode) UnmarshalText(text []byte) error {
	s := stringToQueryReasonTypeCode[string(text)]
	*t = s
	return nil
}

func (t QueryReasonTypeCode) MarshalText() ([]byte, error) {
	return []byte(queryReasonTypeCodeToString[t]), nil
}

func (t QueryReasonTypeCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(queryReasonTypeCodeToDescription[t])
}

func (t QueryReasonTypeCode) String() string {
	return queryReasonTypeCodeToDescription[t]
}

// May be no more than 3 chars long
type RejectionReasonType string

// Must be at least 1 char long
type EnergyMeterIDType string

// Must be at least 1 char long
type EnergyRegisterIDType int

const (
	VerbruikTotaal EnergyRegisterIDType = iota + 1
	VerbruikLaag
	VerbruikNormaal
	TerugleveringTotaal
	TerugleveringLaag
	TerugleveringNormaal
)

var energyRegisterTypeToString = map[EnergyRegisterIDType]string{
	VerbruikTotaal:       "1.8.0",
	VerbruikLaag:         "1.8.1",
	VerbruikNormaal:      "1.8.2",
	TerugleveringTotaal:  "2.8.0",
	TerugleveringLaag:    "2.8.1",
	TerugleveringNormaal: "2.8.2",
}

var energyRegisterTypeToDescription = map[EnergyRegisterIDType]string{
	VerbruikTotaal:       "Verbruik Totaal",
	VerbruikLaag:         "Verbruik Laag",
	VerbruikNormaal:      "Verbruik Normaal",
	TerugleveringTotaal:  "Teruglevering Totaal",
	TerugleveringLaag:    "Teruglevering Laag",
	TerugleveringNormaal: "Teruglevering Normaal",
}

var stringToEnergyRegisterType = map[string]EnergyRegisterIDType{
	"1.8.0": VerbruikTotaal,
	"1.8.1": VerbruikLaag,
	"1.8.2": VerbruikNormaal,
	"2.8.0": TerugleveringTotaal,
	"2.8.1": TerugleveringLaag,
	"2.8.2": TerugleveringNormaal,
}

func (t *EnergyRegisterIDType) UnmarshalText(text []byte) error {
	s := stringToEnergyRegisterType[string(text)]
	*t = s
	return nil
}

func (t EnergyRegisterIDType) MarshalText() ([]byte, error) {
	return []byte(energyRegisterTypeToString[t]), nil
}

func (t EnergyRegisterIDType) MarshalJSON() ([]byte, error) {
	return json.Marshal(energyRegisterTypeToDescription[t])
}

func (t EnergyRegisterIDType) String() string {
	return energyRegisterTypeToDescription[t]
}

type MeasureUnitCode int

const (
	KVAR MeasureUnitCode = iota + 1
	KWH
	KW
	NormaalM3
	M3
	WH
	DM3
	MJ
	GroningenM3
	GroningenM3_per_Hour
	M3_per_hour
)

var measureUnitCodeToString = map[MeasureUnitCode]string{
	KVAR:                 "KVR",
	KWH:                  "KWH",
	KW:                   "KWT",
	NormaalM3:            "M3N",
	M3:                   "MTQ",
	WH:                   "WH",
	DM3:                  "DM3",
	MJ:                   "MJ",
	GroningenM3:          "M3N3517",
	GroningenM3_per_Hour: "M3N3517HR",
	M3_per_hour:          "MQH",
}

var measureUnitCodeToDescription = map[MeasureUnitCode]string{
	KVAR:                 "kVAR",
	KWH:                  "kWh",
	KW:                   "KW",
	NormaalM3:            "NormaalM3",
	M3:                   "m3",
	WH:                   "wh",
	DM3:                  "dm3",
	MJ:                   "MJ",
	GroningenM3:          "m3 (Groningen) ",
	GroningenM3_per_Hour: "m3/h (Groningen)",
	M3_per_hour:          "m3/h",
}

var stringToMeasureUnitCode = map[string]MeasureUnitCode{
	"KVR":       KVAR,
	"KWH":       KWH,
	"KWT":       KW,
	"M3N":       NormaalM3,
	"MTQ":       M3,
	"WH":        WH,
	"DM3":       DM3,
	"MJ":        MJ,
	"M3N3517":   GroningenM3,
	"M3N3517HR": GroningenM3_per_Hour,
	"MQH":       M3_per_hour,
}

func (t *MeasureUnitCode) UnmarshalText(text []byte) error {
	s := stringToMeasureUnitCode[string(text)]
	*t = s
	return nil
}

func (t MeasureUnitCode) MarshalText() ([]byte, error) {
	return []byte(measureUnitCodeToString[t]), nil
}

func (t MeasureUnitCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(measureUnitCodeToDescription[t])
}

func (t MeasureUnitCode) String() string {
	return measureUnitCodeToDescription[t]
}

package messages

import (
	"bytes"
	"encoding/xml"
	"time"
)

func _unmarshalTime(text []byte, t *time.Time, format string) (err error) {
	s := string(bytes.TrimSpace(text))
	*t, err = time.Parse(format, s)
	return err
}

type xsdDateTime time.Time

func (t *xsdDateTime) UnmarshalText(text []byte) error {
	var _t time.Time
	err := _unmarshalTime(text, &_t, "2006-01-02T15:04:05.999Z")
	if err != nil {
		err = _unmarshalTime(text, &_t, "2006-01-02T15:04:05.999-07:00")
	}
	if err != nil {
		return err
	}
	*t = (xsdDateTime)(_t.Local())
	return nil
}

func (t xsdDateTime) MarshalText() ([]byte, error) {
	return []byte((time.Time)(t).UTC().Format("2006-01-02T15:04:05Z")), nil
}

type xsdDate time.Time

func (t *xsdDate) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006-01-02")
}

func (t xsdDate) MarshalText() ([]byte, error) {
	return []byte((time.Time)(t).Format("2006-01-02")), nil
}

func (t xsdDate) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}

func (t xsdDate) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}

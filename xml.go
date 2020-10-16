package go_portal_p4

import (
	"bytes"
	"encoding/xml"
	"io"
)

func ObjToXML(obj interface{}) *bytes.Buffer {
	var buf bytes.Buffer
	enc := xml.NewEncoder(&buf)
	enc.Encode(obj)
	return &buf
}

func ObjFromXML(data io.Reader, target interface{}) error {

	if err := AssertPointer(target, "Non pointer passed to ObjFromXML"); err != nil {
		return err
	}

	dec := xml.NewDecoder(data)
	if err := dec.Decode(&target); err != nil {
		return err
	}
	return nil
}

func ObjFromXMLString(s string, target interface{}) error {

	buffer := bytes.Buffer{}
	buffer.WriteString(s)
	return ObjFromXML(&buffer, target)
}

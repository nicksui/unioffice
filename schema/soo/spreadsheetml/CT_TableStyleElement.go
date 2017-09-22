// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetml

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type CT_TableStyleElement struct {
	// Table Style Type
	TypeAttr ST_TableStyleType
	// Band Size
	SizeAttr *uint32
	// Formatting Id
	DxfIdAttr *uint32
}

func NewCT_TableStyleElement() *CT_TableStyleElement {
	ret := &CT_TableStyleElement{}
	ret.TypeAttr = ST_TableStyleType(1)
	return ret
}

func (m *CT_TableStyleElement) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	attr, err := m.TypeAttr.MarshalXMLAttr(xml.Name{Local: "type"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	if m.SizeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "size"},
			Value: fmt.Sprintf("%v", *m.SizeAttr)})
	}
	if m.DxfIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "dxfId"},
			Value: fmt.Sprintf("%v", *m.DxfIdAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TableStyleElement) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.TypeAttr = ST_TableStyleType(1)
	for _, attr := range start.Attr {
		if attr.Name.Local == "type" {
			m.TypeAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "size" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.SizeAttr = &pt
		}
		if attr.Name.Local == "dxfId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.DxfIdAttr = &pt
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_TableStyleElement: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_TableStyleElement and its children
func (m *CT_TableStyleElement) Validate() error {
	return m.ValidateWithPath("CT_TableStyleElement")
}

// ValidateWithPath validates the CT_TableStyleElement and its children, prefixing error messages with path
func (m *CT_TableStyleElement) ValidateWithPath(path string) error {
	if m.TypeAttr == ST_TableStyleTypeUnset {
		return fmt.Errorf("%s/TypeAttr is a mandatory field", path)
	}
	if err := m.TypeAttr.ValidateWithPath(path + "/TypeAttr"); err != nil {
		return err
	}
	return nil
}
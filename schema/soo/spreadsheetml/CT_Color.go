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

type CT_Color struct {
	// Automatic
	AutoAttr *bool
	// Index
	IndexedAttr *uint32
	// Alpha Red Green Blue Color Value
	RgbAttr *string
	// Theme Color
	ThemeAttr *uint32
	// Tint
	TintAttr *float64
}

func NewCT_Color() *CT_Color {
	ret := &CT_Color{}
	return ret
}

func (m *CT_Color) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.AutoAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "auto"},
			Value: fmt.Sprintf("%d", b2i(*m.AutoAttr))})
	}
	if m.IndexedAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "indexed"},
			Value: fmt.Sprintf("%v", *m.IndexedAttr)})
	}
	if m.RgbAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "rgb"},
			Value: fmt.Sprintf("%v", *m.RgbAttr)})
	}
	if m.ThemeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "theme"},
			Value: fmt.Sprintf("%v", *m.ThemeAttr)})
	}
	if m.TintAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "tint"},
			Value: fmt.Sprintf("%v", *m.TintAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Color) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "auto" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.AutoAttr = &parsed
		}
		if attr.Name.Local == "indexed" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.IndexedAttr = &pt
		}
		if attr.Name.Local == "rgb" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RgbAttr = &parsed
		}
		if attr.Name.Local == "theme" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.ThemeAttr = &pt
		}
		if attr.Name.Local == "tint" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return err
			}
			m.TintAttr = &parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_Color: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_Color and its children
func (m *CT_Color) Validate() error {
	return m.ValidateWithPath("CT_Color")
}

// ValidateWithPath validates the CT_Color and its children, prefixing error messages with path
func (m *CT_Color) ValidateWithPath(path string) error {
	return nil
}
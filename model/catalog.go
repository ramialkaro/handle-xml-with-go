package model

import "encoding/xml"

// Catalog struct
type Catalog struct {
	XMLName xml.Name `xml:"catalog"`
	Books   []Book   `xml:"book"`
}

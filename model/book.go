package model

import "encoding/xml"

// Book struct
type Book struct {
	XMLName     xml.Name `xml:"book"`
	BookID      string   `xml:"id,attr"`
	Author      string   `xml:"author"`
	Title       string   `xml:"title"`
	Genre       string   `xml:"genre"`
	Price       float64  `xml:"price"`
	PublishDate string   `xml:"publish_date"`
	Description string   `xml:"description"`
}

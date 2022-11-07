package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"

	"github.com/ramialkaro/handle-xml-with-go/model"
)

func main() {
	// open the Books.xml file
	xmlFile, err := os.Open("./books.xml")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("\tSuccessfully opened books.xml file")
	defer xmlFile.Close()

	byteValue, _ := io.ReadAll(xmlFile)

	//unmarshal takes a []byte and fills the catalog struct with the values found in the xmlFile
	catalog := &model.Catalog{}
	xml.Unmarshal(byteValue, &catalog)

	for _, book := range catalog.Books {

		fmt.Println(book.Title)
	}
}

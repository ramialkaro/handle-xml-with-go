package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"

	"github.com/ramialkaro/handle-xml-with-go/model"
	"github.com/jedib0t/go-pretty/v6/table"

)

func main() {
	// open the Books.xml file
	xmlFile, err := os.Open("./books.xml")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Successfully opened books.xml file")
	defer xmlFile.Close()

	byteValue, _ := io.ReadAll(xmlFile)

	//unmarshal takes a []byte and fills the catalog struct with the values found in the xmlFile
	catalog := &model.Catalog{}
	xml.Unmarshal(byteValue, &catalog)
 
 	t := table.NewWriter()
  t.SetOutputMirror(os.Stdout)
  t.AppendHeader(table.Row{"#", "Author", "Title", "Genre", "Price", "Publish Date", "Description"})
   
	for _, book := range catalog.Books {

	 t.AppendRows([]table.Row{
        {book.BookID, book.Author, book.Title, book.Genre, book.Price, book.PublishDate, book.Description},
    })
	}

    t.AppendSeparator()
    t.Render()
}

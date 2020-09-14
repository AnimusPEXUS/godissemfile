package main

import (
	"log"

	"github.com/antchfx/htmlquery"

	"github.com/AnimusPEXUS/godissemfile"
)

func main() {
	f := godissemfile.NewDissemFile()
	err := f.LoadFile("0000728889-19-000206.dissem")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Preamble:", f.Preamble)
	log.Println("Preamble (str):", string(f.Preamble))
	log.Println("Attributes:")

	log.Println(htmlquery.OutputHTML(f.Attributes, true))

	r, err := htmlquery.Query(f.Attributes, "//effectiveness-date")
	if err != nil {
		log.Fatalln(err)
	}

	if r != nil {
		log.Println("effectiveness-date", r.FirstChild.Data)
	}

	log.Println("Documents:", len(f.Documents))
}

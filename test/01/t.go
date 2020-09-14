package main

import (
	"log"

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
	for _, i := range f.Attributes {
		log.Println(" ", i.Name, "=", i.Value)
	}
	log.Println("Documents:", len(f.Documents))
}

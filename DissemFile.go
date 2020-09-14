package godissemfile

import (
	"errors"
	"io/ioutil"

	"golang.org/x/net/html"
)

var ERR_INVALID_DOCUMENT = errors.New("invalid document")

type DissemFile struct {
	Preamble   []byte
	Attributes *html.Node
	Documents  []*DissemDocument
}

func NewDissemFile() *DissemFile {
	self := &DissemFile{}
	self.Init()
	return self
}

func (self *DissemFile) Init() {
	// self.Attributes = make([]*DissemAttr, 0)
	self.Documents = make([]*DissemDocument, 0)
}

func (self *DissemFile) LoadFile(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return self.LoadData(data)
}

func (self *DissemFile) LoadData(data []byte) error {
	self.Init()

	err := self.sliceDocuments(data)
	if err != nil {
		return err
	}

	err = self.slicePreamble(data)
	if err != nil {
		return err
	}

	err = self.sliceAttributes(data)
	if err != nil {
		return err
	}

	return nil
}

func (self *DissemFile) slicePreamble(data []byte) error {
	sub0, _ := FindOSubmission(data)
	if sub0 == -1 {
		return ERR_INVALID_DOCUMENT
	}

	new_data := make([]byte, sub0)
	copy(new_data, data[:sub0])

	self.Preamble = new_data

	return nil
}

func (self *DissemFile) sliceAttributes(data []byte) error {

	var subdata []byte

	{
		sub0, sub1 := FindOSubmission(data)
		if sub0 == -1 {
			return ERR_INVALID_DOCUMENT
		}

		first_document, _ := FindODocument(data)
		if first_document == -1 {
			return ERR_INVALID_DOCUMENT
		}

		subdata = data[sub1:first_document]
	}

	a, err := AttributesFromData(subdata)
	if err != nil {
		return err
	}

	self.Attributes = a

	return nil
}

func (self *DissemFile) sliceDocuments(data_o []byte) error {

	self.Init()

	data := make([]byte, len(data_o))
	copy(data, data_o)

	this_is_last := false
	for {
		x0, x1 := FindODocument(data)
		if x0 == -1 {
			break
		}

		data = data[x1:]

		y0, y1 := FindCDocument(data)

		if y0 == -1 {
			this_is_last = true
			y0 = len(data)
		}

		new_data := make([]byte, y0)
		copy(new_data, data[:y0])

		doc := &DissemDocument{}
		err := doc.LoadData(new_data)
		if err != nil {
			return err
		}

		self.Documents = append(self.Documents, doc)

		if this_is_last {
			break
		}

		data = data[y1:]
	}

	return nil
}

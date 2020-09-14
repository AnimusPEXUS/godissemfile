package godissemfile

import "log"

type DissemDocument struct {
	Attributes []*DissemAttr
	Text       []byte
}

func (self *DissemDocument) LoadData(data []byte) error {
	err := self.sliceAttributes(data)
	if err != nil {
		return err
	}

	log.Println("doc attributes")
	for ii, i := range self.Attributes {
		log.Println(ii, i)
	}

	return nil
}

func (self *DissemDocument) sliceAttributes(data []byte) error {

	t0, _ := FindOText(data)
	x := data[:t0]

	self.Attributes = AttributesFromData(x)

	return nil
}

package godissemfile

type DissemDocument struct {
	Attributes []*DissemAttr
	Text       []byte
}

func (self *DissemDocument) LoadData(data []byte) error {
	err := self.sliceAttributes(data)
	if err != nil {
		return err
	}

	// log.Println("doc attributes")
	// for ii, i := range self.Attributes {
	// 	log.Println(ii, i)
	// }

	err = self.sliceText(data)
	if err != nil {
		return err
	}

	return nil
}

func (self *DissemDocument) sliceAttributes(data []byte) error {

	t0, _ := FindOText(data)
	x := data[:t0]

	self.Attributes = AttributesFromData(x)

	return nil
}

func (self *DissemDocument) sliceText(data []byte) error {

	_, t1 := FindOText(data)

	t2, _ := FindCText(data)

	x := data[t1:t2]
	self.Text = make([]byte, len(x))
	copy(self.Text, x)

	return nil
}

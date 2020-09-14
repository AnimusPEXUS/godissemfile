package godissemfile

import (
	"bytes"
	"regexp"
	"strings"
)

var (
	SUBMISSION_O_TXT     = []byte(`<SUBMISSION>`)
	SUBMISSION_O_TXT_LEN = len(SUBMISSION_O_TXT)

	TEXT_O_TXT     = []byte(`<TEXT>`)
	TEXT_O_TXT_LEN = len(TEXT_O_TXT)

	DOCUMENT_O_TXT     = []byte(`<DOCUMENT>`)
	DOCUMENT_O_TXT_LEN = len(DOCUMENT_O_TXT)

	DOCUMENT_C_TXT     = []byte(`</DOCUMENT>`)
	DOCUMENT_C_TXT_LEN = len(DOCUMENT_C_TXT)
)

func FindOSubmission(s []byte) (int, int) {
	ret := bytes.Index(s, SUBMISSION_O_TXT)
	ret2 := ret + SUBMISSION_O_TXT_LEN
	return ret, ret2
}

func FindOText(s []byte) (int, int) {
	ret := bytes.Index(s, TEXT_O_TXT)
	ret2 := ret + TEXT_O_TXT_LEN
	return ret, ret2
}

func FindODocument(s []byte) (int, int) {
	ret := bytes.Index(s, DOCUMENT_O_TXT)
	ret2 := ret + DOCUMENT_O_TXT_LEN
	return ret, ret2
}

func FindCDocument(s []byte) (int, int) {
	ret := bytes.Index(s, DOCUMENT_C_TXT)
	ret2 := ret + DOCUMENT_C_TXT_LEN
	return ret, ret2
}

var RE_ATTR_LINE_C = regexp.MustCompile(`\<(?P<name>[A-Za-z][A-Za-z0-9_\-]*)\>(?P<value>.*)`)

func AttributesFromData(data []byte) []*DissemAttr {

	// TODO: use htmlquery.Parse ?

	ret := make([]*DissemAttr, 0)

	var attrs_block_str_split = strings.Split(string(data), "\n")

	for _, i := range attrs_block_str_split {
		re_f_sm := RE_ATTR_LINE_C.FindStringSubmatch(i)
		if len(re_f_sm) == 0 {
			continue
		}

		var name string
		var value string

		for ii, i := range RE_ATTR_LINE_C.SubexpNames() {
			switch i {
			case "name":
				name = re_f_sm[ii]
			case "value":
				value = re_f_sm[ii]
			}
		}

		if len(re_f_sm) >= 3 {
			a := &DissemAttr{
				Name:  strings.TrimSpace(name),
				Value: strings.TrimSpace(value),
			}
			ret = append(ret, a)
		}
	}

	return ret
}

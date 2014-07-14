package dom

/*
 * CharacterData node implementation
 *
 * Copyright (c) 2009, Rob Russell
 * Copyright (c) 2010, Jeff Schiller
 * Copyright (c) 2011, Robert Johnstone
 */

import (
	"encoding/xml"
)

type _cdata struct {
	*_node
	content []byte
}

func (cd *_cdata) NodeName() (s string) {
	return "#cdata-section"
}

func (cd *_cdata) NodeValue() (s string) {
	return string(cd.content)
}

func (cd *_cdata) OwnerDocument() Document {
	return ownerDocument(cd)
}

func (cd *_cdata) Length() uint32 {
	return uint32(len(cd.content))
}

func (cd *_cdata) GetData() string {
	return cd.NodeValue()
}

func (cd *_cdata) SetData(newData string) {
	cd.content = []byte(newData)
}

func newCData(token xml.CharData) *_cdata {
	n := newNode(CDATA_SECTION_NODE)
	cd := &_cdata{n, token.Copy()}
	n.self = Node(cd)
	return cd
}

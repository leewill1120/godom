package dom

/*
 * Document interface implementation
 *
 * Copyright (c) 2009, Rob Russell
 * Copyright (c) 2010, Jeff Schiller
 */

import (
	"encoding/xml"
	"os"
)

type _doc struct {
	*_node
}

func (d *_doc) NodeValue() string {
	return ""
}

func (d *_doc) AppendChild(c Node) Node {
	return appendChild(d, c)
}

func (d *_doc) RemoveChild(c Node) Node {
	return removeChild(d, c)
}

func (d *_doc) DocumentElement() Element {
	return d.ChildNodes().Item(0).(Element)
}

func (d *_doc) OwnerDocument() Document {
	return ownerDocument(d)
}

func (d *_doc) CreateElement(tag string) Element {
	return newElem(xml.StartElement{xml.Name{"", tag}, nil})
}

func (d *_doc) CreateTextNode(data string) Text {
	return newText(xml.CharData([]byte(data)))
}

func (d *_doc) CreateAttribute(name string) Attr {
	return newAttr(name, "", nil)
}

func (d *_doc) setRoot(r Element) Element {
	// empty the children vector
	if d.ChildNodes().Length() > 0 {
		os.Exit(-1)
	}
	appendChild(d, r)
	return r
}

func (d *_doc) GetElementById(id string) Element {
	return getElementById(d.DocumentElement(), id)
}

func (d *_doc) GetElementsByTagName(tagName string) NodeList {
	return newTagNodeList(d, tagName)
}

func newDoc() *_doc {
	n := newNode(DOCUMENT_NODE)
	d := &_doc{n}
	n.self = Node(n)
	return d
}

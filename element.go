package dom

/*
 * Element implementation
 *
 * Copyright (c) 2009, Rob Russell
 * Copyright (c) 2010, Jeff Schiller
 */

import (
	"encoding/xml"
)

type _elem struct {
	*_node
	attribs map[string]*_attr // attributes of the element
}

func (e *_elem) NodeValue() string {
	return ""
}

func (e *_elem) AppendChild(c Node) Node {
	return appendChild(e, c)
}

func (e *_elem) RemoveChild(c Node) Node {
	return removeChild(e, c)
}

func (e *_elem) OwnerDocument() Document {
	return ownerDocument(e)
}

func (e *_elem) TagName() string {
	return e.NodeName()
}

func (e *_elem) Attributes() NamedNodeMap {
	return newAttrNamedNodeMap(e)
}

func (e *_elem) GetElementById(id string) Element {
	return getElementById(e, id).(Element)
}

func (e *_elem) GetAttribute(name string) string {
	attr, ok := e.attribs[name]
	if ok {
		return attr.GetValue()
	}
	return ""
}

func (e *_elem) GetAttributeNode(attrName string) Attr {
	attr, ok := e.attribs[attrName]
	if ok {
		return attr
	}
	return nil
}

func (e *_elem) SetAttribute(attrName string, attrVal string) {
	attr, ok := e.attribs[attrName]
	if !ok {
		e.attribs[attrName] = newAttr(attrName, attrVal, e)
	} else {
		attr.value = attrVal
	}
}

func (e *_elem) SetAttributeNode(newAttr Attr) Attr {
	oldAttr, ok := e.attribs[newAttr.Name()]
	// New attribute must not have an owner element.
	if newAttr.OwnerElement() == nil {
		var a *_attr = newAttr.(*_attr)
		e.attribs[newAttr.Name()] = a
		if ok {
			return oldAttr
		}
	}
	return nil
}

// http://www.w3.org/TR/DOM-Level-3-Core/core.html#ID-6D6AC0F9
func (e *_elem) RemoveAttribute(name string) {
	delete(e.attribs, name)
}

func (e *_elem) RemoveAttributeNode(oldAttr Attr) Attr {
	for name, attr := range e.attribs {
		if attr == oldAttr {
			delete(e.attribs, name)
			return oldAttr
		}
	}
	return nil
}

// http://www.w3.org/TR/DOM-Level-3-Core/core.html#ID-ElHasAttr
func (e *_elem) HasAttribute(name string) bool {
	_, has := e.attribs[name]
	return has
}

func (e *_elem) GetElementsByTagName(name string) NodeList {
	return newTagNodeList(e, name)
}

func newElem(token xml.StartElement) *_elem {
	n := newNode(ELEMENT_NODE)
	n.n = token.Name
	e := &_elem{n, make(map[string]*_attr)}
	n.self = Node(e)
	return e
}

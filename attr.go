package dom

/*
* Attr implementation
*
* Copyright (c) 2010, Jeff Schiller
*/

import "encoding/xml"

type _attr struct {
  *_node
  value string // value (for attr)
  ownerElement *_elem
}

func (a *_attr) NodeValue() string {
  return a.value
}

func (a *_attr) AppendChild(n Node) Node {
  return n
}

func (a *_attr) RemoveChild(n Node) Node {
  return n
}

func (a *_attr) ParentNode() Node {
  return Node(nil)
}

func (a *_attr) OwnerDocument() Document {
  return ownerDocument(a)
}

func (a *_attr) ChildNodes() NodeList {
  return NodeList(nil)
}

func (a *_attr) Attributes() NamedNodeMap {
  return NamedNodeMap(nil)
}

func (a *_attr) GetValue() string {
  return a.NodeValue()
}

func (a *_attr) SetValue(newValue string) {
  a.value = newValue
}

func (a *_attr) Name() string {
  return a.NodeName()
}

func (a *_attr) OwnerElement() Element {
  if a.ownerElement == nil {
    return nil;
  }
  return a.ownerElement;
}

func newAttr(name string, val string, owner *_elem) (*_attr) {
  node := newNode(ATTRIBUTE_NODE)
  node.n = xml.Name{"", name}
  a := &_attr { _node: node, value: val, ownerElement: owner }
  node.self = Node(a)
  return a
}

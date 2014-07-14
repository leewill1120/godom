package dom

/*
 * NamedNodeMap implementations.
 *
 * Copyright (c) 2010, Jeff Schiller
 */

// used to return the live attributes of a node
type _attrnamednodemap struct {
  e *_elem
}

func (m *_attrnamednodemap) Length() uint {
  return uint(len(m.e.attribs))
}

func (m *_attrnamednodemap) Item(index uint) Node {
  if index >= 0 && index < m.Length() {
    i := uint(0)
    for _, attr := range(m.e.attribs) {
      if i == index {
        return attr
      }
      i += 1
    }
  }
  return Node(nil)
}

func (m *_attrnamednodemap) GetNamedItem(name string) Node {
  for attrName, attr := range(m.e.attribs) {
    if name == attrName {
      return attr
    }
  }
  return Node(nil)
}

func (m *_attrnamednodemap) SetNamedItem(arg Node) Node {
  if arg != nil && arg.NodeType() == ATTRIBUTE_NODE {
    return m.e.SetAttributeNode(arg.(Attr))
  }
  return nil
}

func (m *_attrnamednodemap) RemoveNamedItem(name string) Node {
  if name != "" {
    attr := m.e.GetAttributeNode(name)
    if attr != nil {
      return m.e.RemoveAttributeNode(attr)
    }
  }
  return nil
}

func newAttrNamedNodeMap(e *_elem) (*_attrnamednodemap) {
  nm := new(_attrnamednodemap)
  nm.e = e
  return nm
}

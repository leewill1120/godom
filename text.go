package dom

/*
 * Text node implementation
 *
 * Copyright (c) 2009, Rob Russell
 * Copyright (c) 2010, Jeff Schiller
 */

import (
	"encoding/xml"
)

type _text struct {
	_cdata
}

func (t *_text) NodeName() (s string) {
	return "#text"
}

func newText(token xml.CharData) *_text {
	n := newNode(TEXT_NODE)
	t := &_text{_cdata{n, token.Copy()}}
	n.self = Node(t)
	return t
}

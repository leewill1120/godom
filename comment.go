package dom

/*
 * Comment node implementation
 *
 * Copyright (c) 2011, Jeff Schiller
 */

import (
	"encoding/xml"
)

type _comment struct {
	_cdata
}

func (c *_comment) NodeName() (s string) {
	return "#comment"
}

func newComment(token xml.Comment) *_comment {
	n := newNode(COMMENT_NODE)
	c := &_comment{_cdata{n, token.Copy()}}
	n.self = Node(c)
	return c
}

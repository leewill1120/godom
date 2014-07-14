package dom

/*
 * Part of the xml/dom Go package
 *
 * Declares the constants used to identify various node
 * types.
 *
 * Copyright (c) 2011, Robert Johnstone
 */ 

const (
  ELEMENT_NODE = iota + 1
  ATTRIBUTE_NODE
  TEXT_NODE
  CDATA_SECTION_NODE
  ENTITY_REFERENCE_NODE
  ENTITY_NODE
  PROCESSING_INSTRUCTION_NODE
  COMMENT_NODE
  DOCUMENT_NODE
  DOCUMENT_TYPE_NODE
  DOCUMENT_FRAGMENT_NODE
  NOTATION_NODE
)


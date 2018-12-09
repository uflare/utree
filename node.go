package main

// Node - represents a node
type Node struct {
	Node string `json:"node,omitempty"`
	Tree []Node `json:"tree,omitempty"`
}

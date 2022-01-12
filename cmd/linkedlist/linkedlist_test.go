package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitLinkedList(t *testing.T) {
	res := initLinkedList()
	assert.Equal(t, res.len, 0)
}

func TestAddFrontLinkedList(t *testing.T) {
	ls := initLinkedList()
	ls.AddFront("Bob")
	assert.Equal(t, ls.head.name, "Bob")
	ls.AddFront("John")
	assert.Equal(t, ls.head.name, "John")
	assert.Equal(t, ls.head.next.name, "Bob")
}

func TestAddBackLinkedList(t *testing.T) {
	ls := initLinkedList()
	ls.AddBack("Bob")
	assert.Equal(t, ls.head.name, "Bob")
	ls.AddBack("John")
	assert.Equal(t, ls.head.next.name, "John")
}

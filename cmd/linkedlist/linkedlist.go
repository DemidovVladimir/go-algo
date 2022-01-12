package main

import "fmt"

type ele struct {
	name string
	next *ele
}

type singleList struct {
	head *ele
	len  int
}

func initLinkedList() *singleList {
	return &singleList{}
}

func (s *singleList) AddFront(name string) {
	if s.head == nil {
		s.head = &ele{name: name}
	} else {
		oldHead := s.head
		s.head = &ele{name: name}
		s.head.next = oldHead
	}
	s.len++
}

func (s *singleList) AddBack(name string) {
	if s.head == nil {
		s.head = &ele{name: name}
	} else {
		current := s.head

		for current.next != nil {
			current = current.next
		}
		current.next = &ele{name: name}
	}
	s.len++
}

func (s *singleList) RemoveFront() error {
	if s.head == nil {
		return fmt.Errorf("List is empty")
	}
	s.head = s.head.next
	s.len--
	return nil
}

func (s *singleList) RemoveBack() error {
	if s.head == nil {
		return fmt.Errorf("removeBack: List is empty")
	}
	var prev *ele
	current := s.head
	for current.next != nil {
		prev = current
		current = current.next
	}
	if prev != nil {
		prev.next = nil
	} else {
		s.head = nil
	}
	s.len--
	return nil
}

func (s *singleList) Front() (string, error) {
	if s.head == nil {
		return "", fmt.Errorf("Single List is empty")
	}
	return s.head.name, nil
}

func (s *singleList) Size() int {
	return s.len
}

func (s *singleList) Traverse() error {
	if s.head == nil {
		return fmt.Errorf("TranverseError: List is empty")
	}
	current := s.head
	for current != nil {
		fmt.Println(current.name)
		current = current.next
	}
	return nil
}

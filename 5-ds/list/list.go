package list

import (
	"fmt"
)

// Пакет реализует двусвязный список вместе с базовыми операциями над ним.

// List - двусвязный список.
type List struct {
	root Elem
}

// Elem - элемент списка.
type Elem struct {
	Val        interface{}
	next, prev *Elem
}

// New - конструктор нового списка.
func New() *List {
	var l List
	l.root.next = &l.root
	l.root.prev = &l.root
	return &l
}

// Push вставляет элемент в начало списка.
func (l *List) Push(e Elem) *Elem {
	e.prev = &l.root
	e.next = l.root.next
	l.root.next = &e
	return &e
}

// String реализует интерфейс fmt.Stringer представляя список в виде строки.
func (l *List) String() string {
	el := l.root.next
	var s string
	for el != &l.root {
		s += fmt.Sprintf("%v ", el.Val)
		el = el.next
	}
	s = s[:len(s)-1]
	return s
}

// Pop удаляет первый элемент списка.
func (l *List) Pop() *List {
	return l
}

// Reverse разворачивает список.
func (l *List) Reverse() *List {
	return l
}

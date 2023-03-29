package main

type List struct {
	First *Element
	Last  *Element
}

type Element struct {
	Value interface{}
	Next  *Element
	Prev  *Element
}

func (l *List) PushFront(v interface{}) *Element {
	e := &Element{Value: v}
	if l.First == nil {
		l.First = e
		l.Last = e
		return e
	}
	e.Next = l.First
	l.First.Prev = e
	l.First = e
	return e
}

func (l *List) PushBack(v interface{}) *Element {
	e := &Element{Value: v}
	if l.Last == nil {
		l.First = e
		l.Last = e
		return e
	}
	e.Prev = l.Last
	l.Last.Next = e
	l.Last = e
	return e
}

func (l *List) Remove(e *Element) {
	if e.Prev == nil {
		l.First = e.Next
	} else {
		e.Prev.Next = e.Next
	}
	if e.Next == nil {
		l.Last = e.Prev
	} else {
		e.Next.Prev = e.Prev
	}
}

func (l *List) PopFront() *Element {
	if l.First == nil {
		return nil
	}
	e := l.First
	l.Remove(e)
	return e
}

func (l *List) PopBack() *Element {
	if l.Last == nil {
		return nil
	}
	e := l.Last
	l.Remove(e)
	return e
}

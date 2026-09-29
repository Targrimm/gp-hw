package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	len    int
	pseudo ListItem
}

func (l *list) Init() *list {
	l.len = 0
	l.pseudo.Next = &l.pseudo
	l.pseudo.Prev = &l.pseudo
	return l
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	if l.len == 0 {
		return nil
	}
	return l.pseudo.Next
}

func (l *list) Back() *ListItem {
	if l.len == 0 {
		return nil
	}
	return l.pseudo.Prev
}

func (l *list) PushFront(v interface{}) *ListItem {
	newElement := &ListItem{v, l.pseudo.Next, nil}
	l.pseudo.Next.Prev = newElement
	l.pseudo.Next = newElement
	l.len++
	return newElement
}

func (l *list) PushBack(v interface{}) *ListItem {
	newElement := &ListItem{v, nil, l.pseudo.Prev}
	l.pseudo.Prev.Next = newElement
	l.pseudo.Prev = newElement
	l.len++
	return newElement
}

func (l *list) Remove(i *ListItem) {
	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		l.pseudo.Prev = i.Prev
	}
	if i.Prev != nil {
		i.Prev.Next = i.Next
	} else {
		l.pseudo.Next = i.Next
	}
	i.Next = nil
	i.Prev = nil
	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		l.pseudo.Prev = i.Prev
	}
	if i.Prev != nil {
		i.Prev.Next = i.Next
	} else {
		l.pseudo.Next = i.Next
	}
	i.Next = l.pseudo.Next
	i.Prev = nil
	l.pseudo.Next.Prev = i
	l.pseudo.Next = i
}

func NewList() List {
	return new(list).Init()
}

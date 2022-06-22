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
	front *ListItem
	back  *ListItem
	len   int
}

func NewList() List {
	return &list{
		len: 0,
	}
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Back() *ListItem {
	return l.back
}

func (l *list) Front() *ListItem {
	return l.front
}

func (l *list) PushFront(v interface{}) *ListItem {
	defer func() { l.len++ }()
	item := &ListItem{
		Value: v,
	}
	if l.Front() != nil {
		item.Next = l.Front()
		l.front.Prev = item
		l.front = item
		return item
	}
	if l.len == 0 {
		l.back = item
		l.front = item
		return item
	}
	return item
}

func (l *list) PushBack(v interface{}) *ListItem {
	defer func() { l.len++ }()
	item := &ListItem{
		Value: v,
	}
	if l.Back() != nil {
		item.Prev = l.Back()
		l.back.Next = item
		l.back = item
		return item
	}
	if l.len == 0 {
		l.front = item
		l.back = item
		return item
	}
	return item
}

func (l *list) Remove(i *ListItem) {
	defer func() { l.len-- }()
	if l.front == i {
		l.front = i.Next
	}
	if l.back == i {
		l.back = i.Prev
	}
	if i.Prev != nil {
		i.Prev.Next = i.Next
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	}
}

func (l *list) MoveToFront(i *ListItem) {
	l.Remove(i)
	l.PushFront(i.Value)
}

package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	val, ok := l.items[key]
	if !ok {
		l.items[key] = l.queue.PushFront(key, value)
		if l.capacity < l.queue.Len() {
			delete(l.items, l.queue.Back().Key)
			l.queue.Remove(l.queue.Back())
		}
	} else {
		val.Value = value
		l.queue.MoveToFront(val)
	}
	return ok
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	val, ok := l.items[key]
	if ok {
		l.queue.MoveToFront(val)
		return val.Value, ok
	}
	return nil, false
}

func (l *lruCache) Clear() {
	l.queue = NewList()
	l.items = make(map[Key]*ListItem, l.capacity)
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

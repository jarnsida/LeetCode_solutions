type cMapEl struct {
	val int
	key int
}
type LRUCache struct {
	m    map[int]*list.Element
	caps int
	l    *list.List
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		m:    make(map[int]*list.Element, 2),
		caps: capacity,
		l:    list.New(),
	}
}

func (this *LRUCache) Get(key int) int {
	val, ok := this.m[key]
	if !ok {
		return -1
	}
	this.l.MoveToFront(val)
	return val.Value.(*cMapEl).val
}

func (this *LRUCache) Put(key int, value int) {
	// Check if we already have key in our map
	v, ok := this.m[key]
	listLen := this.l.Len()

	if ok {
		v.Value.(*cMapEl).val = value
		this.l.MoveToFront(v)
		return

	}

	if listLen >= this.caps {
		backNode := this.l.Back()
		backElKey := backNode.Value.(*cMapEl)
		this.l.Remove(backNode)
		delete(this.m, backElKey.key)
	}
	// If we don't have such Key

	el := this.l.PushFront(&cMapEl{value, key}) //push new key as Head of LL
	this.m[key] = el
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
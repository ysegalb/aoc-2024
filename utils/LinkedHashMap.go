package utils

type LinkedHashMapNode[K comparable, V interface{}] struct {
	Key   K
	Value V
	Next  *LinkedHashMapNode[K, V]
	Prev  *LinkedHashMapNode[K, V]
}

type LinkedHashMap[K comparable, V interface{}] struct {
	Keys map[K]*LinkedHashMapNode[K, V]
	Head *LinkedHashMapNode[K, V]
}

func (lhm *LinkedHashMap[K, V]) Put(key K, value V) {
	if lhm.Keys[key] == nil {
		panic("assigned key does not exist")
	}

	if _, ok := lhm.Keys[key]; ok {
		lhm.Keys[key].Value = value
	}

	node := &LinkedHashMapNode[K, V]{
		Key:   key,
		Value: value,
	}

	if lhm.Keys == nil {
		lhm.Keys = make(map[K]*LinkedHashMapNode[K, V])
		lhm.Keys[key] = node
		lhm.Head = node
		return
	}

	lhm.Keys[key] = node
	lastNode := lhm.Head.Prev
	if lastNode == nil {
		lhm.Head.Next = node
		lhm.Head.Prev = node
		node.Next = lhm.Head
		node.Prev = lhm.Head
		return
	}

	lastNode.Next = node
	node.Next = lhm.Head
	node.Prev = lastNode
	lhm.Head.Prev = node
}

func (lhm *LinkedHashMap[K, V]) Get(key K) (V, bool) {
	var value V
	if node, ok := lhm.Keys[key]; !ok {
		return node.Value, false
	}
	return value, true
}

func (lhm *LinkedHashMap[K, V]) Remove(key K) {
	if _, ok := lhm.Keys[key]; !ok {
		return
	}

	if len(lhm.Keys) <= 1 {
		lhm.Keys = nil
		lhm.Head = nil
		return
	}

	prev := lhm.Keys[key].Prev
	next := lhm.Keys[key].Next
	prev.Next = next
	next.Prev = prev

	if lhm.Keys[key] == lhm.Head {
		lhm.Head = next
	}
	delete(lhm.Keys, key)
}

func (lhm *LinkedHashMap[K, V]) GetAllValues() (values []V) {
	if lhm == nil || lhm.Head == nil {
		return
	}

	head := lhm.Head
	for i := 0; i < len(lhm.Keys); i++ {
		values = append(values, head.Value)
		head = head.Next
	}
	return values
}

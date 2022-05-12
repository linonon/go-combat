package hash

const (
	expandFactor = 0.75
)

// Map by myself
type Map struct {
	m   []*KeyPairs
	cap int
	len int
}

// KeyPairs for Map
type KeyPairs struct {
	key   string
	value any
	next  *KeyPairs
}

// NewHashMap returns a new Map
func NewHashMap(cap int) *Map {
	if cap < 16 {
		cap = 16
	}
	hashMap := new(Map)
	hashMap.m = make([]*KeyPairs, cap, cap)
	hashMap.cap = cap
	return hashMap
}

// Index returns the index of hashMap
func (h *Map) Index(key string) uint {
	return BKDRHash(key, h.cap)
}

// BKDRHash calculates the hash of key
func BKDRHash(str string, cap int) uint {
	seed := uint(13)
	hash := uint(0)
	for i := 0; i < len(str); i++ {
		hash = (hash * seed) + uint(str[i])
	}
	return hash % uint(cap)
}

// Put key value pair into hashMap
func (h *Map) Put(key string, value any) {
	index := h.Index(key)
	element := h.m[index]
	if element == nil { // 這個位置沒 value, 直接寫
		h.m[index] = &KeyPairs{key, value, nil}
	} else { // 這個位置已有 value
		// 找到最後一個節點
		for element.next != nil {
			if element.key == key { // 如果是相同的 key 就進行修改
				element.value = value
				return
			}
			element = element.next
		}
		element.next = &KeyPairs{key, value, nil}
	}
	h.len++
	// 需要擴容
	if (float64(h.len) / float64(h.cap)) > expandFactor {
		// 擴兩倍
		newH := NewHashMap(2 * h.cap)
		for _, pairs := range h.m {
			for pairs != nil {
				newH.Put(pairs.key, pairs.value)
			}
		}
		h.cap = newH.cap
		h.m = newH.m
	}
}

// Get returns the value from the given key.
func (h *Map) Get(key string) any {
	index := h.Index(key)
	pairs := h.m[index]
	if pairs.key == key {
		return pairs.value
	}
	for pairs.next != nil {
		if pairs.key == key {
			return pairs.value
		}
		pairs = pairs.next
	}
	return nil
}

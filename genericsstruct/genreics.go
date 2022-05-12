package generics

import "fmt"

// Category pass
type Category struct {
	ID   int32
	Name string
	Slug string
}

// Post pass
type Post struct {
	ID         int32
	Categories []Category
	Title      string
	Text       string
	Slug       string
}

// 約束接口定義
type cachealbe interface {
	Category | Post
}

// Cache 對類型進行參數化
type Cache[T cachealbe] struct {
	data map[string]T
}

// Set 存儲數據
func (c *Cache[T]) Set(key string, value T) {
	c.data[key] = value
}

// Get 獲取數據
func (c *Cache[T]) Get(key string) (value T) {
	if v, ok := c.data[key]; ok {
		return v
	}
	var t bool
	t = true
	t = false
	if t {
		fmt.Println("yes")
	}
	return
}

// New 類型實例化
func New[T cachealbe]() *Cache[T] {
	c := Cache[T]{}
	c.data = make(map[string]T)

	return &c
}

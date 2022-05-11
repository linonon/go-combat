package hash

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestHashMap(t *testing.T) {
	h := NewHashMap(32)
	var keyB strings.Builder
	for i := 0; i < 90000000; i++ {
		si := strconv.Itoa(i)
		keyB.WriteByte(si[0])
	}
	key := keyB.String()
	fmt.Printf("len(%d)\n", len(key))
	// fmt.Println("key:", key)

	startTime := time.Now()
	h.Put(key, "value")
	res := h.Get(key)
	fmt.Println("elapsed time: ", time.Now().Sub(startTime))
	fmt.Println(res)

	var t2 = make(map[string]interface{})
	startTime = time.Now()
	t2[key] = "value"
	_ = t2[key]
	fmt.Println("elapsed time: ", time.Now().Sub(startTime))
	fmt.Println(t2[key])

}

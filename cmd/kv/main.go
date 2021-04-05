package main

import (
	"fmt"

	kv "github.com/tengla/foobar/keyval"
)

func main() {
	cfg := map[string]interface{}{
		"name": "test",
		"age":  "29",
		"city": "Oslo",
	}
	ch := make(chan kv.KeyVal)
	go kv.Fn(ch, cfg)
	for v := range ch {
		fmt.Printf("%s=%s\n", v.Key, v.Val)
	}
}

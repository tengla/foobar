package keyval

import (
	"fmt"
	"time"
)

type KeyVal struct {
	Key string
	Val interface{}
}

func Fn(c chan<- KeyVal, kv map[string]interface{}) {
	for k, v := range kv {
		kv := KeyVal{k, v}
		c <- kv
		time.Sleep(time.Millisecond * 250)
	}
	close(c)
}

func init() {
	fmt.Println("-- Init Keyval --")
}

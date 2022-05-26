package cache

import (
	"fmt"
	"time"
)

type Value struct {
	value    string
	deadline time.Time
}

type Cache struct {
	cashe map[string]Value
}

func NewCache() Cache {
	return Cache{cashe: make(map[string]Value)}
}

//func (receiver) Get(key string) (string, bool) {

//}
func (receiver *Cache) Get(key string) (string, bool) {

	v, ok := receiver.cashe[key]
	if ok == true {
		fmt.Println(v.value, ok)
		return v.value, ok
	}
	if !ok {
		err1 := "not found this value"
		fmt.Println(err1)

	}

	return "", false

}

func (receiver *Cache) Put(key, value string) {
	receiver.cashe[key] = Value{value: value}
}

func (receiver *Cache) Keys() []string {
	output := make([]string, 0)

	for k, _ := range receiver.cashe {
		output = append(output, k)
	}
	return output
}

func (receiver *Cache) PutTill(key, value string, deadline time.Time) {
}

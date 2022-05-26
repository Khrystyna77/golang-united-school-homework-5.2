package cache

import (
	"fmt"
	"time"
)

type Value struct {
	value string
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
		//fmt.Println(v)
		return v.value, ok
	}
	if ok == false {
		err1 := "not found this value"
		fmt.Println(err1)

	}

	return "", false

}

func (receiver *Cache) Put(key, value string) {
	receiver.cashe[key] = Value{value: value}	
	}

}

func (receiver) Keys() []string {
}

func (receiver) PutTill(key, value string, deadline time.Time) {
}

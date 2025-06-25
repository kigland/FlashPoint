package flashmap

import (
	"log"
	"strconv"
	"sync"
	"time"
)

const (
	TypeBinary Type = "bin"
	TypeText   Type = "txt"
	TypeJSON   Type = "json"
)

type Type string

func (t Type) Normalize() Type {
	switch t {
	case "bin":
		return TypeBinary
	case "txt":
		return TypeText
	case "json":
		return TypeJSON
	}
	return TypeText
}

type Value struct {
	Type     Type
	Mime     string
	Value    any
	ExpireAt time.Time
}

type FlashMap struct {
	cache map[string]Value
	lck   sync.RWMutex
}

func NewFlashMap() *FlashMap {
	return &FlashMap{
		cache: make(map[string]Value),
		lck:   sync.RWMutex{},
	}
}

func (f *FlashMap) Set(key string, value any, ttl time.Duration, t Type, mime string) {
	v := Value{
		Value:    value,
		ExpireAt: time.Now().Add(ttl),
		Type:     t,
		Mime:     mime,
	}
	f.lck.Lock()
	defer f.lck.Unlock()

	f.cache[key] = v
}

func (f *FlashMap) Get(key string) (Value, bool) {
	f.lck.RLock()
	value, ok := f.cache[key]
	f.lck.RUnlock()

	if !ok {
		return Value{}, false
	}

	if value.ExpireAt.Before(time.Now()) {
		f.lck.Lock()
		if v, exists := f.cache[key]; exists && v.ExpireAt.Before(time.Now()) {
			delete(f.cache, key)
			log.Println("[GC] key=" + key)
		}
		f.lck.Unlock()
		return Value{}, false
	}

	return value, true
}

func (f *FlashMap) GC() {
	count := 0
	f.lck.Lock()
	defer f.lck.Unlock()

	for key, value := range f.cache {
		if value.ExpireAt.Before(time.Now()) {
			delete(f.cache, key)
			log.Println("[GC] key=" + key)
			count++
		}
	}

	if count > 0 {
		log.Println("[GC] deleted " + strconv.Itoa(count) + " keys")
	}
}

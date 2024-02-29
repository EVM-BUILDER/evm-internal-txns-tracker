package cache

import (
	lru "github.com/hashicorp/golang-lru/v2"
)

var l *lru.Cache[string, any]

func NewMCache() {
	l, _ = lru.New[string, any](10000000)
}

func Add(k string, v interface{}) bool {
	return l.Add(k, v)
}

func Get(k string) (interface{}, bool) {
	return l.Get(k)
}

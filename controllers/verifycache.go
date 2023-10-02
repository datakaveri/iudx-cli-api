package controllers

import (
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
)

func verifyCache(id string, body string) {
	var cache = memcache.New("localhost:11211")

	cacheErr := cache.Set(&memcache.Item{Key: id, Value: []byte(body), Expiration: 2000})

	if cacheErr != nil {
		fmt.Println(cacheErr.Error())
	}

	cache.Close()
}

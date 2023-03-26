package main

import (
	"bitbucket.org/provisota/cache/cache"
	"fmt"
)

func main() {
	cache1 := cache.New()

	cache1.Set("userId", 42)
	fmt.Println(cache1.Get("userId"))

	cache1.Delete("userId")

	fmt.Println(cache1.Get("userId"))
}

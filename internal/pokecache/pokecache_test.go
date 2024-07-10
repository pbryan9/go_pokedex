package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestReportFalseOnCacheMiss(t *testing.T) {
	const testInterval = 5 * time.Second
	cache := NewCache(testInterval)
	if _, ok := cache.Get("test"); ok {
		t.Errorf("should return false on cache miss")
	}
}

func TestReportTrueOnCacheHit(t *testing.T) {
	const testInterval = 5 * time.Second
	cache := NewCache(testInterval)
	cache.Add("test", []byte("test"))
	if _, ok := cache.Get("test"); !ok {
		t.Errorf("should return true on cache hit")
	}
}

func TestAddGet(t *testing.T) {
	const testInterval = 5 * time.Second
	testCases := []struct {
		key string
		val []byte
	}{
		{
			key: "hello",
			val: []byte("there"),
		},
		{
			key: "goodbye",
			val: []byte("farewell"),
		},
		{
			key: "patty.dev/recipes",
			val: []byte("cookies"),
		},
	}
	for i, c := range testCases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(testInterval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("Expected to find key %s", c.key)
			}
			if string(val) != string(c.val) {
				t.Errorf("Expected value at key %s to be %s, but got %s", c.key, c.val, val)
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const testInterval = 5 * time.Millisecond
	const waitTime = testInterval + 5*time.Millisecond
	const testKey = "www.test.com"
	const testVal = "this is a test"
	cache := NewCache(testInterval)
	cache.Add(testKey, []byte(testVal))
	if _, ok := cache.Get(testKey); !ok {
		t.Errorf("expected to find key before reaploop runs")
	}
	time.Sleep(waitTime)
	if _, ok := cache.Get(testKey); ok {
		t.Errorf("expected not to find key after reaploop runs")
	}
}

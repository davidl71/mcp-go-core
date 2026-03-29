package cache

import (
	"testing"
	"time"
)

func TestTTLCacheGetSet(t *testing.T) {
	c := NewTTLCache()
	c.Set("k", []byte("v"), time.Minute)
	got, ok := c.Get("k")
	if !ok || string(got) != "v" {
		t.Fatalf("Get = %q, %v; want v, true", got, ok)
	}
}

func TestTTLCacheExpire(t *testing.T) {
	c := NewTTLCache()
	c.Set("k", []byte("v"), 20*time.Millisecond)
	time.Sleep(40 * time.Millisecond)
	_, ok := c.Get("k")
	if ok {
		t.Fatal("expected miss after TTL")
	}
}

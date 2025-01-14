package pokecache

import (
	"sync"
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	cases := []struct {
		input    map[string]string
		expected Cache
	}{
		{
			input: map[string]string{
				"abcd": "efg",
			},
			expected: Cache{
				map[string]cacheEntry{
					"abcd": {
						time.Now(),
						[]byte("efg"),
					},
				},
				&sync.Mutex{},
				time.Duration(5) * time.Second,
			},
		},
	}

	for _, c := range cases {
		cache := NewCache(5)
		for k, v := range c.input {
			cache.Add(k, []byte(v))
		}
		for k, entry := range c.expected.cache {
			val, exists := cache.Get(k)
			if !exists {
				t.Errorf("Value %s does not exist in cache when it should.", val)
			}
			if string(val) != string(entry.val) {
				t.Logf("Value %s is not the same as expected %s", val, entry.val)
				t.Errorf("\n%v | actual vs\n%v | expected", cache, c.expected)
			}
		}

	}
}
func TestReapLoop(t *testing.T) {
	const baseTime = 5
	const waitTime = baseTime*time.Millisecond + 5*time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}

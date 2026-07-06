package pokecache

import (
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := map[string]struct {
		input    string
		expected []byte
	}{
		"plain example.com": {
			input:    "https://example.com",
			expected: []byte("testdata"),
		},
		"example /path": {
			input:    "https://example.com/path",
			expected: []byte("moretestdata"),
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(tc.input, tc.expected)
			val, ok := cache.Get(tc.input)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(tc.expected) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond

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

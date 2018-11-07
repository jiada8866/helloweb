package redis

import "testing"

func TestSet(t *testing.T) {
	key := "test:key"
	value := "hello world"
	err := Set(key, []byte(value))
	if err != nil {
		t.Error(err)
	}

	b, err := Get(key)
	if err != nil {
		t.Error(err)
	}
	if string(b) != value {
		t.Errorf("get %s, want %s", string(b), value)
	}
}

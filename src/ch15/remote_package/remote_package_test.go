package remote

import (
	"testing"

	cm "github.com/easierway/concurrent_map" // cm为别名
)

func TestConcurrentMap(t *testing.T) {
	m := cm.CreateConcurrentMap(99)
	m.Set(cm.StrKey("key"), 10)
	t.Log(m.Get(cm.StrKey("key")))
}

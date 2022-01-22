package services

import (
	"storage-app/data/cache"
	"testing"
)

func TestSaveFile(t *testing.T) {
	cache := cache.NewCacheManager()
	svc := NewFileSaverService(cache, 0, "/tmp/")
	cache.AddOrUpdate("foo", "bar")
	all := cache.GetAll()

	if len(all) != 1 {
		t.Fatalf(`Count expected is 1 but actual %d`, len(all))
	}
	svc.Save()
	cache.Flush()
	all = cache.GetAll()
	if len(all) != 0 {
		t.Fatalf(`Count expected is 0 but actual %d`, len(all))
	}
	all = svc.FillRecords()
	if len(all) != 1 {
		t.Fatalf(`Count expected is 1 but actual %d`, len(all))
	}
}

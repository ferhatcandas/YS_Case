package services

import (
	"storage-app/data/cache"
	"storage-app/models/request"
	"testing"
)

func TestSaveAndGetExistRecord(t *testing.T) {
	svc := NewStorageService(cache.NewCacheManager())
	expected := request.StorageReq{
		Key:   "foo",
		Value: "bar",
	}
	svc.Save(expected)

	actual, _ := svc.Get("foo")

	if actual != expected.Value {
		t.Fatalf(`Actual %s not matched with Expected %s `, actual, expected.Value)
	}
}
func TestGetNotExistRecord(t *testing.T) {
	svc := NewStorageService(cache.NewCacheManager())

	_, err := svc.Get("foo")
	if err == nil {
		t.Fatalf(`Actual shouldn't equal to be nil`)
	}
}

func TestFlushAll(t *testing.T) {
	svc := NewStorageService(cache.NewCacheManager())
	expected := request.StorageReq{
		Key:   "foo",
		Value: "bar",
	}
	svc.Save(expected)

	all := svc.cache.GetAll()

	if len(all) != 1 {
		t.Fatalf(`Count expected is 1 but actual %d`, len(all))
	}

	svc.RemoveAll()

	all = svc.cache.GetAll()

	if len(all) != 0 {
		t.Fatalf(`Count expected is 0 but actual %d`, len(all))
	}
}

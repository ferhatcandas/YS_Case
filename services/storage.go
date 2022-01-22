package services

import (
	"storage-app/data/cache"
	"storage-app/models/request"
)

type StorageService struct {
	cache cache.CacheMangager
}

func NewStorageService(cache cache.CacheMangager) StorageService {
	return StorageService{
		cache: cache,
	}
}
func (s *StorageService) Save(req request.StorageReq) {
	s.cache.AddOrUpdate(req.Key, req.Value)
}

func (s *StorageService) RemoveAll() {
	s.cache.Flush()
}
func (s *StorageService) Remove(key string) {
	s.cache.Remove(key)
}
func (s *StorageService) Get(key string) (interface{}, error) {
	return s.cache.Get(key)
}

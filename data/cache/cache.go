package cache

import (
	"errors"
)

type CacheMangager struct {
	list map[string]string
}

func NewCacheManager() CacheMangager {
	return CacheMangager{
		list: make(map[string]string),
	}
}
func (c *CacheMangager) AddOrUpdate(key string, value string) {
	c.list[key] = value
}
func (c *CacheMangager) Remove(key string) {
	delete(c.list, key)
}
func (c *CacheMangager) Flush() {
	c.list = make(map[string]string)
}
func (c *CacheMangager) GetAll() map[string]string {
	return c.list
}
func (c *CacheMangager) Get(key string) (string, error) {
	value := c.list[key]
	if value != "" {
		return value, nil
	} else {
		return "", errors.New("Key Not Found")
	}
}

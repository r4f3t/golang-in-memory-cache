package user

import (
	"github.com/patrickmn/go-cache"
	"github.com/r4f3t/webapi/mockdata"
	"strconv"
	"time"
)

// TTL Time to live for cached object as second
const TTL = 10

type UserCacheManager interface {
	GetUser(id int) *mockdata.UserModel
	FillUsers(users []mockdata.UserModel)
}

type cacheManager struct {
	cache *cache.Cache
}

func NewUserCacheManager() UserCacheManager {
	return &cacheManager{
		cache: cache.New(TTL*time.Minute, TTL*time.Minute),
	}
}

func (c *cacheManager) GetUser(id int) *mockdata.UserModel {
	key := strconv.Itoa(id)
	result, err := c.cache.Get(key)
	if !err {
		return nil
	}
	response, ok := result.(mockdata.UserModel)
	if !ok {
		return nil
	}

	return &response
}

func (c *cacheManager) FillUsers(users []mockdata.UserModel) {
	for index, value := range users {
		key := index
		c.cache.Set(strconv.Itoa(key), value, time.Second*TTL)
	}
}

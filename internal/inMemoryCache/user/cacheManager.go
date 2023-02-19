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
}

type cacheManager struct {
	cache    *cache.Cache
	userRepo UserRepository
}

func NewUserCacheManager(ur UserRepository) UserCacheManager {
	return &cacheManager{
		cache:    cache.New(TTL*time.Minute, TTL*time.Minute),
		userRepo: ur,
	}
}

func (c *cacheManager) GetUser(id int) *mockdata.UserModel {
	key := strconv.Itoa(id)
	result, hasItem := c.cache.Get(key)
	if hasItem {
		response, ok := result.(mockdata.UserModel)
		//check is casting ok
		if ok {
			return &response
		}
	}

	var response mockdata.UserModel
	//it is not ok go check db if it has record for id
	users := c.userRepo.GetUsersById(id)
	if users != nil && len(users.Records) > 0 {
		//db has record for id
		response = users.Records[0]
		//set to memory by setting ttl
		c.cache.Set(key, response, time.Second*TTL)
	}

	return &response
}

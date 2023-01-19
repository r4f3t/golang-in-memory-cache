package user

import (
	"github.com/r4f3t/webapi/mockdata"
	"time"
)

// TTL Time to live for cached object as second
const TTL = 10

type UserCacheManager interface {
	GetUser(id int) *mockdata.UserModel
	FillUsers(users []mockdata.UserModel)
}

type cacheManager struct {
	cachedUserList map[int]mockdata.UserModel
	endTime        *time.Time
}

func NewUserCacheManager() UserCacheManager {
	now := time.Now()
	return &cacheManager{
		cachedUserList: map[int]mockdata.UserModel{},
		endTime:        &now,
	}
}

func (c *cacheManager) GetUser(id int) *mockdata.UserModel {
	result := c.cachedUserList[id]
	if result.Id == 0 {
		return nil
	}

	//TTL Check
	now := time.Now()
	if c.endTime != nil && c.endTime.Before(now) {
		setEndTime(c)
		return nil
	}

	return &result
}

func (c *cacheManager) FillUsers(users []mockdata.UserModel) {
	for index, value := range users {
		key := index
		c.cachedUserList[key] = value
	}

	setEndTime(c)
}

func setEndTime(c *cacheManager) {
	endTime := time.Now().Add(time.Second * TTL)
	c.endTime = &endTime
}

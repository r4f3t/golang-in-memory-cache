package user

import (
	"github.com/r4f3t/webapi/mockdata"
)

type UserCacheManager interface {
	GetUser(id int) *mockdata.UserModel
	FillUsers(users []mockdata.UserModel)
}

type cacheManager struct {
	cachedUserList map[int]mockdata.UserModel
}

func NewUserCacheManager() UserCacheManager {

	return &cacheManager{
		cachedUserList: map[int]mockdata.UserModel{},
	}
}

func (c cacheManager) GetUser(id int) *mockdata.UserModel {
	result := c.cachedUserList[id]
	if result.Id == 0 {
		return nil
	}

	return &result
}

func (c cacheManager) FillUsers(users []mockdata.UserModel) {
	for index, value := range users {
		key := index
		c.cachedUserList[key] = value
	}
}

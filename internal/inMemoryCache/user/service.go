package user

import "github.com/r4f3t/webapi/mockdata"

type Service interface {
	GetUserById(id int) *mockdata.UserModel
}

type service struct {
	userCacheManager UserCacheManager
}

func NewService(userCache UserCacheManager) Service {
	return &service{
		userCacheManager: userCache,
	}
}

func (receiver *service) GetUserById(id int) *mockdata.UserModel {
	//check and get cached object
	result := receiver.userCacheManager.GetUser(id)

	return result
}

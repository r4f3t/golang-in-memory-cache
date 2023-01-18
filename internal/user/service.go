package user

import "github.com/r4f3t/webapi/mockdata"

type Service interface {
	GetUserById(id int) mockdata.UserModel
}

type service struct {
	userRepository   UserRepository
	userCacheManager UserCacheManager
}

func NewService(userRepo UserRepository, userCache UserCacheManager) Service {
	return &service{
		userRepository:   userRepo,
		userCacheManager: userCache,
	}
}

func (receiver *service) GetUserById(id int) mockdata.UserModel {
	//check cached object first
	result := receiver.userCacheManager.GetUser(id)
	// if it is null fill object from db
	if result == nil {

		records := receiver.userRepository.GetUsersFromDb()

		receiver.userCacheManager.FillUsers(records.Records)

		result = receiver.userCacheManager.GetUser(id)
	}

	return *result
}

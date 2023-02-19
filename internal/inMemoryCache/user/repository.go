package user

import "github.com/r4f3t/webapi/mockdata"

// repository interface that repository struct implements
type UserRepository interface {
	GetUsersFromDb() *mockdata.TableModel
	GetUsersById(id int) *mockdata.TableModel
}

// repository model that constructor returns
type repository struct {
	Database dbInstance
}

// temp struct it can be any database config
type dbInstance struct {
	database *mockdata.TableModel
}

func NewRepository() UserRepository {

	return &repository{
		Database: dbInstance{},
	}
}

func (receiver *repository) GetUsersFromDb() *mockdata.TableModel {
	return mockdata.GetAll()
}

func (receiver *repository) GetUsersById(id int) *mockdata.TableModel {
	return mockdata.GetById(id)
}

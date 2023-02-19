package mockdata

import (
	"fmt"
	"time"
)

type UserModel struct {
	Id    int
	Name  string
	Email string
}

type TableModel struct {
	Records []UserModel
}

func GetAll() *TableModel {
	table := new(TableModel)

	for i := 0; i < 1000; i++ {
		table.Records = append(table.Records, UserModel{
			Id:    i + 1,
			Name:  fmt.Sprintf("User %d", i),
			Email: fmt.Sprintf("user%d@mail.com", i),
		})
	}

	//wait here ,when data comes from db it will take 2 second for getting it
	time.Sleep(time.Second * 2)

	return table
}

func GetById(id int) *TableModel {
	table := new(TableModel)

	for i := 0; i < 1000; i++ {
		if id == i+1 {
			table.Records = append(table.Records, UserModel{
				Id:    i + 1,
				Name:  fmt.Sprintf("User %d", i),
				Email: fmt.Sprintf("user%d@mail.com", i),
			})
		}
	}

	//wait here ,when data comes from db it will take 2 second for getting it
	time.Sleep(time.Second * 2)

	return table
}

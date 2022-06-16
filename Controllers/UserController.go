package Controllers

import (
	"encoding/json"

	"github.com/MrAmperage/GoWebStruct/WebCore/Modules/ORMModule"
	"github.com/MrAmperage/ReportBoxAdministartionService/ORM"
	"github.com/streadway/amqp"
)

func AddUser(Message amqp.Delivery, ORMs ORMModule.ORMArray) (Data interface{}, Error error) {
	var NewUser ORM.User
	json.Unmarshal(Message.Body, &NewUser)
	ORMElement, Error := ORMs.FindByName("UserORM")
	if Error != nil {

		return
	}
	UserORM := ORMElement.(*ORM.UserORM)
	return "Пользователь добавлен", UserORM.AddUser(NewUser).Error
}

func DeleteUser(Message amqp.Delivery, ORMs ORMModule.ORMArray) (Data interface{}, Error error) {

	Username := string(Message.Body)
	ORMElement, Error := ORMs.FindByName("UserORM")
	if Error != nil {

		return
	}
	UserORM := ORMElement.(*ORM.UserORM)

	return "Пользователь удален", UserORM.DeleteUser(Username).Error
}

func EditUser(Message amqp.Delivery, ORMs ORMModule.ORMArray) (Data interface{}, Error error) {
	var EditUser ORM.User
	json.Unmarshal(Message.Body, &EditUser)
	ORMElement, Error := ORMs.FindByName("UserORM")
	if Error != nil {

		return
	}
	UserORM := ORMElement.(*ORM.UserORM)
	return "Пользователь изменен", UserORM.EditUser(EditUser).Error

}

func GetUsers(Message amqp.Delivery, ORMs ORMModule.ORMArray) (Data interface{}, Error error) {
	ORMElement, Error := ORMs.FindByName("UserORM")
	if Error != nil {

		return
	}
	UserORM := ORMElement.(*ORM.UserORM)
	if len(Message.Body) != 0 {
		return UserORM.GetUser(string(Message.Body))
	} else {
		return UserORM.GetUsers()
	}

}

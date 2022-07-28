package Controllers

import (
	"encoding/json"

	"github.com/MrAmperage/GoWebStruct/WebCore/Modules/ORMModule"
	"github.com/MrAmperage/ReportBoxAdministartionService/ORM"
	"github.com/gofrs/uuid"
	"github.com/streadway/amqp"
)

func AddUser(Message amqp.Delivery, ORMs ORMModule.ORMArray) (Data interface{}, Error error) {
	var NewUser ORM.User
	json.Unmarshal(Message.Body, &NewUser)
	ORMElement, Error := ORMs.FindByName("UsersORM")
	if Error != nil {

		return
	}
	UserORM := ORMElement.(*ORM.UserORM)
	return "Пользователь добавлен", UserORM.AddUser(NewUser).Error
}

func DeleteUser(Message amqp.Delivery, ORMs ORMModule.ORMArray) (Data interface{}, Error error) {

	ORMElement, Error := ORMs.FindByName("UsersORM")
	if Error != nil {

		return
	}
	UserORM := ORMElement.(*ORM.UserORM)
	Uuid, Error := uuid.FromString(string(Message.Body))
	if Error != nil {
		return
	}
	return "Пользователь удален", UserORM.DeleteUser(Uuid).Error
}

func EditUser(Message amqp.Delivery, ORMs ORMModule.ORMArray) (Data interface{}, Error error) {
	var EditUser ORM.User
	json.Unmarshal(Message.Body, &EditUser)
	ORMElement, Error := ORMs.FindByName("UsersORM")
	if Error != nil {

		return
	}
	UserORM := ORMElement.(*ORM.UserORM)
	return UserORM.EditUser(EditUser)

}

func GetUsers(Message amqp.Delivery, ORMs ORMModule.ORMArray) (Data interface{}, Error error) {
	ORMElement, Error := ORMs.FindByName("UsersORM")
	if Error != nil {

		return
	}
	UserORM := ORMElement.(*ORM.UserORM)
	if len(Message.Body) != 0 {
		UUid, Error := uuid.FromString(string(Message.Body))
		if Error != nil {
			return nil, Error
		}
		return UserORM.GetUser(UUid)
	} else {
		return UserORM.GetUsers()
	}

}

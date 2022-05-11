package Controllers

import (
	"encoding/json"

	"github.com/MrAmperage/GoWebStruct/WebCore/Modules/ORMModule"
	"github.com/MrAmperage/ReportBoxAdministartionService/ORM"
	"github.com/streadway/amqp"
)

func AddUser(Message amqp.Delivery, ORMs ORMModule.ORMArray) (Data any, Error error) {
	var NewUser ORM.User
	json.Unmarshal(Message.Body, &NewUser)
	ORMElement, Error := ORMs.FindByName("UserORM")
	if Error != nil {

		return
	}
	UserORM := ORMElement.(*ORM.UserORM)
	return "Пользователь добавлен", UserORM.AddUser(NewUser).Error
}

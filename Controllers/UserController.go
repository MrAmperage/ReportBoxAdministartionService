package Controllers

import (
	"fmt"

	"github.com/MrAmperage/GoWebStruct/WebCore/Modules/ORMModule"
	"github.com/streadway/amqp"
)

func AddUser(Message amqp.Delivery, ORMs ORMModule.ORMArray) (Data any, Error error) {
	fmt.Println("Добавление пользователя")

	return
}

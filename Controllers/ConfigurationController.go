package Controllers

import (
	"github.com/MrAmperage/GoWebStruct/WebCore/Modules/ORMModule"
	"github.com/MrAmperage/ReportBoxAdministartionService/ORM"
	"github.com/streadway/amqp"
)

func Configuration(Message amqp.Delivery, ORMs ORMModule.ORMArray) (Data interface{}, Error error) {

	switch string(Message.Body) {
	case "GetApplicationMenu":

	}
	ORMElement, Error := ORMs.FindByName("TopMenuORM")
	if Error != nil {
		return

	}
	TopMenuORM := ORMElement.(*ORM.TopMenuORM)

	return TopMenuORM.GetTopMenus()

}

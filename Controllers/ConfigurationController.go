package Controllers

import (
	"github.com/MrAmperage/GoWebStruct/WebCore/Modules/ORMModule"
	"github.com/MrAmperage/ReportBoxAdministartionService/ORM"
	"github.com/streadway/amqp"
)

func Configuration(Message amqp.Delivery, ORMs ORMModule.ORMArray) (Data any, Error error) {

	ORMElement, Error := ORMs.FindByName("TopMenuORM")
	if Error != nil {

		return
	}
	TopMenuORM := ORMElement.(*ORM.TopMenuORM)
	switch string(Message.Body) {
	case "GetApplicationMenu":

		TopMenu, Error := TopMenuORM.GetTopMenu()
		return TopMenu, Error

	}
	return

}

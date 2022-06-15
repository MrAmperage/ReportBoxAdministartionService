package Controllers

import (
	"github.com/MrAmperage/GoWebStruct/WebCore/Modules/ORMModule"
	"github.com/MrAmperage/ReportBoxAdministartionService/ORM"
	"github.com/streadway/amqp"
)

func Configuration(Message amqp.Delivery, ORMs ORMModule.ORMArray) (Data interface{}, Error error) {

	ORMElement, Error := ORMs.FindByName("SchemeORM")
	if Error != nil {

		return
	}
	SchemeORM := ORMElement.(*ORM.SchemeORM)
	switch string(Message.Body) {
	case "GetApplicationMenu":

		return SchemeORM.GetSchemeByName("ApplicationMenu")

	}
	return

}

package Controllers

import (
	"github.com/MrAmperage/GoWebStruct/WebCore/Modules/ORMModule"
	"github.com/MrAmperage/ReportBoxAdministartionService/ORM"
	"github.com/streadway/amqp"
)

func Configuration(Message amqp.Delivery, ORMs ORMModule.ORMArray) (Data any, Error error) {

	ORMElement, Error := ORMs.FindByName("SchemeORM")
	if Error != nil {

		return
	}
	SchemeORM := ORMElement.(*ORM.SchemeORM)
	switch string(Message.Body) {
	case "GetApplicationMenu":

		ConfigScheme, Error := SchemeORM.GetSchemeByName("ApplicationMenu")
		return ConfigScheme.Scheme, Error

	}
	return

}

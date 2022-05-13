package Controllers

import (
	"encoding/json"

	"github.com/MrAmperage/GoWebStruct/WebCore/Modules/ORMModule"
	"github.com/MrAmperage/ReportBoxAdministartionService/ORM"
	"github.com/streadway/amqp"
)

func AddScheme(Message amqp.Delivery, ORMs ORMModule.ORMArray) (Data any, Error error) {
	var NewScheme ORM.Scheme

	Error = json.Unmarshal(Message.Body, &NewScheme)
	if Error != nil {
		return
	}
	ORMElement, Error := ORMs.FindByName("SchemeORM")
	if Error != nil {

		return
	}
	SchemeORM := ORMElement.(*ORM.SchemeORM)

	return "Схема добавлена", SchemeORM.AddScheme(NewScheme).Error
}

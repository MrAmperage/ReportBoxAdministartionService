package Controllers

import (
	"encoding/json"

	"github.com/MrAmperage/GoWebStruct/WebCore/Modules/ORMModule"
	"github.com/MrAmperage/ReportBoxAdministartionService/ORM"
	"github.com/gofrs/uuid"
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
func DeleteScheme(Message amqp.Delivery, ORMs ORMModule.ORMArray) (Data any, Error error) {
	SchemeId, Error := uuid.FromString(string(Message.Body))
	if Error != nil {
		return
	}

	ORMElement, Error := ORMs.FindByName("SchemeORM")
	if Error != nil {

		return
	}
	SchemeORM := ORMElement.(*ORM.SchemeORM)

	return "Схема удалена", SchemeORM.DeleteScheme(SchemeId).Error
}

func GetSchemes(Message amqp.Delivery, ORMs ORMModule.ORMArray) (Data any, Error error) {
	ORMElement, Error := ORMs.FindByName("SchemeORM")
	if Error != nil {

		return
	}
	SchemeORM := ORMElement.(*ORM.SchemeORM)

	if len(Message.Body) != 0 {
		SchemeId, Error := uuid.FromString(string(Message.Body))
		if Error != nil {
			return nil, Error
		}
		return SchemeORM.GetScheme(SchemeId)

	} else {
		return SchemeORM.GetShemes()
	}

}
func EditScheme(Message amqp.Delivery, ORMs ORMModule.ORMArray) (Data any, Error error) {
	var EditScheme ORM.Scheme
	json.Unmarshal(Message.Body, &EditScheme)
	ORMElement, Error := ORMs.FindByName("SchemeORM")
	if Error != nil {

		return
	}
	SchemeORM := ORMElement.(*ORM.SchemeORM)
	return "Схема изменена", SchemeORM.EditScheme(EditScheme).Error

}

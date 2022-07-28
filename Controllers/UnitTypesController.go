package Controllers

import (
	"encoding/json"

	"github.com/MrAmperage/GoWebStruct/WebCore/Modules/ORMModule"
	"github.com/MrAmperage/ReportBoxAdministartionService/ORM"
	"github.com/gofrs/uuid"
	"github.com/streadway/amqp"
)

func GetUnitTypes(Message amqp.Delivery, ORMs ORMModule.ORMArray) (Data interface{}, Error error) {
	ORMElement, Error := ORMs.FindByName("UnitTypesORM")
	if Error != nil {

		return
	}
	UnitTypesORM := ORMElement.(*ORM.UnitTypesORM)

	if len(Message.Body) != 0 {
		UnitTypeUUID, Error := uuid.FromString(string(Message.Body))
		if Error != nil {
			return nil, Error
		}

		return UnitTypesORM.GetUnitType(UnitTypeUUID)

	} else {
		return UnitTypesORM.GetUnitTypes()
	}

}
func DeleteUnitType(Message amqp.Delivery, ORMs ORMModule.ORMArray) (Data interface{}, Error error) {
	ORMElement, Error := ORMs.FindByName("UnitTypesORM")
	if Error != nil {

		return
	}
	UnitTypesORM := ORMElement.(*ORM.UnitTypesORM)
	UnitTypeUUID, Error := uuid.FromString(string(Message.Body))
	if Error != nil {
		return
	}
	return "Тип агрегата удален", UnitTypesORM.DeleteUnitType(UnitTypeUUID)
}

func AddUnitType(Message amqp.Delivery, ORMs ORMModule.ORMArray) (Data interface{}, Error error) {
	ORMElement, Error := ORMs.FindByName("UnitTypesORM")
	if Error != nil {

		return
	}
	UnitTypesORM := ORMElement.(*ORM.UnitTypesORM)
	var NewUnitType ORM.UnitType
	Error = json.Unmarshal(Message.Body, &NewUnitType)
	if Error != nil {

		return
	}
	return UnitTypesORM.AddUnitType(NewUnitType)

}

func EditUnitType(Message amqp.Delivery, ORMs ORMModule.ORMArray) (Data interface{}, Error error) {
	ORMElement, Error := ORMs.FindByName("UnitTypesORM")
	if Error != nil {

		return
	}
	UnitTypesORM := ORMElement.(*ORM.UnitTypesORM)
	var NewUnitType ORM.UnitType
	Error = json.Unmarshal(Message.Body, &NewUnitType)
	if Error != nil {

		return
	}
	return UnitTypesORM.EditUnitType(NewUnitType)

}

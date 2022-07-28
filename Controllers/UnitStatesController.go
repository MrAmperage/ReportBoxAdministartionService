package Controllers

import (
	"encoding/json"

	"github.com/MrAmperage/GoWebStruct/WebCore/Modules/ORMModule"
	"github.com/MrAmperage/ReportBoxAdministartionService/ORM"
	"github.com/gofrs/uuid"
	"github.com/streadway/amqp"
)

func GetUnitStates(Message amqp.Delivery, ORMs ORMModule.ORMArray) (Data interface{}, Error error) {
	ORMElement, Error := ORMs.FindByName("UnitStatesORM")
	if Error != nil {

		return
	}
	UnitStatesORM := ORMElement.(*ORM.UnitStatesORM)

	if len(Message.Body) != 0 {
		UnitStateUUID, Error := uuid.FromString(string(Message.Body))
		if Error != nil {
			return nil, Error
		}

		return UnitStatesORM.GetUnitState(UnitStateUUID)

	} else {
		return UnitStatesORM.GetUnitStates()
	}

}

func DeleteUnitState(Message amqp.Delivery, ORMs ORMModule.ORMArray) (Data interface{}, Error error) {
	ORMElement, Error := ORMs.FindByName("UnitStatesORM")
	if Error != nil {

		return
	}
	UnitStatesORM := ORMElement.(*ORM.UnitStatesORM)
	UnitTypeUUID, Error := uuid.FromString(string(Message.Body))
	if Error != nil {
		return
	}
	return "Состояние агрегата удалено", UnitStatesORM.DeleteUnitState(UnitTypeUUID)
}

func EditUnitState(Message amqp.Delivery, ORMs ORMModule.ORMArray) (Data interface{}, Error error) {
	ORMElement, Error := ORMs.FindByName("UnitStatesORM")
	if Error != nil {

		return
	}
	UnitStatesORM := ORMElement.(*ORM.UnitStatesORM)
	var NewUnitState ORM.UnitState
	Error = json.Unmarshal(Message.Body, &NewUnitState)
	if Error != nil {

		return
	}
	return UnitStatesORM.EditUnitType(NewUnitState)

}

func AddUnitState(Message amqp.Delivery, ORMs ORMModule.ORMArray) (Data interface{}, Error error) {
	ORMElement, Error := ORMs.FindByName("UnitStatesORM")
	if Error != nil {

		return
	}
	UnitStatesORM := ORMElement.(*ORM.UnitStatesORM)
	var NewUnitState ORM.UnitState
	Error = json.Unmarshal(Message.Body, &NewUnitState)
	if Error != nil {

		return
	}
	return UnitStatesORM.AddUnitState(NewUnitState)

}

package Controllers

import (
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

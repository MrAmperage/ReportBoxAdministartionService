package Controllers

import (
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
		_, Error := uuid.FromString(string(Message.Body))
		if Error != nil {
			return nil, Error
		}
		return UnitTypesORM.GetUnitTypes()

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

package Controllers

import (
	"github.com/MrAmperage/GoWebStruct/WebCore/Modules/ORMModule"
	"github.com/MrAmperage/ReportBoxAdministartionService/ORM"
	"github.com/gofrs/uuid"
	"github.com/streadway/amqp"
)

func GetCargoTypes(Message amqp.Delivery, ORMs ORMModule.ORMArray) (Data interface{}, Error error) {
	ORMElement, Error := ORMs.FindByName("CargoTypesORM")
	if Error != nil {

		return
	}
	CargoTypesORM := ORMElement.(*ORM.CargoTypesORM)

	if len(Message.Body) != 0 {
		CargoTypeUUID, Error := uuid.FromString(string(Message.Body))
		if Error != nil {
			return nil, Error
		}

		return CargoTypesORM.GetCargoType(CargoTypeUUID)

	} else {
		return CargoTypesORM.GetCargoTypes()
	}

}

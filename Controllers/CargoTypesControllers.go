package Controllers

import (
	"encoding/json"

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

func DeleteCargoType(Message amqp.Delivery, ORMs ORMModule.ORMArray) (Data interface{}, Error error) {
	ORMElement, Error := ORMs.FindByName("CargoTypesORM")
	if Error != nil {

		return
	}
	CargoTypesORM := ORMElement.(*ORM.CargoTypesORM)
	UnitTypeUUID, Error := uuid.FromString(string(Message.Body))
	if Error != nil {
		return
	}
	return "Тип груза удален", CargoTypesORM.DeleteCargoType(UnitTypeUUID)
}

func EditCargoType(Message amqp.Delivery, ORMs ORMModule.ORMArray) (Data interface{}, Error error) {
	ORMElement, Error := ORMs.FindByName("CargoTypesORM")
	if Error != nil {

		return
	}
	CargoTypesORM := ORMElement.(*ORM.CargoTypesORM)
	var NewCargoType ORM.CargoType
	Error = json.Unmarshal(Message.Body, &NewCargoType)
	if Error != nil {

		return
	}
	return CargoTypesORM.EditCargoType(NewCargoType)

}

func AddCargoType(Message amqp.Delivery, ORMs ORMModule.ORMArray) (Data interface{}, Error error) {
	ORMElement, Error := ORMs.FindByName("CargoTypesORM")
	if Error != nil {

		return
	}
	CargoTypesORM := ORMElement.(*ORM.CargoTypesORM)
	var NewCargoType ORM.CargoType
	Error = json.Unmarshal(Message.Body, &NewCargoType)
	if Error != nil {

		return
	}

	return CargoTypesORM.AddCargoType(NewCargoType)

}

package Controllers

import (
	"github.com/MrAmperage/GoWebStruct/WebCore/Modules/ORMModule"
	"github.com/MrAmperage/ReportBoxAdministartionService/ORM"
	"github.com/gofrs/uuid"
	"github.com/streadway/amqp"
)

func GetManufacturers(Message amqp.Delivery, ORMs ORMModule.ORMArray) (Data interface{}, Error error) {
	ORMElement, Error := ORMs.FindByName("ManufacturersORM")
	if Error != nil {

		return
	}
	ManufacturersORM := ORMElement.(*ORM.ManufacturersORM)

	if len(Message.Body) != 0 {
		ManufacturersUUID, Error := uuid.FromString(string(Message.Body))
		if Error != nil {
			return nil, Error
		}

		return ManufacturersORM.GetManufacturer(ManufacturersUUID)

	} else {

		return ManufacturersORM.GetManufacturers()
	}

}

func DeleteManufacturer(Message amqp.Delivery, ORMs ORMModule.ORMArray) (Data interface{}, Error error) {
	ORMElement, Error := ORMs.FindByName("ManufacturersORM")
	if Error != nil {

		return
	}
	ManufacturersORM := ORMElement.(*ORM.ManufacturersORM)
	ManufacturerUUID, Error := uuid.FromString(string(Message.Body))
	if Error != nil {
		return
	}
	return "Производитель удален", ManufacturersORM.DeleteManufacturer(ManufacturerUUID)
}

package Controllers

import (
	"github.com/MrAmperage/GoWebStruct/WebCore/Modules/ORMModule"
	"github.com/MrAmperage/ReportBoxAdministartionService/ORM"
	"github.com/gofrs/uuid"
	"github.com/streadway/amqp"
)

func GetTransportTypes(Message amqp.Delivery, ORMs ORMModule.ORMArray) (Data interface{}, Error error) {
	ORMElement, Error := ORMs.FindByName("TransportTypesORM")
	if Error != nil {

		return
	}
	TransportTypesORM := ORMElement.(*ORM.TransportTypesORM)

	if len(Message.Body) != 0 {
		TransportTypeUUID, Error := uuid.FromString(string(Message.Body))
		if Error != nil {
			return nil, Error
		}

		return TransportTypesORM.GetTransportType(TransportTypeUUID)

	} else {
		return TransportTypesORM.GetTransportTypes()
	}

}

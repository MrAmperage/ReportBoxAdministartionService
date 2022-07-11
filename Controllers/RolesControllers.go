package Controllers

import (
	"github.com/MrAmperage/GoWebStruct/WebCore/Modules/ORMModule"
	"github.com/MrAmperage/ReportBoxAdministartionService/ORM"
	"github.com/gofrs/uuid"
	"github.com/streadway/amqp"
)

func GetRoles(Message amqp.Delivery, ORMs ORMModule.ORMArray) (Data interface{}, Error error) {
	ORMElement, Error := ORMs.FindByName("RolesORM")
	if Error != nil {

		return
	}
	RolesORM := ORMElement.(*ORM.RolesORM)
	if len(Message.Body) != 0 {
		UUid, Error := uuid.FromString(string(Message.Body))
		if Error != nil {
			return nil, Error
		}
		return RolesORM.GetRole(UUid)
	} else {
		return RolesORM.GetRoles()
	}

}

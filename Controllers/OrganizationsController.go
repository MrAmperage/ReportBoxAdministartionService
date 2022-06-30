package Controllers

import (
	"github.com/MrAmperage/GoWebStruct/WebCore/Modules/ORMModule"
	"github.com/MrAmperage/ReportBoxAdministartionService/ORM"
	"github.com/gofrs/uuid"
	"github.com/streadway/amqp"
)

func GetOrganizations(Message amqp.Delivery, ORMs ORMModule.ORMArray) (Data interface{}, Error error) {
	ORMElement, Error := ORMs.FindByName("OrganizationsORM")
	if Error != nil {

		return
	}
	OrganizationsORM := ORMElement.(*ORM.OrganizationsORM)

	if len(Message.Body) != 0 {
		OrganizationUUID, Error := uuid.FromString(string(Message.Body))
		if Error != nil {
			return nil, Error
		}

		return OrganizationsORM.GetOrganization(OrganizationUUID)

	} else {
		return OrganizationsORM.GetOrganizations()
	}

}

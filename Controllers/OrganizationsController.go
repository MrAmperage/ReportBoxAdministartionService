package Controllers

import (
	"encoding/json"

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

func DeleteOrganization(Message amqp.Delivery, ORMs ORMModule.ORMArray) (Data interface{}, Error error) {
	ORMElement, Error := ORMs.FindByName("OrganizationsORM")
	if Error != nil {

		return
	}
	OrganizationsORM := ORMElement.(*ORM.OrganizationsORM)
	OrganizationUUID, Error := uuid.FromString(string(Message.Body))
	if Error != nil {
		return
	}
	return "Организация удалена", OrganizationsORM.DeleteOrganization(OrganizationUUID)
}

func EditOrganization(Message amqp.Delivery, ORMs ORMModule.ORMArray) (Data interface{}, Error error) {
	ORMElement, Error := ORMs.FindByName("OrganizationsORM")
	if Error != nil {

		return
	}
	OrganizationsORM := ORMElement.(*ORM.OrganizationsORM)
	var NewOrganization ORM.Organization
	Error = json.Unmarshal(Message.Body, &NewOrganization)
	if Error != nil {

		return
	}
	return OrganizationsORM.EditOrganization(NewOrganization)

}

func AddOrganization(Message amqp.Delivery, ORMs ORMModule.ORMArray) (Data interface{}, Error error) {
	ORMElement, Error := ORMs.FindByName("OrganizationsORM")
	if Error != nil {

		return
	}
	OrganizationsORM := ORMElement.(*ORM.OrganizationsORM)
	var NewOrganization ORM.Organization
	Error = json.Unmarshal(Message.Body, &NewOrganization)
	if Error != nil {

		return
	}

	return OrganizationsORM.AddOrganization(NewOrganization)

}

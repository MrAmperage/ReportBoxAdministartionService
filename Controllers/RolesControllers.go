package Controllers

import (
	"encoding/json"

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
func AddRole(Message amqp.Delivery, ORMs ORMModule.ORMArray) (Data interface{}, Error error) {
	var NewRole ORM.Role
	json.Unmarshal(Message.Body, &NewRole)
	ORMElement, Error := ORMs.FindByName("RolesORM")
	if Error != nil {

		return
	}
	RolesORM := ORMElement.(*ORM.RolesORM)
	return "Роль добавлена", RolesORM.AddRole(NewRole).Error
}
func EditRole(Message amqp.Delivery, ORMs ORMModule.ORMArray) (Data interface{}, Error error) {
	var EditRole ORM.Role
	json.Unmarshal(Message.Body, &EditRole)
	ORMElement, Error := ORMs.FindByName("RolesORM")
	if Error != nil {

		return
	}
	RolesORM := ORMElement.(*ORM.RolesORM)
	return RolesORM.EditRole(EditRole)

}

package Controllers

import (
	"github.com/MrAmperage/GoWebStruct/WebCore/Modules/ORMModule"
	"github.com/MrAmperage/ReportBoxAdministartionService/ORM"
	"github.com/gofrs/uuid"
	"github.com/streadway/amqp"
)

func GetGroups(Message amqp.Delivery, ORMs ORMModule.ORMArray) (Data interface{}, Error error) {
	ORMElement, Error := ORMs.FindByName("GroupsORM")
	if Error != nil {

		return
	}
	GroupsORM := ORMElement.(*ORM.GroupsORM)

	if len(Message.Body) != 0 {
		GroupUUID, Error := uuid.FromString(string(Message.Body))
		if Error != nil {
			return nil, Error
		}

		return GroupsORM.GetGroup(GroupUUID)

	} else {

		return GroupsORM.GetGroups()
	}

}

func DeleteGroup(Message amqp.Delivery, ORMs ORMModule.ORMArray) (Data interface{}, Error error) {
	ORMElement, Error := ORMs.FindByName("GroupsORM")
	if Error != nil {

		return
	}
	GroupsORM := ORMElement.(*ORM.GroupsORM)
	GroupsUUID, Error := uuid.FromString(string(Message.Body))
	if Error != nil {
		return
	}
	return "Группа удалена", GroupsORM.DeleteGroup(GroupsUUID)
}

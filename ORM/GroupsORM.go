package ORM

import (
	"github.com/MrAmperage/GoWebStruct/WebCore/Modules/ORMModule"
	"github.com/gofrs/uuid"
)

type Group struct {
	Id      uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Caption string    `gorm:"not null;caption"`
}

type GroupsORM struct {
	ORMModule.ORM
}

func (GroupsORM *GroupsORM) GetGroups() (Groups []Group, Error error) {
	return Groups, GroupsORM.ConnectionLink.Find(&Groups).Error
}

func (GroupsORM *GroupsORM) GetGroup(UUID uuid.UUID) (Group Group, Error error) {
	Group.Id = UUID
	return Group, GroupsORM.ConnectionLink.Take(&Group).Error
}

package ORM

import (
	"github.com/MrAmperage/GoWebStruct/WebCore/Modules/ORMModule"
	"github.com/gofrs/uuid"
)

type Group struct {
	Id                 uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Caption            string    `gorm:"not null;caption"`
	ShouldersRound     string    `gorm:"not null;shoulders_round"`
	ShouldersPrecision int       `gorm:"not null;default:1;shoulders_precision"`
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

func (GroupsORM *GroupsORM) DeleteGroup(GroupId uuid.UUID) (Error error) {
	return GroupsORM.ConnectionLink.Delete(&Group{Id: GroupId}).Error
}

func (GroupsORM *GroupsORM) EditGroup(NewGroup Group) (Group, error) {
	return NewGroup, GroupsORM.ConnectionLink.Save(&NewGroup).Error
}

func (GroupsORM *GroupsORM) AddGroup(NewGroup Group) (Group, error) {
	return NewGroup, GroupsORM.ConnectionLink.Create(&NewGroup).Error
}

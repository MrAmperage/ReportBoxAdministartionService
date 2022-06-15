package ORM

import (
	"github.com/MrAmperage/GoWebStruct/WebCore/Modules/ORMModule"
	"github.com/gofrs/uuid"
)

type TopMenu struct {
	Id      uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Caption string    `gorm:"caption"`
	Items   []SubMenu `gorm:"foreignkey:TopMenuId"`
}
type TopMenuORM struct {
	ORMModule.ORM
}

func (TopMenuORM *TopMenuORM) GetTopMenu() (TopMenuElements []TopMenu, Error error) {

	return TopMenuElements, TopMenuORM.ConnectionLink.Preload("Items").Find(&TopMenuElements).Error
}

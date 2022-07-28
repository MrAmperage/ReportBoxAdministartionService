package ORM

import (
	"github.com/MrAmperage/GoWebStruct/WebCore/Modules/ORMModule"
)

type TopMenu struct {
	Caption  string `gorm:"not null;caption"`
	Id       string `gorm:"not null;primary_key;id"`
	LeftMenu []LeftMenu
}
type TopMenuORM struct {
	ORMModule.ORM
}

func (TopMenuORM *TopMenuORM) GetTopMenus() (TopMenus []TopMenu, Error error) {
	TopMenuORM.ConnectionLink.Find(&TopMenus)
	for TopMenuIndex, _ := range TopMenus {

		Error = TopMenuORM.ConnectionLink.Model(TopMenus[TopMenuIndex]).Association("LeftMenu").Find(&TopMenus[TopMenuIndex].LeftMenu)
		if Error != nil {
			return
		}

	}
	return TopMenus, Error
}

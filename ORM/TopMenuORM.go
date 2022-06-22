package ORM

import (
	"github.com/MrAmperage/GoWebStruct/WebCore/Modules/ORMModule"
)

type TopMenu struct {
	Caption  string `gorm:"primary_key;not null;caption"`
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
		for LeftMenuIndex, _ := range TopMenus[TopMenuIndex].LeftMenu {
			Error = TopMenuORM.ConnectionLink.Model(TopMenus[TopMenuIndex].LeftMenu[LeftMenuIndex]).Association("Scheme").Find(&TopMenus[TopMenuIndex].LeftMenu[LeftMenuIndex].Scheme)
		}
	}
	return TopMenus, Error
}

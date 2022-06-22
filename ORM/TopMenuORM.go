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
	for Index, _ := range TopMenus {

		Error = TopMenuORM.ConnectionLink.Debug().Model(TopMenus[Index]).Association("LeftMenu").Find(&TopMenus[Index].LeftMenu)
		if Error != nil {
			return
		}
	}
	return TopMenus, Error
}

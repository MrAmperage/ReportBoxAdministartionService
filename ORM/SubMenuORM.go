package ORM

import "github.com/gofrs/uuid"

type SubMenu struct {
	TopMenuId uuid.UUID `gorm:"type:uuid;column:topmenu_id"`
	Caption   string    `gorm:"not null;caption"`
	Id        uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Scheme    Scheme    `gorm:"foreignKey:SubmenuId"`
}

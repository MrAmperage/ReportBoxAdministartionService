package ORM

import "github.com/gofrs/uuid"

type Area struct {
	Id      uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4();id"`
	Caption string    `gorm:"not null;caption"`
}

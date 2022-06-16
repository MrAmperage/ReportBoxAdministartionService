module github.com/MrAmperage/ReportBoxAdministartionService

go 1.18

replace github.com/MrAmperage/GoWebStruct => ../GoWebStruct

require (
	github.com/MrAmperage/GoWebStruct v0.0.0-20220526062049-b22ed9cfcbc8
	github.com/streadway/amqp v1.0.0
)

require (
	github.com/gofrs/uuid v4.2.0+incompatible
	github.com/jackc/pgtype v1.11.0
	gorm.io/driver/postgres v1.3.7 // indirect
	gorm.io/gorm v1.23.6
)

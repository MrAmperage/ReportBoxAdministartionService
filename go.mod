module github.com/MrAmperage/ReportBoxAdministartionService

go 1.18

replace github.com/MrAmperage/GoWebStruct => ../GoWebStruct

require (
	github.com/MrAmperage/GoWebStruct v0.0.0-20220526062049-b22ed9cfcbc8
	github.com/lib/pq v1.10.6
	github.com/streadway/amqp v1.0.0
)

require (
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.12.1 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/pgx/v4 v4.16.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	golang.org/x/crypto v0.0.0-20220622213112-05595931fe9d // indirect
	golang.org/x/text v0.3.7 // indirect
)

require (
	github.com/gofrs/uuid v4.2.0+incompatible
	github.com/jackc/pgtype v1.11.0 // indirect
	gorm.io/driver/postgres v1.3.8 // indirect
	gorm.io/gorm v1.23.8
)

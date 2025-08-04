package interfaces

import db "github.com/tomoliveirabastos/go-rds-db-connector/db"

type DbInterface interface {
	LoadFromEnv(d *db.DatabaseConnect[DbInterface])
}

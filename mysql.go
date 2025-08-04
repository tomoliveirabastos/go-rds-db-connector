package main

import "github.com/tomoliveirabastos/go-rds-db-connector/interfaces"

type MysqlConnector struct {
	Tls        bool
	DbName     string
	DbUser     string
	DbHost     string
	DbPort     string
	DbPassword string
}

func (m *MysqlConnector) LoadFromEnv(d *db.DatabaseConnect[interfaces.DbInterface]) {
	d.DbHost = m.DbHost
	d.DbPort = m.DbPort
	d.DbName = m.DbName
	d.DbUser = m.DbUser
	d.DbPassword = m.DbPassword
}

package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/rds/auth"
	"github.com/tomoliveirabastos/go-rds-db-connector/interfaces"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DatabaseConnect[T interfaces.DbInterface] struct {
	Tls        bool
	DbName     string
	DbUser     string
	DbHost     string
	DbPort     string
	DbPassword string
}

func (d *DatabaseConnect[T]) Connect(dbInterface T) *gorm.DB {

	dbInterface.LoadFromEnv(d)

	stringTls := "true"

	if !d.Tls {
		stringTls = "false"
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?tls=%s&allowCleartextPasswords=true", d.DbUser, d.DbPassword, d.DbHost, d.DbPort, d.DbName, stringTls)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}

func (d *DatabaseConnect[T]) SetAuthenticationToken(ctx context.Context) {
	cfg, err := config.LoadDefaultConfig(ctx)

	d.Tls = true

	if err != nil {
		panic(err)
	}

	dbEndpoint := fmt.Sprintf("%s:%s", d.DbHost, d.DbPort)

	authenticationToken, err := auth.BuildAuthToken(
		ctx, dbEndpoint, cfg.Region, d.DbUser, cfg.Credentials,
	)

	if err != nil {
		panic(err)
	}

	d.DbPassword = authenticationToken
}

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/rds/auth"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DatabaseConnect struct {
	Tls        bool
	DbName     string
	DbUser     string
	DbHost     string
	DbPort     string
	DbPassword string
}

func NewDatabaseConnect() *DatabaseConnect {
	return &DatabaseConnect{
		Tls:        false,
		DbName:     os.Getenv("DB_NAME"),
		DbUser:     os.Getenv("DB_USER"),
		DbHost:     os.Getenv("DB_HOST"),
		DbPort:     os.Getenv("DB_PORT"),
		DbPassword: os.Getenv("DB_PASSWORD"),
	}
}

func (d *DatabaseConnect) Connect() *gorm.DB {
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

func (d *DatabaseConnect) SetAuthenticationToken(ctx context.Context) {
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

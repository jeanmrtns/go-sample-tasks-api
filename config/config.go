package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitConfigs() error {
	fmt.Println("initializing configurations")

	conn, err := GetSQLite()

	if err != nil {
		return err
	}

	DB = conn

	return nil
}

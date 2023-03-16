package database

import (
	"fmt"

	"github.com/MuShaf-NMS/SV-BE-Test/config"
	"github.com/MuShaf-NMS/SV-BE-Test/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func CreateConnection(config *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True", config.DB_User, config.DB_Pass, config.DB_Host, config.DB_Port, config.DB_Name)
	dialect := mysql.Open(dsn)
	connection, err := gorm.Open(dialect)
	if err != nil {
		panic("DB connection failed!")
	}
	err = connection.AutoMigrate(
		entity.Post{},
	)
	if err != nil {
		panic(err)
	}
	return connection
}

func CloseConnection(connection *gorm.DB) {
	db, err := connection.DB()
	if err != nil {
		panic("Get connection failed")
	}
	err = db.Close()
	if err != nil {
		panic("Close connection failed")
	}
}

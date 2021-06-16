package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

// Sqlite instance new DB object with sqlite driver
func Database() (*gorm.DB, error) {
	var host = os.Getenv("DATABASE_HOST")
	var port = os.Getenv("DATABASE_PORT")
	var user = os.Getenv("DATABASE_USER")
	var password = os.Getenv("DATABASE_PASSWORD")
	var dbname = os.Getenv("DATABASE_NAME")

	//var dbInfo = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslMode)
	var dbInfo = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, dbname)

	db, err := gorm.Open(mysql.Open(dbInfo), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// start sockets listening
	if err := autoMigrate(db); err != nil {
		return nil, err
	}

	return db, nil
}

// autoMigrate create tables for needed models
func autoMigrate(db *gorm.DB) error {
	//Создаем таблицу
	if os.Getenv("DROP_TABLES") == "yes" {
		_ = db.Migrator().DropTable(&User{})
		_ = db.Migrator().DropTable(&Stats{})
		_ = db.Migrator().DropTable(&Invoice{})
		_ = db.Migrator().DropTable(&Rating{})
		_ = db.Migrator().DropTable(&Share{})
		_ = db.Migrator().DropTable(&UserHistory{})
	}
	if os.Getenv("CREATE_TABLE") == "yes" {
		if err := db.AutoMigrate(&User{}); err != nil {
			fmt.Println(err)
		}
		if err := db.AutoMigrate(&Stats{}); err != nil {
			fmt.Println(err)
		}
		if err := db.AutoMigrate(&Invoice{}); err != nil {
			fmt.Println(err)
		}
		if err := db.AutoMigrate(&Rating{}); err != nil {
			fmt.Println(err)
		}
		if err := db.AutoMigrate(&Share{}); err != nil {
			fmt.Println(err)
		}
		if err := db.AutoMigrate(&UserHistory{}); err != nil {
			fmt.Println(err)
		}
	}

	return nil
}

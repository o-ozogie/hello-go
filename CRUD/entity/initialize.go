package entity

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"
	"sync"
)

type Database struct {
	Username   string
	Password   string
	Port       int
	Host       string
	Name       string
}

func (database Database) databaseUrl() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		database.Username,
		database.Password,
		database.Host,
		database.Port,
		database.Name)
}

var instance *Database
var connection *gorm.DB

func New() *Database {
	once := new(sync.Once)
	once.Do(func() {
		instance = &Database{Username: os.Getenv("DATABASE_USERNAME"),
			Password: os.Getenv("DATABASE_PASSWORD"),
			Host: os.Getenv("DATABASE_HOST"),
			Port: func() int {
				p, e := strconv.Atoi(os.Getenv("PORT"))
				if e != nil {
					p = 3306
				}
				return p
			}(),
			Name: os.Getenv("DATABASE_NAME")}
	})

	return instance
}

func (database *Database) Connect() {
	var e error

	connection, e = gorm.Open(mysql.Open(database.databaseUrl()), &gorm.Config{})
	if e != nil {
		log.Panic(e)
	}
}

func (database Database) Migrate() {
	e := connection.AutoMigrate(new(User), new(Post), new(Comment))
	if e != nil {
		log.Panic(e)
	}
}

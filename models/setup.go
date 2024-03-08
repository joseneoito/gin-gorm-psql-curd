package models


import (
    "gorm.io/gorm"
    "gorm.io/driver/postgres"
)

var DB *gorm.DB

func ConnectDatabase(){
    dbconnection := "host=localhost user=jose2007kj password=jose2007kj dbname=kunjamma port=5432 sslmode=disable TimeZone=Asia/Kolkata"
    db, err := gorm.Open(postgres.Open(dbconnection), &gorm.Config{})
    if err != nil {
        panic("Something went wrong")
    }
    err = db.AutoMigrate(&Blog{})
    if err !=nil {
        return
    }
    DB = db
}

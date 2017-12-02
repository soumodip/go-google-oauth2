package controller

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
	"github.com/labstack/echo" 
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

type ObjectGo struct {
	ID int `sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Field_1 string
	Field_2 string
	Field_3 string
	Field_4 string
}

func InsertIntoDB(context echo.Context, db *gorm.DB) error{
	db.AutoMigrate(&ObjectGo{})
	db.Create(&ObjectGo{Field_1: "field", Field_2: "field", Field_3: "field", Field_4: "field"})
	return context.String(http.StatusOK, "OBJECT_ADDED")
}

func ReadFromDB(context echo.Context, db *gorm.DB) error{
	db.AutoMigrate(&ObjectGo{})
	var queriedObject ObjectGo
	db.First(&queriedObject, 1)
	return context.String(http.StatusOK, queriedObject.Field_1)
}

func InsertIntoDBSql(context echo.Context, db *sql.DB) error{
	_, err := db.Exec("INSERT INTO `object_gos` (`id`, `field_1`, `field_2`, `field_3`, `field_4`) VALUES (NULL, 'field', 'field', 'field', 'field');")
	if err!= nil {
		panic(err)
	}
	return context.String(http.StatusOK, "OBJECT INSERTED")
}

func ReadFromDBSql(context echo.Context, db *sql.DB) error{
	res, err := db.Exec("SELECT * FROM `object_gos`;")
	if err!= nil {
		panic(err)
	}
	rowCount, _ := res.RowsAffected()
	str := strconv.Itoa(int(rowCount))
	return context.String(http.StatusOK, str)
}
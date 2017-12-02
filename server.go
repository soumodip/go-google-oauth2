package main

import (
	"github.com/labstack/echo" 
    "database/sql"
	_ "github.com/go-sql-driver/mysql"
	"./controller"
	"./constants"
)

func main() {
	server := echo.New()
	db, _ := sql.Open("mysql", "root:@/meesho")
	defer db.Close()

	//DECLARE THE ROUTES
	server.GET("/google/authenticate", func(context echo.Context) error {
		return controller.RedirectToGoogleOAuth(context)
	})
	server.GET("/google/authenticate/callback", func(context echo.Context) error {
		return controller.ReadGoogleOAuthData(context)
	})
	server.GET("/field/insert/gosql", func(context echo.Context) error {
		return controller.InsertIntoDBSql(context, db)
	})
	server.GET("/field/read/gosql", func(context echo.Context) error {
		return controller.ReadFromDBSql(context, db)
	})

	//START THE SERVER
	var PORT string = constants.ParseConstants()["PORT"]
	server.Logger.Fatal(server.Start(":" + PORT))
}
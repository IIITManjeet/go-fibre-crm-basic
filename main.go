package main

import (
	"fmt"

	"github.com/IIITManjeet/go-fibre-crm-basic/database"
	"github.com/IIITManjeet/go-fibre-crm-basic/lead"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func setupRoutes(app *fiber.App){
	app.Get("/api/v1/lead",lead.GetLeads)
	app.Get("/api/v1/lead/:id",lead.GetLead)
	app.Post("/api/v1/lead",lead.NewLead)
	app.Delete("/api/v1/lead/:id",lead.DeleteLead)
}

func initDatabase(){
 var err error
 database.DBConn, err= gorm.Open("sqlite3","leads.db")
 if err != nil{
	panic("Failed to Connect to Database")
 }
 fmt.Println("Connection opened to Database")
 database.DBConn.AutoMigrate(&lead.Lead{})
 fmt.Println("Database Migrated")
}

func main(){
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(3000)
	defer database.DBConn.Close()
}
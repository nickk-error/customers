package main

import (
	"customer/common"
	"customer/customer"
	"customer/db"
	"customer/logger"
	"time"

	"github.com/labstack/echo/v4"
	config "github.com/spf13/viper"
	"github.com/tylerb/graceful"
)

var log = logger.NewLogger()

func main() {

	transID := common.NewUUID()

	err := common.LoadConfigFile(transID)
	if err != nil {
		log.Error(transID, "LoadConfigFile from yml file error: "+err.Error())
		panic("LoadConfigFile from yml file error: " + err.Error())
	}

	e := echo.New()
	r := e.Group(config.GetString("service.endpoint"))
	r.POST("/createCustomer", customer.CreateCustomer)
	r.GET("/getCustomerByID/:customerID", customer.GetCustomerByID)
	r.PUT("/updateCustomer/:customerID", customer.UpdateCustomer)
	r.DELETE("/deleteCustomer/:customerID", customer.DeleteCustomer)

	createInitialData()

	e.Server.Addr = ":" + config.GetString("service.port")
	graceful.ListenAndServe(e.Server, 5*time.Second)
}

func createInitialData() {
	// Create some initial customer data
	transID := common.NewUUID()
	sqLite := db.ConnectSQLite(transID)
	sqLite.Create(&customer.Customers{Name: "Dollar TheCat", Age: 4})
	sqLite.Create(&customer.Customers{Name: "Million TheCat", Age: 4})
	sqLite.Create(&customer.Customers{Name: "Berlin TheCat", Age: 2})
	sqLite.Create(&customer.Customers{Name: "Vegas TheCat", Age: 2})
	sqLite.Create(&customer.Customers{Name: "Brooklyn TheCat", Age: 2})

	sqLite.Close()
}

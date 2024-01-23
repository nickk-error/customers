package customer

import (
	"customer/common"
	"customer/constant"
	"customer/db"
	"customer/logger"
	"net/http"

	"github.com/labstack/echo/v4"
)

var log = logger.NewLogger()

// --------------------------------------------------------------------------CreateCustomer--------------------------------------------------------------------------
func CreateCustomer(c echo.Context) error {
	transID := common.NewUUID()
	log.Info(transID, "Start Process CreateCustomer")

	response := new(ResponseCustomers)
	request := new(Customers)
	if err := c.Bind(&request); err != nil {
		log.Error(transID, "Can't Bind Json To Model | Exception : ", err.Error())
		return c.JSON(http.StatusInternalServerError, common.MapErrorCode(transID, constant.ApplicationError, "Can't Bind Json To Model | Exception : "+err.Error()))
	}
	log.Info(transID, "CreateCustomer Request :", request)

	sqLite := db.ConnectSQLite(transID)
	sqLite.AutoMigrate(&Customers{})
	sqLite.Create(&request)
	sqLite.Close()

	response.ResponseBean.Code = constant.SuccessCode
	response.ResponseBean.TransID = transID
	response.ResponseBean.Msg = constant.Success
	response.Customers = *request

	log.Info(transID, "CreateCustomer Response ", response)
	return c.JSON(http.StatusOK, response)
}

// --------------------------------------------------------------------------GetCustomerByID--------------------------------------------------------------------------
func GetCustomerByID(c echo.Context) error {
	transID := common.NewUUID()
	log.Info(transID, "Start Process GetCustomerByID")

	response := new(ResponseCustomers)
	var customer Customers

	customerID := c.Param("customerID")
	log.Info(transID, "customerID :", customerID)

	sqLite := db.ConnectSQLite(transID)
	if err := sqLite.First(&customer, customerID).Error; err != nil {
		log.Error(transID, "Query CustomerID : ", err.Error())
		return c.JSON(http.StatusNotFound, common.MapErrorCode(transID, constant.SQLException, "GetCustomerByID Error : "+err.Error()))
	}
	sqLite.Close()

	response.ResponseBean.Code = constant.SuccessCode
	response.ResponseBean.TransID = transID
	response.ResponseBean.Msg = constant.Success
	response.Customers = customer

	log.Info(transID, "GetCustomerByID Response ", response)
	return c.JSON(http.StatusOK, response)
}

// --------------------------------------------------------------------------UpdateCustomer--------------------------------------------------------------------------
func UpdateCustomer(c echo.Context) error {
	transID := common.NewUUID()
	log.Info(transID, "Start Process UpdateCustomer")

	response := new(ResponseCustomers)
	var customer Customers

	customerID := c.Param("customerID")
	log.Info(transID, "customerID :", customerID)

	sqLite := db.ConnectSQLite(transID)
	if err := sqLite.First(&customer, customerID).Error; err != nil {
		log.Error(transID, "Update Customer : ", err.Error())
		return c.JSON(http.StatusNotFound, common.MapErrorCode(transID, constant.SQLException, "UpdateCustomer Error : "+err.Error()))
	}
	c.Bind(&customer)
	sqLite.Save(&customer)
	sqLite.Close()

	response.ResponseBean.Code = constant.SuccessCode
	response.ResponseBean.TransID = transID
	response.ResponseBean.Msg = constant.Success
	response.Customers = customer

	log.Info(transID, "UpdateCustomer Response ", response)
	return c.JSON(http.StatusOK, response)
}

// --------------------------------------------------------------------------DeleteCustomer--------------------------------------------------------------------------
func DeleteCustomer(c echo.Context) error {
	transID := common.NewUUID()
	log.Info(transID, "Start Process DeleteCustomer")

	response := new(ResponseCustomers)
	var customer Customers

	customerID := c.Param("customerID")
	log.Info(transID, "customerID :", customerID)

	sqLite := db.ConnectSQLite(transID)
	if err := sqLite.Where("id = ?", customerID).Delete(&customer).Error; err != nil {
		log.Error(transID, "DeleteCustomer : ", err.Error())
		return c.JSON(http.StatusNotFound, common.MapErrorCode(transID, constant.SQLException, "DeleteCustomer Error : "+err.Error()))
	}
	sqLite.Close()

	response.ResponseBean.Code = constant.SuccessCode
	response.ResponseBean.TransID = transID
	response.ResponseBean.Msg = "id #" + customerID + " is deleted."

	log.Info(transID, "DeleteCustomer Response ", response)
	return c.JSON(http.StatusOK, response)
}

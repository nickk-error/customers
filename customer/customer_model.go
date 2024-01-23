package customer

import "customer/common"

type Customers struct {
	ID   uint   `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type ResponseCustomers struct {
	common.ResponseBean
	Customers Customers `json:"customer"`
}

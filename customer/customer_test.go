package customer

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateCustomer(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/customers", bytes.NewBuffer([]byte(`{"name":"Test User","age":25}`)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, CreateCustomer(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestUpdateCustomer(t *testing.T) {
	e := echo.New()
	reqCreate := httptest.NewRequest(http.MethodPost, "/Customer/createCustomer", bytes.NewBuffer([]byte(`{"name":"Test User","age":25}`)))
	reqCreate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recCreate := httptest.NewRecorder()
	cCreate := e.NewContext(reqCreate, recCreate)
	CreateCustomer(cCreate)

	reqUpdate := httptest.NewRequest(http.MethodPut, "/Customer/updateCustomer/1", bytes.NewBuffer([]byte(`{"name":"Updated User","age":30}`)))
	reqUpdate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recUpdate := httptest.NewRecorder()
	cUpdate := e.NewContext(reqUpdate, recUpdate)

	if assert.NoError(t, UpdateCustomer(cUpdate)) {
		assert.Equal(t, http.StatusOK, recUpdate.Code)
	}
}

func TestGetCustomerByID(t *testing.T) {
	e := echo.New()
	reqCreate := httptest.NewRequest(http.MethodPost, "/Customer/createCustomer", bytes.NewBuffer([]byte(`{"name":"Test User","age":25}`)))
	reqCreate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recCreate := httptest.NewRecorder()
	cCreate := e.NewContext(reqCreate, recCreate)
	CreateCustomer(cCreate)

	reqGet := httptest.NewRequest(http.MethodGet, "/Customer/getCustomerByID/1", nil)
	reqGet.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recGet := httptest.NewRecorder()
	cGet := e.NewContext(reqGet, recGet)

	if assert.NoError(t, UpdateCustomer(cGet)) {
		assert.Equal(t, http.StatusOK, recGet.Code)
	}

}

func TestDeleteCustomer(t *testing.T) {
	e := echo.New()
	reqCreate := httptest.NewRequest(http.MethodPost, "/Customer/createCustomer", bytes.NewBuffer([]byte(`{"name":"Test User","age":25}`)))
	reqCreate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recCreate := httptest.NewRecorder()
	cCreate := e.NewContext(reqCreate, recCreate)
	CreateCustomer(cCreate)

	reqDelete := httptest.NewRequest(http.MethodDelete, "/Customer/deleteCustomer/1", nil)
	recDelete := httptest.NewRecorder()
	cDelete := e.NewContext(reqDelete, recDelete)

	if assert.NoError(t, DeleteCustomer(cDelete)) {
		assert.Equal(t, http.StatusOK, recDelete.Code)
	}
}

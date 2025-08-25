//go:build unit

package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kalilventura/vehicle-management-payment/internal/payments/infrastructure/controllers"
	builders2 "github.com/kalilventura/vehicle-management-payment/test/payments/domain/builders"
	"github.com/kalilventura/vehicle-management-payment/test/payments/domain/commands"
	"github.com/kalilventura/vehicle-management-payment/test/payments/infrastructure/builders"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestProcessPaymentController(t *testing.T) {
	const route = "/v1/payments"
	t.Run("should respond BadRequest when the body is invalid", func(t *testing.T) {
		// given
		data := "{invalid"
		requestBodyBytes, _ := json.Marshal(data)
		reqBody := bytes.NewBuffer(requestBodyBytes)

		command := commands.NewSavePaymentCommandStub()
		controller := controllers.NewProcessPaymentController(command)

		router := echo.New()
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost, route, reqBody)
		request.Header.Set("Content-Type", "application/json")
		router.POST(route, controller.Execute)

		// when
		router.ServeHTTP(recorder, request)

		// then
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("should respond Internal Server Error due an unexpected error", func(t *testing.T) {
		// given
		command := commands.NewSavePaymentCommandStub().WithOnInternalServerError()
		controller := controllers.NewProcessPaymentController(command)

		requestBody := builders.NewProcessPaymentRequestBuilder().
			WithValidData().
			BuildRequest()

		router := echo.New()
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost, route, requestBody)
		request.Header.Set("Content-Type", "application/json")
		router.POST(route, controller.Execute)

		// when
		router.ServeHTTP(recorder, request)

		// then
		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("should respond Accepted when the payment was processed", func(t *testing.T) {
		// given
		payment := builders2.NewPaymentBuilder().Build()
		command := commands.NewSavePaymentCommandStub().WithOnSuccess(&payment)
		controller := controllers.NewProcessPaymentController(command)

		requestBody := builders.NewProcessPaymentRequestBuilder().
			WithValidData().
			BuildRequest()

		router := echo.New()
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost, route, requestBody)
		request.Header.Set("Content-Type", "application/json")
		router.POST(route, controller.Execute)

		// when
		router.ServeHTTP(recorder, request)

		// then
		assert.Equal(t, http.StatusAccepted, recorder.Code)
	})

	t.Run("should return the metadata", func(t *testing.T) {
		// given
		controller := controllers.NewProcessPaymentController(nil)

		// when
		metadata := controller.GetBind()

		// then
		assert.Equal(t, "/payments", metadata.RelativePath)
		assert.Equal(t, http.MethodPost, metadata.Method)
	})
}

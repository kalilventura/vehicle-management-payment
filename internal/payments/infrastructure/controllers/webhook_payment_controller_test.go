//go:build unit

package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kalilventura/vehicle-management-payment/internal/payments/infrastructure/controllers"
	"github.com/kalilventura/vehicle-management-payment/test/payments/domain/builders"
	"github.com/kalilventura/vehicle-management-payment/test/payments/domain/commands"
	"github.com/kalilventura/vehicle-management-payment/test/payments/infrastructure/services"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestWebhookPaymentController(t *testing.T) {
	const route = "/v1/payments"
	t.Run("should respond Bad Request when the request is invalid", func(t *testing.T) {
		// given
		data := "{}"
		requestBodyBytes, _ := json.Marshal(data)
		reqBody := bytes.NewBuffer(requestBodyBytes)

		webhookProcessor := services.NewWebhookProcessorStub().WithError()
		updateCommand := commands.NewUpdatePaymentCommandStub()
		controller := controllers.NewWebhookPaymentController(updateCommand, webhookProcessor)

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

	t.Run("should respond Bad Request when the update payment fails", func(t *testing.T) {
		// given
		data := "{}"
		requestBodyBytes, _ := json.Marshal(data)
		reqBody := bytes.NewBuffer(requestBodyBytes)

		updatePayment := builders.NewUpdatePaymentBuilder().Build()
		webhookProcessor := services.NewWebhookProcessorStub().WithSuccess(&updatePayment)
		updateCommand := commands.NewUpdatePaymentCommandStub().WithOnInternalServerError()
		controller := controllers.NewWebhookPaymentController(updateCommand, webhookProcessor)

		router := echo.New()
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost, route, reqBody)
		request.Header.Set("Content-Type", "application/json")
		router.POST(route, controller.Execute)

		// when
		router.ServeHTTP(recorder, request)

		// then
		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("should respond OK when the request was processed", func(t *testing.T) {
		// given
		data := "{}"
		requestBodyBytes, _ := json.Marshal(data)
		reqBody := bytes.NewBuffer(requestBodyBytes)

		updatePayment := builders.NewUpdatePaymentBuilder().Build()

		webhookProcessor := services.NewWebhookProcessorStub().WithSuccess(&updatePayment)
		updateCommand := commands.NewUpdatePaymentCommandStub().WithOnSuccess()
		controller := controllers.NewWebhookPaymentController(updateCommand, webhookProcessor)

		router := echo.New()
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost, route, reqBody)
		request.Header.Set("Content-Type", "application/json")
		router.POST(route, controller.Execute)

		// when
		router.ServeHTTP(recorder, request)

		// then
		assert.Equal(t, http.StatusOK, recorder.Code)
	})

	t.Run("should return the metadata", func(t *testing.T) {
		// given
		controller := controllers.NewWebhookPaymentController(nil, nil)

		// when
		metadata := controller.GetBind()

		// then
		assert.Equal(t, "/payments/webhook", metadata.RelativePath)
		assert.Equal(t, http.MethodPost, metadata.Method)
	})
}

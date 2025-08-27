//go:build integration

package repositories_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/kalilventura/vehicle-management-payment/internal/payments/infrastructure/repositories"
	"github.com/kalilventura/vehicle-management-payment/internal/shared/infrastructure/configuration"
	"github.com/kalilventura/vehicle-management-payment/test/payments/domain/builders"
	"github.com/kalilventura/vehicle-management-payment/test/shared/infrastructure"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"gorm.io/gorm"
)

type GormPaymentsRepositoryTestSuite struct {
	suite.Suite
	postgresContainer testcontainers.Container
	db                *gorm.DB
}

func (suite *GormPaymentsRepositoryTestSuite) SetupSuite() {
	ctx := context.Background()

	container, err := infrastructure.SetupPostgres(ctx)
	suite.Require().NoError(err)
	suite.postgresContainer = container

	settings, err := infrastructure.CreateDatabaseSettings(ctx, container)
	suite.Require().NoError(err)

	suite.db = configuration.NewDatabaseClient(settings)

	err = infrastructure.CreateDatabaseStructure(suite.db)
	suite.Require().NoError(err)
}

func (suite *GormPaymentsRepositoryTestSuite) TearDownSuite() {
	err := testcontainers.TerminateContainer(suite.postgresContainer)
	suite.Require().NoError(err)
	suite.T().Logf("Stopped postgres container")
}

func (suite *GormPaymentsRepositoryTestSuite) TestSuccessfully() {
	suite.Run("should create a new payment successfully", func() {
		// given
		data := "{'id': 1}"
		rawResponse, _ := json.Marshal(data)

		payTransaction := builders.NewPaymentTransactionBuilder().WithRawResponse(rawResponse).Build()
		payment := builders.NewPaymentBuilder().WithAmount(100000).WithPaymentTransaction(&payTransaction).Build()

		transaction := suite.db.Begin()
		defer transaction.Rollback()

		repository := repositories.NewGormPaymentsRepository(transaction)

		// when
		err := repository.Save(&payment)

		// then
		suite.Require().NoError(err)
	})

	suite.Run("should update the payment successfully", func() {
		// given
		data := "{'id': 1}"
		rawResponse, _ := json.Marshal(data)

		payTransaction := builders.NewPaymentTransactionBuilder().WithRawResponse(rawResponse).Build()

		payment := builders.NewPaymentBuilder().
			WithAmount(100000).
			WithPaymentTransaction(&payTransaction).
			Build()
		update := builders.NewUpdatePaymentBuilder().
			WithGatewayPaymentID(payTransaction.GatewayTransactionID).
			WithPaymentTransaction(&payTransaction).
			Build()

		transaction := suite.db.Begin()
		defer transaction.Rollback()

		repository := repositories.NewGormPaymentsRepository(transaction)
		err := repository.Save(&payment)
		suite.Require().NoError(err)

		// when
		err = repository.Update(&update)

		// then
		suite.Require().NoError(err)
	})
}

func (suite *GormPaymentsRepositoryTestSuite) TestError() {
	suite.Run("should return an error when try to create a payment", func() {
		// given
		payment := builders.NewPaymentBuilder().WithAmount(-90).Build()

		transaction := suite.db.Begin()
		defer transaction.Rollback()

		repository := repositories.NewGormPaymentsRepository(transaction)

		// when
		err := repository.Save(&payment)

		// then
		suite.Require().Error(err)
	})

	suite.Run("should return an error when try to create a payment", func() {
		// given
		payment := builders.NewUpdatePaymentBuilder().Build()

		transaction := suite.db.Begin()
		defer transaction.Rollback()

		repository := repositories.NewGormPaymentsRepository(transaction)

		// when
		err := repository.Update(&payment)

		// then
		suite.Require().Error(err)
	})
}

func TestGormPaymentsRepository(t *testing.T) {
	suite.Run(t, new(GormPaymentsRepositoryTestSuite))
}

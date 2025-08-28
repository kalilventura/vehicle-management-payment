//go:build integration

package services_test

import (
	"context"
	"testing"

	"github.com/kalilventura/vehicle-management-payment/internal/shared/domain/entities"
	"github.com/kalilventura/vehicle-management-payment/internal/shared/infrastructure/configuration"
	"github.com/kalilventura/vehicle-management-payment/internal/shared/infrastructure/services"
	"github.com/kalilventura/vehicle-management-payment/test/shared/infrastructure"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"gorm.io/gorm"
)

type GooseMigrationServiceSuite struct {
	suite.Suite
	ctx               context.Context
	postgresContainer *postgres.PostgresContainer
	settings          *entities.DatabaseSettings
	db                *gorm.DB
}

func (suite *GooseMigrationServiceSuite) SetupSuite() {
	ctx := context.Background()
	suite.ctx = ctx

	container, err := infrastructure.SetupPostgres(ctx)
	suite.Require().NoError(err)
	suite.postgresContainer = container

	settings, err := infrastructure.CreateDatabaseSettings(ctx, container)
	suite.Require().NoError(err)

	suite.settings = settings

	suite.db = configuration.NewDatabaseClient(settings)
}

func (suite *GooseMigrationServiceSuite) TearDownSuite() {
	err := testcontainers.TerminateContainer(suite.postgresContainer)
	suite.Require().NoError(err)
	suite.T().Logf("Stopped postgres container")
}

func (suite *GooseMigrationServiceSuite) TestGormMigrationServiceSuccessfully() {
	suite.Run("should run the migrations successfully", func() {
		// given
		service := services.NewGooseMigrationService(suite.db, suite.settings)

		// when
		err := service.Run("../../../../db/migrations")

		// then
		suite.NoError(err)
	})
}

func (suite *GooseMigrationServiceSuite) TestGormMigrationServiceInvalidPathError() {
	suite.Run("should return an error when the migration path is invalid", func() {
		// given
		service := services.NewGooseMigrationService(suite.db, suite.settings)

		// when
		// Use a directory that does not contain migration files
		err := service.Run("../")

		// then
		suite.Error(err)
		suite.Contains(err.Error(), "failed to run migrations")
	})
}

func (suite *GooseMigrationServiceSuite) TestGormMigrationServiceConnectionError() {
	suite.Run("should return an error when there is a database connection issue", func() {
		// given
		env, err := infrastructure.GetDatabaseEnvSettings(suite.ctx, suite.postgresContainer)
		suite.Require().NoError(err)

		invalidSettings := entities.NewDatabaseSettings(
			env.DbHost,
			env.DbName,
			"123",
			env.DbUser,
			env.DbPassword,
			"disable",
		)
		service := services.NewGooseMigrationService(suite.db, invalidSettings)

		// when
		svcErr := service.Run("../../../../db/migrations")

		// then
		suite.Error(svcErr)
	})
}

func (suite *GooseMigrationServiceSuite) TestGormMigrationServiceInvalidDSNError() {
	suite.Run("should return an error when the database DSN is invalid", func() {
		// given
		invalidSettings := entities.NewDatabaseSettings(
			"localhost malformed",
			"db",
			"5432",
			"user",
			"password",
			"disable",
		)
		service := services.NewGooseMigrationService(suite.db, invalidSettings)

		// when
		err := service.Run("../../../../db/migrations")

		// then
		suite.Error(err)
	})
}

func TestGooseMigrationService(t *testing.T) {
	suite.Run(t, new(GooseMigrationServiceSuite))
}

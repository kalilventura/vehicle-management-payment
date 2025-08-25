//go:build unit

package entities_test

import (
	"fmt"
	"testing"

	"github.com/kalilventura/vehicle-management-payment/internal/shared/domain/entities"
	"github.com/stretchr/testify/assert"
)

func TestDatabaseSettings(t *testing.T) {
	t.Run("should create the DSN successfully", func(t *testing.T) {
		// given
		host := "localhost"
		name := "testdb"
		port := "5432"
		user := "testuser"
		password := "secret"
		dbSSL := "disable"

		settings := entities.NewDatabaseSettings(host, name, port, user, password, dbSSL)
		expectedDSN := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
			host, user, password, name, port, dbSSL,
		)

		// when
		actualDSN := settings.GetDSN()

		// then
		assert.Equal(t, expectedDSN, actualDSN)
	})
}

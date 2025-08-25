//go:build unit

package entities_test

import (
	"testing"

	"github.com/kalilventura/vehicle-management-payment/internal/shared/domain/entities"
	"github.com/stretchr/testify/assert"
)

func TestSettings(t *testing.T) {
	t.Run("should return the port successfully", func(t *testing.T) {
		// given
		settings := entities.Settings{
			Port: 8098,
		}

		// when
		port := settings.GetPort()

		// then
		assert.Equal(t, ":8098", port)
	})
}

//go:build unit

package entities_test

import (
	"testing"

	"github.com/kalilventura/vehicle-management-payment/internal/shared/domain/entities"
	"github.com/stretchr/testify/assert"
)

func TestControllerBind(t *testing.T) {
	t.Run("should return the full path successfully", func(t *testing.T) {
		// given
		version := "/v1"
		relativePath := "/users"
		cb := entities.ControllerBind{
			Version:      version,
			RelativePath: relativePath,
		}

		// when
		path := cb.GetFullPath()

		// then
		expectedPath := "/v1/users"
		assert.Equal(t, expectedPath, path)
	})
}

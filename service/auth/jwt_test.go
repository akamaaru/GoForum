package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateJWT(t *testing.T) {
	t.Run("valid test", func(t *testing.T) {
		token, err := CreateJWT([]byte("secret"), 3)
		if err != nil {
			t.Errorf("failed to create JWT token: %v", err)
		}

		assert.NotNil(t, token)
		assert.NotEqual(t, "", token)
	})
}
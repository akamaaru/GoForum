package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestHashPassword(t *testing.T) {
	t.Run("valid test", func(t *testing.T) {
		passwords := []string {"password1", "password2", "password3"}
		for _, password := range passwords {
			hash, err := HashPassword(password)

			if err != nil {
				t.Errorf("failed hashing password")
				return
			}

			assert.NotEqual(t, hash, password)
			assert.NotEqual(t, hash, "")

			err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) 
			
			assert.Nil(t, err)
		}
	})
}

func TestComparePasswords(t *testing.T) {
	t.Run("valid test", func(t *testing.T) {
		passwords := [][2]string {
			// equal
			{"password1", "password1"}, 
			{"password2", "password2"}, 
			{"password3", "password3"}, 

			// not equal
			{"password1", "password2"}, 
			{"password2", "password3"}, 
			{"password3", "password1"}, 
		}

		for _, pair := range passwords {
			assert.Equal(t, 
				bcrypt.CompareHashAndPassword([]byte(pair[0]), []byte(pair[1])) == nil, 
				ComparePasswords(pair[0], []byte(pair[1])),
			)
		}
	})
}
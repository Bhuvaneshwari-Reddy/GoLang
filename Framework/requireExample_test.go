package framework

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUserCreation(t *testing.T) {
	u, err := CreateUser("John", "john@example.com")
	require.NoError(t, err, "user creation should not fail")
	require.NotNil(t, u, "user should not fail")

	require.Equal(t, "John", u.Name, "user name should match")
	require.Equal(t, "john@example.com", u.Email, "user email should match")
}


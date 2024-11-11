package testutil

import (
	"testing"
)

const (
	DatabaseUrl = "DATABASE_URL"
)

func SetEnv(t *testing.T) string {
	t.Setenv("DATABASE_URL", DatabaseUrl)
	return DatabaseUrl
}

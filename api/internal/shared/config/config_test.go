package config_test

import (
	"testing"

	"api/internal/shared/config"
	"api/internal/shared/testutil"
)

func TestConfig(t *testing.T) {
	databaseUrl := testutil.SetEnv(t)
	config, err := config.NewConfigFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	if config.DatabaseUrl != databaseUrl {
		t.Error("DatabaseUrl is not match")
	}
}

package database_test

import (
	"testing"

	"api/internal/shared/config"
	"api/internal/shared/database"
)

const url = "postgres://sa:Sa01@postgres/push_notification?sslmode=disable"

func TestPool(t *testing.T) {
	t.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "")
	t.Setenv("DATABASE_URL", url)
	config, err := config.NewConfigFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	_, close, err := database.Pool(config)
	defer close()
	if err != nil {
		t.Fatal(err)
	}

}

func TestTx(t *testing.T) {
	t.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "")
	t.Setenv("DATABASE_URL", url)
	config, err := config.NewConfigFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	pool, close, err := database.Pool(config)
	defer close()
	if err != nil {
		t.Fatal(err)
	}
	// Rollback
	tx, err := database.Begin(pool)
	if err != nil {
		t.Fatal(err)
	}
	err = tx.Rollback()
	if err != nil {
		t.Fatal(err)
	}
	// 2回目の実行はエラーになる
	err = tx.Rollback()
	if err == nil {
		t.Fatal("Tx is not released.")
	}
	// Commit
	tx, err = database.Begin(pool)
	if err != nil {
		t.Fatal(err)
	}
	err = tx.Commit()
	if err != nil {
		t.Fatal(err)
	}
	// 2回目の実行はエラーになる
	err = tx.Commit()
	if err == nil {
		t.Fatal("Tx is not released.")
	}
}

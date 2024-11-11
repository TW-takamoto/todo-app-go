package database_test

import (
	"reflect"
	"testing"

	"api/internal/domain/model"
	"api/internal/infrastructure/database"
	"api/internal/shared/config"
	d "api/internal/shared/database"
	"api/internal/shared/testutil"
)

const url = "postgres://sa:Sa01@postgres/todos:postgres?sslmode=disable"

func TestTodoDatabaseRepository(t *testing.T) {
	testutil.SetEnv(t)
	config, err := config.NewConfigFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	pool, close, err := d.Pool(config)
	defer close()
	if err != nil {
		t.Fatal(err)
	}
	tx, err := d.Begin(pool)
	defer tx.Rollback()
	repo := database.NewTodoDatabaseRepository(tx)
	id := int64(1)
	title := "test"
	description := "test"
	completed := false
	builder, err := model.TodoBuilder(id, title, description, completed)
	if err != nil {
		t.Fatal(err)
	}
	todo := builder.Build()
	// 作成
	err = repo.Create(todo)
	if err != nil {
		t.Fatal(err)
	}
	// 取得
	res, err := repo.Find(id)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(res, builder.Build()) {
		t.Fatal("not equal")
	}
	// 更新
	err = repo.Update(todo)
	if err != nil {
		t.Fatal(err)
	}
	// 削除
	err = repo.Delete(id)
	if err != nil {
		t.Fatal(err)
	}
	// 削除確認
	res, err = repo.Find(id)
	if err != nil {
		t.Fatal(err)
	}
	if res.Id() != 0 {
		t.Fatal("not deleted")
	}
}

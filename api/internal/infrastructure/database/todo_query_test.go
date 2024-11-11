package database_test

import (
	"api/internal/shared/config"
	"api/internal/shared/testutil"
	"api/internal/usecase/interfaces"
	"testing"

	"api/internal/domain/model"
	"api/internal/infrastructure/database"
	d "api/internal/shared/database"
	"reflect"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestTodoDatabaseQuery(t *testing.T) {
	testutil.SetEnv(t)
	id := int64(1)
	title := "test"
	description := "test"
	completed := false
	builder, err := model.TodoBuilder(id, title, description, completed)
	if err != nil {
		t.Fatal(err)
	}
	todo := builder

	config, err := config.NewConfigFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	pool, close, err := d.Pool(config)

	// 準備・終了処理
	pool.DB.Exec("INSERT INTO todos (id, title, description, completed) VALUES (?, ?, ?, ?)", id, title, description, completed)
	defer func() {
		_, err := pool.DB.Exec("DELETE FROM todos WHERE id = ?", id)
		if err != nil {
			t.Log(err)
		}
		close()
	}()

	// テスト
	query := database.NewTodoDatabaseQuery(pool)
	res, err := query.Get()
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(res, todo) {
		t.Errorf("expected %v, but got %v", todo, res)
	}
	res_, err := query.Detail(interfaces.TodoDetailRequest{Id: id})
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(res_, todo) {
		t.Errorf("expected %v, but got %v", todo, res_)
	}
}

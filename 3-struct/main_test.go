package main

import (
	"app-demo/api"
	"app-demo/bins"
	"app-demo/config"
	"app-demo/file"
	"app-demo/storage"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const testBaseURL = "https://api.jsonbin.io/v3/b"

func setupTestApp(t *testing.T) (*App, string) {
	t.Helper()

	cfg := config.Load()
	if cfg.Key == "" {
		t.Skip("пропуск теста: переменная KEY1 не задана в key.env")
	}

	tempDir := t.TempDir()
	dbPath := filepath.Join(tempDir, "bins_test.json")

	app := &App{
		api:   api.NewJsonDb(testBaseURL, cfg.Key),
		store: storage.New(file.NewJsonDb(dbPath)),
	}

	return app, tempDir
}

func createTempDataFile(t *testing.T, dir, filename, content string) string {
	t.Helper()

	path := filepath.Join(dir, filename)
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("не удалось создать временный файл данных: %v", err)
	}
	return path
}

func getFirstBin(t *testing.T, store *storage.Storage) bins.Bin {
	t.Helper()

	items, err := store.LoadBins()
	if err != nil {
		t.Fatalf("ошибка чтения хранилища: %v", err)
	}
	for _, item := range items {
		return item
	}

	t.Fatal("ожидался хотя бы 1 bin в хранилище, но список пуст")
	return bins.Bin{}
}

// 1.Тест на создание bin
func TestCreateBin(t *testing.T) {
	app, tempDir := setupTestApp(t)
	dataFile := createTempDataFile(t, tempDir, "create.json", `{"action": "create"}`)

	if err := app.create(dataFile, "test-create"); err != nil {
		t.Fatalf("ошибка при создании bin: %v", err)
	}

	created := getFirstBin(t, app.store)

	t.Cleanup(func() {
		_ = app.api.Delete(created.ID)
	})

	if created.ID == "" {
		t.Error("ID созданного bin не должен быть пустым")
	}
	if created.Name != "test-create" {
		t.Errorf("ожидалось имя 'test-create', получено '%s'", created.Name)
	}
}

// 2.Тест на получение созданного bin
func TestGetBin(t *testing.T) {
	app, tempDir := setupTestApp(t)
	dataFile := createTempDataFile(t, tempDir, "get.json", `{"user": "alice"}`)

	if err := app.create(dataFile, "test-get"); err != nil {
		t.Fatalf("ошибка подготовки (create): %v", err)
	}
	created := getFirstBin(t, app.store)

	t.Cleanup(func() {
		_ = app.api.Delete(created.ID)
	})

	if err := app.get(created.ID); err != nil {
		t.Fatalf("ошибка при получении bin: %v", err)
	}

	body, err := app.api.Get(created.ID)
	if err != nil {
		t.Fatalf("ошибка запроса к API: %v", err)
	}
	if !strings.Contains(string(body), "alice") {
		t.Errorf("ожидалось содержимое с 'alice', получено: %s", string(body))
	}
}

// 3.Тест на обновление bin
func TestUpdateBin(t *testing.T) {
	app, tempDir := setupTestApp(t)
	oldFile := createTempDataFile(t, tempDir, "old.json", `{"status": "draft"}`)
	newFile := createTempDataFile(t, tempDir, "new.json", `{"status": "published"}`)

	if err := app.create(oldFile, "test-update"); err != nil {
		t.Fatalf("ошибка подготовки (create): %v", err)
	}
	created := getFirstBin(t, app.store)

	t.Cleanup(func() {
		_ = app.api.Delete(created.ID)
	})

	if err := app.update(created.ID, newFile); err != nil {
		t.Fatalf("ошибка при обновлении bin: %v", err)
	}

	body, err := app.api.Get(created.ID)
	if err != nil {
		t.Fatalf("ошибка получения обновленного bin: %v", err)
	}
	if !strings.Contains(string(body), "published") {
		t.Errorf("данные не обновились, получено: %s", string(body))
	}
}

// 4.Тест на удаление bin
func TestDeleteBin(t *testing.T) {
	app, tempDir := setupTestApp(t)
	dataFile := createTempDataFile(t, tempDir, "delete.json", `{"temp": true}`)

	if err := app.create(dataFile, "test-delete"); err != nil {
		t.Fatalf("ошибка подготовки (create): %v", err)
	}
	created := getFirstBin(t, app.store)

	t.Cleanup(func() {
		_ = app.api.Delete(created.ID)
	})

	if err := app.delete(created.ID); err != nil {
		t.Fatalf("ошибка при удалении bin: %v", err)
	}

	itemsAfter, err := app.store.LoadBins()
	if err != nil {
		t.Fatalf("ошибка чтения хранилища: %v", err)
	}
	if len(itemsAfter) != 0 {
		t.Errorf("ожидалось 0 записей после удаления, осталось: %d", len(itemsAfter))
	}

	if _, err := app.api.Get(created.ID); err == nil {
		t.Error("ожидалась ошибка при запросе удаленного bin, но API вернул данные")
	}
}

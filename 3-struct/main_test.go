package main

import (
	"app-demo/api"
	"app-demo/config"
	"app-demo/file"
	"app-demo/storage"
	"os"
	"strings"
	"testing"
)

const testBaseURL = "https://api.jsonbin.io/v3/b"

func setupTestApp(t *testing.T) (*App, func()) {
	t.Helper()

	cfg := config.Load()
	if cfg.Key == "" {
		t.Fatal("переменная KEY1 не задана в key.env")
	}

	tmpStorage, err := os.CreateTemp("", "test_bins_*.json")
	if err != nil {
		t.Fatalf("не удалось создать временный файл хранилища: %v", err)
	}
	tmpStorage.Close()

	app := &App{
		api:   api.NewJsonDb(testBaseURL, cfg.Key),
		store: storage.New(file.NewJsonDb(tmpStorage.Name())),
	}

	cleanup := func() {
		os.Remove(tmpStorage.Name())
	}

	return app, cleanup
}

func createTempDataFile(t *testing.T, content string) (string, func()) {
	t.Helper()

	tmpFile, err := os.CreateTemp("", "test_data_*.json")
	if err != nil {
		t.Fatalf("не удалось создать временный файл данных: %v", err)
	}

	if _, err := tmpFile.WriteString(content); err != nil {
		t.Fatalf("ошибка записи во временный файл: %v", err)
	}
	tmpFile.Close()

	cleanup := func() {
		os.Remove(tmpFile.Name())
	}

	return tmpFile.Name(), cleanup
}

// Тест на создание bin
func TestCreateBin(t *testing.T) {
	app, cleanupApp := setupTestApp(t)
	defer cleanupApp()

	dataFile, cleanupFile := createTempDataFile(t, `{"action": "create"}`)
	defer cleanupFile()

	err := app.create(dataFile, "test-create")
	if err != nil {
		t.Fatalf("ошибка при создании bin: %v", err)
	}

	items, err := app.store.LoadBins()
	if err != nil {
		t.Fatalf("ошибка чтения хранилища: %v", err)
	}
	if len(items) != 1 {
		t.Fatalf("ожидался 1 bin в хранилище, получено: %d", len(items))
	}

	createdID := items[0].ID

	defer func() {
		_ = app.api.Delete(createdID)
	}()

	if createdID == "" {
		t.Error("ID созданного bin не должен быть пустым")
	}
	if items[0].Name != "test-create" {
		t.Errorf("ожидалось имя 'test-create', получено '%s'", items[0].Name)
	}
}

// Тест на получение созданного bin
func TestGetBin(t *testing.T) {
	app, cleanupApp := setupTestApp(t)
	defer cleanupApp()

	dataFile, cleanupFile := createTempDataFile(t, `{"user": "alice"}`)
	defer cleanupFile()

	if err := app.create(dataFile, "test-get"); err != nil {
		t.Fatalf("ошибка подготовки (create): %v", err)
	}
	items, _ := app.store.LoadBins()
	binID := items[0].ID

	defer func() {
		_ = app.api.Delete(binID)
	}()

	if err := app.get(binID); err != nil {
		t.Fatalf("ошибка при получении bin: %v", err)
	}

	body, err := app.api.Get(binID)
	if err != nil {
		t.Fatalf("ошибка запроса к API: %v", err)
	}
	if !strings.Contains(string(body), "alice") {
		t.Errorf("ожидалось содержимое с 'alice', получено: %s", string(body))
	}
}

// Тест на обновление bin
func TestUpdateBin(t *testing.T) {
	app, cleanupApp := setupTestApp(t)
	defer cleanupApp()

	oldFile, cleanupOld := createTempDataFile(t, `{"status": "draft"}`)
	defer cleanupOld()

	newFile, cleanupNew := createTempDataFile(t, `{"status": "published"}`)
	defer cleanupNew()

	if err := app.create(oldFile, "test-update"); err != nil {
		t.Fatalf("ошибка подготовки (create): %v", err)
	}
	items, _ := app.store.LoadBins()
	binID := items[0].ID

	defer func() {
		_ = app.api.Delete(binID)
	}()

	if err := app.update(binID, newFile); err != nil {
		t.Fatalf("ошибка при обновлении bin: %v", err)
	}

	body, err := app.api.Get(binID)
	if err != nil {
		t.Fatalf("ошибка получения обновленного bin: %v", err)
	}
	if !strings.Contains(string(body), "published") {
		t.Errorf("данные не обновились, получено: %s", string(body))
	}
}

// Тест на удаление bin
func TestDeleteBin(t *testing.T) {
	app, cleanupApp := setupTestApp(t)
	defer cleanupApp()

	dataFile, cleanupFile := createTempDataFile(t, `{"temp": true}`)
	defer cleanupFile()

	if err := app.create(dataFile, "test-delete"); err != nil {
		t.Fatalf("ошибка подготовки (create): %v", err)
	}
	items, _ := app.store.LoadBins()
	binID := items[0].ID

	defer func() {
		_ = app.api.Delete(binID)
	}()

	if err := app.delete(binID); err != nil {
		t.Fatalf("ошибка при удалении bin: %v", err)
	}

	itemsAfter, err := app.store.LoadBins()
	if err != nil {
		t.Fatalf("ошибка чтения хранилища: %v", err)
	}
	if len(itemsAfter) != 0 {
		t.Errorf("ожидалось 0 записей после удаления, осталось: %d", len(itemsAfter))
	}

	if _, err := app.api.Get(binID); err == nil {
		t.Error("ожидалась ошибка при запросе удаленного bin, но API вернул данные")
	}
}

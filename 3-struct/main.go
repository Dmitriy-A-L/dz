package main

import (
	"app-demo/api"
	"app-demo/bins"
	"app-demo/config"
	"app-demo/file"
	"app-demo/storage"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
)

type App struct {
	api   *api.JsonDb
	store *storage.Storage
}

func main() {
	createFlag := flag.Bool("create", false, "Создать новый bin")
	updateFlag := flag.Bool("update", false, "Обновить bin")
	deleteFlag := flag.Bool("delete", false, "Удалить bin")
	getFlag := flag.Bool("get", false, "Получить bin")
	listFlag := flag.Bool("list", false, "Список bins")

	fileFlag := flag.String("file", "", "Путь к JSON-файлу")
	nameFlag := flag.String("name", "", "Имя bin")
	idFlag := flag.String("id", "", "ID bin")
	flag.Parse()

	cfg := config.Load()
	if cfg.Key == "" {
		log.Fatal("Ошибка: переменная KEY1 не задана в key.env")
	}

	app := &App{
		api:   api.NewJsonDb("https://api.jsonbin.io/v3/b", cfg.Key),
		store: storage.New(file.NewJsonDb("bins.json")),
	}

	var err error
	switch {
	case *createFlag:
		err = app.create(*fileFlag, *nameFlag)
	case *updateFlag:
		err = app.update(*idFlag, *fileFlag)
	case *deleteFlag:
		err = app.delete(*idFlag)
	case *getFlag:
		err = app.get(*idFlag)
	case *listFlag:
		err = app.list()
	default:
		fmt.Println("Укажите команду: --create, --get, --update, --delete или --list")
	}

	if err != nil {
		log.Fatal(err)
	}
}

func (a *App) create(filePath, name string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("ошибка чтения файла (--file): %w", err)
	}
	res, err := a.api.Create(data, name)
	if err != nil {
		return err
	}
	fmt.Printf("Bin создан! ID: %s\n", res.Metadata.ID)
	return a.store.AddBin(bins.Bin{
		ID: res.Metadata.ID, Name: name, Private: res.Metadata.Private, CreatedAt: res.Metadata.CreatedAt,
	})
}

func (a *App) update(id, filePath string) error {
	if id == "" {
		return errors.New("укажите --id")
	}
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("ошибка чтения файла (--file): %w", err)
	}
	resp, err := a.api.Update(id, data)
	if err == nil {
		fmt.Printf("Bin обновлён:\n%s\n", resp)
	}
	return err
}

func (a *App) delete(id string) error {
	if id == "" {
		return errors.New("укажите --id")
	}
	if err := a.api.Delete(id); err != nil {
		return err
	}
	fmt.Printf("Bin %s удалён из облака\n", id)

	if err := a.store.DeleteBin(id); err != nil {
		if errors.Is(err, storage.ErrBinNotFound) {
			fmt.Println("В локальном списке bins.json такой записи не было")
			return nil
		}
		return err
	}

	return nil
}

func (a *App) get(id string) error {
	if id == "" {
		return errors.New("укажите --id")
	}
	data, err := a.api.Get(id)
	if err == nil {
		fmt.Println(string(data))
	}
	return err
}

func (a *App) list() error {
	items, err := a.store.LoadBins()
	if err != nil {
		return err
	}
	for _, b := range items {
		fmt.Printf("ID: %s | Name: %s | CreatedAt: %s\n", b.ID, b.Name, b.CreatedAt.Format("2006-01-02 15:04:05"))
	}
	return nil
}

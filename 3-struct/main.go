package main

import (
	"app-demo/bins"
	"app-demo/file"
	"app-demo/storage"
	"fmt"
	"log"
)

func main() {
	db := file.NewJsonDb("bins.json")
	store := storage.New(db)
	items, err := store.LoadBins()
	if err != nil {
		log.Fatalf("Загрузка: %v", err)
	}
	fmt.Printf("Загружено %d bins\n", len(items))

	newBin := bins.NewBin("1", "hello", false)
	items = append(items, newBin)
	if err := store.SaveBins(items); err != nil {
		log.Fatalf("Сохранение: %v", err)
	}
	fmt.Println("Сохранено: ok")
	fmt.Printf("Создан бин: ID=%s, Name=%s, CreatedAt=%v\n", newBin.ID, newBin.Name, newBin.CreatedAt)
}

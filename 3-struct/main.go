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

	items = append(items, bins.Bin{ID: "1", Payload: []byte("hello")})
	if err := store.SaveBins(items); err != nil {
		log.Fatalf("Сохранение: %v", err)
	}
	fmt.Println("Сохранено: ok")
}

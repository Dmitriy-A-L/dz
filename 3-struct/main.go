package main

import (
	"app-demo/bins"
	"app-demo/file"
	"app-demo/storage"
	"fmt"
	"log"
)

func main() {
	binsList := bins.NewBinList()
	binsList = append(binsList, bins.NewBin("1", "my-json", false))
	binsList = append(binsList, bins.NewBin("2", "notes", true))

	testFilePath := "test.json"
	fmt.Printf("Файл %s является JSON? -> %v\n", testFilePath, file.IsJSON(testFilePath))

	dbPath := "local_db.json"
	err := storage.SaveBins(dbPath, binsList)
	if err != nil {
		log.Fatal("Ошибка сохранения:", err)
	}
	fmt.Println(" Данные успешно сохранены в", dbPath)
	loadedBins, err := storage.LoadBins(dbPath)
	if err != nil {
		log.Fatal("Ошибка загрузки:", err)
	}

	fmt.Println("\n Данные из локальной базы:")
	for _, b := range loadedBins {
		fmt.Printf("ID: %s | Name: %s | Private: %v\n", b.ID, b.Name, b.Private)
	}
}

package main

import (
	"fmt"
	"time"
)

type Bin struct {
	ID        string
	Private   bool
	CreatedAt time.Time
	Name      string
}

type BinList []Bin

func NewBin(id, name string, private bool) Bin {
	return Bin{
		ID:        id,
		Private:   private,
		CreatedAt: time.Now(),
		Name:      name,
	}
}

func NewBinList() BinList {
	return make(BinList, 0)
}

func main() {
	bins := NewBinList()

	bin1 := NewBin("1", "my-json", false)
	bin2 := NewBin("2", "notes", true)

	bins = append(bins, bin1, bin2)

	fmt.Println("Список бинов:")
	for _, b := range bins {

		fmt.Printf("ID: %s | Name: %s | Private: %v | Создан: %v\n",
			b.ID, b.Name, b.Private, b.CreatedAt)
	}
}

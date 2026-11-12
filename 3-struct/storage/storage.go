package storage

import (
	"app-demo/bins"
	"encoding/json"
	"errors"
	"os"
)

type Storage struct {
	db Db
}

func New(db Db) *Storage {
	return &Storage{db: db}
}

func (s *Storage) SaveBins(items []bins.Bin) error {
	data, err := json.Marshal(items)
	if err != nil {
		return err
	}
	return s.db.Write(data)
}

func (s *Storage) LoadBins() ([]bins.Bin, error) {
	data, err := s.db.Read()
	if err != nil {
		// Первый запуск: файла ещё нет — это не ошибка.
		if errors.Is(err, os.ErrNotExist) {
			return []bins.Bin{}, nil
		}
		return nil, err
	}
	if len(data) == 0 {
		return []bins.Bin{}, nil
	}

	var items []bins.Bin
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, err
	}
	return items, nil
}

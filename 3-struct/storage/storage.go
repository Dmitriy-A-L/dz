package storage

import (
	"app-demo/bins"
	"encoding/json"
	"errors"
	"os"
)

var ErrBinNotFound = errors.New("бин не найден в локальном списке")

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

func (s *Storage) AddBin(bin bins.Bin) error {
	items, err := s.LoadBins()
	if err != nil {
		return err
	}
	items = append(items, bin)
	return s.SaveBins(items)
}

func (s *Storage) DeleteBin(id string) error {
	items, err := s.LoadBins()
	if err != nil {
		return err
	}

	updated := make([]bins.Bin, 0, len(items))
	found := false
	for _, item := range items {
		if item.ID == id {
			found = true
			continue
		}
		updated = append(updated, item)
	}

	if !found {
		return ErrBinNotFound
	}

	return s.SaveBins(updated)
}

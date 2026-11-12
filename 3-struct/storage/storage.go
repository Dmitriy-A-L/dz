package storage

import (
	"app-demo/bins"
	"encoding/json"
	"os"
)

func SaveBins(path string, list bins.BinList) error {
	data, err := json.MarshalIndent(list, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

func LoadBins(path string) (bins.BinList, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var list bins.BinList
	err = json.Unmarshal(data, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}

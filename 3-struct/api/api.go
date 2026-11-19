package api

import "errors"

type JsonDb struct {
	endpoint string
	token    string
}

func NewJsonDb(endpoint, token string) *JsonDb {
	return &JsonDb{endpoint: endpoint, token: token}
}

func (db *JsonDb) Read() ([]byte, error) {
	return nil, errors.New("api.JsonDb. Read: ещё не сделано")
}

func (db *JsonDb) Write(content []byte) error {
	return errors.New("api.JsonDb.Write: ещё не сделано")
}

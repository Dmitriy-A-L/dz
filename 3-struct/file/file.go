package file

import "os"

type JsonDb struct {
	path string
}

func NewJsonDb(path string) *JsonDb {
	return &JsonDb{path: path}
}

func (db *JsonDb) Read() ([]byte, error) {
	return os.ReadFile(db.path)
}

func (db *JsonDb) Write(content []byte) error {
	return os.WriteFile(db.path, content, 0o644)
}

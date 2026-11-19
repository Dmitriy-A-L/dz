package storage

type Db interface {
	Read() ([]byte, error)
	Write(content []byte) error
}

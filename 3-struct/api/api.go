package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type JsonDb struct {
	endpoint string
	token    string
}

type CreateResponse struct {
	Metadata struct {
		ID        string    `json:"id"`
		Private   bool      `json:"private"`
		CreatedAt time.Time `json:"createdAt"`
	} `json:"metadata"`
}

func NewJsonDb(endpoint, token string) *JsonDb {
	return &JsonDb{
		endpoint: endpoint,
		token:    token,
	}
}

func (db *JsonDb) doRequest(method, url string, body []byte, extraHeaders map[string]string) ([]byte, error) {
	var bodyReader io.Reader
	if body != nil {
		bodyReader = bytes.NewReader(body)
	}

	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-Master-Key", db.token)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	for k, v := range extraHeaders {
		req.Header.Set(k, v)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("ошибка API (%d): %s", resp.StatusCode, string(respBody))
	}

	return respBody, nil
}

func (db *JsonDb) Create(data []byte, name string) (*CreateResponse, error) {
	headers := map[string]string{}
	if name != "" {
		headers["X-Bin-Name"] = name
	}

	respBody, err := db.doRequest(http.MethodPost, db.endpoint, data, headers)
	if err != nil {
		return nil, err
	}

	var result CreateResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("ошибка разбора ответа: %w", err)
	}

	return &result, nil
}

func (db *JsonDb) Get(id string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s", db.endpoint, id)
	return db.doRequest(http.MethodGet, url, nil, nil)
}

func (db *JsonDb) Update(id string, data []byte) ([]byte, error) {
	url := fmt.Sprintf("%s/%s", db.endpoint, id)
	return db.doRequest(http.MethodPut, url, data, nil)
}

func (db *JsonDb) Delete(id string) error {
	url := fmt.Sprintf("%s/%s", db.endpoint, id)
	_, err := db.doRequest(http.MethodDelete, url, nil, nil)
	return err
}

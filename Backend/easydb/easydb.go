package easydb

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const BASE_URL = "https://app.easydb.io"

type DB struct {
	database string
	token    string
}

func Connect(database string, token string) *DB {
	return &DB{database, token}
}

func (db *DB) doDatabaseRequest(request *http.Request) (*http.Response, error) {
	client := http.Client{}
	request.Header.Set("token", db.token)
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("HTTP error code %d", response.StatusCode))
	}

	return response, nil
}

func (db *DB) Get(key string) ([]byte, error) {
	request, err := http.NewRequest("GET", fmt.Sprintf("%s/database/%s/%s", BASE_URL, db.database, key), nil)
	if err != nil {
		return nil, err
	}
	response, err := db.doDatabaseRequest(request)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(response.Body)
}

func (db *DB) Put(key string, value interface{}) error {
	payload, err := json.Marshal(map[string]interface{}{
		"value": value,
	})
	if err != nil {
		return err
	}
	request, err := http.NewRequest("POST",
		fmt.Sprintf("%s/database/%s/%s", BASE_URL, db.database, key),
		bytes.NewReader(payload),
	)
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		return err
	}
	_, err = db.doDatabaseRequest(request)
	return err
}

func (db *DB) List() (map[string]interface{}, error) {
	request, err := http.NewRequest("GET", fmt.Sprintf("%s/database/%s", BASE_URL, db.database), nil)
	if err != nil {
		return nil, err
	}
	response, err := db.doDatabaseRequest(request)
	if err != nil {
		return nil, err
	}
	result := make(map[string]interface{})

	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (db *DB) Delete(key string) error {
	request, err := http.NewRequest("DELETE", fmt.Sprintf("%s/database/%s/%s", BASE_URL, db.database, key), nil)
	if err != nil {
		return err
	}
	_, err = db.doDatabaseRequest(request)
	return err
}

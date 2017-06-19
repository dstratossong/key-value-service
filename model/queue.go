package model

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func Enqueue(service *Service, id uint64) {
	url := service.Url

	jsonValue, _ := json.Marshal(struct{ Id uint64 }{id})

	resp, err := http.Post(url+"/run", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		panic(err)
	}

	resp.Body.Close()
}

func Dequeue() {

}

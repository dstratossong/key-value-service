package model

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func Enqueue(service string, id uint64) {
	services, found := Services[service]
	if !found {
		panic("service not found")
	}

	url := services[0].Url

	jsonValue, _ := json.Marshal(struct{ Id uint64 }{id})

	resp, err := http.Post(url+"/run", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		panic(err)
	}

	resp.Body.Close()
}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type Guestbook struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func (g *Guestbook) save() error {
	fmt.Println(g.Name)
	return nil
}

func parseJsonGuestbook(jsonString string) (*Guestbook, error) {
	data := Guestbook{}
	err := json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Recieved request...")
	body := r.FormValue("body")
	guestbook, err := parseJsonGuestbook(body)
	if err != nil {
		fmt.Println(err.Error())
	}
	guestbook.save()
}

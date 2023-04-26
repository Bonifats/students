package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
	"students/pkg/storage"
)

type Controller struct {
	Storage storage.Storage
}

func (c *Controller) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var u User
	err := decoder.Decode(&u)
	if err != nil {
		c.setResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	_, err = c.Storage.Add(c.toUserDto(u))
	if err != nil {
		c.setResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	c.setResponse(w, http.StatusCreated, "User was created")
}

func (c *Controller) Attach(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var a Attach
	err := decoder.Decode(&a)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if a.SourceId == 0 {
		c.setResponse(w, http.StatusBadRequest, "Source not found")
		return
	}

	if a.TargetId == 0 {
		c.setResponse(w, http.StatusBadRequest, "Target not found")
		return
	}

	_, err = c.Storage.Attach(a.SourceId, a.TargetId)
	if err != nil {
		c.setResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	c.setResponse(w, http.StatusOK, "Users were attached")
}

func (c *Controller) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var d Delete
	err := decoder.Decode(&d)
	if err != nil {
		c.setResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	_, err = c.Storage.Delete(d.Id)
	if err != nil {
		c.setResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	c.setResponse(w, http.StatusOK, "User was deleted")
}

func (c *Controller) GetFriends(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	id, err := c.getParam(r, "id")
	if err != nil {
		c.setResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	friends, err := c.Storage.GetFriends(id)
	if err != nil {
		c.setResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	var resp ResponseList
	for _, friend := range friends {
		resp.Items = append(resp.Items, c.toResponseItem(friend))
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		c.setResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	c.setResponse(w, http.StatusOK, "")
}

func (c *Controller) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	id, err := c.getParam(r, "id")
	if err != nil {
		c.setResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	decoder := json.NewDecoder(r.Body)
	var u User
	err = decoder.Decode(&u)
	if err != nil {
		c.setResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	dto := c.toUserDto(u)

	_, err = c.Storage.Put(id, dto)
	if err != nil {
		c.setResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	c.setResponse(w, http.StatusOK, "User was updated")
}

func (c *Controller) getParam(r *http.Request, key string) (int, error) {
	param := chi.URLParam(r, key)
	if param == "" {
		return 0, errors.New(fmt.Sprint("User not found"))
	}

	id, err := strconv.Atoi(param)
	if err != nil || id == 0 {
		return 0, errors.New(fmt.Sprint("User not found"))
	}

	fmt.Println(id)

	return id, nil
}

func (c *Controller) setResponse(w http.ResponseWriter, status int, text string) {
	w.WriteHeader(status)
	if text != "" {
		_, err := w.Write([]byte(text))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
}

package controllers

import (
	"encoding/json"
	"net/http"
	"storage-app/models/request"
	"storage-app/server"
	"storage-app/services"
)

type StorageController struct {
	storageService services.StorageService
}

func NewStorageController(svc services.StorageService) StorageController {
	return StorageController{
		storageService: svc,
	}
}
func (s *StorageController) Get(res server.Response, req server.Request) *server.Response {
	value, err := s.storageService.Get(req.Params[0])
	if err != nil {
		return res.JSON(err, http.StatusNotFound)
	} else {
		return res.JSON(value, http.StatusOK)
	}
}

func (s *StorageController) Post(res server.Response, req server.Request) *server.Response {
	var t request.StorageReq
	err := json.Unmarshal(req.Body, &t)
	if err != nil {
		return res.JSON("Bad Request", http.StatusBadRequest)
	}
	s.storageService.Save(t)
	return res.JSON("", http.StatusOK)
}

func (s *StorageController) Flush(res server.Response, req server.Request) *server.Response {
	s.storageService.RemoveAll()
	return res.JSON("", http.StatusOK)
}

func (s *StorageController) Remove(res server.Response, req server.Request) *server.Response {
	s.storageService.Remove(req.Params[0])
	return res.JSON("", http.StatusOK)
}

package main

import (
	"encoding/json"
	"fmt"
	"storage-app/controllers"
	"storage-app/data/cache"
	"storage-app/models/logging"
	"storage-app/server"
	"storage-app/services"
)

func main() {

	srv := server.NewServer()
	cache := cache.NewCacheManager()
	fileSvc := services.NewFileSaverService(cache, 5000, "/tmp/")
	svc := services.NewStorageService(cache)
	storageController := controllers.NewStorageController(svc)
	srv.Middleware(func(res server.Response, req server.Request) *server.Response {
		log := logging.HttpLog{
			Headers:      req.Header,
			IP:           req.RemoteIP,
			Method:       req.Method,
			Status:       res.Status,
			URI:          req.URL,
			RequestBody:  string(req.Body),
			ResponseBody: string(res.Body),
		}
		logMessage, _ := json.Marshal(log)
		fmt.Println(string(logMessage))
		return &res
	})
	srv.GET("/records/:key", storageController.Get)
	srv.POST("/records", storageController.Post)
	srv.DELETE("/records", storageController.Flush)
	srv.DELETE("/records/:key", storageController.Remove)
	fileSvc.FillRecords()
	go fileSvc.Save()
	srv.Run(":3000")

}

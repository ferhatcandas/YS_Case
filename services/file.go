package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"storage-app/data/cache"
	"time"
)

type FileStorageService struct {
	cache      cache.CacheMangager
	intervalMS int
	folderPath string
}

func NewFileSaverService(cache cache.CacheMangager, intervalMS int, folderPath string) FileStorageService {
	fs := FileStorageService{
		cache:      cache,
		intervalMS: intervalMS,
		folderPath: folderPath,
	}
	fs.createDirectoryIfNotExists()
	return fs
}

func (f *FileStorageService) createDirectoryIfNotExists() {
	if _, err := os.Stat(f.folderPath); os.IsNotExist(err) {
		err := os.Mkdir(f.folderPath, os.ModeDir)
		if err != nil {
			panic(err)
		}
	}
}

func (f *FileStorageService) Save() {
	list := f.cache.GetAll()
	byteArr, err := json.Marshal(list)
	if err != nil {
		panic(err)
	}
	file, err := os.OpenFile(f.folderPath+"output.json", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	if _, err = file.WriteString(string(byteArr)); err != nil {
		panic(err)
	}
	fmt.Println("file saved")
	time.Sleep(time.Duration(int(f.intervalMS)) * time.Millisecond)
}
func (f *FileStorageService) Worker() {
	for {
		f.Save()
	}
}

func (f *FileStorageService) FillRecords() map[string]string {
	bytArr, _ := ioutil.ReadFile(f.folderPath + "output.json")
	var cacheData map[string]string
	json.Unmarshal(bytArr, &cacheData)
	fmt.Println("file loaded")
	for key, value := range cacheData {
		f.cache.AddOrUpdate(key, value)
	}
	return cacheData
}

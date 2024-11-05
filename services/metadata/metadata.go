package metadataservice

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MechCarrot/audiometa/internal/interfaces"
	"github.com/MechCarrot/audiometa/storage"
)

type MetaDataService struct {
	Server  *http.Server
	Storage interfaces.Storage
}

func CreateMetadataService(port int, storage interfaces.Storage) *MetaDataService {
	mux := http.NewServeMux()
	MetaDataService := &MetaDataService{
		Server: &http.Server{
			Addr:    fmt.Sprintf(":%v", port),
			Handler: mux,
		},
		Storage: storage,
	}

	mux.HandleFunc("/upload", MetaDataService.UploadHandler)
	mux.HandleFunc("/get", MetaDataService.GetByIdHandler)
	mux.HandleFunc("/list", MetaDataService.ListHandler)

	return MetaDataService
}

func Run(port int) {
	flatfileStorage := storage.FlatFile{}
	service := CreateMetadataService(port, flatfileStorage)
	err := service.Server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

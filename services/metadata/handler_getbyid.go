package metadataservice

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func (mdc *MetaDataService) GetByIdHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if id == "" {
		log.Println("Id param in URL is missing")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Println("Requesting audio by", id)
	audio, err := mdc.Storage.GetById(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "no such") {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	audioJSON, err := audio.JSON()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, audioJSON)
}

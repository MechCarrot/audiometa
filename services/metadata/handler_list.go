package metadataservice

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (mdc *MetaDataService) ListHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("We are getting ListHandler request")
	audioList, err := mdc.Storage.List()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	audioListJSON, err := json.Marshal(audioList)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, audioListJSON, "", "    ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, string(audioListJSON))
}

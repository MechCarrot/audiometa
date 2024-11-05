package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/MechCarrot/audiometa/models"
	"github.com/google/uuid"
)

type FlatFile struct {
	Name string
}

func (f FlatFile) GetById(id string) (*models.Audio, error) {
	dirname, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	metadataFilePath := filepath.Join(dirname, "audiofile", id, "metadata.json")
	if _, err := os.Stat(metadataFilePath); errors.Is(err, os.ErrNotExist) {
		_ = os.Mkdir(metadataFilePath, os.ModePerm)
	}
	file, err := os.ReadFile(metadataFilePath)
	if err != nil {
		return nil, err
	}

	data := models.Audio{}
	err = json.Unmarshal([]byte(file), &data)
	return &data, err
}

func (f FlatFile) SaveMetadata(audio *models.Audio) error {
	dirname, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	metaFilePath := filepath.Join(dirname, "audiofile", audio.Id, "metadata.json")
	file, err := os.Create(metaFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := audio.JSON()
	if err != nil {
		fmt.Println("Err", err)
	}
	_, err = fmt.Fprint(file, data)
	if err != nil {
		return err
	}
	return file.Sync()
}

func (f FlatFile) Upload(bytes []byte, filename string) (string, string, error) {
	id := uuid.New()
	dirname, err := os.UserHomeDir()
	if err != nil {
		return id.String(), "", err
	}

	metadataDirPath := filepath.Join(dirname, "audiofile", id.String())
	if err := os.MkdirAll(metadataDirPath, os.ModePerm); err != nil {
		return id.String(), "", err
	}
	metadataFilePath := filepath.Join(dirname, "audiofile", id.String(), filename)
	err = os.WriteFile(filename, bytes, 0644)
	if err != nil {
		return id.String(), "", err
	}
	return id.String(), metadataFilePath, nil
}

func (f FlatFile) List() ([]*models.Audio, error) {
	dirname, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	metadataListPath := filepath.Join(dirname, "audiofile")
	if _, err := os.Stat(metadataListPath); errors.Is(err, os.ErrNotExist) {
		_ = os.Mkdir(metadataListPath, os.ModePerm)
	}

	files, err := os.ReadDir(metadataListPath)
	if err != nil {
		return nil, err
	}
	audioFiles := make([]*models.Audio, 0)

	for _, file := range files {
		if file.IsDir() {
			name, err := f.GetById(file.Name())
			if err != nil {
				return nil, err
			}
			audioFiles = append(audioFiles, name)
		}
	}
	return audioFiles, err
}

func (f FlatFile) Delete(id string) error {
	fmt.Println("Deleting", id)
	return nil
}

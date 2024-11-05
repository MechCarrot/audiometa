package interfaces

import "github.com/MechCarrot/audiometa/models"

type Storage interface {
	Upload(bytes []byte, filename string) (string, string, error)
	SaveMetadata(audio *models.Audio) error
	List() ([]*models.Audio, error)
	GetById(id string) (*models.Audio, error)
	Delete(id string) error
}

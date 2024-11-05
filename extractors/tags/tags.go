package tags

import (
	"fmt"
	"os"

	"github.com/MechCarrot/audiometa/models"
	"github.com/dhowden/tag"
)

func Extract(m *models.Audio) error {
	f, err := os.Open(m.Path)
	if err != nil {
		return err
	}

	defer f.Close()
	fmt.Println("file: ", m.Path)
	tagMetadata, err := tag.ReadFrom(f)
	if err != nil {
		return err
	}
	fmt.Println("tagMeta:", tagMetadata)
	m.Metadata.Tags = models.Tags{
		Title:       tagMetadata.Title(),
		Album:       tagMetadata.Album(),
		Artist:      tagMetadata.Artist(),
		AlbumArtist: tagMetadata.AlbumArtist(),
		Composer:    tagMetadata.Composer(),
		Genre:       tagMetadata.Genre(),
		Year:        tagMetadata.Year(),
		Lyrics:      tagMetadata.Lyrics(),
		Comment:     tagMetadata.Comment(),
	}

	return nil
}

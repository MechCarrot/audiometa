package models

type Tags struct {
	Title       string `json:"Title"`
	Album       string `json:"album"`
	Artist      string `json:"artist"`
	AlbumArtist string `json:"album_artist"`
	Composer    string `json:"composer"`
	Genre       string `json:"genre"`
	Year        int    `json:"year"`
	Lyrics      string `json:"lyrics"`
	Comment     string `json:"comment"`
}

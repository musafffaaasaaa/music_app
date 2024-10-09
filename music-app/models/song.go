package models

type Song struct {
	ID           int    `json:"id"`
	GroupName    string `json:"group"`
	SongTitle    string `json:"song"`
	ReleaseDate  string `json:"release_date,omitempty"`
	Lyrics       string `json:"lyrics,omitempty"`
	ExternalLink string `json:"link,omitempty"`
}

type SongDetail struct {
	ReleaseDate string `json:"releaseDate"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

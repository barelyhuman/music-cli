package song

// Author consists of the name and other info about the creator of the song
type Author struct {
	Name string `json:"name"`
}

// Duration consists of duration information of the audio content
type Duration struct {
	Seconds int `json:"seconds"`
}

// Song structure definition with json bindings
type Song struct {
	Title    string   `json:"title"`
	Author   Author   `json:"author"`
	Duration Duration `json:"duration"`
	VideoID  string   `json:"videoId"`
}

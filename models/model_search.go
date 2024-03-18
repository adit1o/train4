package models

type SearchModel struct {
	VideoID             string   `json:"video_id"`
	VideoCodeID         string   `json:"video_code_id"`
	VideoCode           string   `json:"video_code"`
	Title               string   `json:"title"`
	TitleJa             string   `json:"title_ja"`
	Description         any      `json:"description"`
	FanArt              string   `json:"fan_art"`
	Poster              string   `json:"poster"`
	Trailer             string   `json:"trailer"`
	Actors              []ActorModel `json:"actors"`
	Genre               []string `json:"genre"`
	ReleaseDate         string   `json:"release_date"`
	ReleaseDateFormated string   `json:"release_date_formated"`
	Duration            int      `json:"duration"`
	DurationFormated    string   `json:"duration_formated"`
	Studio              string   `json:"studio"`
	Label               string   `json:"label"`
	Director            string   `json:"director"`
	Series              any      `json:"series"`
	Screenshot          []string `json:"screenshot"`
	ViewCount           any      `json:"view_count"`
	Score               any      `json:"score"`
}
type ActorModel struct {
	CastID    int    `json:"cast_id"`
	CastNames string `json:"cast_names"`
	CastImage string `json:"cast_image"`
}
package entity

type Film struct {
	Title         string   `json:"title"`
	EpisodeID     int      `json:"episode_id"`
	OpeningCrawl  string   `json:"opening_crawl"`
	Director      string   `json:"director"`
	Producer      string   `json:"producer"`
	CharacterURLs []string `json:"characters"`
	PlanetURLs    []string `json:"planets"`
	StarshipURLs  []string `json:"starships"`
	VehicleURLs   []string `json:"vehicles"`
	SpeciesURLs   []string `json:"species"`
	Created       string   `json:"created"`
	Edited        string   `json:"edited"`
	URL           string   `json:"url"`
}

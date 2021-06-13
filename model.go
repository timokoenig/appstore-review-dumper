package main

type review struct {
	Title      string `json:"title"`
	Text       string `json:"text"`
	Date       int    `json:"date"`
	DateString string `json:"dateString"`
	User       string `json:"user"`
	Rating     int    `json:"rating"`
}

type reviewFile struct {
	URL     string    `json:"url"`
	Reviews []*review `json:"reviews"`
}

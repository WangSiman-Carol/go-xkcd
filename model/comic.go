package model

import (
	"encoding/json"
	"fmt"
)

/*
* ComicResonse struct receives json file from the xkcd endpoint
 */
type ComicResponse struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}

/*
* Comic struct outputs data from the app
 */
type Comic struct {
	Title       string `json:"title"`
	Number      int    `json:"number"`
	Date        string `json:"date"`
	Description string `json:"description"`
	Image       string `json:"imgage"`
}

// FormattedDate() formats individual date elements into a single string
func (cr ComicResponse) FormattedDate() string {
	return fmt.Sprintf("%s-%s-%s", cr.Day, cr.Month, cr.Year)
}

// Comic() transfers the ComicResponse struct into Comic struct
func (cr ComicResponse) Comic() Comic {
	return Comic{
		Title:       cr.Title,
		Number:      cr.Num,
		Date:        cr.FormattedDate(),
		Description: cr.Alt,
		Image:       cr.Img,
	}
}

// PrettyString() output a human-friendly string of the Comic
func (c Comic) PrettyString() string {
	p := fmt.Sprintf(
		"Title: %s\nComic No: %d\nDate: %s\nDescription: %s\nImage: %s\n",
		c.Title, c.Number, c.Date, c.Description, c.Image)
	return p
}

// JSON() transfer Comid struct into JSON file
func (c Comic) JSON() string {
	cJSON, err := json.Marshal(c)
	if err != nil {
		return ""
	}
	return string(cJSON)
}

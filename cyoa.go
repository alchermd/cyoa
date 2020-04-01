package cyoa

import (
	"encoding/json"
	"html/template"
	"io"
	"net/http"
)

type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"paragraphs"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"chapter"`
}

func JsonToStory(r io.Reader) (Story, error) {
	d := json.NewDecoder(r)
	var story Story
	if err := d.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}

func NewHandler(s Story, tpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if tpl == nil {
			tpl = template.Must(template.ParseFiles("views/story.html"))
		}

		chapter := r.URL.Path[1:]
		if chapter == "" {
			chapter = "intro"
		}

		if chapter, ok := s[chapter]; ok {
			err := tpl.Execute(w, chapter)
			if err != nil {
				http.Error(w, "Something went wrong.", http.StatusInternalServerError)
			}
			return
		}

		http.Error(w, "Story not found", http.StatusNotFound)
	}
}

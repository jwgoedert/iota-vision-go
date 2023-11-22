package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type Idea struct {
	ID              string    `json:"id"`
	Title           string    `json:"title"`
	Description     []byte    `json:"description"`
	CreatedAt       time.Time `json:"created_at"`
	SharedObjectIDs []string  `json:"shared_object_ids"`
	SharedTimeIds   []string  `json:"shared_time_ids"`
	SharedActionIds []string  `json:"shared_action_ids"`
}

func (i *Idea) save() error {
	filename := i.Title + ".txt"
	return os.WriteFile(filename, i.Description, 0600)
}

func loadIdea(title string) (*Idea, error) {
	filename := title + ".txt"
	description, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Idea{Title: title, Description: description}, err
}
func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	i, err := loadIdea(title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", i.Title, i.Description)
}

func main() {
	//	i := Idea{
	//		ID:          "1",
	//		Title:       "My Idea",
	//		Description: []byte("This is my idea"),
	//		CreatedAt:   time.Now(),
	//	}
	//
	// i.save()
	// i2, _ := loadIdea("My Idea")
	// println(string(i2.Description))
	http.HandleFunc("/view/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

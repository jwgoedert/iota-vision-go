package main

import (
	"fmt"
	"log"
	"net/http"

	// "os"
	// "time"
	"github.com/gorilla/websocket"
)

// type Idea struct {
// 	ID              string    `json:"id"`
// 	Title           string    `json:"title"`
// 	Description     []byte    `json:"description"`
// 	CreatedAt       time.Time `json:"created_at"`
// 	SharedObjectIDs []string  `json:"shared_object_ids"`
// 	SharedTimeIds   []string  `json:"shared_time_ids"`
// 	SharedActionIds []string  `json:"shared_action_ids"`
// }

// func (i *Idea) save() error {
// 	filename := i.Title + ".txt"
// 	return os.WriteFile(filename, i.Description, 0600)
// }

// func loadIdea(title string) (*Idea, error) {
// 	filename := title + ".txt"
// 	description, err := os.ReadFile(filename)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &Idea{Title: title, Description: description}, err
// }

//	func viewHandler(w http.ResponseWriter, r *http.Request) {
//		title := r.URL.Path[len("/view/"):]
//		i, err := loadIdea(title)
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//		fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", i.Title, i.Description)
// //	}
// func setupRoutes() {
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "Welcome to the Homepage!")
// 	})
// }
// func main() {
// 	setupRoutes()
// 	http.ListenAndServe(":8080", nil)
// i := Idea{
// 	ID:          "1",
// 	Title:       "My Second Idea",
// 	Description: []byte("This is my second idea ever"),
// 	CreatedAt:   time.Now(),
// }

// i.save()
// i2, _ := loadIdea("My Idea")
// println(string(i2.Description))
// http.HandleFunc("/view/", viewHandler)
// log.Fatal(http.ListenAndServe(":8080", nil))
// }
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	// We'll need to check the origin of our connection
	// this will allow us to make requests from our React
	// development server to here.
	// For now, we'll do no checking and just allow any connection
	CheckOrigin: func(r *http.Request) bool { return true },
}

// define a reader which will listen for
// new messages being sent to our WebSocket
// endpoint
func reader(conn *websocket.Conn) {
	for {
		// read in a message
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		// print out that message for clarity
		fmt.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}

	}
}

// define our WebSocket endpoint
func serveWs(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)

	// upgrade this connection to a WebSocket
	// connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	// listen indefinitely for new messages coming
	// through on our WebSocket connection
	reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Simple Server")
	})
	// mape our `/ws` endpoint to the `serveWs` function
	http.HandleFunc("/ws", serveWs)
}

func main() {
	fmt.Println("Chat App v0.01")
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}

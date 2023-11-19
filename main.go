//go:build ignore

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// We'll need to define an Upgrader
// this will require a Read and Write buffer size
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	// We'll need to check the origin of our connection
	// this will allow us to make requests from our React
	// development server to here.
	// For now, we'll do no checking and just allow any connection
	CheckOrigin: func(r *http.Request) bool { return true },
}

// define our WebSocket endpoint
func serveWs(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)
	fmt.Println("testing one two three")

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

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
func setupRoutes() {
	http.HandleFunc("/", handler)
	// mape our `/ws` endpoint to the `serveWs` function
	http.HandleFunc("/ws", serveWs)
}
func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}
func main() {
	fmt.Println("Starting server on port 8080...")
	http.HandleFunc("/ws", serveWs)
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// func main() {

// 	router := mux.NewRouter()

// 	// Define your endpoints and their corresponding handlers
// 	router.HandleFunc("/", RootHandler).Methods("GET")
// 	// router.HandleFunc("/contacts", CreateContact).Methods("POST")
// 	// Add other endpoints for updating, deleting, getting contacts, etc.

// 	log.Fatal(http.ListenAndServe(":8080", router))
// }

func RootHandler(w http.ResponseWriter, r *http.Request) {
	message := "Welcome iotas buds!"
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
	// fmt.Fprintf(w, "Hello World!")
}

// Define your HTTP handlers here (e.g., createContact, updateContact, etc.)
// These functions will handle CRUD operations for contacts
// func createContact(w http.ResponseWriter, r *http.Request) {
// 	// Implement contact creation logic here
// }

// Add other handler functions for CRUD operations

// Contact struct represents your contact model
// type Contact struct {
// 	ID        string
// 	FirstName string
// 	LastName  string
// 	// Add other fields as needed (Twitter, Avatar, Notes, etc.)
// }

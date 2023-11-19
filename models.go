package main

// Contact struct represents your contact model
type Contact struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	// Add other fields as needed (Twitter, Avatar, Notes, etc.)
}

package models

import "time"

type Camera struct {
    ID                  int         `json:"id"`
    Name                string      `json:"name"`
    Username            string      `json:"usr"`
    Password            string      `json:"pass"`
    Address             string      `json:"ip"`
    Created             time.Time   `json:"installed"`
    Updated             time.Time   `json:"inspected"`
    Rating              int         `json:"rating"`
    // As soon as we connect our database we will use
    // a JOIN query to get the location information.
    // Change the line to `Location []Location `json:"location"``
    // to display them.
    Location            []Location  `json:"-"`
}

type CameraModel struct {
    ID                  int         `json:"id"`
    Model               string      `json:"model"`
    Lense               string      `json:"lense"`
    Resolution          string      `json:"res"`
    Created             time.Time   `json:"install"`
    Updated             time.Time   `json:"inspect"`
}

type Location struct {
    ID                  int         `json:"id"`
    LocationName        string      `json:"location"`
    CameraID            string      `json:"cid"`
    LocationID          string      `json:"lid"`
    Model               string      `json:"model"`
    Lense               string      `json:"lense"`
    Resolution          string      `json:"res"`
    Created             time.Time   `json:"install"`
    Updated             time.Time   `json:"inspect"`
}

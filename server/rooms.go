package server

import (
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// participant describe single entity in hashmap
type Participant struct {
	Host bool
	Conn *websocket.Conn
}

// room is the main hashmap
type RoomMap struct {
	Mutex sync.RWMutex
	Map   map[string][]Participant
}

// Initialize room map struct
func (r *RoomMap) Init() {
	r.Map = make(map[string][]Participant)
}

// Get participants in room
func (r *RoomMap) Get(roomID string) []Participant {
	r.Mutex.RLock()
	defer r.Mutex.RUnlock()
	return r.Map[roomID]
}

// create new room
func (r *RoomMap) CreateRoom() string {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	rand.Seed(time.Now().Unix())
	var letters = []rune("abcdefghijklmnopqstuvwxyz1234567890")
	b := make([]rune, 6)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	roomID := string(b)
	r.Map[roomID] = []Participant{}

	return roomID
}

// add client to the room
func (r *RoomMap) InsertInRoom(roomID string, host bool, conn *websocket.Conn) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	p := Participant{host, conn}

	log.Println("Inserting into room with roomID: ", roomID)
	r.Map[roomID] = append(r.Map[roomID], p)
}

// remove delete room
func (r *RoomMap) DeleteRoom(roomID string) {
	r.Mutex.Lock()
	r.Mutex.Unlock()

	delete(r.Map, roomID)
}

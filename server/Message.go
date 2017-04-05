package server
/*
Type Holding data about a message
*/
type Message struct {
	IP	string    `json:"IP"`
	Port	string       `json:"port"`
	Message	string    `json:"message"`

}
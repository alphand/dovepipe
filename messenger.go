package main

import "net/http"

// Messenger - Base FB Messenger Bot Object
type Messenger struct {
	VerifyToken string
	AppSecret   string
	AccessToken string
	PageID      string
}

// Handler - Main HTTP handler for FB Bot integration
func (m *Messenger) Handler(
	rw http.ResponseWriter,
	req *http.Request) {

	rw.Write([]byte("Hello"))
}

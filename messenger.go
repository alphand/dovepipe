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
func (m *Messenger) Handler() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodGet {
			m.verifyToken(rw, req)
		} else {
			rw.WriteHeader(http.StatusUnauthorized)
			rw.Write([]byte("Hello"))
		}
	})
}

func (m *Messenger) verifyToken(
	rw http.ResponseWriter,
	req *http.Request,
) {

	query := req.URL.Query()
	verifyToken := query.Get("hub.verify_token")

	if verifyToken != m.VerifyToken {
		rw.WriteHeader(http.StatusUnauthorized)
	} else {
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(query.Get("hub.challenge")))
	}
	return
}

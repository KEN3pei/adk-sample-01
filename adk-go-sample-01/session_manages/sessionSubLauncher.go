package session_manages

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"google.golang.org/adk/cmd/launcher"
	"google.golang.org/adk/session"
)

type SessionLauncher struct {
	SessionService session.Service
}

type CreateSessionRequest struct {
	AppName string `json:"appName"`
	UserID  string `json:"userId"`
}

func (s *SessionLauncher) Keyword() string { return "session" }

func (s *SessionLauncher) SimpleDescription() string {
	return "Session management sublauncher"
}

func (s *SessionLauncher) CommandLineSyntax() string {
	return "session management endpoints"
}

func (s *SessionLauncher) Parse(args []string) ([]string, error) {
	// No specific command-line arguments to parse for this sublauncher
	return args, nil
}

func (s *SessionLauncher) SetupSubrouters(router *mux.Router, _ *launcher.Config) error {
	router.HandleFunc("/sessions", s.HandleCreateSession).Methods("POST")
	return nil
}

func (s *SessionLauncher) UserMessage(webURL string, printer func(v ...any)) {
	printer("Session management endpoints available at:", webURL+"/sessions")
}

func (s *SessionLauncher) HandleCreateSession(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var req CreateSessionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if req.UserID == "" {
		http.Error(w, "userId is required", http.StatusBadRequest)
		return
	}
	if req.AppName == "" {
		http.Error(w, "AppName is required", http.StatusBadRequest)
		return
	}

	resp, err := s.SessionService.Create(r.Context(), &session.CreateRequest{
		AppName: req.AppName,
		UserID:  req.UserID,
		State:   nil,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"sessionId": resp.Session.ID()})
}

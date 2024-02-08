package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Server struct {
	storage *Storage
	logger  *log.Logger
}

func NewServer() *Server {
	return &Server{
		logger:  log.New(os.Stdout, "logger: ", log.Lshortfile),
		storage: NewStorage(),
	}
}

type CreateEventRequest struct {
	UserId int `json:"user_id"`
	Id     int `json:"id"`
	Event  Event
}

func CreateEventFormData(r *http.Request) (CreateEventRequest, error) {
	event := CreateEventRequest{}

	userId := r.FormValue("user_id")
	if userId == "" {
		return event, fmt.Errorf("user id not found")
	}
	uI, err := strconv.Atoi(userId)
	if err != nil {
		return event, err
	}

	id := r.FormValue("id")
	if id == "" {
		return event, fmt.Errorf("id not found")
	}
	i, err := strconv.Atoi(id)
	if err != nil {
		return event, err
	}

	title := r.FormValue("title")
	if title == "" {
		return event, fmt.Errorf("title not found")
	}

	content := r.FormValue("content")
	if content == "" {
		return event, fmt.Errorf("content not found")
	}

	d := r.FormValue("date")
	date, err := time.Parse("2006-01-02", d)
	if err != nil {
		return event, fmt.Errorf("incorrect date: %s", d)
	}

	return CreateEventRequest{
		UserId: uI,
		Id:     i,
		Event: Event{
			Title:   title,
			Date:    date,
			Content: content,
		},
	}, nil
}

type DeleteEventRequest struct {
	UserId int `json:"user_id"`
	Id     int `json:"id"`
}

func (s *Server) HandleCreateEvent(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var cEvent CreateEventRequest
		cEvent, err := CreateEventFormData(r)
		if err != nil {
			s.logger.Printf("Could not read request body: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		s.storage.CreateNewEvent(cEvent.UserId, cEvent.Id, cEvent.Event)
		s.logger.Printf("Created event: %d", cEvent.Id)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (s *Server) HandleUpdateEvent(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var cEvent CreateEventRequest
		cEvent, err := CreateEventFormData(r)
		if err != nil {
			s.logger.Printf("Could not read request body: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		s.storage.UpdateEvent(cEvent.UserId, cEvent.Id, cEvent.Event)
		s.logger.Printf("Updated event: %d", cEvent.Id)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (s *Server) HandleDeleteEvent(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var dEvent DeleteEventRequest
		err := json.NewDecoder(r.Body).Decode(&dEvent)
		if err != nil {
			s.logger.Printf("Could not read request body: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		s.storage.DeleteEvent(dEvent.UserId, dEvent.Id)
		s.logger.Printf("Deleted event: %d", dEvent.Id)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (s *Server) HandleGetEventsForDay(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		userId := r.URL.Query().Get("user_id")
		id, err := strconv.Atoi(userId)
		if err != nil {
			s.logger.Printf("Wrong user id: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		events := s.storage.GetEventsForDay(id)
		msg, err := json.Marshal(events)
		if err != nil {
			s.logger.Printf("Could not marshal request: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		s.logger.Printf("Get events for day for user: %s", userId)
		w.Header().Add("Content-Type", "application/json")
		w.Write(msg)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (s *Server) HandleGetEventsForMonth(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		userId := r.URL.Query().Get("user_id")
		id, err := strconv.Atoi(userId)
		if err != nil {
			s.logger.Printf("Wrong user id: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		events := s.storage.GetEventsForMonth(id)
		msg, err := json.Marshal(events)
		if err != nil {
			s.logger.Printf("Could not marshal request: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		s.logger.Printf("Get events for month for user: %s", userId)
		w.Header().Add("Content-Type", "application/json")
		w.Write(msg)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (s *Server) HandleGetEventsForWeek(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		userId := r.URL.Query().Get("user_id")
		id, err := strconv.Atoi(userId)
		if err != nil {
			s.logger.Printf("Wrong user id: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		events := s.storage.GetEventsForWeek(id)
		msg, err := json.Marshal(events)
		if err != nil {
			s.logger.Printf("Could not marshal request: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		s.logger.Printf("Get events for week for user: %s", userId)
		w.Header().Add("Content-Type", "application/json")
		w.Write(msg)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

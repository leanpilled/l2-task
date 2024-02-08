package main

import (
	"sync"
	"time"
)

type Event struct {
	Date    time.Time `json:"date"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
}

type Storage struct {
	m map[int]map[int]Event
	*sync.RWMutex
}

func NewStorage() *Storage {
	return &Storage{
		m:       make(map[int]map[int]Event),
		RWMutex: new(sync.RWMutex),
	}
}

func (s *Storage) CreateNewEvent(userId int, id int, event Event) {
	s.RLock()
	defer s.RUnlock()
	if _, ok := s.m[userId]; !ok {
		s.m[userId] = make(map[int]Event)
	}
	s.m[userId][id] = Event{
		Date:    event.Date,
		Title:   event.Title,
		Content: event.Content,
	}
}

func (s *Storage) DeleteEvent(userId int, id int) {
	s.RLock()
	defer s.RUnlock()
	delete(s.m[userId], id)
}

func (s *Storage) UpdateEvent(userId int, id int, event Event) {
	s.RLock()
	defer s.RUnlock()
	e, ok := s.m[userId][id]
	if ok {
		e.Content = event.Content
		e.Title = event.Title
		e.Content = event.Content
	}
}

func (s *Storage) GetEventsForDay(userId int) []Event {
	currentTime := time.Now()
	newTime := currentTime.Add(24 * time.Hour)
	res := []Event{}
	for _, event := range s.m[userId] {
		if event.Date.After(currentTime) && event.Date.Before(newTime) {
			res = append(res, event)
		}
	}
	return res
}

func (s *Storage) GetEventsForMonth(userId int) []Event {
	currentTime := time.Now()
	newTime := currentTime.Add(24 * 30 * time.Hour)
	res := []Event{}
	for _, event := range s.m[userId] {
		if event.Date.After(currentTime) && event.Date.Before(newTime) {
			res = append(res, event)
		}
	}
	return res
}

func (s *Storage) GetEventsForWeek(userId int) []Event {
	currentTime := time.Now()
	newTime := currentTime.Add(24 * 7 * time.Hour)
	res := []Event{}
	for _, event := range s.m[userId] {
		if event.Date.After(currentTime) && event.Date.Before(newTime) {
			res = append(res, event)
		}
	}
	return res
}

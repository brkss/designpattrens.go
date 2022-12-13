package observer

import (
	"fmt"
	"sync"
	"time"
)

// Observe Fib function and how long it take to produce new value

type Event struct {
  Data int
}

type Observer interface {
  NotifyCallback(Event)
}

type Subject interface {
  AddListener(Observer)
  RemoveListener(Observer)
  Notify(Event)
}

type EventObserver struct {
  Id int
  Time time.Time
}

type EventSubject struct {
  Observers sync.Map
}

func (e *EventObserver) NotifyCallback(event Event){
  fmt.Printf("Observer : %d, \tRecieved: %d \tafter %v\n", e.Id, event.Data, time.Since(e.Time))
}

func (s *EventSubject) AddListener(obs Observer){
  s.Observers.Store(obs, struct{}{})
}

func (s *EventSubject) RemoveListener(obs Observer){
  s.Observers.Delete(obs)
}

func (s *EventSubject) Notify(event Event){
  s.Observers.Range(func(key interface{}, value interface{}) bool {
    if key == nil || value == nil {
      return false;
    }

    key.(Observer).NotifyCallback(event)
    return true;
  })
}

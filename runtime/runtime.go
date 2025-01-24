package runtime

import "sync"

type Event struct {
	ID      int
	Payload string
}

type EventHandler func(event Event) string

type Runtime struct {
	eventQueue chan Event
	handlers   map[string]EventHandler
	stopChan   chan struct{}
	reportChan chan string
	wg         sync.WaitGroup
}

func NewRuntime(bufferSize int) *Runtime {
	return &Runtime{
		eventQueue: make(chan Event, bufferSize),
		handlers:   make(map[string]EventHandler),
		stopChan:   make(chan struct{}),
		reportChan: make(chan string),
	}
}

func (r *Runtime) RegisterHandler(eventType string, handler EventHandler) {
	r.handlers[eventType] = handler
}

func (r *Runtime) Emit(event Event) {
	r.wg.Add(1)
	r.eventQueue <- event
}

func (r *Runtime) Start() {
	go func() {
		for {
			select {
			case event := <-r.eventQueue:
				handler, ok := r.handlers[event.Payload]
				if ok {
					r.reportChan <- handler(event)
				} else {
					r.reportChan <- "No handler registered for event type: " + event.Payload
				}
				r.wg.Done()
			case <-r.stopChan:
				close(r.reportChan)
				return
			}
		}
	}()
}

func (r *Runtime) Stop() {
	r.wg.Wait()
	close(r.stopChan)
}

func (r *Runtime) ListenResponses() <-chan string {
	return r.reportChan
}

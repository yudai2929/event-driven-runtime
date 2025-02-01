package runtime

type Event struct {
	ID  int
	Key string
}

type EventHandler func(event Event) string

type Runtime struct {
	eventQueue chan Event
	handlers   map[string]EventHandler
	stopChan   chan struct{}
	reportChan chan string
}

func NewRuntime(bufferSize int) *Runtime {
	return &Runtime{
		eventQueue: make(chan Event, bufferSize),
		handlers:   make(map[string]EventHandler),
		stopChan:   make(chan struct{}),
		reportChan: make(chan string),
	}
}

func (r *Runtime) RegisterHandler(key string, handler EventHandler) {
	r.handlers[key] = handler
}

func (r *Runtime) Emit(event Event) {
	r.eventQueue <- event
}

func (r *Runtime) Start() {
	go func() {
		for {
			select {
			case event := <-r.eventQueue:
				handler, ok := r.handlers[event.Key]
				if ok {
					r.reportChan <- handler(event)
				} else {
					r.reportChan <- "No handler registered for event type: " + event.Key
				}
			case <-r.stopChan:
				close(r.reportChan)
				return
			}
		}
	}()
}

func (r *Runtime) Stop() {
	close(r.stopChan)
}

func (r *Runtime) ListenResponses() <-chan string {
	return r.reportChan
}

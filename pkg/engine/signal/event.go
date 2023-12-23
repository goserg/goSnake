package signal

type connection interface {
	IsDisposed() bool
}

type removedConnection struct{}

func (r *removedConnection) IsDisposed() bool { return true }

var theRemovedConnection = &removedConnection{}

type eventHandler[T any] struct {
	conn     connection
	callback func(T)
}

type Event[T any] struct {
	handlers []eventHandler[T]
}

func (e *Event[T]) Reset() {
	e.handlers = e.handlers[:0]
}

func (e *Event[T]) Connect(conn connection, slot func(arg T)) {
	e.handlers = append(e.handlers, eventHandler[T]{
		conn:     conn,
		callback: slot,
	})
}

func (e *Event[T]) Disconnect(conn connection) {
	for i, handler := range e.handlers {
		if handler.conn == conn {
			e.handlers[i].conn = theRemovedConnection
		}
	}
}

func (e *Event[T]) Emit(arg T) {
	length := 0
	for _, h := range e.handlers {
		if h.conn != nil && h.conn.IsDisposed() {
			continue
		}
		h.callback(arg)
		e.handlers[length] = h
		length++
	}
	e.handlers = e.handlers[:length]
}

func (e *Event[T]) IsEmpty() bool {
	return len(e.handlers) == 0
}

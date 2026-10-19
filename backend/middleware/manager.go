package middleware

import "net/http"

type Middleware func(next http.Handler) http.Handler

type Manager struct {
	globalMiddlewares []Middleware
}

func NewManager() *Manager {
	return &Manager{
		globalMiddlewares: make([]Middleware, 0),
	}
}

func (mngr *Manager) With(next http.Handler, middlewares ...Middleware) http.Handler {
	nextToPass := next

	for _, middleware := range middlewares {
		nextToPass = middleware(nextToPass)
	}

	return nextToPass
}

func (mngr *Manager) Use(middlewares ...Middleware) {
	mngr.globalMiddlewares = append(mngr.globalMiddlewares, middlewares...)
}

func (mngr *Manager) WrapMux(next http.Handler, middlewares ...Middleware) http.Handler {
	nextToPass := next

	for _, middleware := range mngr.globalMiddlewares {
		nextToPass = middleware(nextToPass)
	}
	return nextToPass
}



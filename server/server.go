package server

import (
	"context"
	"fmt"
	"github.com/IMQS/gateway-store/globals"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Server struct {
	http.Server
	globals        *globals.Globals
	migrationMutex *sync.RWMutex
}

func (s *Server) migrationBlockingHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.migrationMutex.RLock()
		h.ServeHTTP(w, r)
		s.migrationMutex.RUnlock()
	})
}

func (s *Server) loggingHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		s.globals.Log.Debugf("[%s] %q %.3fms", r.Method, r.URL.String(), float64(time.Since(start).Nanoseconds())/float64(1e6))
		h.ServeHTTP(w, r)
	})
}

func (s *Server) noCacheHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Cache-Control", "no-cache, no-store, must-revalidate")
		w.Header().Add("Pragma", "no-cache")
		w.Header().Add("Expires", "0")
		h.ServeHTTP(w, r)
	})
}

func (s *Server) ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"Timestamp": %v}`, time.Now().Unix())
}

// The HandlerFunc type is an adapter to allow the use of
// ordinary functions as HTTP handlers. If f is a function
// with the appropriate signature, HandlerFunc(f) is a
// Handler that calls f.
type HandlerFuncUser func(http.ResponseWriter, *http.Request, int)

// NewServer returns the Server instance
func NewServer(g *globals.Globals) *Server {
	r := mux.NewRouter()

	s := &Server{
		Server: http.Server{
			Addr:    g.Config.Server.Port,
			Handler: r,
		},
		globals:        g,
		migrationMutex: &sync.RWMutex{},
	}

	// Alice creates a middleware chain which groups commonly used http handlers to be
	// executed in a sequential order for each http endpoint and only executes
	// the intended method after all of the methods before has completed.
	defaultHandler := alice.New(s.migrationBlockingHandler, s.loggingHandler, s.noCacheHandler)

	pingHandler := defaultHandler.ThenFunc(s.ping)
	r.Handle("/ping", pingHandler).Methods(http.MethodGet)

	//Client
	createClientHandler := defaultHandler.ThenFunc(s.createClient)
	r.Handle("/client/{id:[0-9]+}/{name:[A-Za-z0-9]+}", createClientHandler).Methods(http.MethodPost)

	readClientAllHandler := defaultHandler.ThenFunc(s.readAllClient)
	r.Handle("/client", readClientAllHandler).Methods(http.MethodGet)

	readClient := defaultHandler.ThenFunc(s.readClient)
	r.Handle("/client/{id:[0-9]+}", readClient).Methods(http.MethodGet)

	deleteClient := defaultHandler.ThenFunc(s.deleteClient)
	r.Handle("/client/{id:[0-9]+}", deleteClient).Methods(http.MethodDelete)

	//Type
	createTypeHandler := defaultHandler.ThenFunc(s.createType)
	r.Handle("/client/{id:[0-9]+}/{name:[A-Za-z0-9]+}", createTypeHandler).Methods(http.MethodPost)

	readTypeAllHandler := defaultHandler.ThenFunc(s.readAllType)
	r.Handle("/client", readTypeAllHandler).Methods(http.MethodGet)

	readType := defaultHandler.ThenFunc(s.readType)
	r.Handle("/client/{id:[0-9]+}", readType).Methods(http.MethodGet)

	deleteType := defaultHandler.ThenFunc(s.deleteType)
	r.Handle("/client/{id:[0-9]+}", deleteType).Methods(http.MethodDelete)

	//Message
	createHandler := defaultHandler.ThenFunc(s.create)
	r.Handle("/{clientId:[0-9]+}/{type:[A-Za-z0-9]+}", createHandler).Methods(http.MethodPost)

	readAllHandler := defaultHandler.ThenFunc(s.readAll)
	r.Handle("/", readAllHandler).Methods(http.MethodGet)

	readHandler := defaultHandler.ThenFunc(s.read)
	r.Handle("/{id:[0-9]+}", readHandler).Methods(http.MethodGet)

	delete := defaultHandler.ThenFunc(s.delete)
	r.Handle("/{id:[0-9]+}", delete).Methods(http.MethodDelete)

	return s
}

// RunHTTPServer sets up the ListenAndServe closer and then starts the http ListenAndServe worker
func (s *Server) RunHTTPServer() error {
	go s.close()

	err := s.ListenAndServe()
	if err != http.ErrServerClosed {
		s.globals.Log.Errorf("Failed to start http listen and serve: %v", err)
		return err
	}
	return nil
}

func (s *Server) close() {
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, syscall.SIGINT, syscall.SIGTERM)
	<-sigint
	s.killSwitch()
}

func (s *Server) killSwitch() {
	err := s.Shutdown(context.Background())
	if err != nil {
		s.globals.Log.Errorf("Failed to shutdown http server: %v", err)
	}
	err = s.globals.Close()
	if err != nil {
		s.globals.Log.Errorf("Failed to stop globals: %v", err)
	}
	s.globals.Log.Error("Process killed")
}

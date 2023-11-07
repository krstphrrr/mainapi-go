package main
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddr string
	store Storage
}
// * pointer to the apiserver type 
// & pointer to the value of a variable 
func NewAPIServer(listenAddr string, store Storage) *APIServer{
	return &APIServer{
		listenAddr: listenAddr,
		store: store,
	}
}

func (s *APIServer) Run(){
	router := mux.NewRouter()

	router.HandleFunc("/gap", makeHTTPHandleFunc(s.handleRequest))
	router.HandleFunc("/gap/{id}", makeHTTPHandleFunc(s.handleGetGap))

	log.Println("JSON API server running on port: ", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) handleRequest(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetGap(w, r)
	}
	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleGetGap(w http.ResponseWriter, r *http.Request) error{
	// return nil
	// vars := mux.Vars(r)
	return WriteJSON(w,http.StatusOK, &Gap{})
}

// HELPER FUNCTIONS
// json header writer: goes inside makehttphandlefunc
func WriteJSON(w http.ResponseWriter, status int, v any) error {
	// write the header into the response, write a header AND
	// encode json message 
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	// new encoder takes a io writer (http.responsewriter)
	return json.NewEncoder(w).Encode(v)
}



type apiFunc func(http.ResponseWriter, *http.Request) error
type ApiError struct {
	Error string
}

// takes apifunc( which takes responsewriter AND request)
func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		if err := f(w,r); err != nil {
		// handle errors
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}
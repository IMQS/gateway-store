package server

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (s *Server) createClient(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) readAllClient(w http.ResponseWriter, r *http.Request) {
	res, err := s.globals.DB.GetAllClients()

	if err != nil {
		s.globals.Log.Errorf("Failed to find data for %s", err)
		http.Error(w, "Failed to get data : "+err.Error(), http.StatusBadRequest)
		return
	}

	jData, err := json.Marshal(&res)
	if err != nil {
		http.Error(w, "Could not marshal JSON", http.StatusInternalServerError)
		return
	}

	n := bytes.IndexByte(jData, 0)
	str := string(jData[:n])
	s.globals.Log.Tracef("%v", str)

	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)
}

func (s *Server) readClient(w http.ResponseWriter, r *http.Request) {
	var rowid int64
	var id int32
	var res interface{}
	var err error

	vars := mux.Vars(r)
	idStr := vars["id"]
	if len(idStr) != 0 {
		if rowid, err = strconv.ParseInt(idStr, 10, 32); err != nil {
			s.globals.Log.Error("Failed to parse 'id' parameter to integer")
			http.Error(w, "Failed to parse 'id' parameter to integer", http.StatusBadRequest)
			return
		}
		id = int32(rowid)
	} else {
		http.Error(w, "'id' parameter missing", http.StatusBadRequest)
		return
	}

	res, err = s.globals.DB.GetClient(id)

	if err != nil {
		s.globals.Log.Errorf("Failed to find data for %s", err)
		http.Error(w, "Failed to get data : "+err.Error(), http.StatusBadRequest)
		return
	}

	jData, err := json.Marshal(res)
	if err != nil {
		http.Error(w, "Could not marshal JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)

}

func (s *Server) updateClient(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) deleteClient(w http.ResponseWriter, r *http.Request) {

}

package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (s *Server) createClient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	name := vars["name"]

	res, err := s.globals.DB.CreateClient(idStr, name)

	if err != nil {
		s.globals.Log.Errorf("Failed to find data for %s", err)
		http.Error(w, "Failed to get data : "+err.Error(), http.StatusBadRequest)
		return
	}

	clientsJSON, err := json.Marshal(res)
	if err != nil {
		http.Error(w, "Could not marshal JSON"+err.Error(), http.StatusInternalServerError)
		return
	}
	s.globals.Log.Debugf("S. Client JSON %v", string(clientsJSON))

	w.Header().Set("Content-Type", "application/json")
	w.Write(clientsJSON)
}

func (s *Server) readAllClient(w http.ResponseWriter, r *http.Request) {

	var res interface{}
	var err error

	res, err = s.globals.DB.GetAllClients()

	if err != nil {
		s.globals.Log.Errorf("Failed to find data for %s", err)
		http.Error(w, "Failed to get data : "+err.Error(), http.StatusBadRequest)
		return
	}

	clientsJSON, err := json.Marshal(res)
	if err != nil {
		http.Error(w, "Could not marshal JSON"+err.Error(), http.StatusInternalServerError)
		return
	}
	s.globals.Log.Debugf("S. Client JSON %v", string(clientsJSON))

	w.Header().Set("Content-Type", "application/json")
	w.Write(clientsJSON)
}

func (s *Server) readClient(w http.ResponseWriter, r *http.Request) {
	var rowid int64
	var id int
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
		id = int(rowid)
	} else {
		http.Error(w, "'id' parameter missing", http.StatusBadRequest)
		return
	}

	res, err = s.globals.DB.GetClient(id)

	if err != nil {
		s.globals.Log.Errorf("Failed to find data for %s", err)
		http.Error(w, "Failed to get data : "+err.Error(), http.StatusNotFound)
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

	var err error
	vars := mux.Vars(r)
	idStr := vars["id"]
	nameStr := vars["name"]
	statusStr := vars["status"]

	if len(idStr) != 0 {
		if _, err = strconv.ParseInt(idStr, 10, 32); err != nil {
			s.globals.Log.Error("Failed to parse 'id' parameter to integer")
			http.Error(w, "Failed to parse 'id' parameter to integer", http.StatusBadRequest)
			return
		}
	} else {
		http.Error(w, "'id' parameter missing", http.StatusBadRequest)
		return
	}

	_, err = s.globals.DB.UpdateClient(idStr, nameStr, statusStr)

	if err != nil {
		s.globals.Log.Errorf("Failed to find data for %s", err)
		http.Error(w, "Failed to get data : "+err.Error(), http.StatusNotFound)
		return
	}
	return
}

func (s *Server) deleteClient(w http.ResponseWriter, r *http.Request) {
	var rowid int64
	var id int
	//	var res interface{}
	var err error

	vars := mux.Vars(r)
	idStr := vars["id"]
	if len(idStr) != 0 {
		if rowid, err = strconv.ParseInt(idStr, 10, 32); err != nil {
			s.globals.Log.Error("Failed to parse 'id' parameter to integer")
			http.Error(w, "Failed to parse 'id' parameter to integer", http.StatusBadRequest)
			return
		}
		id = int(rowid)
	} else {
		http.Error(w, "'id' parameter missing", http.StatusBadRequest)
		return
	}

	_, err = s.globals.DB.DeleteClient(id)

	if err != nil {
		s.globals.Log.Errorf("Failed to find data for %s", err)
		http.Error(w, "Failed to get data : "+err.Error(), http.StatusNotFound)
		return
	}
	return

}

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/IMQS/gateway-store/config"
	"github.com/IMQS/gateway-store/globals"
	server2 "github.com/IMQS/gateway-store/server"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

const originURL = "http://localhost:2011"

var g *globals.Globals
var id int
var count int32 = 1

var data = [...]string{
	`ALTER TABLE public.messages DROP CONSTRAINT messages_clientid_fkey;`,
	`ALTER TABLE public.messages DROP CONSTRAINT messages_type_fkey;`,
	`TRUNCATE TABLE public.client;`,
	`TRUNCATE TABLE public.type;`,
	`TRUNCATE TABLE public.messages;`,
	`ALTER TABLE public.messages
  		ADD CONSTRAINT messages_clientid_fkey FOREIGN KEY (clientid)
      	REFERENCES public.client (clientid) MATCH SIMPLE
      	ON UPDATE NO ACTION ON DELETE NO ACTION;`,
	`ALTER TABLE public.messages
  		ADD CONSTRAINT messages_type_fkey FOREIGN KEY (type)
      REFERENCES public.type (name) MATCH SIMPLE
      ON UPDATE NO ACTION ON DELETE NO ACTION;`,
	`INSERT INTO public.type(name) VALUES ('DEPRECIATION');`,
	`INSERT INTO public.client( clientid, name, status) VALUES ('2' , 'ADM', 'ACTIVE');`,
	`INSERT INTO public.client( clientid, name, status) VALUES ('3' , 'COU', 'ACTIVE');`,
	`INSERT INTO public.messages( clientId, type, message) 
	VALUES ('2', 'DEPRECIATION', 
	'{ "MANDT": "101", "BUKRS": "1000", "ASSET": "000000010000", "ANLN2": "0000", "GJAHR": 2018, "PERAF": 7, 
		"INT_ID": 8573,"BATCHNO": 1, "KOSTL": "FX20006800", "NAFAZ": -907963,"DATETO": "2018-01-31"}'
	);`,
}

func setup() (*globals.Globals, *sql.DB, error) {

	c, err := config.NewConfig("./example/gateway-store.json")
	if err != nil {
		return nil, nil, fmt.Errorf("Error config init %v", err)
	}

	g, err = globals.NewGlobals(c)

	server := server2.NewServer(g)

	//Create a database connection
	pgurl := fmt.Sprintf("user=%s password=%s database=%s host=%s port=%s sslmode=disable",
		g.Config.Db.Username, g.Config.Db.Password, g.Config.Db.Dbname, g.Config.Db.Host, g.Config.Db.Port)
	pgdb, err := sql.Open("postgres", pgurl)

	if err != nil {
		return nil, nil, err
	}

	//Ping the database
	err = pgdb.Ping()
	if err != nil {
		return nil, nil, err
	}

	execute(g, pgdb, data[0])
	execute(g, pgdb, data[1])
	execute(g, pgdb, data[2])
	execute(g, pgdb, data[3])
	execute(g, pgdb, data[4])
	execute(g, pgdb, data[5])
	execute(g, pgdb, data[6])
	execute(g, pgdb, data[7])
	execute(g, pgdb, data[8])
	execute(g, pgdb, data[9])
	execute(g, pgdb, data[10])

	//Run the http server
	run := func() {
		err := server.RunHTTPServer()
		if err != nil {
			g.Log.Errorf("Error running HTTP server: %v", err)
		}
	}

	go run()

	return g, pgdb, nil
}

func execute(g *globals.Globals, pgdb *sql.DB, dt string) {
	_, err := pgdb.Exec(dt)
	if err != nil {
		g.Log.Errorf("Could not execute the SQL %v", err)
		os.Exit(1)
	}
}

func teardown(pgdb *sql.DB) error {
	return nil
}

func doRequest(verb, url string) (*http.Response, error) {
	request, err := http.NewRequest(verb, url, nil)
	if err != nil {
		return nil, err
	}
	return http.DefaultClient.Do(request)
}

func checkHTTPResponseCode(t *testing.T, expected, actual int, status string) {
	if expected != actual {
		t.Errorf("Expected http response code %d . Got %v\n", expected, status)
	}
}

func TestMain(m *testing.M) {

	g, db, err := setup()

	if err != nil {
		fmt.Errorf("Error config init %v", err)
		os.Exit(1)
	}

	if g.Config.Db.Dbname != "gateway_test" {
		fmt.Errorf("Expected gateway_test. Got %s\n", g.Config.Db.Dbname)
		os.Exit(1)
	}

	retCode := m.Run()
	if err := teardown(db); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	os.Exit(retCode)
}

func TestPing(t *testing.T) {
	r, err := doRequest("GET", fmt.Sprintf("%v/ping", originURL))
	if err != nil {
		t.Fatal(err)
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Fatal(err)
	}
	var res map[string]interface{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Print(res)

	checkHTTPResponseCode(t, http.StatusOK, r.StatusCode, r.Status)
}

func TestReadAllClient(t *testing.T) {
	r := get(t, "client")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Fatal(err)
	}

	var clientsMap map[string]interface{}
	err = json.Unmarshal(body, &clientsMap)
	if err != nil {
		t.Fatal(err)
	}

	clients := clientsMap["clients"].([]interface{})
	client1 := clients[0].(map[string]interface{})
	name := client1["name"]
	g.Log.Debugf("T. Name : %v", name)
	if name != "ADM" {
		t.Fatalf("Returned id mismatch, expected '%v' got '%v'", "ADM", name)
	}
}

func TestReadClient(t *testing.T) {

	row := g.DB.QueryRow("2")
	row.Scan(&id)
	g.Log.Debugf(" id %d", id)

	r := get(t, "client/"+fmt.Sprint(id))
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Fatal(err)
	}
	var client map[string]interface{}
	err = json.Unmarshal(body, &client)
	if err != nil {
		t.Fatal(err)
	}
	name := client["name"]
	g.Log.Debugf("T. Name : %v", name)
	if name != "ADM" {
		t.Fatalf("Returned id mismatch, expected '%v' got '%v'", "ADM", name)
	}
}

func TestDeleteClient(t *testing.T) {
	row := g.DB.QueryRow("2")
	row.Scan(&id)
	g.Log.Debugf(" id %d", id)

	r := delete(t, "client/"+fmt.Sprint(id))
	_, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Fatal(err)
	}

	r, err = doRequest("GET", fmt.Sprintf("%v/%v", originURL, "client/"+fmt.Sprint(id)))
	if err != nil {
		t.Fatal(err)
	}
	checkHTTPResponseCode(t, http.StatusOK, r.StatusCode, r.Status)
}

func TestReadAllMessage(t *testing.T) {
	get(t, "")
}

func get(t *testing.T, url string) *http.Response {
	r, err := doRequest("GET", fmt.Sprintf("%v/%v", originURL, url))
	if err != nil {
		t.Fatal(err)
	}
	checkHTTPResponseCode(t, http.StatusOK, r.StatusCode, r.Status)
	return r
}

func delete(t *testing.T, url string) *http.Response {
	r, err := doRequest("DELETE", fmt.Sprintf("%v/%v", originURL, url))
	if err != nil {
		t.Fatal(err)
	}
	checkHTTPResponseCode(t, http.StatusOK, r.StatusCode, r.Status)
	return r
}

package main

import (
	"database/sql"
	"encoding/json"
	"errors"
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

var data = [...]string{
	`INSERT INTO public.type(name) VALUES ('DEPRECIATION');`,
	`INSERT INTO public.client( clientid, name) VALUES ('2' , 'ADM');`,
	`INSERT INTO public.client( clientid, name) VALUES ('3' , 'COU');`,
	`INSERT INTO public.messages( clientId, type, message) 
	VALUES ('2', 'DEPRECIATION', 
	'{ "MANDT": "101", "BUKRS": "1000", "ASSET": "000000010000", "ANLN2": "0000", "GJAHR": 2018, "PERAF": 7, "INT_ID": 8573,"BATCHNO": 1, "KOSTL": "FX20006800", "NAFAZ": -907963,"DATETO": "2018-01-31"}'
	);`,
}

var count int32 = 1

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
	defer pgdb.Close()

	//Ping the database
	err = pgdb.Ping()
	if err != nil {
		return nil, nil, err
	}

	err = execute(g, pgdb, data[0])
	err = execute(g, pgdb, data[1])
	err = execute(g, pgdb, data[2])

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

func execute(g *globals.Globals, pgdb *sql.DB, dt string) error {
	_, err := pgdb.Exec(dt)
	if err != nil {
		g.Log.Errorf("Could not insert the data %v", err)
		return errors.New("Setup failed. Data insertion failed")
	}
	return nil
}

func teardown(pgdb *sql.DB) error {
	_, err := pgdb.Exec("DROP DATABASE gateway_test;")
	if err != nil {
		g.Log.Errorf("Could not insert the data %v", err)
		return errors.New("Setup failed. Could not delete the database")
	}
	return nil
}

func doRequest(verb, url string) (*http.Response, error) {
	request, err := http.NewRequest(verb, url, nil)
	if err != nil {
		return nil, err
	}
	return http.DefaultClient.Do(request)
}

func checkHTTPResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected http response code %d. Got %d\n", expected, actual)
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

	checkHTTPResponseCode(t, http.StatusOK, r.StatusCode)
}

func TestReadAllClient(t *testing.T) {
	r := get(t, "client")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Fatal(err)
	}

	var clientsMap map[string][]map[string]string
	err = json.Unmarshal(body, &clientsMap)
	if err != nil {
		t.Fatal(err)
	}
	g.Log.Debugf("Clients List : %v", clientsMap)

	clients := clientsMap["clients"]

	client := clients[1]

	name := client["name"]
	if name != "ADM" {
		t.Fatalf("Returned id mismatch, expected '%v' got '%v'", "ADM", name)
	}
}

func TestReadAllMessage(t *testing.T) {
	get(t, "")
}

func get(t *testing.T, url string) *http.Response {
	r, err := doRequest("GET", fmt.Sprintf("%v/%v", originURL, url))
	if err != nil {
		t.Fatal(err)
	}
	checkHTTPResponseCode(t, http.StatusOK, r.StatusCode)
	return r
}

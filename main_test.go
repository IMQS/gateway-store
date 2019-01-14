package main

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/IMQS/gateway-store/config"
	"github.com/IMQS/gateway-store/database"
	"github.com/IMQS/gateway-store/globals"
	server2 "github.com/IMQS/gateway-store/server"
	"net/http"
	"os"
	"testing"
)

const originURL = "http://localhost:2011"

var g *globals.Globals

var data = [...]string{
	`INSERT INTO public.messages( clientId, type, message) 
	VALUES ('2', 'DEPRECIATION', 
	'{ "MANDT": "101", "BUKRS": "1000", "ASSET": "000000010000", "ANLN2": "0000", "GJAHR": 2018, "PERAF": 7, "INT_ID": 8573,"BATCHNO": 1, "KOSTL": "FX20006800", "NAFAZ": -907963,"DATETO": "2018-01-31"}'
	);`,
}

func setup() (*globals.Globals, error) {

	c, err := config.NewConfig("./example/gateway-store.json")
	if err != nil {
		return nil, fmt.Errorf("Error config init %v", err)
	}

	g, err := globals.NewGlobals(c)

	server := server2.NewServer(g)

	//Create a database connection
	pgurl := fmt.Sprintf("user=%s password=%s database=%s host=%s port=%s sslmode=disable",
		g.Config.Db.Username, g.Config.Db.Password, g.Config.Db.Dbname, g.Config.Db.Host, g.Config.Db.Port)
	pgdb, err := sql.Open("postgres", pgurl)

	if err != nil {
		return nil, err
	}
	defer pgdb.Close()

	//Ping the database
	err = pgdb.Ping()
	if err != nil {
		return nil, err
	}

	_, err = pgdb.Exec(data[0])
	if err != nil {
		g.Log.Errorf("Could not insert the data %v", err)
		return nil, errors.New("Setup failed")
	}

	//Run the http server
	run := func() {
		err := server.RunHTTPServer()
		if err != nil {
			g.Log.Errorf("Error running HTTP server: %v", err)
		}
	}

	go run()

	return g, nil
}

func teardown() error {
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

	g, err := setup()

	if err != nil {
		fmt.Errorf("Error config init %v", err)
		os.Exit(1)
	}

	if g.Config.Db.Dbname != "gateway_test" {
		fmt.Errorf("Expected gateway_test. Got %s\n", g.Config.Db.Dbname)
		os.Exit(1)
	}

	g.Log.Debug("Created the log")

	retCode := m.Run()
	if err := teardown(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	os.Exit(retCode)
}

func TestDatabase(t *testing.T) {
	g, err := setup()

	if err != nil {
		t.Errorf("Error config init %v", err)
	}

	db, err := database.NewDBServer(g.Config.Db, g.Log)

	defer db.Close()
}

func TestPing(t *testing.T) {
	r, err := doRequest("GET", fmt.Sprintf("%v/ping", originURL))
	if err != nil {
		t.Fatal(err)
	}
	checkHTTPResponseCode(t, http.StatusOK, r.StatusCode)
}

func TestGet(t *testing.T) {
	r, err := doRequest("GET", fmt.Sprintf("%v/", originURL))
	if err != nil {
		t.Fatal(err)
	}
	checkHTTPResponseCode(t, http.StatusOK, r.StatusCode)
}

package main

import (
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/juju/mgosession"
	mgo "gopkg.in/mgo.v2"
)

func main() {
	err := godotenv.Load("config/.env")
	checkErr(err)

	session, err := mgo.Dial(os.Getenv("MONGODB_HOST"))
	checkErr(err)
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	cPool, err := strconv.Atoi(os.Getenv("MONGODB_CONNECTION_POOL"))
	checkErrMongo(err)
	mPool := mgosession.NewPool(nil, session, cPool)
	defer mPool.Close()

	r := mux.NewRouter()
	http.Handle("/", r)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

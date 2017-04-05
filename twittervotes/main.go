package main

import (
	"gopkg.in/mgo.v2"
	"log"
)

func main() {

}

// TODO:  find an elegant way to make the location of the MongoDB instance configurable so that we don't need to run it locally.
var db *mgo.Session
func dialdb() error {
	var err error
	log.Println("dialing mongodb: localhost")
	db, err = mgo.Dial("localhost")
	return err
}
func closedb() {
	db.Close()
	log.Println("closed database connection")
}

type poll struct {
	Options []string
}
func loadOptions() ([]string, error) {
	var options []string
	iter := db.DB("ballots").C("polls").Find(nil).Iter()
	var p poll
	for iter.Next(&p) {
		options = append(options, p.Options...)
	}
	iter.Close()
	return options, iter.Err()
}

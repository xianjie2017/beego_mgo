package controllers

import (
	"gopkg.in/mgo.v2"
)

const (
	URL    = "127.0.0.1:27017"
	DBNAME = "runoob"
)

var (
	session *mgo.Session
)

func getSession() *mgo.Session {
	if session == nil {
		var err error
		session, err = mgo.Dial(URL)
		if err != nil {
			panic(err)
		}
	}
	return session.Clone()
}

func Execute(collection string, f func(*mgo.Collection) error) error {
	session := getSession()
	defer session.Close()
	c := session.DB(DBNAME).C(collection)
	return f(c)
}

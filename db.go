package main

import (
	"fmt"
	"net/http"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// DB - database object
type DB struct {
	url string
}

func (db DB) Connect() *mgo.Session {
	session, err := mgo.Dial(db.url)
	if err != nil {
		panic(err)
	}

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	return session
}

func (db DB) Create(database string, collection string, data FormData) (result bson.M, code int, err error) {
	code = http.StatusInternalServerError
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("create: %v", r)
		}
	}()

	session := db.Connect()
	defer session.Close()

	coll := session.DB(database).C(collection)
	data.Data["_id"] = bson.NewObjectId()
	if ret := coll.Insert(data.Data); ret != nil {
		// code = http.StatusNotFound
		panic(ret)
	}
	result = data.Data
	code = http.StatusOK

	return
}

func (db DB) DeleteByID(database string, collection string, id string) (code int, err error) {
	code = http.StatusInternalServerError
	defer func() {
		if r := recover(); r != nil {
			// fmt.Printf("[DEBUG] = %#v \n", r)
			err = fmt.Errorf("findOne: %v", r)
			// runtime.Breakpoint()
		}
	}()
	session := db.Connect()
	defer session.Close()

	coll := session.DB(database).C(collection)

	// err2 := c.Find(bson.M{"_id": bson.ObjectIdHex("58593d1d6aace357b32bb3a1")}).One(&data)
	if ret := coll.Remove(bson.M{"_id": bson.ObjectIdHex(id)}); ret != nil {
		code = http.StatusNotFound
		panic(ret)
	}
	code = http.StatusOK

	return
}

func (db DB) FindAll(database string, collection string) (results []bson.M, code int, err error) {
	code = http.StatusInternalServerError
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("findAll: %v", r)
		}
	}()

	session := db.Connect()
	defer session.Close()

	coll := session.DB(database).C(collection)

	if ret := coll.Find(nil).All(&results); ret != nil {
		code = http.StatusNotFound
		panic(ret)
	}
	code = http.StatusOK

	return
}

func (db DB) FindByID(database string, collection string, id string) (result bson.M, code int, err error) {
	code = http.StatusInternalServerError
	defer func() {
		if r := recover(); r != nil {
			// fmt.Printf("[DEBUG] = %#v \n", r)
			err = fmt.Errorf("findOne: %v", r)
			// runtime.Breakpoint()
		}
	}()
	session := db.Connect()
	defer session.Close()

	coll := session.DB(database).C(collection)

	// err2 := c.Find(bson.M{"_id": bson.ObjectIdHex("58593d1d6aace357b32bb3a1")}).One(&data)
	if ret := coll.FindId(bson.ObjectIdHex(id)).One(&result); ret != nil {
		code = http.StatusNotFound
		panic(ret)
	}
	code = http.StatusOK

	return
}

func (db DB) FindOne(database string, collection string, key string, value string) (result bson.M, code int, err error) {
	code = http.StatusInternalServerError
	defer func() {
		if r := recover(); r != nil {
			// fmt.Printf("[DEBUG] = %#v \n", r)
			err = fmt.Errorf("findOne: %v", r)
			// runtime.Breakpoint()
		}
	}()
	session := db.Connect()
	defer session.Close()

	coll := session.DB(database).C(collection)

	if ret := coll.Find(bson.M{key: value}).One(&result); ret != nil {
		code = http.StatusNotFound
		panic(ret)
	}
	code = http.StatusOK

	return
}

func (db DB) UpdateByID(database string, collection string, id string, data bson.M) (code int, err error) {
	code = http.StatusInternalServerError
	defer func() {
		if r := recover(); r != nil {
			// fmt.Printf("[DEBUG] = %#v \n", r)
			err = fmt.Errorf("findOne: %v", r)
			// runtime.Breakpoint()
		}
	}()
	session := db.Connect()
	defer session.Close()

	coll := session.DB(database).C(collection)

	if ret := coll.Update(bson.M{"_id": bson.ObjectIdHex(id)}, bson.M{"$set": data}); ret != nil {
		code = http.StatusNotFound
		panic(ret)
	}
	code = http.StatusOK

	return
}

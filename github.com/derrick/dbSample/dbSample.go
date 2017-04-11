package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"time"
)

type Person struct {
	Name string
	Phone string
}

func main () {
	mangoSample()
}

/**
使用第三方package mgo
*/

func mangoSample(){
	//获取session
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	//关闭session
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	//insert
	c := session.DB("test").C("people")
	err = c.Insert(&Person{"Tom", "+86 133 1234 1234"},
		&Person{"Jerry", "+86 133 5678 5678"})
	if err != nil {
		log.Fatal(err)
	}

	//find one
	result := Person{}
	err = c.Find(bson.M{"name": "Tom"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Phone:", result.Phone)

	// Update
	colQuerier := bson.M{"name": "Tom"}
	change := bson.M{"$set": bson.M{"phone": "+86 99 8888 7777", "timestamp": time.Now()}}
	err = c.Update(colQuerier, change)
	if err != nil {
		panic(err)
	}
	result1 := Person{}
	err = c.Find(bson.M{"name": "Tom"}).One(&result1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Phone:", result1.Phone)
}
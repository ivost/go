package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gocql/gocql"
	"github.com/kristoiv/gocqltable"
	"github.com/kristoiv/gocqltable/recipes"
)

// https://github.com/gocql/gocql
/* Before you execute the program, Launch `cqlsh` and execute:
create keyspace example with replication = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };
create table example.tweet(timeline text, id UUID, text text, PRIMARY KEY(id));
create index on example.tweet(timeline);
*/
func TryGocql() {
	//justDriver()
	tryGocqlTable()
}

func justDriver() {

	log.Print("TryGocql")

	// connect to the cluster
	cluster := gocql.NewCluster("localhost")
	cluster.Port = 9042
	cluster.Keyspace = "example"
	cluster.Consistency = gocql.Quorum
	// need this for c* 3.0
	//cluster.CQLVersion = "4.0.0"
	cluster.ProtoVersion = 4

	session, _ := cluster.CreateSession()
	defer session.Close()

	log.Print("Inserting tweet")
	// insert a tweet
	timeline := "other"
	if err := session.Query(`INSERT INTO tweet (timeline, id, text) VALUES (?, ?, ?)`,
		timeline, gocql.TimeUUID(), "hello Ivo !!!").Exec(); err != nil {
		log.Fatal(err)
	}
	var id gocql.UUID
	var text string

	log.Print("select tweets from me")

	//if err := session.Query(`SELECT id, text FROM tweet WHERE timeline = ? LIMIT 1`, timeline).
	//	Consistency(gocql.One).Scan(&id, &text); err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("Tweet:", id, text)

	// list all tweets
	iter := session.Query(`SELECT id, text FROM tweet WHERE timeline = ?`, timeline).Iter()
	for iter.Scan(&id, &text) {
		fmt.Println("Tweet:", id, text)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}

}

func tryGocqlTable() {

	log.Print("tryGocqlTable")

	// connect to the cluster
	cluster := gocql.NewCluster("localhost")
	cluster.Port = 9042
	cluster.Keyspace = "example"
	cluster.Consistency = gocql.Quorum
	// need this for c* 3.0
	//cluster.CQLVersion = "4.0.0"
	cluster.ProtoVersion = 4

	session, err := cluster.CreateSession()
	defer session.Close()

	// Tell gocqltable to use this session object as the default for new objects
	gocqltable.SetDefaultSession(session)


	// Now we're ready to create our first keyspace. We start by getting a keyspace object
	keyspace := gocqltable.NewKeyspace("gocqltable_test")

	// Now lets create that in the database using the simple strategy and durable writes (true)
	err = keyspace.Create(map[string]interface{}{
		"class": "SimpleStrategy",
		"replication_factor": 1,
	}, true)

	//if err != nil { // If something went wrong we print the error and quit.
	//	log.Fatalln(err)
	//}


	// Now that we have a very own keyspace to play with, lets create our first table.

	// First we need a Row-object to base the table on. It will later be passed to the table wrapper
	// to be used for returning row-objects as the answer to fetch requests.
	type User struct{
		Email string // Our primary key
		Password string `password`     // Use Tags to rename fields
		Active bool     `cql:"active"` // If there are multiple tags, use `cql:""` to specify what the table column will be
		Created time.Time
	}

	// Let's define and instantiate a table object for our user table
	userTable := struct{
		recipes.CRUD    // If you looked at the base example first, notice we replaced this line with the recipe
	}{
		recipes.CRUD{ // Here we didn't replace, but rather wrapped the table object in our recipe, effectively adding more methods to the end API
			keyspace.NewTable(
				"users",            // The table name
				[]string{"email"},  // Row keys
				nil,                // Range keys
				User{},             // We pass an instance of the user struct that will be used as a type template during fetches.
			),
		},
	}

	// Lets create this table in our cassandra database
	err = userTable.Create()
	//if err != nil {
	//	log.Fatalln(err)
	//}


	// Now that we have a keyspace with a table in it: lets make a few rows! In the base example we had to write out the CQL manually, this time
	// around, however, we can insert entire User objects.

	// Lets instantiate a user object, set its values and insert it
	user1 := User{
		Email: "foo@example.com",
		Password: "123456",
		Active: true,
		Created: time.Now().UTC(),
	}
	err = userTable.Insert(user1)
	if err != nil {
		log.Fatalln(err)
	}


	// With our database filled up with users, lets query it and print out the results (containing all users in the database).
	rowset, err := userTable.List() // Our rowset variable is a "interface{}", and here we type assert it to a slice of pointers to "User"
	for _, user := range rowset.([]*User) {
		fmt.Println(user)
	}
	if err != nil {
		log.Fatalln(err)
	}


	// You can also fetch a single row, obviously
	row, err := userTable.Get("1@example.com")
	if err != nil {
		log.Fatalln(err)
	}
	user := row.(*User)


	// Lets update this user by changing his password
	user.Password = "654321"
	err = userTable.Update(user)
	if err != nil {
		log.Fatalln(err)
	}


	//// Lets delete user 1@example.com
	//err = userTable.Delete(user)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//// Lets clean up after ourselves by dropping the keyspace.
	//keyspace.Drop()

}




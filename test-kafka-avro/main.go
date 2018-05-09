package main

import (
	"log"
	"github.com/linkedin/goavro"
	"github.com/dangkaka/go-kafka-avro"
	"github.com/bsm/sarama-cluster"
	"fmt"
)

// go get github.com/dangkaka/go-kafka-avro
// 	"github.com/Shopify/sarama"

var   brokers = []string{"localhost:9092"}
var   registries = []string{"http://localhost:8081"}

const group = "G1"

// schema name aka subject
const schemaName = "test"

const topicName = "test"

var codec      *goavro.Codec
var err error

var verbose = true
// if true will delete ALL schemas
var cleanup = false

type TestObject struct {
	Codec      *goavro.Codec
	Subject    string
	Id         int
	Count      int
}

var testObj *TestObject

func init() {
	codec, err = goavro.NewCodec(`
        {
          "type": "record",
          "name": "test",
          "fields" : [
            {"name": "val", "type": "int", "default": 0}
          ]
        }`)
	checkError(err)
	testObj = makeTestObject("test", 1)
}

func main() {
	//registry()
	//go produce()
	produce()
	consume()
}

func produce() {
	prod, err := kafka.NewAvroProducer(brokers, registries)
	checkError(err)
	defer prod.Close()
	for i := 0; i < 10; i++ {
		key := []byte("key")
		val := fmt.Sprintf("{\"val\":%d}",i)
		err = prod.Add(topicName, testObj.Codec.Schema(), key, []byte(val))
		checkError(err)
		log.Printf("message %v sent OK", i)
		// artificial delay - just for testing
		// side effect - some strange things with rebalance - to investigate
		//time.Sleep(10 * time.Millisecond)
	}
}

func consume() {

	callbacks := &kafka.ConsumerCallbacks{
		OnDataReceived:func(m kafka.Message) {
			log.Printf("message received: %v", string(m.Value))
		},
		OnError:func(err error) {
			log.Printf("error: %v", err.Error())
		},
		OnNotification: func(notification *cluster.Notification) {
			log.Printf("cluster.Notification: %v", notification.Type.String())
		},
	}

	cons, err := kafka.NewAvroConsumer(brokers, registries, schemaName, group, *callbacks)
	checkError(err)
	// this will block - callbacks will trigger
	cons.Consume()
	log.Printf("after consume")
}


func registry() {
	log.Println("registry 1.0 start...")
	r := kafka.NewSchemaRegistryClientWithRetries(registries, 2)
	subjects, err := r.GetSubjects()
	checkError(err)
	log.Printf("subjects: %v", subjects)

	if cleanup {
		for _, s := range(subjects) {
			log.Printf("deleting %v", s)
			err = r.DeleteSubject(s)
			if err != nil {
				log.Printf("error %v", err)
			}
		}
		return
	}

	// IsSchemaRegistered tests if the schema is registered, if so it returns the unique id of that schema
	id, err := r.IsSchemaRegistered(schemaName, codec)
	//checkError(err)
	if err != nil || id == 0 {
		log.Printf("Registering schema %v", schemaName)
		id, err := r.CreateSubject(schemaName, codec)
		checkError(err)
		log.Printf("Registered schema %v, id %v", schemaName, id)
	} else {
		log.Printf("Schema %v already registered, id %v", schemaName, id)
	}
	if verbose {
		for _, s := range(subjects) {
			// sc is avro codec
			sc, _ := r.GetLatestSchema(s)
			log.Printf("%v schema: %v", s, sc.CanonicalSchema())
			vv, err := r.GetVersions(s)
			checkError(err)
			log.Printf("%v versions: %v", s, vv)
			//for _, id := range(vv) {
			//	sc, err = r.GetSchemaByVersion(s, id)
			//	checkError(err)
			//}

		}
		log.Println("registry end")
	}
}


func checkError(err error) {
	if err != nil {
		log.Fatalf("Error %v", err)
	}
}

func makeTestObject(subject string, id int) *TestObject {
	obj := &TestObject{}
	obj.Subject = subject
	obj.Id = id
	obj.Codec = codec
	return obj
}

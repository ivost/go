package main

import (
	"log"
	"github.com/linkedin/goavro"
	"github.com/dangkaka/go-kafka-avro"
)
// go get github.com/dangkaka/go-kafka-avro

const regUrl = "http://localhost:8081"

// schema name aka subject
const schemaName = "S2"

type TestObject struct {
	Codec      *goavro.Codec
	Subject    string
	Id         int
	Count      int
}

func main() {
	log.Println("Test 1.0 start...")

	codec, err := goavro.NewCodec(`
        {
          "type": "record",
          "name": "` + schemaName + `",
          "fields" : [
            {"name": "id", "type": "int"},
            {"name": "firstName", "type": "string", "default": "John"},
            {"name": "lastName", "type": "string", "default": "Doe"},
            {"name": "age", "type": "int"}
          ]
        }`)

	checkError(err)

	r := kafka.NewSchemaRegistryClientWithRetries([]string{regUrl}, 2)

	subjects, err := r.GetSubjects()
	checkError(err)
	log.Printf("subjects: %v", subjects)

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
	log.Println("Test end")
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("Error %v", err)
	}
}

/*testObject := &TestObject{}
testObject.Subject = "subject"
testObject.Id = 123
testObject.Count = 42
testObject.Codec = codec
*/
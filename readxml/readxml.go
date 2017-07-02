package main

// https://golang.org/pkg/encoding/xml/#example_Unmarshal

// http://blog.davidsingleton.org/parsing-huge-xml-files-with-go/

import (
	"encoding/xml"
	"time"
	"log"
	"os"
	"fmt"
	"flag"
	"strings"
)
//var inputFile = flag.String("infile", "/Users/ivo/github/go/data/test.xml", "Input file path")
//var inputFile = flag.String("infile", "/Users/stoyaiv/github/go/data/test.xml", "Input file path")
var inputFile = flag.String("infile", "/Users/stoyaiv/Documents/SFTP/NPHASE_SLTX_20160927_222101_0071926_FRM3.xml", "Input file path")

type UpdateFunc func(result []string) (updated []byte, err error)

var fields = map[string]int{
	"servedIMSI": 1,
	}

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	flag.Parse()
	// callback - called on every record with slice of field values
	// in fixed order
	process := func(fields []string) {
		log.Printf("= %v", fields)
	}
	xmlStreamRead(process);
}


func xmlStreamRead(callback func(fields []string)) {
	defer elapsed(time.Now())
	xmlFile, err := os.Open(*inputFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer xmlFile.Close()

	decoder := xml.NewDecoder(xmlFile)
	field := ""
	value := ""
	size  := len(fields)
	collect := make([]string, size)
	index := 0
	count := 0
	for {
		// Read tokens from the XML document in a stream.
		token, _ := decoder.Token()
		if token == nil {
			break
		}
		// Inspect the type of the token just read.
		switch element := token.(type) {
		case xml.StartElement:
			field = element.Name.Local
			if _, present := fields[field]; !present {
				//log.Printf("field: %v", field)
				field = ""
			}
		case xml.CharData:
			if len(field) > 0 {
				value = string([]byte(element))
				value = strings.TrimSpace(value)
				if len(value) > 0 {
					//log.Printf("%v: %v", field, value)
					collect[index] = value
					index++
					if index == size {
						callback(collect)
						index = 0
						count++;
					}
				}
				field = ""
			}
		case xml.EndElement:
			//field = ""
		default:
			//log.Printf("t %v", string( ([] byte) t))
		}
	}
	log.Printf("count %v", count)
}

func elapsed(start time.Time) {
	log.Printf("=== elapsed time: %s", time.Since(start))
}

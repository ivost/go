package store

import (
	"net/http"
	"fmt"
	"os"
	"io"
	"log"
	"bufio"
	// don't use - problem on mac with \r line ending "encoding/csv"
	// "encoding/csv"
	"strings"
	"server/model"
	"strconv"
	"server/database"
)

func Upload(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("upload - method:", r.Method)
	if r.Method != "POST" {
		return
	}

	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("file")
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()

	filename := "./data/" + handler.Filename

	//f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	io.Copy(f, file)
	f.Close()
	fmt.Fprintf(w, "File %s uploaded", filename)
	Process(filename)
}

func Process(filename string) (err1 error) {
	log.Printf("process %s", filename)

	// parse
	recs, err := Parse(filename)
	if err != nil {
		return err
	}
	// store to db
	for _, rec := range recs {
		err := database.AddPoi(&rec)
		if err != nil {
			err1 = err
			log.Printf("err: %v\n", err)
		}
	}
	return err1
}

func Parse(filename string) (res [] model.POI, err error) {
	// on mac
	lineEnd := byte('\r')

	res = make([] model.POI, 0, 100)
	log.Printf("parse %s", filename)
	file, err := os.Open(filename)
	if err != nil {
		return res, err
	}
	defer file.Close()
	file.Seek(0, 0)
	reader := bufio.NewReader(file)

	for i := 0; ; i++ {
		// \r on mac - will be included in line
		line, err := reader.ReadString(lineEnd)
		if err != nil {
			//log.Printf("err: %v\n", err)
			break
		}
		// skip 1st line - header
		if i == 0 || len(line) < 2 {
			continue
		}
		line = line[:len(line)-1]
		f := strings.Split(line, "\t")
		if len(f) < 9 {
			log.Printf("invalid record: %v\n", line)
			continue
		}
		lat := ToFloat64(f[6])
		lon := ToFloat64(f[7])
		radius := ToFloat64(f[8])

		rec := model.POI{f[0], f[1], f[2], f[3], f[4], f[5],
				   lat, lon, radius}
		//log.Printf("record: %s - lat: %f, lon: %f", rec.Name, rec.Lat, rec.Lng)
		res = append(res, rec)
	}

	return res, nil
}


func ToFloat64(s string) (res float64) {
	var err error
	res, err = strconv.ParseFloat(s, 64)
	if err != nil {
		log.Printf("Invalid number: %v", s)
		res = 0.0
	}
	return res
}


package main

import (
	"bufio"
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/gocql/gocql"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var BUCKET_MODULO int32 = 1600

var count int = 0
var total_count int = 0

// configuration
// having strange problems when trying lowercase keys (reflection?)

type MyConfig struct {
	Host        string
	Keyspace    string
	Consistency string
	Filename    string
	Foreignid   string
	Tenantid    string

	Truncate bool
	Random   bool
	Repeat   bool
	Count    int
	Progress int
	Delay    int
	Port     int
	Ttl      int
}

var myconfig MyConfig

type Device struct {
	deviceid             string
	devicebucket         int32
	foreignid            string
	tenantid             string
	mobilenumber         uint64 // mdn/msisdn
	iccid                string
	esn                  string
	imei                 string
	imsi                 uint64
	billingcyclestartday int
	timezone             string
	mac                  string
	qrcode               string
	serial               string
	refid                string
	createdon            string
	lastupdated          string
}

type Extract struct {
	costcenter   string //  0
	ecpd         string //  1
	custid       string //  2
	mtasacnt     string //  3
	iccid        string //  5
	esn          string //  6
	mobilenumber string //  7 msisdn?
	ipaddr       string //  8
	priceplan    string //  9
	status       string // 13
	reasoncode   string // 14
	actdate      string // 15
}

//type ExtractRowHandler func(line string) (extract Extract, err error)
type ExtractRowHandler func(line string)

//var cluster *gocql.ClusterConfig
var session *gocql.Session

var insert *gocql.Query

var findByMdn *gocql.Query

var findByDeviceId *gocql.Query

var test_mdn uint64

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const lettersLen = len(letters)

const digits = "0123456789"
const hexdigits = "0123456789ABCDEF"
const version = "v.1.0.11.22"

func init() {
	log.SetOutput(os.Stdout)
	flag.Parse()

	log.Printf("datagen " + version)
	println("datagen " + version)

	myconfig = readConfig()

	rand.Seed(42)

	// connect to the cluster
	cluster := gocql.NewCluster(myconfig.Host)
	cluster.Port = myconfig.Port
	cluster.Keyspace = myconfig.Keyspace
	if myconfig.Consistency == "one" {
		cluster.Consistency = gocql.One
		log.Print("Using consisteny: one")
	} else {
		log.Print("Using consisteny: quorum")
		cluster.Consistency = gocql.Quorum
	}

	cluster.CQLVersion = "3.4.0"
	cluster.ProtoVersion = 4
	ses, err := cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
	}
	session = ses
	// prepare insert
	// createdon 2016-11-18 09:01:02
	insert = session.Query(`INSERT INTO deviceauxinfo
		(devicebucket, deviceid, foreignid, tenantid, mobilenumber, iccid, esn, createdon)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?) USING TTL ?`)

	// prepare selects
	findByDeviceId = session.Query(`SELECT mobilenumber FROM deviceauxinfo
		WHERE deviceid = ? AND devicebucket = ?`)
	findByMdn = session.Query(`SELECT deviceid, devicebucket FROM deviceauxinfo_by_mobilenumber
		WHERE mobilenumber = ? `)
}

var file_log = true

///////////////////////
func main() {
	//flag.Parse()
	if file_log {
		f, err := os.OpenFile("datagen.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err == nil {
			println("Logging to datagen.log")
			log.SetOutput(f)
			defer f.Close()
		}
	}
	defer session.Close()

	if myconfig.Truncate {
		truncateData()
	}

	if myconfig.Random {
		loadRandom()
	} else {
		readExtract(handler)
	}
	if myconfig.Repeat {
		for {
			//d Duration = 5 * time.Second
			//time.Sleep(int(myconfig.Delay) * time.Second)
			d := 5 * time.Second
			time.Sleep(d)
			count = 0
			loadRandom()
			total_count += count
			if total_count%myconfig.Progress == 0 {
				log.Print("Total inserted: ", total_count)
				println("Total inserted: ", total_count)
			}
			sanityCheck()
		}
	}
	sanityCheck()
	log.Print("Exit, devices inserted: ", count)
	println("Exit, devices inserted: ", count)
}

///////////////////////

func sanityCheck() {
	// sanity check - refactor, call on every iteration
	// add update, delete, tickle
	if test_mdn == 0 {
		return
	}

	id, idbucket := findMdn(test_mdn)
	//log.Printf("mdn: %d -> found id: %s, devicebucket: %d ", test_mdn, id, idbucket)

	mdn2 := findMobileNumber(id, idbucket)
	//log.Printf("mdn from base table: %d", mdn2)
	if test_mdn != mdn2 {
		log.Fatal("*** Sanity check failed!!!")
		println("*** Sanity check failed!!!")
	}
	//log.Print("Passed sanity check.")
	test_mdn = 0
}

func readConfig() (conf MyConfig) {
	var configfile = "datagen.config"
	_, err := os.Stat(configfile)
	if err != nil {
		log.Fatal("Config file is missing: ", configfile)
	}
	// default values
	conf.Host = "localhost"
	conf.Port = 9042
	conf.Keyspace = "sandbox"
	conf.Consistency = "one"

	if _, err := toml.DecodeFile(configfile, &conf); err != nil {
		log.Fatal(err)
	}
	log.Printf("%v", conf)
	return conf
}

func loadRandom() {
	for count = 0; count < myconfig.Count; count++ {
		device := randomDevice()
		err := insertDevice(&device)
		// save one mdn for sanity check after load
		if err == nil && test_mdn == 0 {
			test_mdn = device.mobilenumber
		}
	}
}

var handler ExtractRowHandler = func(line string) {
	//log.Print("line: ", line)
	s := strings.Split(line, "|")
	extract := Extract{s[0], s[1], s[2], s[3], s[5], s[6], s[7], s[8], s[9], s[13], s[14], s[15]}
	//log.Printf("Extract %v", extract)

	// insert device
	device := deviceFromFeed(extract)
	count++
	err := insertDevice(&device)
	// save one mdn for sanity check after load
	if err == nil && test_mdn == 0 {
		test_mdn = device.mobilenumber
	}
	//return extract, nil
}

func safenum(input string) uint64 {
	//log.Print("safenum: ", input)
	if len(input) == 0 {
		return 0
	}
	num, err := strconv.ParseUint(input, 10, 64)
	if err != nil {
		log.Print("err: ", err)
		return 0
	}
	return uint64(num)
}

func deviceFromFeed(feed Extract) (device Device) {
	//log.Printf("deviceFromFeed %v", feed)
	//var device Device //{id, idbucket, foreignid, tenantid}
	device.deviceid = gocql.TimeUUID().String()
	device.devicebucket = calcBucket(device.deviceid)
	device.foreignid = myconfig.Foreignid
	device.tenantid = myconfig.Tenantid
	device.mobilenumber = safenum(feed.mobilenumber)
	if len(feed.iccid) > 0 {
		device.iccid = feed.iccid
	}
	device.esn = feed.esn
	return device
}

func randomDevice() (device Device) {
	device.deviceid = gocql.TimeUUID().String()
	device.devicebucket = calcBucket(device.deviceid)
	device.foreignid = myconfig.Foreignid
	device.tenantid = myconfig.Tenantid
	device.imei = randomNumber(20)
	device.imsi = uint64(rand.Uint32())
	device.mobilenumber = randomLong(10)
	device.esn = randomHexString(11)
	device.iccid = randomNumber(2)
	return device
}

func randomLong(n int) uint64 {
	s := "6" + randomNumber(n-1)
	return safenum(s)
}

func randomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(lettersLen)]
	}
	return string(b)
}

func randomHexString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = hexdigits[rand.Intn(16)]
	}
	return string(b)
}

func randomNumber(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = digits[rand.Intn(10)]
	}
	return string(b)
}

func elapsed(start time.Time) {
	log.Printf("=== elapsed time: %s", time.Since(start))
}

func truncateData() {
	log.Print("Truncating data")
	session.Query(`TRUNCATE deviceauxinfo`).Exec()
}

func readExtract(lineHandler ExtractRowHandler) {

	log.Print("Loading customer extract: ", myconfig.Filename)

	defer elapsed(time.Now())

	file, err := os.Open(myconfig.Filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// read text file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		lineHandler(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

//public int hashCode() {
//long hilo = mostSigBits ^ leastSigBits;
//return ((int)(hilo >> 32)) ^ (int) hilo;
//}

/**
Calculate hash code of UUID
copied from scala
todo: confirm same in java
*/
func hash(uuid []byte) int32 {
	var msb int64
	var lsb int64

	for i := 0; i < 8; i++ {
		msb = (msb << 8) | int64(uuid[i]&0xff)
	}

	for i := 8; i < 16; i++ {
		lsb = (lsb << 8) | int64(uuid[i]&0xff)
	}
	hilo := msb ^ lsb
	return (int32(hilo>>32) ^ int32(hilo))
}

/**
Calculate bucket from UUID
simply hash code (int) of uuid modulo 10000
*/
func calcBucket(uuid string) int32 {
	id, _ := gocql.ParseUUID(uuid)
	bucket := hash(id.Bytes()) % BUCKET_MODULO
	return bucket
}

func findMdn(mdn uint64) (string, int) {
	var id gocql.UUID
	var idbucket int
	//log.Print("lookup mdn: ", mdn)
	iter := findByMdn.Bind(mdn).Iter()
	//iter := session.Query(`SELECT id FROM device_by_mobilenumber`).Iter()
	//iter := session.Query(`SELECT id FROM device`).Iter()

	// make sure to close iterator
	defer iter.Close()

	for iter.Scan(&id, &idbucket) {
		return id.String(), idbucket
	}
	return "", 0
}

func findMobileNumber(id string, idbucket int) uint64 {
	var result uint64
	iter := findByDeviceId.Bind(id, idbucket).Iter()

	// make sure to close iterator
	defer iter.Close()

	for iter.Scan(&result) {
		return result
	}
	return 0
}

/*
	INSERT ROW
*/
func insertDevice(device *Device) error {
	//log.Print("insert device", device)
	err := insert.Bind(
		device.devicebucket, device.deviceid,
		device.foreignid, device.tenantid,
		device.mobilenumber, device.iccid, device.esn,
		time.Now(),
		myconfig.Ttl).Exec()
	if err != nil {
		log.Print(err)
	}
	return err
}

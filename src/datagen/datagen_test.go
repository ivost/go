package main

import (
	"github.com/smartystreets/goconvey/convey"
	"log"
	"os"
	"testing"
)

func init() {
	log.SetOutput(os.Stdout)
}

func TestDatagen(t *testing.T) {

	convey.Convey("Test uuid1", t, func() {
		id := "10000001-dead-1000-beef-000000000001"
		convey.So(calcBucket(id), convey.ShouldEqual, 64)
	})

	convey.Convey("Test uuid2", t, func() {
		id := "22a37bd9-ac47-11e6-bec9-f45c898a2447"
		convey.So(calcBucket(id), convey.ShouldEqual, -1372)
	})
}

// iconv -f UTF-16LE -t UTF-8  ~/Documents/SFTP/extract_VIP_NPHASE_201609251041.10k.txt > /tmp/extr10k.txt

//type cats struct {
//	Plato  string
//	Cauchy string
//}
//
//type simple struct {
//	Age     int
//	Colors  [][]string
//	Pi      float64
//	YesOrNo bool
//	Now     time.Time
//	Andrew  string
//	Kait    string
//	My      map[string]cats
//}

//func TestDecodeSimple(t *testing.T) {
//
//	var testSimple = `
//age = 250
//andrew = "gallant"
//kait = "brady"
//now = 1987-07-05T05:45:00Z
//yesOrNo = true
//pi = 3.14
//colors = [
//	["red", "green", "blue"],
//	["cyan", "magenta", "yellow", "black"],
//]
//
//[My.Cats]
//plato = "cat 1"
//cauchy = "cat 2"
//`
//
//
//	var val simple
//	_, err := toml.Decode(testSimple, &val)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	now, err := time.Parse("2006-01-02T15:04:05", "1987-07-05T05:45:00")
//	if err != nil {
//		panic(err)
//	}
//	var answer = simple{
//		Age:     250,
//		Andrew:  "gallant",
//		Kait:    "brady",
//		Now:     now,
//		YesOrNo: true,
//		Pi:      3.14,
//		Colors: [][]string{
//			{"red", "green", "blue"},
//			{"cyan", "magenta", "yellow", "black"},
//		},
//		My: map[string]cats{
//			"Cats": {Plato: "cat 1", Cauchy: "cat 2"},
//		},
//	}
//	if !reflect.DeepEqual(val, answer) {
//		t.Fatalf("Expected\n-----\n%#v\n-----\nbut got\n-----\n%#v\n",
//			answer, val)
//	}
//}

// https://github.com/gocql/gocql
/*
Some tips for getting more performance from the driver:

Use the TokenAware policy
Use many goroutines when doing inserts, the driver is asynchronous
but provides a synchronous API, it can execute many queries concurrently
Tune query page size
Reading data from the network to unmarshal will incur a large amount of allocations,
this can adversely affect the garbage collector, tune GOGC
Close iterators after use to recycle byte buffers
*/

/*
idbucket int,		-- how to calculate?
id uuid,				-- device uuid

foreignid uuid,		-- account id
tenantid uuid,		--?

mobilenumber text,	-- msisdn?
--    iccid bigint,	-- 22 digits
imei bigint,		-- 19 digits
imsi bigint,		-- 15 digits
meid text,		-- hex
esn bigint		-- 3g


extr10.txt

case class Feed(
    cost_center: String,    //  0
    ecpd: String,           //  1
    customer_id: String,    //  2
    bill_acnt_no: String,   //  3
    iccid: String,          //  5
    esn: String,            //  6
    mdn: String,            //  7
    ip_addr: String,        //  8
    price_plan_id: String,  //  9
    status: String,         // 13
    reason_code: String,    // 14
    act_date: String        // 15
    )
*/

//public int hashCode() {
//long hilo = mostSigBits ^ leastSigBits;
//return ((int)(hilo >> 32)) ^ (int) hilo;
//}

//id := "10000001-dead-1000-beef-000000000001"
//
//bucket := calcBucket(id)
//log.Print(bucket)
//if bucket != 7664 {
//	log.Fatal("hash error ", bucket)
//}
//
//id = "22a37bd9-ac47-11e6-bec9-f45c898a2447"
//bucket = calcBucket(id)
//log.Print(bucket)
//if bucket != -172 {
//	log.Fatal("hash error ", bucket)
//}
//}

/*
// ClusterConfig is a struct to configure the default cluster implementation
// of gocql. It has a variety of attributes that can be used to modify the
// behavior to fit the most common use cases. Applications that require a
// different setup must implement their own cluster.
type ClusterConfig struct {
	Hosts             []string          // addresses for the initial connections
	CQLVersion        string            // CQL version (default: 3.0.0)
	ProtoVersion      int               // version of the native protocol (default: 2)
	Timeout           time.Duration     // connection timeout (default: 600ms)
	Port              int               // port (default: 9042)
	Keyspace          string            // initial keyspace (optional)
	NumConns          int               // number of connections per host (default: 2)
	Consistency       Consistency       // default consistency level (default: Quorum)
	Compressor        Compressor        // compression algorithm (default: nil)
	Authenticator     Authenticator     // authenticator (default: nil)
	RetryPolicy       RetryPolicy       // Default retry policy to use for queries (default: 0)
	SocketKeepalive   time.Duration     // The keepalive period to use, enabled if > 0 (default: 0)
	MaxPreparedStmts  int               // Sets the maximum cache size for prepared statements globally for gocql (default: 1000)
	MaxRoutingKeyInfo int               // Sets the maximum cache size for query info about statements for each session (default: 1000)
	PageSize          int               // Default page size to use for created sessions (default: 5000)
	SerialConsistency SerialConsistency // Sets the consistency for the serial part of queries, values can be either SERIAL or LOCAL_SERIAL (default: unset)
	SslOpts           *SslOptions
	DefaultTimestamp  bool // Sends a client side timestamp for all requests which overrides the timestamp at which it arrives at the server. (default: true, only enabled for protocol 3 and above)
	// PoolConfig configures the underlying connection pool, allowing the
	// configuration of host selection and connection selection policies.
	PoolConfig PoolConfig

	// If not zero, gocql attempt to reconnect known DOWN nodes in every ReconnectSleep.
	ReconnectInterval time.Duration

	// The maximum amount of time to wait for schema agreement in a cluster after
	// receiving a schema change frame. (deault: 60s)
	MaxWaitSchemaAgreement time.Duration

	// HostFilter will filter all incoming events for host, any which don't pass
	// the filter will be ignored. If set will take precedence over any options set
	// via Discovery
	HostFilter HostFilter

	// If IgnorePeerAddr is true and the address in system.peers does not match
	// the supplied host by either initial hosts or discovered via events then the
	// host will be replaced with the supplied address.
	//
	// For example if an event comes in with host=10.0.0.1 but when looking up that
	// address in system.local or system.peers returns 127.0.0.1, the peer will be
	// set to 10.0.0.1 which is what will be used to connect to.
	IgnorePeerAddr bool

	// If DisableInitialHostLookup then the driver will not attempt to get host info
	// from the system.peers table, this will mean that the driver will connect to
	// hosts supplied and will not attempt to lookup the hosts information, this will
	// mean that data_centre, rack and token information will not be available and as
	// such host filtering and token aware query routing will not be available.
	DisableInitialHostLookup bool

	// Configure events the driver will register for
	Events struct {
		// disable registering for status events (node up/down)
		DisableNodeStatusEvents bool
		// disable registering for topology events (node added/removed/moved)
		DisableTopologyEvents bool
		// disable registering for schema events (keyspace/table/function removed/created/updated)
		DisableSchemaEvents bool
	}

	// DisableSkipMetadata will override the internal result metadata cache so that the driver does not
	// send skip_metadata for queries, this means that the result will always contain
	// the metadata to parse the rows and will not reuse the metadata from the prepared
	// statement.
	//
	// See https://issues.apache.org/jira/browse/CASSANDRA-10786
	DisableSkipMetadata bool

	// internal config for testing
	disableControlConn bool
}


// Device (Global or Account Owned)
type Device struct {

	// Detail
	ProviderId string `json:"providerid,omitempty" db:"data.device.providerid" validate:"omitempty,uuid,id=ts.provider"`
	State      string `json:"state,omitempty" db:"..state" validate:"required"`
	QRCode     string `json:"qrcode,omitempty" db:"..qrcode"`
	Esn        uint64 `json:"esn,omitempty" db:"..esn"`
	Serial     string `json:"serial,omitempty" db:"..serial"`
	Mac        string `json:"mac,omitempty" db:"..mac"`
	Imei       uint64 `json:"imei,omitempty" db:"..imei"`
	Imsi       uint64 `json:"imsi,omitempty" db:"..imsi"`
	MsIsdn     string `json:"msisdn,omitempty" db:"..msisdn"`
	MeId       string `json:"meid,omitempty" db:"..meid"`
	IccId      string `json:"iccid,omitempty" db:"..iccid"`
	RefId      string `json:"refid,omitempty" db:"..refid"`

	// Foreign Identity
	rest.Foreign

	// Identity
	rest.Identity
}


// resource foreign identity
type Foreign struct {
	ForeignId string   `json:"foreignid,omitempty" db:"..foreignid,key" immutable:"-" validate:"uuid,required"`
	TagIds    []string `json:"tagids,omitempty" db:"..tags" validate:"dive,uuid,omitempty"`
}

// resource identity
type Identity struct {
	Id          string    `json:"id,omitempty" db:"..id,key" immutable:"-" validate:"uuid"`
	Kind        string    `json:"kind,omitempty" db:"..kind" validate:"kind"`
	Version     string    `json:"version,omitempty" db:"..version" validate:"required"`
	VersionId   string    `json:"versionid,omitempty" db:"..versionid" validate:"required"`
	CreatedOn   time.Time `json:"createdon,omitempty" db:"..createdon,key" immutable:"-" validate:"required"`
	LastUpdated time.Time `json:"lastupdated,omitempty" db:"..lastupdated" immutable:"-" validate:"required"`
	Name        string    `json:"name,omitempty" db:"..name"`
	Description string    `json:"description,omitempty" db:"..description"`
	// TODO
	// IsDeleted bool `json:"isdeleted,omitempty" db:"..isdeleted"`
}

extr10.txt

case class Feed(
    cost_center: String,    //  0
    ecpd: String,           //  1
    customer_id: String,    //  2
    bill_acnt_no: String,   //  3
    iccid: String,          //  5
    esn: String,            //  6
    mdn: String,            //  7
    ip_addr: String,        //  8
    price_plan_id: String,  //  9
    status: String,         // 13
    reason_code: String,    // 14
    act_date: String        // 15
    )


*/

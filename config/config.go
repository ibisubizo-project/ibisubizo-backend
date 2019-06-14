package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/globalsign/mgo"
)

const (
	DATABASE               = "ibusizo"
	USERCOLLECTION         = "users"
	COMMENTSCOLLECTION     = "comments"
	PROBLEMSCOLLECTION     = "problems"
	COMMENTLIKESCOLLECTION = "commentlikes"
	PROBLEMLIKESCOLLECTION = "problemlikes"
	METRICSCOLLECTION      = "metrics"
	LIMITS                 = 6
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
	config     Config
)

type Config struct {
	AppName  string `json:"app_name"`
	MongoURL string `json:"mongo_url"`
	Session  *mgo.Session
	JWTKey   string `json:"jwt_secret"`
	Port     string `json:"port"`
}

//Init Initializes the configuration struct
func Init() {
	log.Println("[INIT]")
	config = Config{}
	port := os.Getenv("IBUSIBUZO_PORT")
	if len(port) == 0 {
		log.Println("len(port) != 0")
		config.Port = "8000"
	} else {
		log.Println("[len(port) != 0] else")
		fmt.Println(port)
		config.Port = port
	}

	databaseURL := os.Getenv("IBUSIBUZO_DATABASE_URL")
	if len(databaseURL) == 0 {
		log.Println("len(databaseURL) != 0")
		config.MongoURL = "127.0.0.1:27017"
	} else {
		log.Println("[len(databaseURL) != 0] else")
		config.MongoURL = databaseURL
	}

	session, err := mgo.Dial(config.MongoURL)
	// session, err := mgo.DialWithInfo(&mgo.DialInfo{
	// 	Addrs:    config.MongoURL,
	// 	Username: "",
	// 	Password: "",
	// })

	if err != nil {
		log.Println(err)
		log.Panic("Error connecting to database")
	}
	config.Session = session

}

//Get Returns a config file
func Get() *Config {
	return &config
}

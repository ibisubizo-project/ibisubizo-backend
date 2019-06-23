package config

import (
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
	DEFAULT_PORT           = "8000"
	DEFAULT_DB_URL         = "127.0.0.1:27017"
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
	config = Config{}
	port := os.Getenv("IBUSIBUZO_PORT")
	if len(port) == 0 {
		config.Port = DEFAULT_PORT
	} else {
		config.Port = port
	}

	databaseURL := os.Getenv("IBUSIBUZO_DATABASE_URL")
	if len(databaseURL) == 0 {
		config.MongoURL = DEFAULT_DB_URL
	} else {
		config.MongoURL = databaseURL
	}

	session, err := mgo.Dial(config.MongoURL)
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

package config

import (
	"log"
	"path/filepath"
	"runtime"

	"github.com/globalsign/mgo"
)

const (
	DATABASE           = "problemApp"
	USERCOLLECTION     = "users"
	COMMENTSCOLLECTION = "comments"
	PROBLEMSCOLLECTION = "problems"
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
	//Read in configuration from a file
	// configFile, err := ioutil.ReadFile("../config.json")
	// if err != nil {
	// 	log.Println(err)
	// 	log.Panic("There was an error loading the config file")
	// }

	// if err := json.Unmarshal(configFile, config); err != nil {
	// 	log.Panic("Error decoding into struct. Please check the config file")
	// }

	config.Port = "8000"

	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{"127.0.0.1:27017", "127.0.0.1:27018"},
		Username: "",
		Password: "",
	})
	if err != nil {
		log.Panic("Error connecting to database")
	}
	config.Session = session

}

//Get Returns a config file
func Get() *Config {
	return &config
}

package main

import (
    "encoding/json"
    "io/ioutil"
    "strconv"
    "time"
)

var Config map[string]interface{}
var persistence bool
var port string
var gcInterval time.Duration
var diskWriteInterval time.Duration

// read the config file, if there's any
func readConfig() {
    Config = make(map[string]interface{})
    file, e := ioutil.ReadFile("/etc/mumbo-conf.json")

    if e != nil {
        extendConfig(Config)
    } else {
        json.Unmarshal(file, &Config)

        // extend config with defaults
        extendConfig(Config)
    }

    // initialize the system
    initializer()
}

// will initialize the system with/ without persistence
func initializer() {

    if persistence {
        // initialize disk storage
        initializeDiskStorage()
    }


    // initialize the basic structure, load the persistence values etc
    initializeStore()

    // intialize random garbage collection
    go collectionGarbageCycle()
}

// will extend the config will defaults , if props not specified
func extendConfig(conf map[string]interface{}) {
    _, ok := conf["persistence"]

    if ok {
        persistence = Config["persistence"].(bool)
    } else {
        persistence = false
    }

    _, ok = conf["port"]

    if ok {
        Fport := Config["port"].(float64)
        Iport := int(Fport)
        port = ":" + strconv.Itoa(int(Iport))
    } else {
        port = ":" + "2700"
    }

    _, ok = conf["gcInterval"]

    if ok {
        gcInterval = time.Duration(int(Config["gcInterval"].(float64)))
    } else {
        gcInterval = 100
    }

    _, ok = conf["diskWriteInterval"]

    if ok {
        diskWriteInterval = time.Duration(int(Config["gcInterval"].(float64)))
    } else {
        diskWriteInterval = 100
    }
}
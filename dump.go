package main

import (
    "log"
    "os"
    "path/filepath"
    "github.com/syndtr/goleveldb/leveldb"
    "encoding/json"
    "time"
    "fmt"
)

var disk *leveldb.DB
var batch *leveldb.Batch
var written bool

// Initialize the disk storage
func initializeDiskStorage() {
    path := cwd()
    disk, _ = leveldb.OpenFile(path, nil)
    batch = new(leveldb.Batch)
    written = true
    go flushingActivity()
}

// This will dump the keys to the disk
func flush() {
    if !written  {
        fmt.Println("Dumping to disk")
        disk.Write(batch, nil)
    }
}

// will load the data from the disk on start up
func iterate() {
    fmt.Println("Server is loading your data")
    fmt.Println("Please be patient")
    iter := disk.NewIterator(nil, nil)
    for iter.Next() {
        k := iter.Key()
        v := iter.Value()
        key := string(k)
        err, value := decodeBytes(v)

        if !err {
            data[key] = value
        }
    }
    iter.Release()
    _ = iter.Error()
    fmt.Println("Server has been initialized. We're ready to go!")
}

// puts data on the batch
func diskPut(key string, value interface{}) {
    err, bytes := GetBytes(value)
    if !err {
        batch.Put([]byte(key), bytes)
        written = false
    }
}

//delete from batch
func diskDel(key string) {
    batch.Delete([]byte(key))
    written = false
}

// get current working directory
func cwd() string {
    dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
    if err != nil {
            log.Fatal(err)
    }
    os.Mkdir(dir + "/data", os.FileMode(0777))
    return dir + "/data"
}

// converts json to bytes
func GetBytes(value interface{}) (bool, []byte) {
    b, err := json.Marshal(value)

    if err == nil {
        return false, b
    }

    return true, nil
}

// decodes the bytes to json
func decodeBytes(value []byte) (bool, values) {
    var v values
    err := json.Unmarshal(value, &v)
    if err == nil {
        return false, v
    }
    return true, values{}
}

// will flush to disk every 5 minutes(configurable value)
func flushingActivity() {
    for _ = range time.Tick(300000*time.Millisecond) {
        flush()
        written = true
    }
}
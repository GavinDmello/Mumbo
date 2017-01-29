/*
 * Mumbo - A fast  in-memory key value store
 * Copyright(c) 2016-present @GavinDmello
 * BSD Licensed
 */

package main

import (
    "sync"
    "time"
)

var mutex = &sync.Mutex{}
var data map[interface{}]interface{}

type keys struct {
    Key interface{} `json:"key"`
    Value interface{} `json:"value"`
    Exists bool `json:"exists"`
}

type values struct {
    Value interface{} `json:"value"`
    Locked bool `json:"locked"`
    Ttl  interface{}  `json:"ttl"`
}

// initialize in memory store
func initializeStore() {
    data = make(map[interface{}]interface{})
}

// set value for specified key
func setVal(key interface{}, value interface{}, ttl interface{}) interface{} {
    data[key] = values{
        Value : value,
        Locked : false,
        Ttl : ttl,
    }
    return data[key]
}

// get value for a specfied key
func getVal(key interface{}) (bool, interface{}){
    value, ok := data[key]

    if !ok {
        return true, nil
    } else {
        return false, value
    }
}

// deletes a value from the store
func delVal(key interface{}) {
    delete(data, key)
}

// will append an item to the list
func listPush(key interface{}, item interface{}) (bool, interface{}) {
    res, ok := data[key]

    if !ok {
        return true, "Item not found"
    }

    castResult, _ := res.(values)

    if value, ok := castResult.Value.([]interface{}); ok {
        mutex.Lock()
        value = append(value, item)
        castResult.Value = value
        data[key] =  castResult
        mutex.Unlock()
        return false, value
    } else {
        return true, "Item is not of type list"
    }
}

// will remove an item from the list by value
func listRemove(key interface{}, item interface{}) (bool, interface{}) {
    index := -1
    res, ok := data[key]

    if !ok {
        return true, "Item not found"
    }

    castResult, _ := res.(values)

    if value, ok := castResult.Value.([]interface{}); ok {

        mutex.Lock()
        for k, val := range value {
            if (val == item) {
                index = k
                break
            }
        }

        if index >= 0 {
            value = append(value[:index], value[index+1:]...)
            castResult.Value = value
            data[key] = castResult
        }

        mutex.Unlock()

        return false, value
    } else {
        return true, "Item is not of type list"
    }

}

// gets a list a values in a single call
func batchGet(keylist interface{}) []keys {

    list := keylist.([]interface{})
    pairs := make([]keys, 0) // array & void values


    for _, key := range list {
        res, ok := data[key]
        if ok {
            castResult, _ := res.(values)
            pairs = append(pairs, keys{
                Key : key,
                Value: castResult.Value,
                Exists : true,
            })
        } else {
            pairs = append(pairs, keys{
                Key : key,
                Exists : false,
            })
        }
    }

    return pairs
}

// returns the timestamp in milliseconds
func returnTimestamp() int64 {
    return time.Now().UnixNano() / int64(time.Millisecond)
}

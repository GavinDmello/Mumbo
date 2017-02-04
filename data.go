/*
 * Mumbo - A fast  in-memory key value store
 * Copyright(c) 2016-present @GavinDmello
 * BSD Licensed
 */

package main

import (
    "sync"
    "time"
    "math/rand"
)

var mutex = &sync.Mutex{}
var data map[interface{}]interface{}

// structure for batch keys
type keys struct {
    Key interface{} `json:"key"`
    Value interface{} `json:"value"`
    Exists bool `json:"exists"`
}

// structure for values
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
func setVal(key, value interface{}, ttl interface{}) interface{} {
    var castTtl, newTtlValue int64


    if ttl != nil {
        castTtl = int64(ttl.(float64))
        newTtlValue = returnTimestamp() + castTtl
    } else {
        newTtlValue = -1 // infinite ttl
    }

    data[key] = values{
        Value : value,
        Locked : false,
        Ttl : newTtlValue,
    }

    return data[key]
}

// get value for a specfied key
func getVal(key interface{}) (bool, interface{}){

    v, ok := data[key]

    if !ok {
        return true, nil
    }

    value := v.(values)
    // expired := deleteIfExpired(key, value)

    // if expired {
    //     return expired, nil
    // }

    return false, value

}

// deletes a value from the store
func delVal(key interface{}) {
    delete(data, key)
}

// will append an item to the list
func listPush(key, item interface{}) (bool, interface{}) {
    res, ok := data[key]


    if !ok {
        return true, "Item not found"
    }

    castResult, _ := res.(values)
    expired := deleteIfExpired(key, castResult)

    if expired {
        return true, "Item not found"
    }

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
func listRemove(key, item interface{}) (bool, interface{}) {
    index := -1
    res, ok := data[key]


    if !ok {
        return true, "Item not found"
    }

    castResult, _ := res.(values)
    expired := deleteIfExpired(key, castResult)

    if expired {
        return true, "Item not found"
    }

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
            // lazy deletion of keys
            expired := deleteIfExpired(key, castResult)

            if  !expired {
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

// checks if the key has expired
func deleteIfExpired(key interface{}, value values) bool {

    ttl := value.Ttl.(int64)
    currentTimestamp := returnTimestamp()

    if ttl > -1 && currentTimestamp > ttl {
        delVal(key)
        return true
    }

    return false
}

// get random keys
func random(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}

// check for dead keys and delete
func collectionGarbageCycle() {

    // todo, make this configurable
    randomIndex := 0
    currentIndex := 0
    for _ = range time.Tick(20*time.Millisecond) {
        if len(data) == 0 {
            continue
        }
        randomIndex = random(1, len(data))

        currentIndex = 0

        //mutex.RLock()
        for k, v := range data {
            currentIndex++
           if currentIndex % randomIndex == 0 {
                value := v.(values)
                deleteIfExpired(k, value)
            }
        }
        //mutex.RUnlock()
    }
}
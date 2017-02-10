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

var mutex = &sync.RWMutex{}
var data map[string]interface{}
var totalKeys int
var ttlKeys []string

// structure for batch keys
type keys struct {
    Key string `json:"key"`
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
    data = make(map[string]interface{})
    totalKeys = 0
    ttlKeys = make([]string, 0)
    iterate()
}

// set value for specified key
func setVal(key string, value interface{}, ttl interface{}) interface{} {
    var castTtl, newTtlValue int64


    if ttl != nil {
        castTtl = int64(ttl.(float64))
        newTtlValue = returnTimestamp() + castTtl
        ttlKeys = append(ttlKeys, key)
        totalKeys++
    } else {
        newTtlValue = -1 // infinite ttl
    }

    mutex.Lock()
    data[key] = values{
        Value : value,
        Locked : false,
        Ttl : newTtlValue,
    }
    mutex.Unlock()

    diskPut(key, values{
        Value : value,
        Locked : false,
        Ttl : newTtlValue,
    })

    return data[key]
}

// get value for a specfied key
func getVal(key string) (bool, interface{}){

    mutex.RLock()
    v, ok := data[key]
    mutex.RUnlock()

    if !ok {
        return true, nil
    }

    value := v.(values)
    expired := deleteIfExpired(key, value)

    if expired {
        return expired, nil
    }

    return false, value

}

// deletes a value from the store
func delVal(key string) {
    mutex.Lock()
    delete(data, key)
    go deleteTTLKeys(key)
    mutex.Unlock()

    // delete from disk
    diskDel(key)
}

// will append an item to the list
func listPush(key string, item interface{}) (bool, interface{}) {
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
func listRemove(key string, item interface{}) (bool, interface{}) {
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
    var stringKey string

    for _, key := range list {
        mutex.RLock()
        stringKey = key.(string)
        res, ok := data[stringKey]
        mutex.RUnlock()

        if ok {
            castResult, _ := res.(values)
            // lazy deletion of keys
            expired := deleteIfExpired(stringKey, castResult)

            if  !expired {
                pairs = append(pairs, keys{
                    Key : stringKey,
                    Value: castResult.Value,
                    Exists : true,
                })
            } else {
                pairs = append(pairs, keys{
                    Key : stringKey,
                    Exists : false,
                })
            }

        } else {
            pairs = append(pairs, keys{
                Key : stringKey,
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
func deleteIfExpired(key string, value values) bool {

    ttl, ok := value.Ttl.(int64)

    if !ok {
        floatTtl , _ := value.Ttl.(float64)
        ttl = int64(floatTtl)
    }

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

// delete keys from TTL lists
func deleteTTLKeys(key string) {
    mutex.RLock()
    index := -1
    for k, val := range ttlKeys {
        if (val == key) {
            index = k
            break
        }
    }

    if index >= 0 {
        ttlKeys = append(ttlKeys[:index], ttlKeys[index+1:]...)
        totalKeys--
    }

    mutex.RUnlock()
}

// check for dead keys and delete
func collectionGarbageCycle() {
    randomIndex := 0
    key := ""
    var value values
    // todo, make this configurable
    for _ = range time.Tick(100*time.Millisecond) {
        if totalKeys > 0 {
            randomIndex = random(0, totalKeys)

            if randomIndex < totalKeys {
                mutex.RLock()
                key = ttlKeys[randomIndex]
                v, ok := data[key]
                mutex.RUnlock()

                if ok {
                    value = v.(values)
                    deleteIfExpired(key, value)
                }
            }
        }
    }
}
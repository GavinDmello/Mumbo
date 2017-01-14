package main

var data map[interface{}]interface{}

type keys struct {
    Key interface{} `json:"key"`
    Value interface{} `json:"value"`
    Exists bool `json:"exists"`
}

// initialize in memory store
func initializeStore() {
    data = make(map[interface{}]interface{})
}

// set value for specified key
func setVal(key interface{}, value interface{}) interface{} {
    data[key] = value
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

// gets a list a values in a single call
func batchGet(keylist interface{}) []keys {

    list := keylist.([]interface{})
    pairs := make([]keys, 0)


    for _, key := range list {
        value, ok := data[key]
        if ok {
            pairs = append(pairs, keys{
                Key : key,
                Value: value,
                Exists : true,
            })
        } else {
            pairs = append(pairs, keys{
                Key : key,
                Exists : false,
                Value : nil,
            })
        }
    }

    return pairs
}
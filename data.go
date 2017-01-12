package main

var data map[interface{}]interface{}

func initializeStore() {
    data = make(map[interface{}]interface{})
}


func setVal(key interface{}, value interface{}) interface{} {
    data[key] = value
    return data[key]
}

func getVal(key interface{}) (bool, interface{}){
    value, ok := data[key]

    if !ok {
        return true, nil
    } else {
        return false, value
    }
}
package main

func setValidation(data map[string]interface{}) bool {
    _, ok := data["cmd"]

    if !ok {
        return false
    }

    _, ok = data["key"]

    if !ok {
        return false
    }

    _, ok = data["value"]

    if !ok {
        return false
    }

    return true

}

func getValidation(data map[string]interface{}) bool {
    _, ok := data["cmd"]

    if !ok {
        return false
    }

    _, ok = data["key"]

    if !ok {
        return false
    }

    return true

}


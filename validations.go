package main

// validation for all set commands
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

// validation for all get commands
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


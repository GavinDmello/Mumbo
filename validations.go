/*
 * Mumbo - A fast  in-memory key value store
 * Copyright(c) 2016-present @GavinDmello
 * BSD Licensed
 */

package main

// validation for all set commands
func setValidation(data map[string]interface{}) bool {
    _, ok := data["key"]

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
    _, ok := data["key"]

    if !ok {
        return false
    }

    return true
}

// validation for all batchGet commands
func batchGetValidation(data map[string]interface{}) bool {
    _, ok := data["keylist"]

    if !ok {
        return false
    }

    return true
}


// validation for all batchGet commands
func listOpValidation(data map[string]interface{}) bool {
    _, ok := data["key"]

    if !ok {
        return false
    }

    _, ok = data["item"]

    if !ok {
        return false
    }

    return true
}



/*
 * Mumbo - A fast  in-memory key value store
 * Copyright(c) 2016-present @GavinDmello
 * BSD Licensed
 */

package main

import (
    "fmt"
)

// check command
func check(data map[string]interface{}) map[string]interface{} {
    responseMap := make(map[string]interface{})
    responseMap["cmd"] = data["cmd"]

    switch data["cmd"] {

        case "set" :
            isValid := setValidation(data)

            if !isValid {
                formValidationErrorResp(responseMap)
                return responseMap
            }

            key := data["key"]
            value := data["value"]
            ttl := data["ttl"]
            result := setVal(key, value, ttl)
            responseMap["value"] = result
            formSuccessResponse(responseMap)

            return responseMap

        case "exist":
            fallthrough

        case "get" :
            isValid := getValidation(data)

            if !isValid {
                formValidationErrorResp(responseMap)
                return responseMap
            }

            key := data["key"]
            err, result := getVal(key)

            if err {
                missingKey(responseMap)
            } else {
                responseMap["value"] = result
                formSuccessResponse(responseMap)
            }

            return responseMap

        case "del" :
            isValid := getValidation(data)

            if !isValid {
                formValidationErrorResp(responseMap)
                return responseMap
            }

            key := data["key"]
            delVal(key)
            formSuccessResponse(responseMap)

            return responseMap

        case "batchget" :
            isValid := batchGetValidation(data)

            if !isValid {
                formValidationErrorResp(responseMap)
                return responseMap
            }

            keyList := data["keylist"]
            result := batchGet(keyList)
            responseMap["value"] = result
            formSuccessResponse(responseMap)

            return responseMap

        case "listpush" :
            isValid := listOpValidation(data)

            if !isValid {
                formValidationErrorResp(responseMap)
                return responseMap
            }

            item := data["item"]
            key  := data["key"]
            err, result := listPush(key, item)

            if err {
                listError(responseMap, result)
                return responseMap
            }

            responseMap["value"] = result
            formSuccessResponse(responseMap)

            return responseMap

        case "listremove" :
            isValid := listOpValidation(data)

            if !isValid {
                formValidationErrorResp(responseMap)
                return responseMap
            }

            item := data["item"]
            key  := data["key"]

            err, result := listRemove(key, item)

            if err {
                listError(responseMap, result)
                return responseMap
            }

            responseMap["value"] = result
            formSuccessResponse(responseMap)

            return responseMap

        default :
            fmt.Println("invalid command")
            formValidationErrorResp(responseMap)
            return responseMap
    }

}

// missing key packet
func missingKey(errorResponse map[string]interface{}) {
    errorResponse["status"] = 401
    errorResponse["message"] = "key does not exist"
}

// success packet
func formSuccessResponse(success map[string]interface{}) {
    success["status"] = 200
    success["message"] = "success"
}


// validation error packet
func formValidationErrorResp(errorResponse map[string]interface{}) {
    errorResponse["status"] = 404
    errorResponse["message"] = "validation error"
}

// returns list errors
func listError(errorResponse map[string]interface{}, message interface{}) {
    errorResponse["status"] = 402
    errorResponse["message"] = message
}

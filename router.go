package main

import (
    "fmt"
)

func check(data map[string]interface{}) map[string]interface{} {
    returnMap := make(map[string]interface{})

    switch data["cmd"] {

        case "set" :
            isValid := setValidation(data)

            if !isValid {
                formValidationErrorResp(returnMap)
            } else {
                key := data["key"]
                value := data["value"]
                result := setVal(key, value)
                returnMap["value"] = result
                formSuccessResponse(returnMap)
            }
            return returnMap

        case "get" :
            fmt.Println("get")
            isValid := getValidation(data)

            if !isValid {
                formValidationErrorResp(returnMap)
            } else {
                key := data["key"]
                err, result := getVal(key)
                if err {
                    missingKey(returnMap)
                } else {
                    returnMap["value"] = result
                    formSuccessResponse(returnMap)
                }

            }
            return returnMap


        default :
            fmt.Println("invalid command")
            formValidationErrorResp(returnMap)

    }
    return returnMap
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

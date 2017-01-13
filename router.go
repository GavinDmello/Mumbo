package main

import (
    "fmt"
)

func check(data map[string]interface{}) map[string]interface{} {
    responseMap := make(map[string]interface{})

    switch data["cmd"] {

        case "set" :
            isValid := setValidation(data)

            if !isValid {
                formValidationErrorResp(responseMap)
            } else {
                key := data["key"]
                value := data["value"]
                result := setVal(key, value)
                responseMap["value"] = result
                formSuccessResponse(responseMap)
            }
            return responseMap

        case "get" :
        case "exist":
            isValid := getValidation(data)

            if !isValid {
                formValidationErrorResp(responseMap)
                return responseMap
            } else {
                key := data["key"]
                err, result := getVal(key)
                if err {
                    missingKey(responseMap)
                } else {
                    responseMap["value"] = result
                    formSuccessResponse(responseMap)
                }

            }
            return responseMap

        default :
            fmt.Println("invalid command")
            formValidationErrorResp(responseMap)

    }
    return responseMap
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

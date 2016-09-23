package utils

import "log"

func LogError(o interface{}, err error) interface{} {
        if err != nil {
                log.Printf("Got error [%v]", err)
        }
        return o
}


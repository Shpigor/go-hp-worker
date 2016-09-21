package model

import (
        "encoding/json"
        "encoding/xml"
        "fmt"
)

type Instance struct {
        Id   uint32 `json:"id" xml:"id"`
        Name string `json:"name" xml:"name"`
}

func (instance *Instance) ToJsonString() (string, error) {
        var js string
        bytes, err := json.Marshal(instance)
        if err != nil {
                fmt.Println(err)
        } else {
                js = string(bytes)
        }
        return js, err
}

func (instance *Instance) ToXML() (xmlString string, err error) {
        bytes, err := xml.Marshal(instance)
        if err != nil {
                fmt.Println(err)
        } else {
                xmlString = string(bytes)
        }
        return
}


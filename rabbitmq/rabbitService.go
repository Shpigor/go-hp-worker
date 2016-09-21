package rabbitmq

import (
        "log"
        "github.com/michaelklishin/rabbit-hole"
)

var client rabbithole.Client

func Init() {
        log.Println("Init rabbit service.")
        client, err := rabbithole.NewClient("", "", "")
        if err != nil {

        }
        client.DeclareBinding(nil, nil)
}
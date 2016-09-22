package main

import (
        "log"
        "os"
        "./config"
)

func main() {
        if len(os.Args) < 2 {
                log.Fatalln("incorrent parameters")
        }
        log.Println("Start worker...")
        config.Load(os.Args[1])
}
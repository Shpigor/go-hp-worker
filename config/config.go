package config

import (
        "log"
        "encoding/json"
        "os"
)

type RabbitOpenStackConfig struct {
        Host            string  `json:"host"`
        Port            int16   `json:"port"`
        Username        string  `json:"username"`
        Password        string  `json:"password"`
        NovaQueuePrefix string  `json:"novaQueuePrefix"`
        NovaTopic       string  `json:"novaTopic"`
}

type WorkerConfig struct {
        VnfInfoUrl                 string       `json:"vnfInfoUrl"`
        OperationTimeout           int32        `json:"operationTimeout"`
        Hostname                   string       `json:"hostname"`
        FwConfigPollInerwal        int32        `json:"fwConfigPollInerwal"`
        AccessIpCacheRefreshPeriod int32        `json:"accessIpCacheRefreshPeriod"`
        RabbitOpenStackConfig      RabbitOpenStackConfig `json:"rabbitOpenStackConfig"`
}

func Load(path string) (config *WorkerConfig) {
        file, openErr := os.Open(path)
        if (openErr != nil) {
                log.Fatalf("Can't read configuration file: %s [%v]", path, openErr)
        }
        decoder := json.NewDecoder(file)
        decodeErr := decoder.Decode(&config)
        if decodeErr != nil {
                log.Fatalf("Can't decode configuration from file [%v]", decodeErr)
        }
        log.Println(config)
        return
}
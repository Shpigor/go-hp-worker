package main

import (
        "log"
        "github.com/rgbkrk/libvirt-go"
)

func main() {

        vc, err := libvirt.NewVirConnectionReadOnly("qemu+ssh://dev@192.168.1.211/system")
        if err != nil {
                log.Printf("Got error during connect to libvrt. [%v]", err)
        }
        version, err := vc.GetLibVersion()
        if err != nil {
                log.Printf("Failed to retrieve version: %v", err)
        }

        log.Printf("Libvirt version is %d\n", version)
        domains, err := vc.ListAllDomains(0);
        if err != nil {
                log.Printf("Failed to retrieve domains: %v", err)
        }
        callback := libvirt.DomainEventCallback(func(c *libvirt.VirConnection, d *libvirt.VirDomain, event interface{}, f func()) int {
                log.Println("Inside callback")
                return 0
        })
        for _, domain := range domains {
                vc.DomainEventRegister(domain, libvirt.VIR_DOMAIN_EVENT_ID_LIFECYCLE, &callback, func() {})
                vdi, _ := domain.GetInfo()
                log.Printf("Domain name: %s \t uuid: %s\n", logError(domain.GetName()), logError(domain.GetUUIDString()))
                log.Printf("CPU time: %d\n \t Memory: %d", vdi.GetCpuTime(), vdi.GetMemory())
        }
}

func logError(o interface{}, err error) interface{} {
        if err != nil {
                log.Printf("Got error [%v]", err)
        }
        return o
}



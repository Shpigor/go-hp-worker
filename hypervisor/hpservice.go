package main

import (
        "net"
        "time"
        "log"
        "github.com/digitalocean/go-libvirt"
        "sync"
)

func connectToLibvirt() *libvirt.Libvirt {
        c, err := net.DialTimeout("tcp", "192.168.1.211:16509", 2 * time.Second)
        //c, err := net.DialTimeout("unix", "/var/run/libvirt/libvirt-sock", 2 * time.Second)
        if err != nil {
                log.Fatalf("Failed to dial libvirt: [%v]", err)
        }

        l := libvirt.New(c)
        if err := l.Connect(); err != nil {
                log.Fatalf("Failed to connect: [%v]", err)
        }
        return l
}

func main() {
        var wg sync.WaitGroup
        l := connectToLibvirt()
        domains, err := l.Domains()
        if err != nil {
                log.Fatalf("Failed to retrieve domains: %v", err)
        }

        log.Println("ID\tName\t\tUUID")
        log.Printf("--------------------------------------------------------\n")
        for _, d := range domains {
                log.Printf("%d\t%s\t%x\n", d.ID, d.Name, d.UUID)
                //wg.Add(1)
                func() {
                        for {
                                dmEvent, err := l.Events(d.Name)
                                if (err != nil) {
                                        log.Printf("Failed to read from events channel: %v", err)
                                }
                                log.Println(<-dmEvent)
                        }
                }()
        }
        wg.Wait()
}

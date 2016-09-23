package main

import (
        "log"
        //"../utils"
        "github.com/rgbkrk/libvirt-go"
        "time"
        "sync"
)

type ActiveDomainInfo struct {
        IsInitialized bool
        libVrtUrl     string
        VrtConn       *libvirt.VirConnection
        ActiveDomains map[string]*libvirt.VirDomain
}

func (activeDomainInfo *ActiveDomainInfo) init() {
        vc, err := libvirt.NewVirConnectionReadOnly(activeDomainInfo.libVrtUrl)
        if err != nil {
                log.Printf("Failed to connect to libvrt. [%v]", err)
        }
        version, err := vc.GetLibVersion()
        if err != nil {
                log.Printf("Failed to retrieve version: %v", err)
        }
        log.Printf("Libvirt version is %d\n", version)
        activeDomainInfo.VrtConn = &vc
        activeDomainInfo.IsInitialized = true
}

func (activeDomainInfo *ActiveDomainInfo) StartCollectDomainsData(periodSeconds time.Duration, wg *sync.WaitGroup) chan bool {
        ticker := time.NewTicker(periodSeconds * time.Second)
        quit := make(chan bool)
        go func() {
                for {
                        select {
                        case <-ticker.C:
                                activeDomainInfo.fetchDomains()
                        case <-quit:
                                ticker.Stop()
                                wg.Done()
                                return
                        }
                }
        }()
        return quit
}

func (activeDomainInfo *ActiveDomainInfo) fetchDomains() {
        vc := activeDomainInfo.VrtConn;
        domains, err := vc.ListAllDomains(0);
        if err != nil {
                log.Printf("Failed to retrieve domains: %v", err)
        }
        for _, domain := range domains {
                uuid, _ := domain.GetUUIDString()
                if activeDomainInfo.ActiveDomains[uuid] == nil {
                        callback := libvirt.DomainEventCallback(func(c *libvirt.VirConnection, d *libvirt.VirDomain, event interface{}, f func()) int {
                                log.Println("Inside callback.")
                                return 0
                        })
                        activeDomainInfo.ActiveDomains[uuid] = &domain
                        result := vc.DomainEventRegister(domain, libvirt.VIR_DOMAIN_EVENT_ID_CONTROL_ERROR, &callback, func() {
                                log.Printf("Inside event registration function.")
                        })
                        log.Printf("Result: %d", result)
                }
                //vdi, _ := domain.GetInfo()
                //log.Printf("Domain uuid: %s \t cpu: %d\t memory: %d\n\n", utils.LogError(domain.GetUUIDString()), vdi.GetCpuTime(), vdi.GetMemory())
        }
}

// Create new ActiveDomainInfo structure and initialize it
// create empty map
func NewActiveDomainInfo(libVrtUrl string) *ActiveDomainInfo {
        activeDomainInfo := ActiveDomainInfo{ActiveDomains:make(map[string]*libvirt.VirDomain), libVrtUrl:libVrtUrl}
        activeDomainInfo.init()
        return &activeDomainInfo
}

func main() {
        wg := sync.WaitGroup{}
        activeDomainInfo := NewActiveDomainInfo("qemu+ssh://dev@192.168.1.211/system")
        wg.Add(1)
        control := activeDomainInfo.StartCollectDomainsData(time.Duration(5), &wg)
        time.Sleep(time.Duration(30 * time.Minute))
        control <- true
        wg.Wait()
}
package main

import (
    "github.com/HikoQiu/go-eureka-client/eureka"
)

func main() {

    config := eureka.GetDefaultEurekaClientConfig()
    config.UseDnsForFetchingServiceUrls = true
    config.Region = "region-cn-hd-1"
    config.AvailabilityZones = map[string]string{
        "region-cn-hd-1": "zone-cn-hz-1",
    }
    config.EurekaServerDNSName = "dev.ms-registry.xf.io"
    config.EurekaServerUrlContext = "eureka"
    config.EurekaServerPort = "9001"

    // custom logger
    //eureka.SetLogger(func(format string, a ...interface{}) {
    //    fmt.Println("[custom logger] " + format, a)
    //})

    // run eureka client async
    eureka.DefaultClient.Config(config).
        Register("APP_ID_CLIENT_FROM_DNS", 9000).
        Run()

    select {}
}

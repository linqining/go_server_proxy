package main

import (
    "log"
    "github.com/zalando/skipper"
    "github.com/zalando/skipper/filters"
    "fmt"
)

var Config Cfg

func main() {
    Config = GetConfig()
    fmt.Println(Config)
    log.Fatal(skipper.Run(skipper.Options{
        Address:  ":"+Config.Port,
        CertPathTLS: Config.CertPathTLS,
        KeyPathTLS: Config.KeyPathTLS,
        RoutesFile: "routes.eskip",
        CustomFilters: []filters.Spec{&helloSpec{}}}))
}
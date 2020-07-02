package main

import (
    "log"
    "github.com/zalando/skipper"
    "github.com/zalando/skipper/filters"
)

func main() {
    log.Fatal(skipper.Run(skipper.Options{
        Address: ":9090",
        RoutesFile: "routes.eskip",
        CustomFilters: []filters.Spec{&helloSpec{}}}))
}
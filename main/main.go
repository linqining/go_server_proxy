package main

import (
    "log"

    "github.com/zalando/skipper"
    "github.com/zalando/skipper/filters"
    "github.com/zalando/skipper/routing"
)

func main() {
    log.Fatal(skipper.Run(skipper.Options{
        Address: ":9090",
        RoutesFile: "routes.eskip",
        CustomPredicates: []routing.PredicateSpec{&randomSpec{}},
        CustomFilters: []filters.Spec{&helloSpec{}}}))
}
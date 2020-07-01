package main
import (
    "github.com/zalando/skipper/routing"
    "math/rand"
    "net/http"
)

type randomSpec struct {}

type randomPredicate struct {
    chance float64
}

func (s *randomSpec) Name() string { return "Random" }

func (s *randomSpec) Create(args []interface{}) (routing.Predicate, error) {
    p := &randomPredicate{.5}
    if len(args) > 0 {
        if c, ok := args[0].(float64); ok {
            p.chance = c
        }
    }

    return p, nil
}

func (p *randomPredicate) Match(_ *http.Request) bool {
    return rand.Float64() < p.chance
}
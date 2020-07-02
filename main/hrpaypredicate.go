package main
import (
	"github.com/zalando/skipper/routing"
	"math/rand"
	"net/http"
)

type hrpaySpec struct {}

type hrpayPredicate struct {
	chance float64
}

func (s *hrpaySpec) Name() string { return "hrpay" }

func (s *hrpaySpec) Create(args []interface{}) (routing.Predicate, error) {
	p := &randomPredicate{.5}
	if len(args) > 0 {
		if c, ok := args[0].(float64); ok {
			p.chance = c
		}
	}

	return p, nil
}

func (p *hrpayPredicate) Match(_ *http.Request) bool {
	return rand.Float64() < p.chance
}
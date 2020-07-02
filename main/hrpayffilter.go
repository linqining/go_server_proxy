package main

import (
	"fmt"
	"github.com/zalando/skipper/filters"
	"strings"
)

type hrpayfSpec struct {}

type hrpayfFilter struct {
	who string
}

func (s *hrpayfSpec) Name() string { return "hrpayf" }

func (s *hrpayfSpec) CreateFilter(config []interface{}) (filters.Filter, error) {
	if len(config) == 0 {
		return nil, filters.ErrInvalidFilterParameters
	}

	if who, ok := config[0].(string); ok {
		return &hrpayfFilter{who}, nil
	} else {
		return nil, filters.ErrInvalidFilterParameters
	}
}

func (f *hrpayfFilter) Request(ctx filters.FilterContext) {
	res_url :=ctx.Request().URL
	rep_str := strings.Replace(res_url.Path,"/hrpay","",1)
	var err error
	ctx.Request().URL, err= ctx.Request().URL.Parse(rep_str)
	if err!=nil{
		fmt.Println(err)
	}
}

func (f *hrpayfFilter) Response(ctx filters.FilterContext) {
	ctx.Response().Header.Set("X-hrpay", fmt.Sprintf("hrpay, %s!", f.who))
}
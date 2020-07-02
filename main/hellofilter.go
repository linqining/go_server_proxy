package main

import (
    "compress/gzip"
    "github.com/PuerkitoBio/goquery"
    "strings"
    "fmt"
    "github.com/zalando/skipper/filters"
    "io"
    "io/ioutil"
    "regexp"
    "log"
)

type helloSpec struct {}

type helloFilter struct {
    who string
}

func (s *helloSpec) Name() string { return "hello" }

func (s *helloSpec) CreateFilter(config []interface{}) (filters.Filter, error) {
    if len(config) == 0 {
        return nil, filters.ErrInvalidFilterParameters
    }

    if who, ok := config[0].(string); ok {
        return &helloFilter{who}, nil
    } else {
        return nil, filters.ErrInvalidFilterParameters
    }
}

func (f *helloFilter) Request(ctx filters.FilterContext) {

}

func (f *helloFilter) Response(ctx filters.FilterContext) {
    request_url := ctx.Request().URL
    app_reg := regexp.MustCompile(".*app\\..*\\.js")

    if app_reg.MatchString(request_url.Path){
        bodyRes, err := readRes(ctx)

        if err!=nil{
            fmt.Printf("读取错误%s", err)
        }

        replace_str := replaceDomain(bodyRes)

        _, werr:=ctx.ResponseWriter().Write(replace_str)
        if werr!=nil{
            log.Print(werr)
        }
    }

    if ctx.Response().Header.Get("Content-Type")=="text/html" {
        bodyRes, err := readRes(ctx)

        fmt.Println(ctx.Response().Header.Get("Content-Type"))

        if err!=nil{
            fmt.Printf("读取错误%s", err)
        }

        replace_str := replaceDomain(bodyRes)

        doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(replace_str)))
        if err!=nil{

        }
        if doc!=nil{
            head := doc.Find("head")
            head.AppendHtml(customCSS())
            head.Find("title").SetText("法律咨询")
        }
        var result string
        result, _ = doc.Html()

        _, err =ctx.ResponseWriter().Write([]byte(result))
        if err!=nil{
            log.Print(err)
        }
    }
}

func replaceDomain(res []byte) []byte {
    reg := regexp.MustCompile("ai.12348.gov.cn")
    replace_str := reg.ReplaceAll(res,[]byte("www.taofa.cn:9090"))

    reg = regexp.MustCompile("hrpay.laway.cn")
    replace_str = reg.ReplaceAll(replace_str,[]byte("www.taofa.cn:9090/hrpay"))

    reg = regexp.MustCompile("newsystem.laway.cn")
    replace_str = reg.ReplaceAll(replace_str,[]byte("www.taofa.cn:9090/newsystem"))

    return replace_str
}

func readRes(ctx filters.FilterContext)([]byte,error){
    var reader io.ReadCloser
    var err error
    if ctx.Response().Header.Get("Content-Encoding") == "gzip" {
        reader, err = gzip.NewReader(ctx.Response().Body)
    } else {
        reader = ctx.Response().Body
    }
    bodyRes, err := ioutil.ReadAll(reader)
    return bodyRes, err
}

func customCSS()string{
    return "<style>.advisory-info-main,.page-header,.pc-header,.pc-footer,.page-header-con,.footer-container,.advisory-main li:nth-child(n+17){display:none!important}</style>"+
        "<style>.page-header-home,.page-banner,.tec,.advisory-icon-list li:nth-child(n+17){display:none !important}</style>"
}

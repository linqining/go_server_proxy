package main
import (
    "bytes"
    //"bytes"
    "fmt"
    "github.com/zalando/skipper/filters"
    "io/ioutil"
    "io"
    "compress/gzip"
    "regexp"
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
    //fmt.Printf("请求的url是%s", request_url)
    app_reg := regexp.MustCompile(".*app\\..*\\.js")
    if app_reg.MatchString(request_url.Path){
        fmt.Printf("找到匹配的url%s", request_url)

    }

    if ctx.Response().Header.Get("Content-Type")=="text/html" || app_reg.MatchString(request_url.Path){
        //var reader io.ReadCloser
        //var err error
        //if ctx.Response().Header.Get("Content-Encoding") == "gzip" {
        //    reader, err = gzip.NewReader(ctx.Response().Body)
        //    if err != nil {
        //        return
        //    }
        //} else {
        //    reader = ctx.Response().Body
        //}
        //bodyRes, err := ioutil.ReadAll(reader)
        bodyRes, err := readRes(ctx)

        fmt.Println(ctx.Response().Header.Get("Content-Type"))

        if err!=nil{
            fmt.Printf("读取错误%s", err)
        }
        reg := regexp.MustCompile("https://ai.12348.gov.cn|https://newsystem.laway.cn")

        replace_str := reg.ReplaceAll(bodyRes,[]byte("http://localhost:9090"))

        reg = regexp.MustCompile("ai.12348.gov.cn|newsystem.laway.cn")

        replace_str = reg.ReplaceAll(replace_str,[]byte("localhost:9090"))

        //reg2 := regexp.MustCompile("newsystem.laway.cn")

        ctx.ResponseWriter().Write(replace_str)
    }
}

func gzipFast(a *[]byte) []byte {
    var b bytes.Buffer
    gz := gzip.NewWriter(&b)
    if _, err := gz.Write(*a); err != nil {
        gz.Close()
        panic(err)
    }
    gz.Close()
    return b.Bytes()
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

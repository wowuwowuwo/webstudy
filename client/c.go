package main

import (
    "net/http"
    "strings"
    "fmt"
    "io/ioutil"
    "time"
)

// 用于读取resp的body
func helpRead(resp *http.Response)  {
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("ERROR2!: ", err)
    }
    // fmt.Printf("response: %v\n", resp)
    // fmt.Printf("response headers: %s\n", resp.Header)
    fmt.Printf("body: %s", string(body))
    resp.Body.Close()
}

func main() {
    // 下面测试binding数据
    // 首先测试binding-JSON,
    // 注意Body中的数据必须是JSON格式
    var resp *http.Response

    resp,_ = http.Get("http://0.0.0.0:8888/test6?name=BBB&passwd=CCC")
    helpRead(resp)
    resp,_ = http.Post("http://0.0.0.0:8888/test7?name=DDD&passwd=EEE", "",strings.NewReader(""))
    helpRead(resp)

    resp,_ = http.Post("http://0.0.0.0:8888/bindJSON", "application/json", strings.NewReader("{\"user\":\"TAO\", \"password\": \"123\"}"))
    helpRead(resp)

    // 下面测试bind FORM数据
    resp,_ = http.Post("http://0.0.0.0:8888/bindForm", "application/x-www-form-urlencoded", strings.NewReader("user=TAO&password=123"))
    helpRead(resp)

    // 下面测试接收JSON和XML数据
    resp,_ = http.Get("http://0.0.0.0:8888/someJSON")
    helpRead(resp)
    resp,_ = http.Get("http://0.0.0.0:8888/moreJSON")
    helpRead(resp)
    resp,_ = http.Get("http://0.0.0.0:8888/someXML")
    helpRead(resp)

    // index
    // create dir
    for i := 0; i < 10; i++ {
        go func(i int) {
            for j := 0; j < 1000; j++ {
                fmt.Printf("============== %d: %d ======================\n", i, j)
                resp, _ = http.Post("http://0.0.0.0:8888/?op=create&name=test_dir&current_dir=root", "",strings.NewReader(""))
                helpRead(resp)
                time.Sleep(10 * time.Millisecond)
            }
        }(i)
    }
    time.Sleep(3000 * time.Second)
}

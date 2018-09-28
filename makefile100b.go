package main


import (
    "time"
    "fmt"
    "os"
    "math/rand"
    "io/ioutil"
    "sync"
)


var (
    //  定义待存储文件目录名
    _dir = "./result100b"

    // 定义文件个数
    F int64 = 1000000

    // 生成的文件最大字节数
    // var MAX_BYTE int64 = 1000000
    MAX_BYTE int64 = 100
    wg sync.WaitGroup

    // 文件名前缀
    FILE_PRIFIX string = "MASS"

    // 文件后缀名
    FILE_POSTFIX string = ".bcp"
)


func dosomeing(filef *os.File, j int, FILE int64, wa *sync.WaitGroup) {
    for i:=1; int64(i) <= FILE; i++{
        file_name := _dir + "/" + FILE_PRIFIX + "_" + fmt.Sprintf("%d", i) + "_" + fmt.Sprintf("%d", j) + FILE_POSTFIX
        // fmt.Println("file_name:", file_name)

        // 生成随机数种子
        rand.Seed(time.Now().UnixNano())

        // 返回64位的随机数
        _NUM := rand.Int63n(MAX_BYTE)+1
        // fmt.Println("_NUM", _NUM)
        buffer := make([]byte,_NUM)
        // fmt.Println("buffer", string(buffer))

        filef.Seek(0, 0)
        filef.Read(buffer)
        // fmt.Println("buffer", string(buffer))

        //  写文件
        err := ioutil.WriteFile(file_name, buffer, 0644)
        if err != nil {
            fmt.Println(err)
        }

        if i%100 == 0 {
            fmt.Println(i, )
        }

        wa.Done()

    }
}


func main() {
    now := time.Now()
    var FILE int64 = F / 10

    //  创建目录
    err := os.Mkdir(_dir, 0777)
    if err != nil {
        fmt.Println(err)
    }

    // 读取文件
    file, err := os.Open("/etc/services")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()

    for j :=1; j <= 10; j++ {
        wg.Add(int(FILE))
        go dosomeing(file, j, FILE, &wg)
    }
    wg.Wait()

    fmt.Printf("总耗时：", time.Now().Sub(now))

}



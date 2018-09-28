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
    //  Define file directory name to be stored
    _dir = "./result100b"

    // Number of files 
    F int64 = 1000000

    // The maximum number of bytes generated by files
    MAX_BYTE int64 = 100
    wg sync.WaitGroup

    // Prefix of file name
    FILE_PRIFIX string = "MASS"

    // Suffix after file
    FILE_POSTFIX string = ".bcp"
)


func dosomeing(filef *os.File, j int, FILE int64, wa *sync.WaitGroup) {
    for i:=1; int64(i) <= FILE; i++{
        file_name := _dir + "/" + FILE_PRIFIX + "_" + fmt.Sprintf("%d", i) + "_" + fmt.Sprintf("%d", j) + FILE_POSTFIX
        // fmt.Println("file_name:", file_name)

        // Generating random number seeds
        rand.Seed(time.Now().UnixNano())

        // Return the random number of 64 bits, The value is not equal to 0.
        _NUM := rand.Int63n(MAX_BYTE)+1
        // fmt.Println("_NUM", _NUM)
        buffer := make([]byte,_NUM)
        // fmt.Println("buffer", string(buffer))

        filef.Seek(0, 0)
        filef.Read(buffer)
        // fmt.Println("buffer", string(buffer))

        //  Write data to file
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

    //  Create a directory
    err := os.Mkdir(_dir, 0777)
    if err != nil {
        fmt.Println(err)
    }

    // Read file data
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

    fmt.Printf("Total time consuming: ", time.Now().Sub(now))

}



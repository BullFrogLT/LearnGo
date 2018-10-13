package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"sync"
	"time"
)

var (
	//  Define file directory name to be stored
	_dir = "./result100b"

	// Number of files
	F int64 = 1000000

	// The maximum number of bytes generated by files
	MaxByte int64 = 10
	wg	sync.WaitGroup

	// Prefix of file name
	FilePrefix = "MASS"

	// Suffix after file
	FilePostfix = ".bcp"
)


func WhiteField(fief *os.File, j int, FILE int64, wa *sync.WaitGroup) {
	for i:=1; int64(i) <= FILE; i++{
		fileName := _dir + "/" + FilePrefix + "_" + fmt.Sprintf("%d", i) + "_" + fmt.Sprintf("%d", j) + FilePostfix
		// fmt.Println("file_name:", file_name)

		// Generating random number seeds
		rand.Seed(time.Now().UnixNano())

		// Return the random number of 64 bits, The value is not equal to 0.
		_NUM := rand.Int63n(MaxByte)+1
		// fmt.Println("_NUM", _NUM)
		buffer := make([]byte,_NUM)
		// fmt.Println("buffer", string(buffer))

		fief.Seek(0, 0)
		fief.Read(buffer)
		// fmt.Println("buffer", string(buffer))

		//  Write data to file
		err := ioutil.WriteFile(fileName, buffer, 0644)
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
	var FILE = F / 10

	//  Create a directory
	err := os.Mkdir(_dir, 0777)
	if err != nil {
		fmt.Println(err)
	}

	// Read file data
	// Linux and windows
	//file, err := os.Open("/etc/services")
	file, err := os.Open("C:\\Windows\\System32\\drivers\\etc\\hosts")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	for j :=1; j <= 10; j++ {
		wg.Add(int(FILE))
		go WhiteField(file, j, FILE, &wg)
	}
	wg.Wait()

	fmt.Printf("Total time consuming: %s", time.Now().Sub(now))

}



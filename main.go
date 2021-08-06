package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)
var (
	FilePath string
	SizeToPump int
)


func main() {
	FilePath = os.Args[1]
	v := os.Args[2]
	i, _ := strconv.Atoi(v)
	SizeToPump = i
	var wantedSize = int64(SizeToPump * 1024 * 1024) //Makes a MB

	fmt.Println("PumpItUp")

	fi, _ := os.Stat(FilePath)

	fmt.Println("Starting Size", fi.Size(), "bytes")

	toPump, err := os.OpenFile(FilePath, os.O_RDWR, 0644)
	if err != nil {
		log.Fatalf("Error Opening File: %s", err)
	}
	defer toPump.Close()

	_, err = toPump.WriteAt([]byte{0}, fi.Size()+wantedSize)
	if err != nil {
		log.Fatalf("Error Writing to File: %s", err)
	}

	fi, _ = os.Stat(FilePath)
	fmt.Println("Pumped Size", fi.Size() / 1024 /1024, "MB")
}

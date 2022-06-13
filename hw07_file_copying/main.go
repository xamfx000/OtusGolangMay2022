package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	from, to      string
	limit, offset int64
)

func init() {
	flag.StringVar(&from, "from", "", "file to read from")
	flag.StringVar(&to, "to", "", "file to write to")
	flag.Int64Var(&limit, "limit", 0, "limit of bytes to copy")
	flag.Int64Var(&offset, "offset", 0, "offset in input file")
}

func main() {
	flag.Parse()
	if from == "" || to == "" {
		fmt.Println("need to specify -from and -to params")
		os.Exit(0)
	}
	err := Copy(from, to, offset, limit)
	if err != nil {
		fmt.Printf("%v\n", err.Error())
		return
	}
	println("Done!")
}

package main

import (
	"log"
	"os"
)

type Query struct {
	path string
	err  error
}

//var ErrFileNotFound = errors.New("not in the root directory")

//var ErrnotFound = errors.New("not found")

func (q *Query) Error() string {
	return q.path + " : " + " input " + q.err.Error()
}

func main() {
	_, err := openFile1("test.txt", "root")
	if err != nil {
		log.Println(err)
		return
	}
}

func openFile1(fileName string, dirName string) (*os.File, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, &Query{
			path: "cd drive",
			err:  err,
		}

	}
	return f, nil
}

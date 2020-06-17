package main

import (
	"log"
	"os"

	"github.com/pkg/errors"
)

func setupLogging() {
	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()
	log.SetOutput(file)
}

//Config is a
type Config struct {
	//parsing the fields
}

func readFile(path string) (*Config, error) {
	file, error := os.Open(path)
	//fmt.Println(file.Read())
	if error != nil {
		return nil, errors.Wrap(error, "Error in opening. Please place the file in the folder")
	}
	defer file.Close()
	cfg := &Config{}
	return cfg, nil
}
func main() {
	_, err := readFile("/home/vathirajan/project/Golang/g.go")

	if err != nil {
		//	fmt.Fprintln(os.Stderr, "error: ", err)
		//	fmt.Println("*** Printing out logs **")

		log.Printf("error : %+v", err)
		os.Exit(1)
	}
	//	fmt.Println(_)
}

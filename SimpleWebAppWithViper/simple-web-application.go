package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	cnf        *Config
	configPath string
)

func initConfig() error {
	flag.StringVar(&configPath, "c", "config.yaml", "Configuration File")
	flag.Parse()

	c, err := NewCfg(configPath)
	if err != nil {
		return err
	}
	cnf = c

	return err
}

func main() {
	err := initConfig()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, you've requested: %s\n", r.URL.Path)
	})
	fmt.Println(fmt.Sprintf("Your server are running on %s:%d", cnf.HttpCfg().Host, cnf.HttpCfg().Port))
	http.ListenAndServe(fmt.Sprintf("%s:%d", cnf.HttpCfg().Host, cnf.HttpCfg().Port), nil)
}

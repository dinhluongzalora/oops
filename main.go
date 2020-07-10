package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"gopkg.in/yaml.v2"
)

type shell struct {
	Shell string `yaml:"shell"`
}

type runhttp struct {
	Method         string            `yaml:"method"`
	Url            string            `yaml:"url"`
	Headers        map[string]string `yaml:"headers"`
	QueryParams    string            `yaml:"query_params"`
	RequestPayload string            `yaml:"request_payload"`
	ResponseOut    string            `yaml:"response_out"`
	Asserts        shell             `yaml:"asserts"`
}

type run struct {
	Http runhttp `yaml:"http"`
}

type test struct {
	Env     map[string]string `yaml:"env"`
	Run     []run             `yaml:"run"`
	Asserts shell             `yaml:"asserts"`
}

func main() {
	log.SetFlags(0)
	log.SetOutput(os.Stdout)

	c := exec.Command("sh", "-c", "/home/f14t/gopath/src/github.com/flowerinthenight/oops/sh")
	c.Env = os.Environ()
	b, err := c.CombinedOutput()
	if err != nil {
		log.Println(err)
	}

	log.Printf("%v <-- val", string(b))

	yamlFile, err := ioutil.ReadFile("/home/f14t/gopath/src/github.com/flowerinthenight/oops/test.yaml")
	if err != nil {
		log.Fatal(err)
	}

	var t test
	err = yaml.Unmarshal(yamlFile, &t)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v", t)

	d1 := []byte(t.Run[0].Http.Asserts.Shell)
	err = ioutil.WriteFile("/tmp/dat1", d1, 0777)
	c = exec.Command("sh", "-c", "/tmp/dat1")
	c.Env = os.Environ()
	b, err = c.CombinedOutput()
	if err != nil {
		log.Printf("failed: %v", err)
	}

	log.Printf("%v", string(b))
}
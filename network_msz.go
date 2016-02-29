package network_msz

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct{
  Port string
  Url string
}

func (p Page) Status(){
  res, err := http.Get(p.Url)
	if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}

func (p Page) StatusHead(){
  res, err := http.Head(p.Url)
	if err != nil {
		log.Fatal(err)
	}
	// robots, err := ioutil.ReadAll(res.Body)
	// res.Body.Close()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	fmt.Printf("%s", res)
}


func (p Page) IsIp(){
  fmt.Printf(p.Url)
  return
}

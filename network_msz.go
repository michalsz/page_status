package network_msz

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/tatsushid/go-fastping"
	"os"
	"net"
	"time"
)

type Page struct{
  Port string
  Url string
	Ip string
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
  fmt.Printf("Checking: %s\n", p.Url)
	res, err := http.Head(p.Url)
	if err != nil {
		fmt.Println("Serwer nie dziala\n")

		makePing(p.Ip)
	}else{
		if res.StatusCode == 200 {
		  fmt.Println("Serwer dziala")
		}
	}
}

func makePing(ip string){
	fmt.Println("Spradzam po adresie IP\n")
	fmt.Println(ip)

	p := fastping.NewPinger()
  ra, err := net.ResolveIPAddr("ip4:icmp", ip)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
	p.AddIPAddr(ra)
  p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
    fmt.Printf("IP Addr: %s receive, RTT: %v\n", addr.String(), rtt)
  }
  p.OnIdle = func() {
    fmt.Println("finish")
  }
  err = p.Run()
  if err != nil {
    fmt.Println(err)
  }
}

func (p Page) IsIp(){
  fmt.Printf("Checking: %s\n", p.Url)
}

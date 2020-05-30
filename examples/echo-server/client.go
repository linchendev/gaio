package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"time"
)

var addr = flag.String("addr", "", "server addr")

func main() {
	flag.Parse()

	wait := make(chan int)
	go func() {
		conn, err := net.Dial("tcp", *addr)
		if err != nil {
			log.Fatal("connect fail:", err)
		}
		go func() {
			rx := make([]byte, 20)
			for {
				if n, err := conn.Read(rx); err != nil {
					log.Fatal("read fail:", err)
				} else {
					fmt.Println("read:", string(rx), n)
				}
			}
		}()
		ticker := time.NewTicker(time.Second * 5)
		for {
			select {
			case <-ticker.C:
				//push := []byte(fmt.Sprintf("%s\n", t))
				if n, err := conn.Write([]byte("hello world\n")); err != nil {
					log.Fatal("writer fail:", err)
				} else {
					fmt.Println("write:", n)
				}
			}
		}
		fmt.Println("exit")
		conn.Close()
		close(wait)
	}()
	<-wait
}

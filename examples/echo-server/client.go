package main

import (
	"flag"
	//"fmt"
	"log"
	"net"
	"time"
)

var addr = flag.String("addr", "", "server addr")

func main() {
	flag.Parse()

	for i := 0; i < 1; i++ {
		go func() {
			conn, err := net.Dial("tcp", *addr)
			if err != nil {
				log.Fatal("connect fail:", err)
			}

			//for i := 0; i < 1; i++ {
			//	tx := []byte("hello world:")
			//	tx = append(tx, []byte(fmt.Sprintf("%d", i))...)
			//	rx := make([]byte, len(tx))
			//	_, err = conn.Write(tx)
			//	if err != nil {
			//		log.Fatal(err)
			//	}

			//	log.Println("tx:", string(tx))
			//	_, err = conn.Read(rx)
			//	if err != nil {
			//		log.Fatal("read fail:", err)
			//	}
			//	log.Println("rx:", string(rx))
			//}
			conn.Close()
		}()
	}
	time.Sleep(time.Second * 5)
}

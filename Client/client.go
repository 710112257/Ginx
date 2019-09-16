package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	fmt.Println("test client...")
	time.Sleep(time.Second * 3)

	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	dealErr(err)

	for {
		_, err := conn.Write([]byte("hello"))
		dealErr(err)
		buf := make([]byte, 512)
		cut, err := conn.Read(buf)
		dealErr(err)

		fmt.Printf("date:%s ,cut:%d \n", buf, cut)
		time.Sleep(time.Second * 1)
	}

}
func dealErr(err error) {
	if err != nil {
		error.Error(err)
	}
}

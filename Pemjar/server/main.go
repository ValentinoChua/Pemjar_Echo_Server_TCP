package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatal("gagal membuka port 9001: ", err)
	}
	defer ln.Close()
	log.Println("server echo mendengarkan di port 9001")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("gagal menerima sambungan:", err)
			continue
		}
		go layani(conn)
	}
}

func layani(conn net.Conn) {
	defer conn.Close()
	alamat := conn.RemoteAddr().String()
	log.Println("klien tersambung:", alamat)
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		baris := scanner.Text()
		log.Printf("dari %s: %q", alamat, baris)
		if _, err := fmt.Fprintf(conn, "%s\n", baris); err != nil {
			log.Println("gagal mengirim ke", alamat, err)
			return
		}
	}
	log.Println("klien terputus:", alamat)
}

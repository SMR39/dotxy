package tcp

import (
	"log"
	"net"

	"github.com/dotxy/dns"
)

// ServeTCP will listen to TCP requests and redirects them to the DNS resolver
func ServeTCP(listenAddr, resolverAddr string, verbose bool) error {

	laddr, err := net.ResolveTCPAddr("tcp", listenAddr)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.ListenTCP("tcp", laddr)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := conn.Accept()
		if err != nil {
			continue
		}
		go handleClient(verbose, conn, resolverAddr)
	}
}

func handleClient(verbose bool, conn net.Conn, resolverAddr string) error {
	defer conn.Close()
	query := make([]byte, 2000)
	n, err := conn.Read(query)
	if err != nil {
		return err
	}
	query = query[2:n]
	resp, err := dns.ResolveDNSQuery(verbose, resolverAddr, query)
	if err != nil {
		log.Println("error", err)
	}
	_, err = conn.Write([]byte(resp))
	if err != nil {
		log.Println("error", err)
	}
	if verbose {
		log.Println("sent results to", conn.RemoteAddr())
	}
	return nil
}

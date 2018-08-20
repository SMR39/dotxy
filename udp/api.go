package udp

import (
	"log"
	"net"

	"github.com/SMR39/dotxy/dns"
)

//ServeUDP will listen to UDP requests and redirects them to the DNS resolver
func ServeUDP(listenAddr, resolverAddr string, verbose bool) error {

	laddr, err := net.ResolveUDPAddr("udp", listenAddr)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.ListenUDP("udp", laddr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for {
		query := make([]byte, 1232)

		n, addr, err := conn.ReadFromUDP(query)
		if err != nil {
			log.Println("error", err)
			continue
		}
		if verbose {
			log.Println("received query from", addr)
		}
		query = query[:n]
		resp, err := dns.ResolveDNSQuery(verbose, resolverAddr, query)
		if err != nil {
			log.Println("error", err)
			continue
		}
		_, err = conn.WriteToUDP(resp, addr)
		if err != nil {
			log.Println("error", err)
			continue
		}
		if verbose {
			log.Println("sent results to", addr)
		}
	}
}
